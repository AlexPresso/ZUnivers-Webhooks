package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

const NewEventEvent = "new_event"
const EventChangedEvent = "event_changed"
const EventRemovedEvent = "event_removed"

func checkEvents(db *gorm.DB, embeds *[]discord.Embed) {
	if utils.EventsAllDisabled([]string{NewEventEvent, EventChangedEvent, EventRemovedEvent}) {
		return
	}

	events, resSpec, err := services.FetchEvents()
	if err != nil {
		utils.Log("An error occurred while fetching events: " + err.Error())
		return
	}

	checkResponse(db, embeds, resSpec)

	eventsMap := make(map[string]*structures.Event)
	for _, event := range events {
		event := event
		eventsMap[event.EventID] = &event
	}

	var dbEvents []structures.Event
	db.Find(&dbEvents)
	dbEventsMap := make(map[string]*structures.Event)
	for _, event := range dbEvents {
		event := event
		dbEventsMap[event.EventID] = &event
	}

	for i := 0; i < len(events); i++ {
		event := &events[i]
		dbEvent := dbEventsMap[event.EventID]

		if dbEvent != nil {
			event.ID = dbEvent.ID

			if utils.AreDifferent(*event, *dbEvent) {
				services.MakeEmbed(EventChangedEvent, *dbEvent, *event, embeds)
			}
		} else if len(dbEvents) > 0 {
			services.MakeEmbed(NewEventEvent, nil, *event, embeds)
		}
	}

	db.Save(&events)

	for _, event := range dbEvents {
		event := event
		if eventsMap[event.EventID] == nil {
			db.Delete(&event)
			services.MakeEmbed(EventRemovedEvent, nil, event, embeds)
		}
	}
}
