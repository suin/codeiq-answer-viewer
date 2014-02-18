package main

import (
	"encoding/json"
	"fmt"
	"github.com/hoisie/web"
)

func routeGetAnswer(ctx *web.Context, no string) (result string) {
	ctx.ContentType("json")
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
