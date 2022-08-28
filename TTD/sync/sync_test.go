package main

import(
	"testing"
	"sync"
)

func TestCounter(t *testing.T){
	t.Run("increment counter by 3 times leaves it 3", func(t *testing.T){
		c := NewCounter()
		c.Inc()
		c.Inc()
		c.Inc()
		assertCounter(t,c,3)	
	})
	t.Run("safety check concurrently", func(t * testing.T){
		wantedCount := 1000
		c := NewCounter()
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i:=0;i<wantedCount;i++{
			go func(){
				c.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, c, wantedCount)
	})
	
}//TestCounter

func assertCounter(t testing.TB, got *Counter, want int){
	t.Helper()
	if got.Value() != want{
		t.Errorf("got %d, but wanted %d",got.Value(), want)
	}//if
}


