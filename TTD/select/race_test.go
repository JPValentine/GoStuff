package main
 
import(
	"testing"
	"time"
	"net/http"
	"net/http/httptest"
)

func makeDelayedServer(delay time.Duration) *httptest.Server{
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}//makeDelayedServer

func TestRacer(t *testing.T){
	t.Run("race check" , func(t *testing.T){
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(1 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()
		
		slowURL := slowServer.URL 
		fastURL := fastServer.URL

		want := fastURL
		got,err := Racer(slowURL, fastURL)

		if got != want{
			t.Errorf("got %q, want %q", got, want)
		}

		if err != nil{
			t.Fatalf("got an error but did not expect one")
		}
	})
	t.Run("timeout check", func(t *testing.T){
		x := makeDelayedServer(25 * time.Millisecond)
	
		defer x.Close()
		
		_, err := ConfigRacer(x.URL, x.URL, 20*time.Millisecond)
		if err== nil{
			t.Error("expected an error but didn't get one")
		}
	})
}

