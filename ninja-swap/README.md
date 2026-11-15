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

## 3. Journal d'optimisation

_Une ligne par version. Ne supprimez jamais une ligne : le chemin compte autant que l'arrivée._

| Version | Stratégie | Big-O (jutsus) | Moy. 100 | Moy. 500 | Ce que nous avons appris |
|---|---|---|---|---|---|
| v1 | Rock Lee (brute-force) |O(2ⁿ) | — | — |Le brute-force trouve toujours la solution la plus courte, mais le temps de calcul explose dès que le rouleau grossit. Il est utile pour vérifier nos solutions faites à la main sur les petites piles, mais inutilisable au-delà de 8 archives. |
| v2 | Patrouille | | | | |
| v3 | | | | | |

## 4. Questions de réflexion

### Mission 1 — Rock Lee
- _Que devient le temps de calcul du brute-force entre 5 et 7 archives ? Pourquoi ?_
Il explose. À chaque étape on essaie 11 jutsus, donc pour une solution de longueur k il faut examiner jusqu'à 11^k séquences. Plus il y a d'archives, plus la solution est longue, et chaque jutsu de plus multiplie le travail par environ 11

- _Quelle est la complexité (Big-O) de votre brute-force ?_
O(11^k), où k est la longueur de la solution la plus courte. La table des états déjà vus évite de refaire des rouleaux identiques, mais le nombre d'états possibles reste énorme.

### Mission 2 — La patrouille
- _Pourquoi cette méthode est-elle en O(n²) ?_
- _Le programme est-il lent à calculer, ou seulement « cher » en jutsus ?_

### Mission 3 — Shikamaru
- _Pourquoi la normalisation simplifie-t-elle autant le problème ?_
- _Comment avez-vous choisi la taille de vos escouades ? Montrez vos mesures._

### Mission 4 — Le contrôle du chakra
- _Que coûte en calcul la recherche de l'archive la moins chère à chaque tour ? Quelle complexité totale ?_
- _Quelle amélioration vous a fait gagner le plus ? Laquelle vous a déçus ?_

## 5. Le travail d'équipe

_Qui a fait quoi ? Comment avez-vous alterné pilote et copilote à chaque séance ?_

## 6. Missions secrètes (facultatif)

_Pour chaque mission secrète tentée : ce qui fonctionne, vos mesures, ce que vous avez appris._



test
