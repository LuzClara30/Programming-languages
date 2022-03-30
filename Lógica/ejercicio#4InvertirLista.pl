invertir([],[]).
invertir([A|Tail],R2):-
    invertir(Tail,R),
    añadir(A,R,R2).
añadir(E,[],[E]).
añadir(E,[A|Tail],[A|R]):-
añadir(E,Tail,R).
