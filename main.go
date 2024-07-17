package main

import (
	"log"
	"strconv"

	"github.com/gerardo02/dsgo/ds"
)

func main() {
	otra := ds.List[string]{}
	lista_perrona := ds.List[string]{}

	err := lista_perrona.InstertAt(0, "roomba")
	if err != nil {
		log.Println(err)
	}

	_, err = lista_perrona.Remove("ropaa")

	if err != nil {
		log.Println(err)
	}

	lista_perrona.Prepend("bro")
	lista_perrona.Prepend("compilla")
	lista_perrona.Prepend("brosqui")
	lista_perrona.Prepend("raza")
	lista_perrona.Prepend("compa")

	err = lista_perrona.InstertAt(10, "compu")
	if err != nil {
		log.Println(err)
	}

	err = lista_perrona.InstertAt(lista_perrona.Length, "amd")
	if err != nil {
		log.Println(err)
	}

	err = lista_perrona.InstertAt(0, "ropa")
	if err != nil {
		log.Println(err)
	}

	_, err = lista_perrona.Remove("ropa")

	if err != nil {
		log.Println(err)
	}

	_, err = lista_perrona.RemoveAt(6)

	if err != nil {
		log.Println(err)
	}

	otra.Append("otra brother")
	otra.Append("ramses")
	otra.Append("calamity")
	otra.Prepend("colas")
	otra.Prepend("rifadas")
	otra.Prepend("no pues")
	// elementos := lista_perrona.GetAll()
	// for i := 0; i < len(elementos); i++ {
	// 	println("Elements en el loop: " + elementos[i])
	// }
	println("Lista perrona: ")
	println(lista_perrona.IsEmpty())
	elementos := lista_perrona.GetAll()
	for i := 0; i < len(elementos); i++ {
		println(strconv.Itoa(i) + ": " + elementos[i])
	}
	println("=======================")
	// println("Lista otra: ")
	// elementosOtra := otra.GetAll()
	// for i := 0; i < len(elementosOtra); i++ {
	// 	println(strconv.Itoa(i) + ": " + elementosOtra[i])
	// }

	lista_perrona.Clear()

	println(lista_perrona.IsEmpty())

}
