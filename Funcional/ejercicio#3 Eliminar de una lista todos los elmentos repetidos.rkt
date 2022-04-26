(define (concatenar List1 List2) ; concatena la lista 1 con la lista 2
(if (null? List1)
List2
(cons (car List1) (concatenar (cdr List1) List2))))

(define (member? X List) ;permite saber si X es miembro de la lista
(if (null? List)
#f
(if (equal? X (car List))
#t
(member? X (cdr List)))))

(define (filtrar List); Construya la funcion filtrar. Esta funcion recibe una lista y retorna otra lista con los mismos elementos de la primera, pero eliminando los elementos repetidos.
(cond
((null? List) List)
((member? (car List) (cdr List)) (filtrar (cdr List)))
(else (concatenar (list (car List)) (filtrar (cdr List))))))