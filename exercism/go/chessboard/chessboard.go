package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	total:=0
	f:=cb[file]
	for _,v:= range f{
		if v{total++}
	}
	return total
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	total:=0
	for k,_ := range cb{
		for i,v := range cb[k]{
			if ((rank-1) == i)&&v==true{
				total++
			}
		}
	}
	return total
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	total:=0
	for k,_ := range cb{
		for range cb[k]{
			total++
		}
	}
	return total
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	total:=0
	for k,_ := range cb{
		for _,v := range cb[k]{
			if v{total++}
		}
	}
	return total

}
