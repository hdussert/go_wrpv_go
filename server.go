package main

import (
	"fmt"
	"net/http"
	"github.com/dimfeld/httptreemux"
)

func main() {
	router := httptreemux.New()
	
	group := router.NewGroup("/api")
	group.GET("/v1/:id", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
			id := params["id"]
			fmt.Fprintf(w, "GET /api/v1/%s", id)
	})
	
	http.ListenAndServe(":8080", router)
}