package _911

type vote struct {
	time   int
	leader int
}

type TopVotedCandidate struct {
	votes       []vote
	personVotes map[int]int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	var votes []vote
	personVotes := map[int]int{}
	currentLeader := -1
	currentLeaderVotes := -1
	for i := 0; i < len(times); i++ {
		if len(votes) == 0 || votes[len(votes)-1].time != times[i] {
			votes = append(votes, vote{time: times[i]})
		}
		currentVote := &votes[len(votes)-1]
		personVotes[persons[i]] += 1
		if personVotes[persons[i]] >= currentLeaderVotes {
			currentLeader = persons[i]
			currentLeaderVotes = personVotes[persons[i]]
		}
		currentVote.leader = currentLeader
		//fmt.Println(votes)
		//fmt.Println(personVotes)
	}
	return TopVotedCandidate{votes: votes}
}

func (this *TopVotedCandidate) Q(t int) int {
	left := 0
	right := len(this.votes)
	for left < right-1 {
		mid := (left + right) >> 1
		if this.votes[mid].time == t {
			return this.votes[mid].leader
		} else if this.votes[mid].time < t {
			left = mid
		} else {
			right = mid
		}
	}
	return this.votes[left].leader
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
