package handler

import (
	"github.com/valyala/fasthttp"
)

func Initialize(pr string) func(*fasthttp.RequestCtx) {
	h := handler{prefix: pr}
	mux := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			h.redirect(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}
	return mux
}

type handler struct {
	prefix string
}

func (handler)  redirect(ctx *fasthttp.RequestCtx)  {
	if !ctx.IsGet() {
		 ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}
	//code:= r.URL.Path[len("/"):]
	//Add the pull from db then redirect to correct url
	ctx.Redirect("http://www.google.com", 301)
}
