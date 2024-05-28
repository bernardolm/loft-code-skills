package main

import (
	"time"

	"github.com/k0kubun/pp/v3"
)

func Feed(realties *[]Realty, visitors *[]Visitor, a *Agenda) error {
	printDataUntilNow := true

	defer func() {
		if printDataUntilNow {
			pp.Println("realties", realties)
			pp.Println("visitors", visitors)
			pp.Println("schedules", schedules)
		}
	}()

	*realties = []Realty{
		{ID: 10},
		{ID: 20},
	}

	*visitors = []Visitor{
		{ID: 30},
		{ID: 40},
	}

	schedule1 := Schedule{
		Realty:  (*realties)[0],
		Visitor: (*visitors)[0],
		Begin:   time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
	}
	schedule1.End = schedule1.Begin.Add(1 * time.Hour)
	if err := a.Add(schedule1); err != nil {
		return err
	}

	schedule2 := Schedule{
		Realty:  (*realties)[0],
		Visitor: (*visitors)[1],
		Begin:   time.Date(2024, 05, 28, 15, 20, 0, 0, time.Local),
	}
	schedule2.End = schedule2.Begin.Add(1 * time.Hour)
	if err := a.Add(schedule2); err != nil {
		return err
	}

	schedule3 := Schedule{
		Realty:  (*realties)[1],
		Visitor: (*visitors)[0],
		Begin:   time.Date(2024, 05, 28, 18, 20, 0, 0, time.Local),
	}
	schedule3.End = schedule1.Begin.Add(1 * time.Hour)
	if err := a.Add(schedule3); err != nil {
		return err
	}

	printDataUntilNow = false

	return nil
}
