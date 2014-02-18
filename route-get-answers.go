package main

import (
	"encoding/json"
	"fmt"
	"github.com/hoisie/web"
)

func routeGetAnswers(ctx *web.Context) string {
	ctx.ContentType("json")

	list := []*Answer{}

	answers, _ := answerRepository.Answers()

	for _, answer := range answers {
		list = append(list, answer)
	}

	json, _ := json.Marshal(list)

	return fmt.Sprintf("%s", json)
}
