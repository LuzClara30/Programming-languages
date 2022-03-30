subconj([],[]).
subconj([A|T1],[A|T2]):-
    subconj(T2,T1).
subconj(T2,[_|T1]):-
    subconj(T2,T1).
