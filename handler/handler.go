package handler

import (
	"UrlShortener/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/valyala/fasthttp"
)

//Initialize receives a prefix string and create a mux route handler and passing the request to the handlers
func Initialize(pr string, storage *storage.Service) func(*fasthttp.RequestCtx) {
	h := handler{prefix: pr, db: *storage}
	mux := func(ctx *fasthttp.RequestCtx) {
		//	r := bytes.NewReader(ctx.Path())
		//	path := r
		switch string(string(ctx.Path())) {
		case "/":
			h.redirect(ctx)
		case "/save/":
			url, _, err := h.saveURL(ctx)
			if err != nil {
				ctx.Error("not found", fasthttp.StatusBadRequest)
			}
			urlb, err := json.Marshal(url)
			ctx.Response.SetBody(urlb)
		default:
			fmt.Println(string(ctx.Path()))
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}
	return mux
}

type handler struct {
	prefix string
	db     storage.Service
}

func (h *handler) redirect(ctx *fasthttp.RequestCtx) {
	if !ctx.IsGet() {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}
	//code:= r.URL.Path[len("/"):]
	//Add the pull from db then redirect to correct url
	ctx.Redirect("http://www.google.com", 301)
}

func (h *handler) saveURL(ctx *fasthttp.RequestCtx) (interface{}, int, error) {
	if !ctx.IsPost() {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return nil, fasthttp.StatusBadRequest, fmt.Errorf("method %s not allowed", ctx.Method())
	}
	var input struct {
		URL string `json:"url"`
	}

	if err := json.Unmarshal(ctx.PostBody(), &input); err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("Unable to decode JSON request body: %v", err)
	}

	url := strings.TrimSpace(input.URL)
	if url == "" {
		return nil, http.StatusBadRequest, fmt.Errorf("URL is empty")
	}

	if !strings.Contains(url, "http") {
		url = "http://" + url
	}
	c, err := h.db.Save(url)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Could not store in database: %v", err)
	}

	return h.prefix + c, http.StatusCreated, nil
}
