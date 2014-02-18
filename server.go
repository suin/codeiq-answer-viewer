package main

import (
	"encoding/json"
	"fmt"
	"github.com/hoisie/web"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

var answerRepository = NewOnMemoryAnswerRepository()

func index(ctx *web.Context) string {
	return `<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
<script src="/assets/script.js"></script><noscript>Turn on JavaScript</noscript>`
}

func upload(ctx *web.Context) {
	file, _, err := ctx.Request.FormFile("file")

	if errorOnRedirect(ctx, err, "/") != nil {
		return
	}

	tmp, err := ioutil.TempFile("/tmp", "codeiq")

	if errorOnRedirect(ctx, err, "/") != nil {
		return
	}

	defer tmp.Close()
	defer os.Remove(tmp.Name())

	io.Copy(tmp, file) // 一旦実ファイルにしないと codeiq.NewReader が動かないため

	csv, err := NewReader(tmp.Name())

	if errorOnRedirect(ctx, err, "/") != nil {
		return
	}

	defer csv.Close()

	answerRepository.Purge()

	inAnser := false

	for {
		record, err := csv.Read() // 1行読み出す
		if err == io.EOF {
			break
		} else {
			if errorOnRedirect(ctx, err, "/") != nil {
				return
			}
		}

		if inAnser == true {
			answer := NewAnswerFromRecord(record)
			answerRepository.AddAnswer(answer)
		}

		if record[0] == "No" {
			inAnser = true
		}
	}

	ctx.Redirect(301, "/")
}

func errorOnRedirect(ctx *web.Context, err error, path string) (returnErr error) {
	returnErr = err

	if err != nil {
		log.Printf("[error] %s", err)
		setCookie(ctx, "error", fmt.Sprintf("%s", err))
		ctx.Redirect(301, path)
	}

	return
}

func setCookie(ctx *web.Context, name string, value string) {
	ctx.SetCookie(web.NewCookie(name, url.QueryEscape(value), 0))
}

func getCookie(ctx *web.Context, name string) (value string) {
	cookie, err := ctx.Request.Cookie(name)
	if err != nil {
		return
	}

	value, err = url.QueryUnescape(cookie.Value)

	if err != nil {
		return
	}

	return value
}

func getAnswerLists(ctx *web.Context) string {
	ctx.SetHeader("Content-Type", "application/json; charset=UTF-8", true)

	list := []*Answer{}

	answers, _ := answerRepository.Answers()

	for _, answer := range answers {
		list = append(list, answer)
	}

	json, _ := json.Marshal(list)

	return fmt.Sprintf("%s", json)
}

func getAnswer(ctx *web.Context, no string) (result string) {
	ctx.SetHeader("Content-Type", "application/json; charset=UTF-8", true)
	answer, err := answerRepository.NoOf(no)

	if err == AnswerNotFound {
		ctx.NotFound("Answer not found")
		return
	}

	if err != nil {
		ctx.Abort(500, "Something wrong")
		return
	}

	json, _ := json.Marshal(answer)
	return fmt.Sprintf("%s", json)
}

func main() {
	web.Get("/", index)
	web.Post("/upload", upload)
	web.Get("/answers", getAnswerLists)
	web.Get("/answers/(.*)", getAnswer)

	web.Get("/assets/(.*)", func(ctx *web.Context, file string) (contents string) {
		if len(file) == 0 {
			ctx.NotFound("File not found")
			return
		}

		pwd, _ := os.Getwd()
		filename := pwd + "/assets/" + file

		if fileExists(filename) == true {
			// for developments
			file, err := os.Open(filename)

			if err != nil {
				ctx.Abort(500, errorToString(err))
				return
			}

			ctx.ContentType(filepath.Ext(filename))

			contentsBytes, err := ioutil.ReadAll(file)

			if err != nil {
				ctx.Abort(500, errorToString(err))
				return
			}

			contents = fmt.Sprintf("%s", contentsBytes)

			return
		}

		if _, ok := assets[file]; ok {
			ctx.ContentType(filepath.Ext(filename))
			contents = assets[file]

			return
		}

		ctx.NotFound("File not found")

		return
	})

	web.Run("0.0.0.0:51019")
}

func errorToString(err error) string {
	return fmt.Sprintf("%s", err)
}
