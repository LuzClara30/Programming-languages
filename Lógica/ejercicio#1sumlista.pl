sumar([],0).
sumar([A|Tail],S):-
		sumar(Tail,S2),
		S is S2+A.

