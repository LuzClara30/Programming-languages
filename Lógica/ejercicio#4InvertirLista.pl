invertir([],[]).
invertir([A|Tail],R2):-
    invertir(Tail,R),
    a�adir(A,R,R2).
a�adir(E,[],[E]).
a�adir(E,[A|Tail],[A|R]):-
a�adir(E,Tail,R).
