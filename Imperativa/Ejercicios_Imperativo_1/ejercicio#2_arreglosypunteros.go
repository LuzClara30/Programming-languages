package main

import "fmt"

func sustitucionPerdida(array *[10]int, pos int) {
	(*array)[pos] = 100
}
func sustsinPerdida(array *[10]int, element int, pos int, size int) {
	var slice = array[:]
	fmt.Println("\n capacidad:", cap(slice), " largo:", len(slice))
	j := (size - 1)
	for range slice {
		if j == (size - 1) {
			*&slice = append(slice, slice[j]) //9, 8, 7, 6, 100, 5, 4, 3, 2, 1,0
		}
		(slice[j+1]) = slice[j]
		if j == pos {
			(slice[j]) = element
			break
		}
		j--
	}
	fmt.Println(slice)
	fmt.Println("\n capacidad:", cap(slice), " largo:", len(slice))
}
func main() {
	const size = 10
	var arreglo [10]int = [size]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	arreglo2 := arreglo
	slice := arreglo2[:]
	sustitucionPerdida(&arreglo, 5)
	fmt.Println(arreglo)
	fmt.Println(slice)
	sustsinPerdida(&arreglo2, 100, 4, 10)

}
