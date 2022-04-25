(define (vector-inverso N lista)
  (cond ((null? lista)
         '())
        (else
         (cons (- N (car lista))(vector-inverso N(cdr lista)))
         )
        ))