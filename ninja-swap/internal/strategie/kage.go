package strategie

import "ninja-swap/internal/rouleau"

type rotation struct {
	coups int
	sens  int
}

func optionsRotation(n, index int) (rotation, rotation) {
	if n < 2 || index == 0 {
		return rotation{}, rotation{}
	}
	return rotation{coups: index, sens: 1}, rotation{coups: n - index, sens: -1}
}

func indiceDebutDescente(b []int) int {
	n := len(b)
	if n == 0 {
		return 0
	}

	for i := 0; i < n; i++ {
		if b[i] < b[(i+1)%n] {
			return (i + 1) % n
		}
	}

	return 0
}

func emplacementInsertion(b []int, x int) int {
	if len(b) == 0 {
		return 0
	}

	for i, v := range b {
		precedent := b[(i-1+len(b))%len(b)]
		if v <= x && precedent >= x {
			return i
		}
	}

	return indiceDebutDescente(b)
}

func choisirCandidat(a, b []int) (rotation, rotation) {
	meilleurCout := int(^uint(0) >> 1)
	meilleurX := 0
	var meilleurA, meilleurB rotation

	for i, x := range a {
		optionsA1, optionsA2 := optionsRotation(len(a), i)
		j := emplacementInsertion(b, x)
		optionsB1, optionsB2 := optionsRotation(len(b), j)

		for _, ra := range [2]rotation{optionsA1, optionsA2} {
			for _, rb := range [2]rotation{optionsB1, optionsB2} {
				cout := 1
				if ra.coups > 0 && rb.coups > 0 && ra.sens == rb.sens {
					cout += max(ra.coups, rb.coups)
				} else {
					cout += ra.coups + rb.coups
				}

				if cout < meilleurCout || (cout == meilleurCout && x < meilleurX) {
					meilleurCout = cout
					meilleurX = x
					meilleurA, meilleurB = ra, rb
				}
			}
		}
	}

	return meilleurA, meilleurB
}

func executerRotations(r *rouleau.Rouleaux, ra, rb rotation, lancer func(string)) {
	if ra.coups > 0 && rb.coups > 0 && ra.sens == rb.sens {
		commun := min(ra.coups, rb.coups)
		for k := 0; k < commun; k++ {
			if ra.sens > 0 {
				lancer("rr")
			} else {
				lancer("rrr")
			}
		}

		if ra.coups > rb.coups {
			for k := commun; k < ra.coups; k++ {
				if ra.sens > 0 {
					lancer("ra")
				} else {
					lancer("rra")
				}
			}
		} else if rb.coups > ra.coups {
			for k := commun; k < rb.coups; k++ {
				if rb.sens > 0 {
					lancer("rb")
				} else {
					lancer("rrb")
				}
			}
		}
		return
	}

	for k := 0; k < ra.coups; k++ {
		if ra.sens > 0 {
			lancer("ra")
		} else {
			lancer("rra")
		}
	}
	for k := 0; k < rb.coups; k++ {
		if rb.sens > 0 {
			lancer("rb")
		} else {
			lancer("rrb")
		}
	}
}

func Kage(r *rouleau.Rouleaux) []string {
	jutsus := make([]string, 0, 3*len(r.A))

	lancer := func(nom string) {
		_ = r.Appliquer(nom)
		jutsus = append(jutsus, nom)
	}

	for len(r.A) > 0 {
		ra, rb := choisirCandidat(r.A, r.B)
		executerRotations(r, ra, rb, lancer)
		lancer("pb")
	}

	if len(r.B) > 0 {
		indexMax := indiceDebutDescente(r.B)

		haut, bas := optionsRotation(len(r.B), indexMax)
		if haut.coups <= bas.coups {
			for k := 0; k < haut.coups; k++ {
				lancer("rb")
			}
		} else {
			for k := 0; k < bas.coups; k++ {
				lancer("rrb")
			}
		}

		for len(r.B) > 0 {
			lancer("pa")
		}
	}

	return jutsus
}
   
