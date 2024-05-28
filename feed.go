package main

import (
	"time"
)

func feed(a Agenda) error {
	schedule1 := Schedule{
		IDRealty:  10,
		IDVisitor: 20,
		Begin:     time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
	}
	schedule1.End = schedule1.Begin.Add(1 * time.Hour)
	if err := a.Add(schedule1); err != nil {
		panic(err)
	}

	schedule2 := Schedule{
		IDRealty:  30,
		IDVisitor: 40,
		Begin:     time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
	}
	schedule2.End = schedule2.Begin.Add(1 * time.Hour)
	if err := a.Add(schedule2); err != nil {
		panic(err)
	}

	schedule3 := Schedule{
		IDRealty:  30,
		IDVisitor: 50,
		Begin:     time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
	}
	schedule3.End = schedule1.Begin.Add(1 * time.Hour)
	if err := a.Add(schedule3); err != nil {
		panic(err)
	}

	return nil
}
