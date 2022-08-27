package main
 
import(
	"testing"
	"time"
	"net/http"
//	"net/http/httptest"
)

func makeDelayedServer(delay time.Duration) *httptest.Server{
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}//makeDelayedServer

func TestRacer(t *testing.T){
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.google.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want{
		t.Errorf("got %q, want %q", got, want)
	}
}

