package http

import(
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
)

func StartServer(port string) error {
	r := mux.NewRouter()

	r.HandleFunc("/knockknock", KnockKnock)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func KnockKnock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	w.Write([]byte("Who's there?"))
}