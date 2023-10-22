package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

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
		w.Write([]byte(calculateCal()))
	}
}
func calculateCal() (string){
	cal1, e := os.OpenFile("./testcal1.ics", os.O_RDONLY, 0)
	if e != nil {
		fmt.Println(e)
	}
	cal2, e := os.OpenFile("./testcal2.ics", os.O_RDONLY, 0)
	if e != nil {
		fmt.Println(e)
	}
	cal1_parsed, e := ics.ParseCalendar(cal1)
	if e != nil {
		fmt.Println(e)
	}
	cal2_parsed, e := ics.ParseCalendar(cal2)
	if e != nil {
		fmt.Println(e)
	}
	// fmt.Printf("%+v",cal1_parsed.Events())
	// fmt.Printf("%+v",cal2_parsed.Events())
	result_cal := subtractCalendarTimeslots(*cal1_parsed, *cal2_parsed)
	for _,e := range result_cal.Events() {
		start, _ :=  e.GetStartAt()
		end, _ :=  e.GetEndAt()
		fmt.Printf("%+v - %+v\n", start, end)
	}
	return result_cal.Serialize()
}
func createCal() (ics.Calendar){
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)

	// thunderbird seems to have a lower limit of 30 min
	cal.SetRefreshInterval("PT1M")
	return *cal
}

func main() {
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
