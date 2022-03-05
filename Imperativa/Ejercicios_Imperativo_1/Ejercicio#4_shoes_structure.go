package main

import "fmt"

//Defina una estructura para un inventario de una tienda que vende zapatos. Cada zapato debe contar con la información de su modelo (marca), su precio y la talla del mismo que debe ir únicamente del 34 al 44.

//Escriba un programa que lea la información anterior para 10 zapatos del inventario y los almacene en un arreglo.
//Escriba una función que reciba de parámetro dicho arreglo y dos tallas minimo y máximo,
//y que retorne un nuevo arreglo con los zapatos que concuerden con un el rango de tallas entre dichos mínimo y máximo.
//Documente la estrategia utilizada para crear un nuevo conjunto de elementos en tiempo de ejecución para ser retornado por la función.
//OBJETIVO: manejo de estructuras, ciclos, arreglos y slices (consideren el uso de slices para este ejercicio)
type Zapatos struct {
	marca  string
	precio int
	talla  int
}

var Shoes []Zapatos

func informacion(marca string, precio int, talla int) []Zapatos {
	if talla < 34 || talla > 44 {
		print("No puede agregarse el zapato debido a que la talla es menor a 34 o mayor a 44\n")
	} else {
		zapato := Zapatos{marca, precio, talla}
		Shoes = append(Shoes, zapato)
	}
	return Shoes
}

//La estrategia utilizada para crear un nuevo conjunto de elementos en tiempo de ejecución para ser retornada por la función
//fue un slices que puede verse como arreglos de longitud dinámica, los cuales no tienen una capacidad y longitud limitada
//pues tiene varias funciones como .append que permite agregar datos y aumnetar la capacidad de este si es necesario.
func MinMax(shoes []Zapatos, min int, max int) []Zapatos {
	var slice []Zapatos
	for i := 0; i < len(shoes); i++ {
		if shoes[i].talla >= min && shoes[i].talla <= max {
			slice = append(slice, shoes[i])
		}
	}
	return slice
}
func main() {
	informacion("Kolosh", 35000, 35)
	informacion("SODA", 12200, 43)
	informacion("Kolosh", 45000, 37)
	informacion("SODA", 10500, 38)
	informacion("SODA", 12450, 40)
	informacion("VIZZANO", 18700, 41)
	informacion("VIZZANO", 15655, 37)
	informacion("SODA", 9578, 39)
	informacion("Kolosh", 32450, 38)
	informacion("Ramarim", 22180, 36)
	fmt.Println(MinMax(Shoes, 36, 40))
}
