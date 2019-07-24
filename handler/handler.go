package handler

import (
	"UrlShortener/storage"
	"encoding/json"
	"fmt"
	"log"
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
			h.saveURL(ctx)
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
	log.Println("saving to db")
	if !ctx.IsPost() {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return nil, fasthttp.StatusBadRequest, fmt.Errorf("method %s not allowed", ctx.Method())
	}
	log.Println("saving to db 2")
	var input struct {
		URL string `json:"url"`
	}

	if err := json.Unmarshal(ctx.PostBody(), &input); err != nil {
		log.Println(err)
		return nil, http.StatusBadRequest, fmt.Errorf("Unable to decode JSON request body: %v", err)
	}
	//	log.Println(json.NewDecoder(string((ctx.PostBody()))
	// if err := json.NewDecoder(r).Decode(&input); err != nil {
	// 	log.Println(err)
	// 	return nil, http.StatusBadRequest, fmt.Errorf("Unable to decode JSON request body: %v", err)
	// }
	log.Println("inserting ur   l :  to db", input.URL)

	url := strings.TrimSpace(input.URL)
	if url == "" {
		return nil, http.StatusBadRequest, fmt.Errorf("URL is empty")
	}

	if !strings.Contains(url, "http") {
		url = "http://" + url
	}
	log.Println("inserting url : to db  ", url)
	c, err := h.db.Save(url)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Could not store in database: %v", err)
	}

	return h.prefix + c, http.StatusCreated, nil
}
