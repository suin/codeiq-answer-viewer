package main

import (
	"github.com/hoisie/web"
)

func routeGetindex(ctx *web.Context) string {
	return `<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
<script src="/assets/script.js"></script><noscript>Turn on JavaScript</noscript>`
}
