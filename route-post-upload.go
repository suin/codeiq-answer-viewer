package main

import (
	"github.com/hoisie/web"
	"io"
	"io/ioutil"
	"os"
)

func routePostUpload(ctx *web.Context) {
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
