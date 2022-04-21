(ns untitled.core)
(defn factorial [x]
  (cond(= x 1)
    1
    :else (* x (factorial (- x 1))))
  )
