package main

import (
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
	defer func() {
		pp.Println("realties", realties)
		pp.Println("visitors", visitors)
		pp.Println("schedules", schedules)
	}()

	if err := Feed(&realties, &visitors, &schedules); err != nil {
		panic(err)
	}
}
