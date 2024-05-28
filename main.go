package main

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
)

// req 0: registrar 3 visitas e exibí-las
// req 1: garantir que não seja possível agendar uma visita no mesmo apartamento no mesmo horário
// req 2: permitir que uma visita agendada seja cancelada
// req 3: registrar que uma visita agendada foi realmente realizada

var (
	realties = []Realty{
		{ID: 10},
		{ID: 20},
	}
	visitors = []Visitor{
		{ID: 30},
		{ID: 40},
	}
)

func main() {
	fmt.Println(realties)
	fmt.Println(visitors)

	schedules := make(Agenda)

	if err := Feed(realties, visitors, schedules); err != nil {
		panic(err)
	}

	pp.Printf("schedules:\n%#v\n", schedules)
}
