package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

func checkChallenges(db *gorm.DB, embeds *[]discord.Embed) {
	chProgress, err := services.FetchChallenges()
	if err != nil {
		utils.Log("An error occurred while fetching challenges: " + err.Error())
		return
	}

	dbChallengesMap := make(map[string]*structures.Challenge)
	var challenges []*structures.Challenge

	for i := 0; i < len(chProgress); i++ {
		challenge := &chProgress[i].Challenge
		challenges = append(challenges, *challenge)
		dbChallenge := dbChallengesMap[(*challenge).ChallengeID]

		if dbChallenge != nil {
			(*challenge).ID = dbChallenge.ID

			if utils.AreDifferent(**challenge, *dbChallenge) {
				*embeds = append(*embeds, *services.MakeEmbed("challenge_changed", *dbChallenge, **challenge))
			}
		} else if len(dbChallengesMap) > 0 {
			*embeds = append(*embeds, *services.MakeEmbed("new_challenge", nil, **challenge))
		}
	}

	db.Save(&challenges)
}
