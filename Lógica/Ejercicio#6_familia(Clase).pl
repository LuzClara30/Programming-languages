%Regla para consultar a tía
madre(petra, juan).
madre(petra, rosa).
madre(carmen, pedro).
madre(maria, ana).
madre(maria, enrique).
madre(rosa, raul).
madre(rosa, alfonso).
madre(rosa, cande).
%Juan y Rosa son hermanos y Juan es el Padre de ana, por tanto Rosa
% es la tía de ana.
padre(angel, juan).
padre(angel, rosa).
padre(pepe, pedro).
padre(juan, ana).
padre(juan, enrique).
padre(pedro, raul).
padre(pedro, alfonso).
padre(pedro, cande).
%%%tia(rosa,ana).
progenitor(X, Y) :- padre(X, Y).
progenitor(X, Y) :- madre(X, Y).

abuelo(X, Y) :- padre(X, Z), progenitor(Z, Y).
abuela(X, Y) :- madre(X, Z), progenitor(Z, Y).
hermano(X,Y):- padre(Z,X),padre(Z,Y).
hermano(X,Y):- madre(Z,X),madre(Z,Y).
sobrino(X,Y):- hermano(Z,Y), padre(Z,X).
sobrino(X,Y):- hermano(Z,Y), madre(Z,X).
tia(X,Y):-sobrino(Y,X).


