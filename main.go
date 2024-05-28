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

type Imovel struct {
	id int
}

type Visitante struct {
	id int
}

type Agendamento struct {
	IDImovel    int
	IDVisitante int
	Inicio      time.Time
	Fim         time.Time
	Realizado   bool
}

func (v Agendamento) ID() string {
	return fmt.Sprintf("%d_%s_%s",
		v.IDImovel,
		v.Inicio.Format("060102-0304"),
		v.Fim.Format("060102-0304"))
}

func (v Agendamento) InicioF() string {
	return v.Inicio.Format("02/01/2006 03h04")
}

func (v Agendamento) FimF() string {
	return v.Fim.Format("02/01/2006 03h04")
}

func (v Agendamento) String() string {
	return fmt.Sprintf("im: %d | vs: %d | in: %s | fm: %s\n",
		v.IDImovel, v.IDVisitante, v.InicioF(), v.FimF())
}

type Agendamentos map[string]Agendamento

func (a Agendamentos) Add(novo Agendamento) error {
	if novo.Inicio.Equal(novo.Fim) {
		return fmt.Errorf("horário %s às %s inválido\n",
			novo.InicioF(), novo.FimF())
	}

	if a, ok := a[novo.ID()]; ok {
		return fmt.Errorf("horário %s às %s no imóvel %d já agendado para o visitante %d\n",
			novo.InicioF(), novo.FimF(), novo.IDImovel, a.IDVisitante)
	}

	for _, agendamento := range a {
		if agendamento.IDImovel == novo.IDImovel {
			if novo.Inicio.Equal(agendamento.Inicio) ||
				(novo.Inicio.After(agendamento.Inicio) && novo.Inicio.Before(agendamento.Fim)) ||
				(novo.Inicio.Before(agendamento.Inicio) && novo.Fim.After(agendamento.Inicio)) {
				return fmt.Errorf("horário %s às %s no imóvel %d já agendado para o visitante %d\n",
					novo.InicioF(), novo.FimF(), novo.IDImovel, novo.IDVisitante)
			}
		}
	}

	novoCopy := novo

	a[novoCopy.ID()] = novoCopy

	return nil
}

func (a Agendamentos) Remove(agendamento Agendamento) error {
	if _, ok := a[agendamento.ID()]; !ok {
		return fmt.Errorf("não existe agendamento no horário de %s às %s no imóvel %d\n",
			agendamento.InicioF(), agendamento.FimF(), agendamento.IDImovel)
	}

	delete(a, agendamento.ID())

	return nil
}

func (a Agendamentos) Confirm(agendamento Agendamento) error {
	if _, ok := a[agendamento.ID()]; !ok {
		return fmt.Errorf("não existe agendamento no horário de %s às %s no imóvel %d\n",
			agendamento.InicioF(), agendamento.FimF(), agendamento.IDImovel)
	}

	ag := a[agendamento.ID()]
	ag.Realizado = true
	a[agendamento.ID()] = ag

	return nil
}

var (
	apartamentos = []Imovel{}
	visitantes   = []Visitante{}
)

func feed(a Agendamentos) error {
	agendamento1 := Agendamento{
		IDImovel:    10,
		IDVisitante: 20,
		Inicio:      time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
	}
	agendamento1.Fim = agendamento1.Inicio.Add(1 * time.Hour)
	if err := a.Add(agendamento1); err != nil {
		return fmt.Errorf("não foi possível adicionar o agendamento: %#v", agendamento1)
	}

	agendamento2 := Agendamento{
		IDImovel:    30,
		IDVisitante: 40,
		Inicio:      time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
	}
	agendamento2.Fim = agendamento2.Inicio.Add(1 * time.Hour)
	if err := a.Add(agendamento2); err != nil {
		return fmt.Errorf("não foi possível adicionar o agendamento: %#v", agendamento2)
	}

	agendamento3 := Agendamento{
		IDImovel:    30,
		IDVisitante: 50,
		Inicio:      time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
	}
	agendamento3.Fim = agendamento1.Inicio.Add(1 * time.Hour)
	if err := a.Add(agendamento3); err != nil {
		return fmt.Errorf("não foi possível adicionar o agendamento: %#v", agendamento3)
	}

	return nil
}

func main() {
	fmt.Println(apartamentos)
	fmt.Println(visitantes)

	agendamentos := make(Agendamentos)

	if err := feed(agendamentos); err != nil {
		panic(err)
	}

	pp.Printf("%#v\n", agendamentos)
}
