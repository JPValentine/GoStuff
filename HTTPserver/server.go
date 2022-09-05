package main

import(
	"net/http"
	"fmt"
	"strings"
)

type PlayerStore interface{
	GetPlayerScore(name string) int
}

type PlayerServer struct{
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request)(){
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.store.GetPlayerScore(player)

	if score == 0{
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(name string) string{
	if name == "Pepper"{
		return "20"
	}
	if name == "Bob"{
		return "10"
	}
	return ""
}//GetPlayerScore

