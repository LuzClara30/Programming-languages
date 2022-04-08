entrada(paella).
entrada(gazpacho).
entrada(consomé).

carne(filete_de_cerdo).
carne(pollo_asado).
pescado(trucha).
pescado(bacalao).

postre(flan).
postre(nueces_con_miel).
postre(naranja).

calorias(paella, 200).
calorias(gazpacho,150).
calorias(consomé, 300).
calorias(filete_de_cerdo, 400).
calorias(pollo_asado, 280).
calorias(trucha, 160).
calorias(bacalao, 300).
calorias(flan, 200).
calorias(nueces_con_miel, 500).
calorias(naranja, 50).
% 1) Un plato principal es aquel que contiene o carne o pescado al
% menos.
plato_principal(X):- carne(X).
plato_principal(X):-pescado(X).
% 2) Una comida completa se compone de una entrada, un plato principal y
% un postre.
comidaCompleta(Entrada, Plato, Postre):- entrada(Entrada),plato_principal(Plato),
    postre(Postre).
% Valor calorico para la entrada X, plato Y y el postre Z, la suma de
% sus calorias será el valor calorico. T, será el total de calorias.
valor_calorico(X,Y,Z,T):-calorias(X,VCE),
    calorias(Y,VCP),
    calorias(Z,VCPo),
    T is VCE+VCP+VCPo,
    comidaCompleta(X,Y,Z).
%Máximo de calorias.
maximo(X,Y,Z,Maximo) :- valor_calorico(X,Y,Z,V),
    V<Maximo.
