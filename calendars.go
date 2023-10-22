package main

import (
	"fmt"
	"net/http"

	ics "github.com/arran4/golang-ical"
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
func setEvents(events []*ics.VEvent) (ics.Calendar) {
    cal := ics.NewCalendar()
    cal.SetMethod(ics.MethodRequest)
    cal.SetRefreshInterval("PT1H")

    for i, e := range events {
	cal.AddVEvent(events[i])
	start, _ :=  e.GetStartAt()
	end, _ :=  e.GetEndAt()
	fmt.Printf("(0) %+v = %+v - %+v\n", e.Id(), start, end)
    }
    for _,e := range cal.Events() {
	start, _ :=  e.GetStartAt()
	end, _ :=  e.GetEndAt()
	fmt.Printf("(-1) %+v = %+v - %+v\n", e.Id(), start, end)
    }

    return *cal

}
func remove(array []*ics.VEvent, element *ics.VEvent) []*ics.VEvent {
    var idx *int
    for i, e := range array {
	if (element == e) {
	    idx = &i
	}
    }
    if (idx != nil) {
	array[*idx] = array[len(array)-1]
	return array[:len(array)-1]
    }
    return array
}

func subtractCalendarTimeslots(base ics.Calendar, sub ics.Calendar) ics.Calendar {
    base_events := base.Events()
    for _, base_elem := range base_events {
	for _, sub_elem := range sub.Events() {
	    for _,e := range subtractEventTimes(*base_elem, *sub_elem){
		start, _ :=  e.GetStartAt()
		end, _ :=  e.GetEndAt()
		fmt.Printf("(2) %+v - %+v\n", start, end)
		base_events = append(base_events,&e)
		remove(base_events,base_elem)
	    }
	}
    }
    for _,e := range base_events {
	start, _ :=  e.GetStartAt()
	end, _ :=  e.GetEndAt()
	fmt.Printf("(1) %+v - %+v\n", start, end)
    }
    return setEvents(base_events)
}

func subtractEventTimes(base_elem ics.VEvent, sub_elem ics.VEvent) []ics.VEvent {
    start1,_ := base_elem.GetStartAt()
    end1,_ := base_elem.GetEndAt()
    start2,_ := sub_elem.GetStartAt()
    end2,_ := sub_elem.GetEndAt()
    r := TimeDuration{start: Time{time:start1}, end: Time{time:end1}}.subtractTimeDuration(TimeDuration{start: Time{time:start2}, end: Time{time: end2}})
    for _,e := range r {
	fmt.Printf("output: %+v - %+v\n", e.start.time.String(), e.end.time.String())
    }
    switch len(r) {
    case 2: 
	event1 := ics.NewEvent("28093479023490283")
	event1.SetStartAt(r[0].start.time)
	event1.SetEndAt(r[0].end.time)
	event1.SetSummary("Summary")
	event2 := ics.NewEvent("19090859239203984")
	event2.SetStartAt(r[1].start.time)
	event2.SetEndAt(r[1].end.time)
	event2.SetSummary("Summary")
	ret := []ics.VEvent{*event1, *event2}
	return ret
    case 1:
	event1 := ics.NewEvent("free")
	event1.SetStartAt(r[0].start.time)
	event1.SetEndAt(r[0].end.time)
	return []ics.VEvent{*event1}
    default:
	return []ics.VEvent{}
}
}
