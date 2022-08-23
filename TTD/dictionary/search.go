package main

import(
	"errors"
)
type Dictionary map[string]string

var ErrNotFound = errors.New("couldn't find the the word you're looking for")

func (d Dictionary) Search(word string) (string,error){
	definition,isOkay := d[word]
	if !isOkay{
		return "", ErrNotFound
	}//if
	return definition,nil
}//Search

