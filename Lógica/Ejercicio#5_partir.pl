%%partir([1,2,3,4,5],3,X,Y).
%Debe estar incluido el Umbral en la lista para que funcione.
%Este ejercicio funciona por medio de recursividad, 
% Cuando la cabeza de la lista es igual al umbral, la cola de la lista
% pasa a hacer la lista de los mayores y se devuelve en la de los menr
partir([H|T],H,[],T).
partir([H|T],U,[H|T2],S) :-
    partir(T,U,T2,S).
