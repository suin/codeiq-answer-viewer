package main

import (
	"fmt"
	"github.com/hoisie/web"
	"io/ioutil"
	"os"
	"path/filepath"
)

func routeGetAssets(ctx *web.Context, file string) (contents string) {
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
}
