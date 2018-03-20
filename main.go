package main

import "log"

type Team struct {
	Id         string
	Challenges *[]Challenge
}
type Challenge struct {
	Id        string
	Histories *[]History
}
type History struct {
	Id string
}

func main() {
	challenge := &Challenge{
		Id: "1",
	}
	challenges := []Challenge{
		*challenge,
	}
	team := &Team{
		Id:         "1",
		Challenges: &challenges,
	}

	newChallenge := append(*team.Challenges,
		Challenge{
			Id: "2",
		})
	team.Challenges = &newChallenge
	currentChallenges := *team.Challenges
	for i, v := range currentChallenges {
		if v.Id == "2" {
			var newHistories []History
			if v.Histories != nil {
				newHistories = append(*v.Histories,
					History{
						Id: "1",
					})
			} else {
				newHistories = []History{
					History{
						Id: "1",
					},
				}
			}
			currentChallenges[i].Histories = &newHistories
		}
	}
	team.Challenges = &currentChallenges

	log.Println(team.Id)
	log.Println("Histories:" + (*(*team.Challenges)[1].Histories)[0].Id)
}
