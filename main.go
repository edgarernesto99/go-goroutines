package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"./administrador"
)

func Proceso(id uint64) {
	i := uint64(0);
	for {
		fmt.Printf("id %d: %d", id, i)
		i = i + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	procesos := administrador.AdmProcesos{}
	procesos.Mostrar()
	scaner := bufio.NewScanner(os.Stdin)
	var op string
	var ID uint64 = 0

	for true {
		fmt.Println(" ============================ ")
		fmt.Println("|            MENU            |")
		fmt.Println("|============================|")
		fmt.Println("| 1. Agregar proceso         |")
		fmt.Println("| 2. Mostrar procesos        |")
		fmt.Println("| 3. Eliminar proceso        |")
		fmt.Println("| 0. Salir                   |")
		fmt.Println(" ============================ ")
		fmt.Print("Opción: ")
		scaner.Scan()
		op = scaner.Text()

		if op == "1" {
			p := administrador.Proceso{Id: ID}
			go p.Iniciar()
			procesos.AdmProcesos = append(procesos.AdmProcesos, &p)
			ID = ID +1
		} else if op == "2" {
			procesos.Mostrar()
			scaner.Scan()
			procesos.Ocultar()
		} else if op == "3" {
			var id uint64
			fmt.Print("ID: ")
			fmt.Scanln(&id)
			procesos.Eliminar(id)
		} else if op == "0" {
			break
		} else {
			fmt.Println("Opcion no valida")
		}
	}
}