package main

import "log"

type Team struct {
	Id         string
	Challenges *[]Challenge
}
type Challenge struct {
	Id        string
	Histories *[]History
	Logs      []Log
}
type History struct {
	Id string
}
type Log struct {
	Message string
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
			// newHistories = append(*v.Histories, History{Id: "1"})  // This code cause null pointer panic
			newLogs := append(v.Logs, Log{
				Message: "New Message",
			})

			currentChallenges[i].Histories = &newHistories
			currentChallenges[i].Logs = newLogs
		}
	}
	team.Challenges = &currentChallenges

	log.Println(team.Id)
	log.Println("Histories:" + (*(*team.Challenges)[1].Histories)[0].Id)
	log.Println("Logs:" + (*team.Challenges)[1].Logs[0].Message)
}
