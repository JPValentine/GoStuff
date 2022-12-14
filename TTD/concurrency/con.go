package main

type WebChecker func(string) bool
type result struct{
	string
	bool
}
func CheckWebsites(wc WebChecker, urls []string) map[string]bool{

	results := make(map[string]bool)
	resChannel:= make(chan result)
	for _, url := range urls{
		go func(u string){
			resChannel <- result{u, wc(u)}
		}(url)
	}
	for i:=0;i<len(urls);i++{
		r:= <-resChannel
		results[r.string] = r.bool
	}
	return results
}//CheckWebsites

