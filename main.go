package main

import (
	"log"
	"strconv"

	"github.com/gerardo02/dsgo/ds"
)

func main() {
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

	println("Lista perrona: ")
	println(lista_perrona.IsEmpty())
	elementos := lista_perrona.GetAll()
	for i := 0; i < len(elementos); i++ {
		println(strconv.Itoa(i) + ": " + elementos[i])
	}
	println("=======================")

	_, err = lista_perrona.Remove("ropa")

	if err != nil {
		log.Println(err)
	}
	println("Lista pero con Get: ")
	for i := 0; i < lista_perrona.Length; i++ {
		val, _ := lista_perrona.Get(i)
		println(strconv.Itoa(i) + ": " + val)
	}
	println("=======================")

	lista_perrona.Clear()

	println(lista_perrona.IsEmpty())

}
