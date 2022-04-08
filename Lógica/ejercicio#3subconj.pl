%%Debe ingresar subconj([Lista],[Subconjunto]).

subconj([],[]).
subconj([A|T1],[A|T2]):-
    subconj(T1,T2).
subconj([_|T1],T2):-
    subconj(T1,T2).
