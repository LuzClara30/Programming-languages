package main

import "fmt"

type People struct {
	name   string
	id     [8]int
	number int
}

func muestreo(personas []People, muestra [8]int) []People {
	n := 0
	indice := 0
	cantidad := 0
	for i := 0; i < len(personas); i++ {
		n = 0
		indice = 0
		for y := 0; y < len(muestra); y++ {
			if muestra[y] == personas[i].id[indice] {
				n++
			}
			indice++
			if y == (len(muestra) - 1) {
				personas[i].number = n
			}
		}
	}
	for i := 0; i < len(personas); i++ {
		indice = 0
		cantidad = 0
		for indice < len(personas) {
			if personas[i].number < personas[indice].number {
				cantidad++
			}
			indice++
			if cantidad >= 3 {
				personas = append(personas[:i], personas[i+1:]...)
				i--

			}
		}
	}
	return personas[0:3]
}

func main() {
	arreglo := [8]int{1, 1, 0, 1, 1, 0, 1, 0}
	arreglo1 := [8]int{0, 0, 0, 1, 0, 1, 0, 1}
	arreglo2 := [8]int{0, 0, 1, 0, 0, 1, 1, 0}
	arreglo3 := [8]int{1, 0, 1, 1, 0, 1, 0, 1}
	arreglo4 := [8]int{1, 0, 1, 0, 1, 1, 1, 1}

	muestra := [8]int{0, 1, 0, 0, 1, 0, 0, 1}
	george := People{"George", arreglo, 0}
	Lucia := People{"Lucia", arreglo1, 0}
	jose := People{"Jose", arreglo2, 0}
	marta := People{"Marta", arreglo3, 0}
	antonio := People{"Antonio", arreglo4, 0}
	personas := []People{george, Lucia, jose, marta, antonio}
	print("Las tres personas con mayor coincidencia son: \n")
	muestreoFinal := muestreo(personas, muestra)
	i := 0
	for i < len(muestreoFinal) {
		fmt.Println(muestreoFinal[i].name)
		i++
	}

}
