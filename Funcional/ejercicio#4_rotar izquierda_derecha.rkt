 #lang racket

(define (derecha n lst)
  (let ((x (min n (length lst))))
    (append (take-right lst x)
            (drop-right lst x))))


(define (izquierda n lst)
  (let ((x (min n (length lst))))
    (append (drop lst x)
            (take lst x))))
(define (rotar dir n lst) 
  (if (equal? dir "i") ; "i" es igual a izquierda
      (izquierda n lst)
      (derecha n lst)))