package main
import (
	"fmt"
	"time"
)

type Time struct {
	time time.Time
}

func (a Time) lessThan(b Time) bool {
	return a.time.Unix() < b.time.Unix()
}

type TimeDuration struct {
    start Time
    end Time
}

func (a TimeDuration) subtractTimeDuration(b TimeDuration) []TimeDuration {
	fmt.Printf("%+v - %+v\n",a.start.time.String(), a.end.time.String())
	fmt.Printf("%+v - %+v\n",b.start.time.String(), b.end.time.String())
	i := a.start.lessThan(b.start)
	j := a.end.lessThan(b.end)
	k := a.start.lessThan(b.end)
	l := a.end.lessThan(b.start)
	b_is_fully_inside_a := i && !j && k && !l 
	overlap_at_start_of_a := !i && !j && k && !l 
	overlap_at_end_of_a := i && j && k && !l 
	a_is_fully_inside_b := !i && j && k && !l 
	fmt.Printf("%t - %t - %t - %t\n", i,j,k,l)
	fmt.Printf("%t - %t - %t - %t\n", b_is_fully_inside_a, overlap_at_start_of_a, overlap_at_end_of_a, a_is_fully_inside_b)
	if b_is_fully_inside_a {
		return []TimeDuration{
			{start: a.start, end:b.start},
			{start: b.end, end:a.end},
		}
	} else if overlap_at_start_of_a {
		return []TimeDuration{
			{start: b.end, end:a.end},
		}
	} else if overlap_at_end_of_a {
		return []TimeDuration{
			{start: a.start, end:b.start},
		}
	} else if a_is_fully_inside_b {
		return []TimeDuration{}
	}
	// b doesn't overlap with a, so we just return a
	return []TimeDuration{a}
}
