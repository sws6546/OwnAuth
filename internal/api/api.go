package api

import (
	"fmt"
	"net/http"

	"github.com/sws6546/OwnAuth/config"
)

func ServeApi() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello")
	})

	fmt.Printf("Runnin at port %d", config.Port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", config.Port), nil)
}
