package main

import (
	"reflect"

	"github.com/k0kubun/pp/v3"
)

// req 0: registrar 3 visitas e exibí-las
// req 1: garantir que não seja possível agendar uma visita no mesmo apartamento no mesmo horário
// req 2: permitir que uma visita agendada seja cancelada
// req 3: registrar que uma visita agendada foi realmente realizada

var (
	realties  = []Realty{}
	visitors  = []Visitor{}
	schedules = make(Agenda)
)

func main() {
	if err := Feed(&realties, &visitors, &schedules); err != nil {
		panic(err)
	}

	pp.Println("initial schedules", schedules)

	keys := reflect.ValueOf(schedules).MapKeys()

	toRemove := schedules[keys[1].String()]
	if err := schedules.Remove(toRemove); err != nil {
		panic(err)
	}

	toConfirm := schedules[keys[0].String()]
	if err := schedules.Confirm(toConfirm); err != nil {
		panic(err)
	}

	pp.Println("final schedules", schedules)
}
