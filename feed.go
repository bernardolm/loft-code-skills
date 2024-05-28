package main

import (
	"time"
)

func Feed(realties []Realty, visitors []Visitor, a Agenda) error {
	schedule1 := Schedule{
		Realty:  realties[0],
		Visitor: visitors[0],
		Begin:   time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
	}
	schedule1.End = schedule1.Begin.Add(1 * time.Hour)
	if err := a.Add(schedule1); err != nil {
		panic(err)
	}

	schedule2 := Schedule{
		Realty:  realties[1],
		Visitor: visitors[0],
		Begin:   time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
	}
	schedule2.End = schedule2.Begin.Add(1 * time.Hour)
	if err := a.Add(schedule2); err != nil {
		panic(err)
	}

	schedule3 := Schedule{
		Realty:  realties[1],
		Visitor: visitors[1],
		Begin:   time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
	}
	schedule3.End = schedule1.Begin.Add(1 * time.Hour)
	if err := a.Add(schedule3); err != nil {
		panic(err)
	}

	return nil
}
