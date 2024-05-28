package main

import (
	"fmt"
	"time"

	"github.com/k0kubun/pp/v3"
)

// req 0: registrar 3 visitas e exibí-las
// req 1: garantir que não seja possível agendar uma visita no mesmo apartamento no mesmo horário
// req 2: permitir que uma visita agendada seja cancelada
// req 3: registrar que uma visita agendada foi realmente realizada

type Realty struct {
	id int
}

type Visitor struct {
	id int
}

type Schedule struct {
	IDRealty  int
	IDVisitor int
	Begin     time.Time
	End       time.Time
	Realized  bool
}

func (v Schedule) ID() string {
	return fmt.Sprintf("%d_%s_%s",
		v.IDRealty,
		v.Begin.Format("060102-0304"),
		v.End.Format("060102-0304"))
}

func (v Schedule) BeginF() string {
	return v.Begin.Format("02/01/2006 03h04")
}

func (v Schedule) EndF() string {
	return v.End.Format("02/01/2006 03h04")
}

func (v Schedule) String() string {
	return fmt.Sprintf("im: %d | vs: %d | in: %s | fm: %s\n",
		v.IDRealty, v.IDVisitor, v.BeginF(), v.EndF())
}

type Agenda map[string]Schedule

func (a Agenda) Add(new Schedule) error {
	if new.Begin.Equal(new.End) {
		return fmt.Errorf("schedule from %s to %s is invalid\n",
			new.BeginF(), new.EndF())
	}

	errorMsg := "schedule from %s to %s at realty %d already scheduled to visitor %d\n"

	if a, ok := a[new.ID()]; ok {
		return fmt.Errorf(errorMsg, new.BeginF(), new.EndF(), new.IDRealty, a.IDVisitor)
	}

	for _, schedule := range a {
		if schedule.IDRealty == new.IDRealty {
			if new.Begin.Equal(schedule.Begin) ||
				(new.Begin.After(schedule.Begin) && new.Begin.Before(schedule.End)) ||
				(new.Begin.Before(schedule.Begin) && new.End.After(schedule.Begin)) {
				return fmt.Errorf(errorMsg, new.BeginF(), new.EndF(), new.IDRealty, new.IDVisitor)
			}
		}
	}

	newCopy := new

	a[newCopy.ID()] = newCopy

	return nil
}

func (a Agenda) Remove(schedule Schedule) error {
	if _, ok := a[schedule.ID()]; !ok {
		return fmt.Errorf("there isn't a schedule from %s to %s at realty %d\n",
			schedule.BeginF(), schedule.EndF(), schedule.IDRealty)
	}

	delete(a, schedule.ID())

	return nil
}

func (a Agenda) Confirm(schedule Schedule) error {
	if _, ok := a[schedule.ID()]; !ok {
		return fmt.Errorf("there isn't a schedule from %s to %s at realty %d\n",
			schedule.BeginF(), schedule.EndF(), schedule.IDRealty)
	}

	ag := a[schedule.ID()]
	ag.Realized = true
	a[schedule.ID()] = ag

	return nil
}

var (
	realties = []Realty{}
	visitors = []Visitor{}
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

func main() {
	fmt.Println(realties)
	fmt.Println(visitors)

	schedules := make(Agenda)

	if err := feed(schedules); err != nil {
		panic(err)
	}

	pp.Printf("schedules:\n%#v\n", schedules)
}
