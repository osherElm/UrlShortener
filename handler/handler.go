package handler

import (
	"fmt"

	"../storage"

	"github.com/valyala/fasthttp"
)

//Initialize receives a prefix string and create a mux route handler and passing the request to the handlers
func Initialize(pr string, storage storage.Service) func(*fasthttp.RequestCtx) {
	h := handler{prefix: pr, db: storage}
	mux := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			h.redirect(ctx)
		case "/save/":

		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}
	return mux
}

type handler struct {
	prefix string
	db     storage.Service
}

func (handler) redirect(ctx *fasthttp.RequestCtx) {
	if !ctx.IsGet() {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}
	//code:= r.URL.Path[len("/"):]
	//Add the pull from db then redirect to correct url
	ctx.Redirect("http://www.google.com", 301)
}

func (handler) saveUrl(ctx *fasthttp.RequestCtx) (interface{}, int, error) {
	if !ctx.IsPost() {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return nil, fasthttp.StatusBadRequest, fmt.Errorf("method %s not allowed", ctx.Method())
	}
}
