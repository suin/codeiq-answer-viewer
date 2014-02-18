package main

import (
	"github.com/hoisie/web"
)

var answerRepository = NewOnMemoryAnswerRepository()

func main() {
	web.Get("/", routeGetindex)
	web.Post("/upload", routePostUpload)
	web.Get("/answers", routeGetAnswers)
	web.Get("/answers/(.*)", routeGetAnswer)
	web.Get("/assets/(.*)", routeGetAssets)
	web.Run("0.0.0.0:51019")
}
