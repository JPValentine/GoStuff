package main

import(
	"time"
	"net/http"
	"fmt"
	"testing"
)

type Store interface{
	Fetch(ctx context.Context) (string, error)
}

type SpyStore struct{
	response string
	t *testing.T
}

func (s *SpyStore) AssertWasCancelled(){
	s.t.Helper()
	if !s.cancelled{
		s.t.Error("Store was not told to cancel")
	}
}

func (s *SpyStore) AssertWasNotCancelled(){
	s.t.Helper()
	if s.cancelled{
		s.t.Error("Store was told to cancel")
	}
}

func (s *SpyStore) Fetch(ctx context.Context) (string,error){
	data := make(chan string, 1)
	go func(){
		var result string
		for _, c := range s.response{
			select{
			case<-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(20 * Millisecond)
				result += string(c)
			}
		}		
		data<-result
	}()
	select{
		case <- ctx.Done():
			return"", ctx.Err()
		case res := <-data:
			return res, nil
	}
}

func (s *SpyStore) Cancel(){
	s.cancelled = true
}

func Server(store Store) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		//	
	}
}


