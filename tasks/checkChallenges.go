package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

const NewChallengeEvent = "new_challenge"
const ChallengeChangedEvent = "challenge_changed"

func checkChallenges(db *gorm.DB, embeds *[]discord.Embed) {
	if utils.EventsAllDisabled([]string{NewChallengeEvent, ChallengeChangedEvent}) {
		return
	}

	chProgress, resSpec, err := services.FetchChallenges()
	if err != nil {
		utils.Log("An error occurred while fetching challenges: " + err.Error())
		return
	}

	checkResponse(db, embeds, resSpec)

	var challenges []*structures.Challenge
	var dbChallenges []structures.Challenge
	dbChallengesMap := make(map[string]*structures.Challenge)

	db.Find(&dbChallenges)
	for _, chall := range dbChallenges {
		chall := chall
		dbChallengesMap[chall.ChallengeID] = &chall
	}

	for i := 0; i < len(chProgress); i++ {
		challenge := &chProgress[i].Challenge
		challenges = append(challenges, *challenge)
		dbChallenge := dbChallengesMap[(*challenge).ChallengeID]

		if dbChallenge != nil {
			(*challenge).ID = dbChallenge.ID

			if utils.AreDifferent(**challenge, *dbChallenge) {
				services.MakeEmbed(ChallengeChangedEvent, *dbChallenge, **challenge, embeds)
			}
		} else if len(dbChallengesMap) > 0 {
			services.MakeEmbed(NewChallengeEvent, nil, **challenge, embeds)
		}
	}

	db.Save(&challenges)
}
