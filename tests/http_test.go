package testsss


import (
	"testing"
	"net/http"
)

func TestHttpServer(t *testing.T) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Logf(r.URL.Host)
	})

	http.ListenAndServe(":8080", nil)
}