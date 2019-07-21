package handler

import "net/http"

func Initialize(pr string) *http.ServeMux {
	h := handler{prefix: pr}

	mux := http.NewServeMux()
	mux.HandleFunc("/", h.redirect)

	return mux
}

type handler struct {
	prefix string
}

func (handler) redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//code:= r.URL.Path[len("/"):]
	//Add the pull from db then redirect to correct url
	http.Redirect(w, r, "http://www.google.com", 301)
}
