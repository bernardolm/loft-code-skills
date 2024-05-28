package main

import (
	"testing"
	"time"

	"github.com/k0kubun/pp/v3"
)

func TestAgendamentos_Add(t *testing.T) {
	fixture := make(Agendamentos)
	a := Agendamento{
		IDImovel:    99,
		IDVisitante: 88,
		Inicio:      time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
		Fim:         time.Date(2024, 05, 28, 14, 20, 0, 0, time.Local),
	}
	if err := fixture.Add(a); err != nil {
		t.Errorf("erro não esperado: deveria permitir agendamento para: %s", a)
	}

	a.IDVisitante = 77
	if err := fixture.Add(a); err == nil {
		t.Errorf("erro esperado: não deveria permitir agendamento para: %s", a)
	}

	a.Inicio = time.Date(2024, 05, 28, 17, 20, 0, 0, time.Local)
	a.Fim = time.Date(2024, 05, 28, 17, 50, 0, 0, time.Local)
	if err := fixture.Add(a); err != nil {
		t.Errorf("erro não esperado: deveria permitir agendamento para: %s", a)
	}

	a.Inicio = time.Date(2024, 05, 28, 13, 25, 0, 0, time.Local)
	a.Fim = time.Date(2024, 05, 28, 14, 15, 0, 0, time.Local)
	if err := fixture.Add(a); err == nil {
		t.Errorf("erro esperado: não deveria permitir agendamento para: %s", a)
	}

	a.Inicio = time.Date(2024, 05, 28, 14, 15, 0, 0, time.Local)
	a.Fim = time.Date(2024, 05, 28, 15, 15, 0, 0, time.Local)
	if err := fixture.Add(a); err == nil {
		t.Errorf("erro esperado: não deveria permitir agendamento para: %s", a)
	}

	a.Inicio = time.Date(2024, 05, 28, 14, 20, 0, 0, time.Local)
	a.Fim = time.Date(2024, 05, 28, 15, 20, 0, 0, time.Local)
	if err := fixture.Add(a); err != nil {
		t.Errorf("erro não esperado: deveria permitir agendamento para: %s", a)
	}

	a.Inicio = time.Date(2024, 05, 28, 20, 20, 0, 0, time.Local)
	a.Fim = time.Date(2024, 05, 28, 20, 20, 0, 0, time.Local)
	if err := fixture.Add(a); err == nil {
		t.Errorf("erro esperado: não deveria permitir agendamento para: %s", a)
	}

	a.Inicio = time.Date(2024, 05, 28, 13, 00, 0, 0, time.Local)
	a.Fim = time.Date(2024, 05, 28, 14, 00, 0, 0, time.Local)
	if err := fixture.Add(a); err == nil {
		t.Errorf("erro esperado: não deveria permitir agendamento para: %s", a)
	}

	a.Inicio = time.Date(2024, 05, 28, 9, 00, 0, 0, time.Local)
	a.Fim = time.Date(2024, 05, 28, 10, 00, 0, 0, time.Local)
	if err := fixture.Add(a); err != nil {
		t.Errorf("erro não esperado: deveria permitir agendamento para: %s", a)
	}

	pp.Println("agendamentos finais", fixture)
}

func TestAgendamentos_Remove(t *testing.T) {
	fixture := make(Agendamentos)
	a := Agendamento{
		IDImovel:    99,
		IDVisitante: 88,
		Inicio:      time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
		Fim:         time.Date(2024, 05, 28, 14, 20, 0, 0, time.Local),
	}
	_ = fixture.Add(a)

	if err := fixture.Remove(a); err != nil {
		t.Errorf("erro não esperado: deveria permitir remover o agendamento: %s", a)
	}

	a.IDImovel = 100
	if err := fixture.Add(a); err == nil {
		t.Errorf("erro esperado: não deveria permitir remover o agendamento ausente: %s", a)
	}
}

func TestAgendamentos_Confirm(t *testing.T) {
	fixture := make(Agendamentos)
	a := Agendamento{
		IDImovel:    99,
		IDVisitante: 88,
		Inicio:      time.Date(2024, 05, 28, 13, 20, 0, 0, time.Local),
		Fim:         time.Date(2024, 05, 28, 14, 20, 0, 0, time.Local),
	}
	_ = fixture.Add(a)

	if err := fixture.Confirm(a); err != nil {
		t.Errorf("erro não esperado: deveria permitir confirmar o agendamento: %s", a)
	}

	a.IDImovel = 100
	if err := fixture.Confirm(a); err == nil {
		t.Errorf("erro esperado: não deveria permitir confirmar o agendamento ausente: %s", a)
	}
}
