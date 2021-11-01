package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

func checkEvents(db *gorm.DB) {
	events, err := services.FetchEvents()
	if err != nil {
		utils.Log("An error occurred while fetching events: " + err.Error())
		return
	}
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
				services.DispatchEvent("event_changed", *dbEvent, *event)
			}
		} else if len(dbEvents) > 0 {
			services.DispatchEvent("new_event", nil, *event)
		}
	}

	db.Save(&events)

	for _, event := range dbEvents {
		event := event
		if eventsMap[event.EventID] == nil {
			db.Delete(&event)
			services.DispatchEvent("removed_event", nil, event)
		}
	}
}
