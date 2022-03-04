package main
//Escriba un programa que a partir de una cadena (sea predefinida o leida del usuario),
//busque mediante iteracviones en dicha cadena, una palabra cualquiera.
//Debe indicar si se encuentra o no dicha palabra. OBJETIVO: Uso de variables tipo cadena de datos y ciclos.
func main() {
	cadenaOriginal := "La clase de lenguajes es los viernes en la mañana"
	subCadena := "viernes"
	tamañosub := int(len(subCadena))
	indice := 0
	for i := 0; i < len(cadenaOriginal); i++ {
		if indice == 0 {
			if string(cadenaOriginal[i]) == string(subCadena[indice]) {
				indice++
			} else if i == len(cadenaOriginal)-1 {
				print("\nNo coinciden las palabras")
			}
		} else if indice > 0 {
			if string(cadenaOriginal[i]) != string(subCadena[indice]) {
				indice = 0
				if indice != tamañosub {
					print("\nNo coinciden las palabras")
				}
			} else {
				indice++
				if indice == tamañosub {
					print("\nLa subcadena se encuentra dentro de la cadena de texto")
					break
				}

			}
		}
	}
}
