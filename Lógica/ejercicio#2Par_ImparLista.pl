largopar([]).
largopar([_|L]):-
    largoimpar(L).
largoimpar([_]).
largoimpar([_|L]):-
    largopar(L).
