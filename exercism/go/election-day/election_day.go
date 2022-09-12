package electionday

import(
	"fmt"
)

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	var x *int
	x = &initialVotes
	return x
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter==nil{
		return 0
	}
	var x int
	x = *counter
	return x
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
//	fmt.Println("pointer: ",counter)
//	fmt.Println("value: ",*counter)
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var x *ElectionResult
	x = &ElectionResult{Name: candidateName, Votes: votes}
	return x
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] = results[candidate]-1
}
