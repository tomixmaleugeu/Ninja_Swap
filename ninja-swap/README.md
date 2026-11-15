# Rouleau de l'équipe n° __

> **Consigne :** ce README est le « Rouleau de l'équipe ». Il est noté (documentation et analyse de complexité).
> Remplacez chaque bloc de consignes en italique par votre propre contenu, et supprimez les consignes à la fin.
>
> **Rendu :** vendredi 23 octobre à 11 h précises, avec `git tag rendu` puis `git push --tags`.
> Poussez votre travail à la fin de CHAQUE séance, par les deux membres de l'équipe.
> Donnez accès à ce dépôt au Hokage et à Rémy dès la première séance.

**Genin :** _Ethan Minaud_ · _Tom Degorce_

---

## 1. Compiler, lancer, tester

_Donnez les commandes exactes pour :_

- _compiler `ninja_swap` et `juge` ;_
go build -o ninja_swap ./cmd/ninja_swap
go build -o juge ./cmd/juge

- _lancer un exemple et le vérifier avec le juge ;_
./ninja_swap $ARG                 
./ninja_swap $ARG | ./juge $ARG
./ninja_swap $ARG | wc -l  

## 2. Notre stratégie finale

- 3 archives : les 6 ordres possibles sont traités au cas par cas, 2 jutsus au maximum.
- 4 archives : on cherche la plus petite, on la fait remonter par le chemin le plus court (ra ou rra), on la scelle (pb), on trie les 3 restantes, puis on invoque (pa).
- 5 archives : même idée, mais on scelle les 2 plus petites avant de trier les 3 restantes, puis pa pa.
- Brute-force : parcours en largeur de toutes les suites de jutsus, avec une table des états déjà vus pour ne pas refaire deux fois le même rouleau. La première suite qui trie est forcément la plus courte.

Kage (6 archives et plus)
L'idée : on vide A dans B en gardant B trié en ordre décroissant (le plus grand en haut). À la fin, on invoque tout B vers A : les archives arrivent dans le bon ordre.
À chaque tour, on ne prend pas l'archive du haut. On teste toutes les archives de A et on calcule, pour chacune :
combien de rotations il faut dans A pour l'amener en haut (ra ou rra, selon le chemin le plus court) ;
combien de rotations il faut dans B pour que sa place soit en haut de B (rb ou rrb) ;
le coût total = ces rotations + 1 pour le pb.
On joue l'archive la moins chère. Si A et B doivent tourner dans le même sens, on les fusionne : rr ou rrr coûte 1 jutsu au lieu de 2.
À la fin, on tourne B pour mettre le plus grand en haut, puis on fait pa jusqu'à vider B.

## 3. Journal d'optimisation

_Une ligne par version. Ne supprimez jamais une ligne : le chemin compte autant que l'arrivée._

| Version | Stratégie | Big-O (jutsus) | Moy. 100 | Moy. 500 | Ce que nous avons appris |
|---|---|---|---|---|---|
| v1 | Rock Lee (brute-force) |O(2ⁿ) | — | — |Le brute-force trouve toujours la solution la plus courte, mais le temps de calcul explose dès que le rouleau grossit. Il est utile pour vérifier nos solutions faites à la main sur les petites piles, mais inutilisable au-delà de 8 archives. |
| v2 | Patrouille |O(n²) |	≈ 1 456 |≈ 32 260 |Trouver le minimum puis le sceller est simple et toujours correct, mais coûteux : 5 fois plus d'archives donnent plus de 20 fois plus de jutsus. |
| v3 |Radix |O(n log n) |1 084 |6 784 |La normalisation en rangs rend le tri par bits très régulier. Le nombre de jutsus ne dépend presque pas de l'ordre de départ |
| v4 |Kage |O(n log n) |576 |5 138 |Choisir à chaque tour l'archive la moins chère divise presque par 2 le coût du radix à 100 archives. |

## 4. Questions de réflexion

### Mission 1 — Rock Lee
- _Que devient le temps de calcul du brute-force entre 5 et 7 archives ? Pourquoi ?_
Il explose. À chaque étape on essaie 11 jutsus, donc pour une solution de longueur k il faut examiner jusqu'à 11^k séquences. Plus il y a d'archives, plus la solution est longue, et chaque jutsu de plus multiplie le travail par environ 11

- _Quelle est la complexité (Big-O) de votre brute-force ?_
O(11^k), où k est la longueur de la solution la plus courte. La table des états déjà vus évite de refaire des rouleaux identiques, mais le nombre d'états possibles reste énorme.

### Mission 2 — La patrouille
- _Pourquoi cette méthode est-elle en O(n²) ?_
Pour chacune des n archives, on cherche le minimum  et on fait jusqu'à n/2 rotations. On fait donc de l'ordre de n × n opérations.
- _Le programme est-il lent à calculer, ou seulement « cher » en jutsus ?_ Il est rapide à calculer mais cher en jutsus. Ce sont deux mesures différentes : le temps de calcul est le temps que met notre programme à trouver la suite, le nombre de jutsus est la longueur de cette suite.

### Mission 3 — Shikamaru
- _Pourquoi la normalisation simplifie-t-elle autant le problème ?_
Après normalisation, les nombres vont de 0 à n-1 sans trou, quelles que soient les valeurs de départ . On sait donc exactement où chaque archive doit aller, et on peut regarder ses bits un par un.
- _Comment avez-vous choisi la taille de vos escouades ? Montrez vos mesures._

### Mission 4 — Le contrôle du chakra
- _Que coûte en calcul la recherche de l'archive la moins chère à chaque tour ? Quelle complexité totale ?_
À chaque tour, pour chaque archive de A (jusqu'à n), on cherche sa place dans B  : ≈ n² opérations par tour. Il y a n tours, donc le temps de calcul est en O(n³). Pour 500 archives cela reste de l'ordre de la seconde ou moins. Pour le nombre de jutsus, la mesure donne ≈ 576 pour 100 et ≈ 5 138 pour 500 : 5 fois plus d'archives donnent environ 9 fois plus de jutsus, donc une croissance entre O(n log n) et O(n²).
- _Quelle amélioration vous a fait gagner le plus ? Laquelle vous a déçus ?_
Le plus gros gain vient du passage du radix (1 084) à « l'archive la moins chère à chaque tour » . La fusion des rotations (rr, rrr) fait encore gagner de ≈ 652 à ≈ 576 à 100 archives et de ≈ 5 952 à ≈ 5 138 à 500. À 500 archives c'est même elle qui fait passer sous la barre des 5 500.



test
