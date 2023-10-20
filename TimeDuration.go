package main
import (
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
	i := a.start.lessThan(b.start)
	j := a.end.lessThan(b.end)
	k := a.start.lessThan(b.end)
	l := a.end.lessThan(b.start)
	b_is_fully_inside_a := i && !j && k && !l
	overlap_at_start_of_a := !i && !j && k && !l
	overlap_at_end_of_a := (!i && j && !k && !l) || (i && j && k && !l)
	a_is_fully_inside_b := (i && j && !k && !l) || (!i && j && k && !l)
	if b_is_fully_inside_a {
		return []TimeDuration{
			TimeDuration{start: a.start, end:b.start},
			TimeDuration{start: b.end, end:a.end},
		}
	} else if overlap_at_start_of_a {
		return []TimeDuration{
			TimeDuration{start: b.end, end:a.end},
		}
	} else if overlap_at_end_of_a {
		return []TimeDuration{
			TimeDuration{start: a.start, end:b.start},
		}
	} else if a_is_fully_inside_b {
		return []TimeDuration{}
	}
	// b doesn't overlap with a, so we just return a
	return []TimeDuration{a}
}
