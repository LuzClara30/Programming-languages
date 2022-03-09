#include <stdio.h>
#include <malloc.h>
const int SIZE_ARRAY =10;
//punteros
void print_array(int *arreglo, int size){
    int* end_ptr = arreglo+size;
    while(arreglo< end_ptr){
        printf("%i ",*arreglo);
        arreglo++;
    }
    printf("\n");
}
void copy_array(int *source_ptr, int *dest_ptr, int size){
    int* end_ptr = source_ptr+size;
    while(source_ptr<end_ptr){
        *dest_ptr = *source_ptr;
        source_ptr++;
        dest_ptr++;
    }
}
int insert_array(int* source_ptr, int ele, int size, int pos){
    source_ptr = (int*)realloc(source_ptr,(size+1)*sizeof(int));//reasize a la location
    //*(source_ptr+size) = ele;
    int* limit_ptr = source_ptr+pos;
    int* end_ptr = source_ptr+size;
    // Se implementó ciclo para insertar en una posición específica que contemple mover todos los elementos posteriores... SE INCLUYÓ EL PARÁMETRO pos
    while(end_ptr>limit_ptr){
        *end_ptr = *(end_ptr-1);
        end_ptr--;
    }
    *(source_ptr+pos)= ele;


    // implementar ciclo para insertar en una posición específica que contemple mover todos los elementos posteriores... DEBE INCLUIR EL PARÁMETRO pos
    return size+1;
}
int main() {
    int arreglo[SIZE_ARRAY];
    int *arreglo2 =  (int*) malloc(SIZE_ARRAY*sizeof(int));
    arreglo[0] =9; arreglo[1]=8;arreglo[2] =7; arreglo[3] =6;arreglo[4] =5;
    arreglo[5] =4;arreglo[6] =3;arreglo[7] =2;arreglo[8] =1;arreglo[9] =0;
    printf("Primer arreglo sin puntero\n");
    print_array(arreglo, SIZE_ARRAY);
    printf("Copia del arreglo original a uno con puntero\n");
    copy_array(arreglo, arreglo2, SIZE_ARRAY);
    print_array(arreglo2, SIZE_ARRAY);
    int new_size =0;
    //Trabajo extra clase: Se ingresa un elemento en una posicion y se corren los demás elementos.
    printf("Resultado de ingresar un nuevo elemento en una posicion especifica\n");
    new_size = insert_array(arreglo2,100,SIZE_ARRAY,4);
    print_array(arreglo2, new_size);
    return 0;
}