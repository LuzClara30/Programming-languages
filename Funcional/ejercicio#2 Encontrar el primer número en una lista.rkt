(define (primer-numero lista)
  (cond((null? lista)
        null)
       ((number?(car lista))
       (car lista))
       (else
   (primer-numero(cdr lista))
   )
       ))
  