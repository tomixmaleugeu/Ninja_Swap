
package strategie

import "ninja-swap/internal/rouleau"


func Patrouille(r *rouleau.Rouleaux) []string {
	jutsus := make([]string, 0, 2*len(r.A))

	lancer := func(nom string) {
		_ = r.Appliquer(nom) 
		jutsus = append(jutsus, nom)
	}

	for len(r.A) > 0 {
		minimum := 0
		for i := 1; i < len(r.A); i++ {
			if r.A[i] < r.A[minimum] {
				minimum = i
			}
		}

		n := len(r.A)
		if minimum <= n/2 {
			for i := 0; i < minimum; i++ {
				lancer("ra")
			}
		} else {
			for i := 0; i < n-minimum; i++ {
				lancer("rra")
			}
		}

		lancer("pb")
	}

	for len(r.B) > 0 {
		lancer("pa")
	}

	return jutsus
}
