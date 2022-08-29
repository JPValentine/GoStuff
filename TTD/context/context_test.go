package main

import(
	"time"
	"testing"
	"net/http"
	"net/http/httptest"
	"context"
)

func TestServer(t *testing.T){
	t.Run("returns data from store", func(t *testing.T){	
		data:= "hello world"
		store := &SpyStore{response: data,t: t}
		svr:= Server(store)

		request:= httptest.NewRequest(http.MethodGet, "/", nil)
		response:= httptest.NewRecorder()

		svr.ServeHTTP(response,request)

		if response.Body.String() != data{
			t.Errorf("got %s, want %s",response.Body.String(), data)
		}//if

		if store.cancelled{
			t.Errorf("should not have cancelled the store!")
		}
		store.AssertWasNotCancelled()
	})
	t.Run("tells store to cancel is a task is cancelled", func(t *testing.T){
		data:= "hello world"
		store:= &SpyStore{response: data,t: t}
		svr:= Server(store)

		request:= httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancelled{
			t.Error("store was not told to cancel")
		}
		store.AssertWasCancelled()
	})	
}

