(define (remove-if fun? lista);elimina las listas vacias de una lista
  (apply append
         (map (lambda(x) (cond ((fun? x) '())
                               (else (list x))))
              lista)))
(define (Indic lista);genera los indices de una lista
  (cond((= (length lista) 0)'())
       (else
        (build-list (length lista) values))))
;recibe la n-esima posición, y la lista donde se busca el elemento.
;(ejercicio5 4 '(10 14 3 2 11 7 6 5))
(define (ejercicio5 n lista)
  (remove-if null?(map (lambda (x y) (cond ((equal? n y)
                            x)
                           (else'())
                           )
         )lista (Indic lista)))
  )