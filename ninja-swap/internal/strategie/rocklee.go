// =====================================================================
//  internal/strategie/rocklee.go — L'ENTRAÎNEMENT DE ROCK LEE
//  Mission 1 (Genin · Rang D)
// =====================================================================
//
//  Rock Lee ne maîtrise aucun ninjutsu : il compense par l'effort pur.
//  Ce fichier contient deux choses.
//
//  PARTIE A — Les petites piles, résolues à la main
//
//   1. 3 archives : il n'existe que 6 ordres possibles.
//      Trouvez SUR PAPIER la meilleure suite de jutsus pour chacun,
//      puis codez-les. Objectif : 2 jutsus maximum.
//
//   2. 5 archives : objectif 12 jutsus maximum.
//      Idée de Gai-sensei : sceller les 2 plus petites dans B (pb),
//      trier les 3 restantes avec la partie 1, puis les invoquer (pa).
//      4 archives : même idée, en ne scellant que la plus petite.
//
//      Le rang Kage exige 8 jutsus au maximum pour 5 archives : c'est le
//      vrai minimum. Pour l'atteindre, utilisez le brute-force (partie B).
//
//  PARTIE B — Le brute-force
//
//   3. Essayer toutes les suites de jutsus de longueur 1, puis 2, puis 3…
//      jusqu'à en trouver une qui trie. C'est forcément la plus courte.
//      Indice : une file (parcours en largeur) évite de tout réessayer.
//
//   4. MESURER le temps de calcul de 3 à 8 archives
//      et noter les résultats dans le journal d'optimisation du README.
//
//  QUESTIONS DE RÉFLEXION (réponses dans le README) :
//   - Que devient le temps de calcul entre 5 et 8 archives ? Pourquoi ?
//   - Quelle est la complexité (Big-O) de votre brute-force ?
//
// =====================================================================

package strategie

import (
	"fmt"
	"ninja-swap/internal/rouleau"
)

func Trier3(r *rouleau.Rouleaux) []string {
	suite := []string{}

	if len(r.A) != 3 {
		return suite
	}

	a := r.A[0]
	b := r.A[1]
	c := r.A[2]

	if a < b && b < c {
	} else if a > b && c > a && c > b {
		r.Sa()
		suite = append(suite, "sa")
	} else if a > b && b > c {
		r.Sa()
		suite = append(suite, "sa")
		r.Rra()
		suite = append(suite, "rra")
	} else if a > b && b < c && c < a {
		r.Ra()
		suite = append(suite, "ra")
	} else if a < b && b > c && c > a {
		r.Sa()
		suite = append(suite, "sa")
		r.Ra()
		suite = append(suite, "ra")
	} else if a < b && b > c && c < a {
		r.Rra()
		suite = append(suite, "rra")
	}

	return suite
}

func PossMin(r *rouleau.Rouleaux) int {
	longeur := len(r.A)
	indexMin := 0

	for i := 1; i < longeur; i++ {
		if r.A[i] < r.A[indexMin] {
			indexMin = i
		}
	}

	return indexMin
}

func Sceller(r *rouleau.Rouleaux) []string {
	suite := []string{}
	indexMin := PossMin(r)

	if indexMin <= len(r.A)/2 {
		for i := 0; i < indexMin; i++ {
			r.Ra()
			suite = append(suite, "ra")
		}
	} else {
		for i := 0; i < len(r.A)-indexMin; i++ {
			r.Rra()
			suite = append(suite, "rra")
		}
	}

	r.Pb()
	suite = append(suite, "pb")
	return suite
}

func Trier4(r *rouleau.Rouleaux) []string {
	suite := []string{}

	if len(r.A) != 4 {
		return suite
	}

	suite = append(suite, Sceller(r)...)
	suite = append(suite, Trier3(r)...)
	r.Pa()
	suite = append(suite, "pa")

	return suite
}

func Trier5(r *rouleau.Rouleaux) []string {
	suite := []string{}

	if len(r.A) != 5 {
		return suite
	}

	suite = append(suite, Sceller(r)...)
	suite = append(suite, Sceller(r)...)
	suite = append(suite, Trier3(r)...)
	r.Pa()
	r.Pa()
	suite = append(suite, "pa", "pa")

	return suite
}
func RockLee(r *rouleau.Rouleaux) []string {
	suite := []string{}

	if len(r.A) == 2 && r.A[0] > r.A[1] {
		r.Sa()
		suite = append(suite, "sa")
	} else if len(r.A) == 3 {
		suite = Trier3(r)
	} else if len(r.A) == 4 {
		suite = Trier4(r)
	} else if len(r.A) == 5 {
		suite = Trier5(r)
	}

	return suite
}

func BruteForce(r *rouleau.Rouleaux) []string {
	jutsus := []string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"}

	if r.EstTrie() {
		return []string{}
	}

	file := []rouleau.Rouleaux{*r}
	suites := [][]string{{}}

	vus := map[string]bool{}
	vus[fmt.Sprint(r.A, r.B)] = true

	for len(file) > 0 {
		actuel := file[0]
		suiteActuelle := suites[0]
		file = file[1:]
		suites = suites[1:]

		for _, j := range jutsus {
			copieA := append([]int{}, actuel.A...)
			copieB := append([]int{}, actuel.B...)
			nouveau := rouleau.Rouleaux{A: copieA, B: copieB}
			nouveau.Appliquer(j)

			nouvelleSuite := append(append([]string{}, suiteActuelle...), j)

			if nouveau.EstTrie() {
				return nouvelleSuite
			}

			cle := fmt.Sprint(nouveau.A, nouveau.B)
			if !vus[cle] {
				vus[cle] = true
				file = append(file, nouveau)
				suites = append(suites, nouvelleSuite)
			}
		}
	}

	return []string{}
}
