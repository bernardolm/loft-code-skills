package main

import (
	"testing"
	"time"

	"github.com/k0kubun/pp/v3"
)

func TestAgenda_Add(t *testing.T) {
	fixture := make(Agenda)
	a := Schedule{
		IDRealty:  99,
		IDVisitor: 88,
		Begin:     time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
		End:       time.Date(2024, 05, 28, 14, 20, 0, 0, time.Local),
	}
	if err := fixture.Add(a); err != nil {
		t.Errorf("unexpected failure: should be able to schedule: %s", a)
	}

	a.IDVisitor = 77
	if err := fixture.Add(a); err == nil {
		t.Errorf("expected failure: shouldn't be able to schedule: %s", a)
	}

	a.Begin = time.Date(2024, 05, 28, 17, 20, 0, 0, time.Local)
	a.End = time.Date(2024, 05, 28, 17, 50, 0, 0, time.Local)
	if err := fixture.Add(a); err != nil {
		t.Errorf("unexpected failure: should be able to schedule: %s", a)
	}

	a.Begin = time.Date(2024, 05, 28, 13, 25, 0, 0, time.Local)
	a.End = time.Date(2024, 05, 28, 14, 15, 0, 0, time.Local)
	if err := fixture.Add(a); err == nil {
		t.Errorf("expected failure: shouldn't be able to schedule: %s", a)
	}

	a.Begin = time.Date(2024, 05, 28, 14, 15, 0, 0, time.Local)
	a.End = time.Date(2024, 05, 28, 15, 15, 0, 0, time.Local)
	if err := fixture.Add(a); err == nil {
		t.Errorf("expected failure: shouldn't be able to schedule: %s", a)
	}

	a.Begin = time.Date(2024, 05, 28, 14, 20, 0, 0, time.Local)
	a.End = time.Date(2024, 05, 28, 15, 20, 0, 0, time.Local)
	if err := fixture.Add(a); err != nil {
		t.Errorf("unexpected failure: should be able to schedule: %s", a)
	}

	a.Begin = time.Date(2024, 05, 28, 20, 20, 0, 0, time.Local)
	a.End = time.Date(2024, 05, 28, 20, 20, 0, 0, time.Local)
	if err := fixture.Add(a); err == nil {
		t.Errorf("expected failure: shouldn't be able to schedule: %s", a)
	}

	a.Begin = time.Date(2024, 05, 28, 13, 00, 0, 0, time.Local)
	a.End = time.Date(2024, 05, 28, 14, 00, 0, 0, time.Local)
	if err := fixture.Add(a); err == nil {
		t.Errorf("expected failure: shouldn't be able to schedule: %s", a)
	}

	a.Begin = time.Date(2024, 05, 28, 9, 00, 0, 0, time.Local)
	a.End = time.Date(2024, 05, 28, 10, 00, 0, 0, time.Local)
	if err := fixture.Add(a); err != nil {
		t.Errorf("unexpected failure: should be able to schedule: %s", a)
	}

	pp.Println("final agenda", fixture)
}

func TestAgenda_Remove(t *testing.T) {
	fixture := make(Agenda)
	a := Schedule{
		IDRealty:  99,
		IDVisitor: 88,
		Begin:     time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
		End:       time.Date(2024, 05, 28, 14, 20, 0, 0, time.Local),
	}
	_ = fixture.Add(a)

	if err := fixture.Remove(a); err != nil {
		t.Errorf("unexpected failure: should allow you to remove the schedule: %s", a)
	}

	a.IDRealty = 100
	if err := fixture.Remove(a); err == nil {
		t.Errorf("expected failure: shouldn't allow you to remove a schedule that doesn't exist: %s", a)
	}
}

func TestAgenda_Confirm(t *testing.T) {
	fixture := make(Agenda)
	a := Schedule{
		IDRealty:  99,
		IDVisitor: 88,
		Begin:     time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
		End:       time.Date(2024, 05, 28, 14, 20, 0, 0, time.Local),
	}
	_ = fixture.Add(a)

	if err := fixture.Confirm(a); err != nil {
		t.Errorf("unexpected failure: should allow you to confirm the schedule: %s", a)
	}

	a.IDRealty = 100
	if err := fixture.Confirm(a); err == nil {
		t.Errorf("expected failure: shouldn't allow you to confirm a schedule that doesn't exist: %s", a)
	}
}
