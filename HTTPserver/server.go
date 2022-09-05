package main

import(
	"net/http"
	"fmt"
	"strings"
)

type PlayerStore interface{
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct{
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request)(){
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method{
	case http.MethodPost:
		p.processWin(w,player)
	case http.MethodGet:
		p.showScore(w,r,player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request, player string){
	score := p.store.GetPlayerScore(player)

	if score == 0{
		w.WriteHeader(http.StatusNotFound)
	}//if

	fmt.Fprint(w, score)
}//showScore

func (p *PlayerServer) processWin(w http.ResponseWriter, player string){
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}//processWin

func GetPlayerScore(name string) string{
	if name == "Pepper"{
		return "20"
	}
	if name == "Bob"{
		return "10"
	}
	return ""
}//GetPlayerScore


