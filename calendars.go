package main
import (
    "net/http"
	"github.com/arran4/golang-ical"
)

func fetchCalendar(url string) (*ics.Calendar, error) {
    requestURL := "http://localhost:3333/calendar"
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

func subtractCalendarTimeslots(base ics.Calendar, sub ics.Calendar) {
    base_local := base.Events()
    for _, base_elem := range base_local {
	for _, sub_elem := range sub.Events() {
	}
	
    }
}
