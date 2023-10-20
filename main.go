package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"github.com/arran4/golang-ical"
)

func routeRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello\n")
}
func routeCalendar(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		w.WriteHeader(http.StatusOK)
	case "PROPFIND":
		w.WriteHeader(http.StatusOK)
	default:
		w.Header().Set("Content-Type", "text/calendar")
		w.Write([]byte(createDummyCalSerialized()))
	}
}
func createDummyCalSerialized() (string){
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	event := cal.AddEvent("test event")
	event.SetCreatedTime(time.Now())
	event.SetDtStampTime(time.Now())
	event.SetModifiedAt(time.Now())
	event.SetStartAt(time.Now())
	event.SetEndAt(time.Now())
	event.SetSummary("another 30 min test - 15:00")
	event.SetLocation("Address")
	event.SetDescription("Description")
	event.SetURL("https://URL/")
	event.AddRrule(fmt.Sprintf("FREQ=YEARLY;BYMONTH=%d;BYMONTHDAY=%d", time.Now().Month(), time.Now().Day()))
	event.SetOrganizer("sender@domain", ics.WithCN("This Machine"))
	event.AddAttendee("reciever or participant", ics.CalendarUserTypeIndividual, ics.ParticipationStatusNeedsAction, ics.ParticipationRoleReqParticipant, ics.WithRSVP(true))
	// thunderbird seems to have a lower limit of 30 min
	cal.SetRefreshInterval("PT1M")
	return cal.Serialize()
}

func main() {
	cal, _ := fetchCalendar("http://www.htwk-stundenplan.de/353b419d/")
		// for _,e := range cal.Components {
			fmt.Printf("%#v\n", cal)
		// }
	fmt.Print("test")
	http.HandleFunc("/", routeRoot)
	http.HandleFunc("/calendar", routeCalendar)
	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
