package main
import (
    "net/http"
	"github.com/arran4/golang-ical"
)

func fetchCalendar(requestURL string) (*ics.Calendar, error) {
	res, err := http.Get(requestURL)
	if err != nil {
        return nil, err
	}
    cal, err := ics.ParseCalendar(res.Body)
    if err != nil {
        return nil, err
    }
    return cal, nil
}
func setEvents(base ics.Calendar, events []ics.Component) (ics.Calendar) {
    components := []ics.Component{}
    for _, e := range base.Components {
    switch e.(type) {
	    case *ics.VEvent:
		continue
	    default:
		components = append(components, e)
	}
    }
    for _, e := range events {
	components = append(components, e)
    }

    return ics.Calendar{CalendarProperties: base.CalendarProperties, Components: components}

}

func subtractCalendarTimeslots(base ics.Calendar, sub ics.Calendar) {
    base_events := base.Events()
    new_events := []ics.VEvent{}
    for _, base_elem := range base_events {
	for _, sub_elem := range sub.Events() {
	    new_events = append(new_events, subtractEventTimes(base_elem, sub_elem)[:]...)
	}
    }
}

func subtractEventTimes(base_elem *ics.VEvent, sub_elem *ics.VEvent) []ics.VEvent {
    TimeDuration{start: base_elem.GetStartAt(), end: base_elem.GetEndAt()}.subtractTimeDuration(TimeDuration{start:sub_elem.GetStartAt(), end:sub_elem.GetEndAt()})
}
