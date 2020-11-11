package administrador

import (
	"fmt"
	"time"
)

type AdmProcesos struct {
	AdmProcesos []Procesos
}

func (a *AdmProcesos) Mostrar() {
	for _, v := range a.AdmProcesos {
		v.Mostrar()
	}
}

func (a *AdmProcesos) Ocultar() {
	for _, v := range a.AdmProcesos {
		v.Ocultar()
	}
}

func (a *AdmProcesos) Eliminar(id uint64) {
	for i, v := range a.AdmProcesos {
		if id == v.GetID() {
			v.Eliminar()
			a.AdmProcesos = append(a.AdmProcesos[:i], a.AdmProcesos[i+1:]...)
		}
	}
}

type Procesos interface {
	Iniciar()
	Mostrar()
	Ocultar()
	Eliminar()
	GetID() uint64
}

type Proceso struct {
	Id              uint64
	I               uint64
	banderaMostrar  bool
	banderaEliminar bool
}

func (p *Proceso) Iniciar() {
	p.I = 0
	p.banderaMostrar = false
	p.banderaEliminar = false
	for {
		if p.banderaMostrar {
			fmt.Printf("id %d: %d", p.Id, p.I)
			fmt.Println()
		}
		p.I = p.I + 1
		time.Sleep(time.Millisecond * 500)
		if p.banderaEliminar {
			return
		}
	}
}

func (p *Proceso) Mostrar() {
	p.banderaMostrar = true
	//c<-"id" +strconv.Itoa(p.id)+": "+strconv.Itoa(p.i)
}

func (p *Proceso) Ocultar() {
	p.banderaMostrar = false
}

func (p *Proceso) Eliminar() {
	p.banderaEliminar = true
}

func (p *Proceso) GetID() uint64 {
	return p.Id
}
