package main

//Escriba un programa que utilice métodos que reciban como parámetros dos cadenas de caracteres,
//y que muestre por pantalla un mensaje que indique si la primera de ellas es una subcadena de la segunda.
//Para dicho ejercicio, intente utilizar alguna función predefinida para dicho fin e intente implementar otra versión que NO utilice ninguna función predefinida.
//OBJETIVO: uso de librerías, métodos, paso de parámetros, retornos, ciclos.
import (
	"strings"
)
func main() {
	cadena := "hola me llamo Juan"
	subcadena := "llamo"
	print("\nComparación sin libreria")
	if comparar2(cadena, subcadena) == true {
		print("\nLa subcadena se encuentra dentro de la cadena")
	} else {
		print("\nLa subcadena no se encuentra dentro de la cadena")
	}
	cadena = "hola soy pablo"
	subcadena = "padre"
	print("\nComparación con libreria")
	if comparar(cadena, subcadena) == true {
		print("\nLa subcadena se encuentra dentro de la cadena")
	} else {
		print("\nLa subcadena no se encuentra dentro de la cadena")
	}

}

func comparar(cadena string, subcadena string) bool {
	if strings.Contains(strings.ToLower(cadena), strings.ToLower(subcadena)) {
		return true
	}
	return false
}
func comparar2(cadena, subcadena string) bool {
	var Indice int = 0
	var status bool = false
	for i := 0; i < len(cadena); i++ {
		if Indice == len(subcadena) {
			status = true
			break
		}
		if cadena[i] != subcadena[Indice] {
			if Indice >= 1 {
				Indice = 0
			}
			continue
		} else {
			Indice++
		}
	}
	return status
}
