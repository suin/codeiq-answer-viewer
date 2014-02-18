package main

import (
	"fmt"
	"github.com/hoisie/web"
	"log"
	"net/url"
)

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
