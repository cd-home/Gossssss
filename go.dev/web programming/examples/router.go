package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RouterServer() {
	// thirdParty MUX
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/{user}/{name}/", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Println(vars["user"], vars["name"])
	})

	fs := http.FileServer(http.Dir("static/"))
	// Need user New router to handle
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", router)
}
