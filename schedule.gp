package main

import (
	"fmt"
	"time"

	"github.com/k0kubun/pp/v3"
)



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


