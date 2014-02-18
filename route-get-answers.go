package main

import (
	"encoding/json"
	"fmt"
	"github.com/hoisie/web"
)

func routeGetAnswers(ctx *web.Context) string {
	ctx.SetHeader("Content-Type", "application/json; charset=UTF-8", true)

	list := []*Answer{}

	answers, _ := answerRepository.Answers()

	for _, answer := range answers {
		list = append(list, answer)
	}

	json, _ := json.Marshal(list)

	return fmt.Sprintf("%s", json)
}
