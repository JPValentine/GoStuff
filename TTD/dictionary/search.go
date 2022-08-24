package main

type Dictionary map[string]string
type DictionaryErr string

const(
	ErrNotFound =DictionaryErr("couldn't find the the word you're looking for")
	ErrWordExists =DictionaryErr("Cannot add a word because it already exsists")
	ErrWordDoesNotExist = DictionaryErr("Cannot update this word because it doesn't exist")
)//const

func (e DictionaryErr) Error() string{
	return string(e)
}//Error

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
func (d Dictionary) Update(word string, newDef string) error{

	_, err := d.Search(word)
	switch err{
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDef
	default:
		return err
	}//switch

	return nil
}

func (d Dictionary) Search(word string) (string,error){
	definition,isOkay := d[word]
	if !isOkay{
		return "", ErrNotFound
	}//if
	return definition,nil
}//Search

func (d Dictionary) Add(word string, definition string) error{
	_, err := d.Search(word)
	switch err{
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}//switch

	return nil
}

