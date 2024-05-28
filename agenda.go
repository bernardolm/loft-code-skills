package main

import (
	"fmt"
)

type Agenda map[string]Schedule

func (a Agenda) Add(new Schedule) error {
	if new.Begin.Equal(new.End) {
		return fmt.Errorf("schedule from %s to %s is invalid\n",
			new.BeginF(), new.EndF())
	}

	errorMsg := "schedule from %s to %s at realty %d already scheduled to visitor %d\n"

	if a, ok := a[new.ID()]; ok {
		return fmt.Errorf(errorMsg, new.BeginF(), new.EndF(), new.Realty.ID, a.Visitor.ID)
	}

	for _, schedule := range a {
		if schedule.Realty.ID == new.Realty.ID {
			if new.Begin.Equal(schedule.Begin) ||
				(new.Begin.After(schedule.Begin) && new.Begin.Before(schedule.End)) ||
				(new.Begin.Before(schedule.Begin) && new.End.After(schedule.Begin)) {
				return fmt.Errorf(errorMsg, new.BeginF(), new.EndF(), new.Realty.ID, new.Visitor.ID)
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
			schedule.BeginF(), schedule.EndF(), schedule.Realty.ID)
	}

	delete(a, schedule.ID())

	return nil
}

func (a Agenda) Confirm(schedule Schedule) error {
	if _, ok := a[schedule.ID()]; !ok {
		return fmt.Errorf("there isn't a schedule from %s to %s at realty %d\n",
			schedule.BeginF(), schedule.EndF(), schedule.Realty.ID)
	}

	ag := a[schedule.ID()]
	ag.Realized = true
	a[schedule.ID()] = ag

	return nil
}
