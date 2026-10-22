// =====================================================================
//  internal/rouleau/rouleau.go — LES DEUX ROULEAUX ET LES 11 JUTSUS
//  Mission 0 (Académie)
// =====================================================================
//
//  Rôle : représenter les deux piles et coder chaque jutsu une fois pour
//  toutes. ninja_swap, le juge et toutes vos stratégies utiliseront CE
//  package : ne réécrivez jamais un jutsu ailleurs.
//
//  À FAIRE :
//
//   1. Choisir une représentation des rouleaux.
//      Convention conseillée : une slice d'entiers, l'indice 0 = le HAUT.
//
//   2. Prévoir une structure qui regroupe le Rouleau du Ciel (A) et le
//      Rouleau de la Terre (B), et qui COMPTE les jutsus lancés.
//      (Utile pour vos stratégies et vos mesures.)
//
//   3. Coder les 11 jutsus, une fonction (ou méthode) chacun :
//
//      sa  / sb  / ss    Kawarimi      échange les 2 archives du haut
//      pa                Kuchiyose     haut de B → haut de A
//      pb                Fûin          haut de A → haut de B
//      ra  / rb  / rr    Rasengan      le haut passe tout en bas
//      rra / rrb / rrr   Rasengan inv. le bas remonte tout en haut
//
//      Les jutsus « en duo » (ss, rr, rrr) comptent pour UN seul jutsu.
//
//   4. Chaque jutsu doit être SANS DANGER :
//      sur un rouleau vide ou à 1 archive, il ne fait rien… et ne plante pas.
//
//   5. Prévoir un moyen d'appliquer un jutsu à partir de son nom ("sa", "pb"…).
//      Le juge en aura besoin. Nom inconnu → renvoyer une erreur.
//
//   6. Prévoir une fonction qui dit si A est trié et B vide.
//
//  Testez tout dans rouleau_test.go AVANT de passer aux stratégies.
//
// =====================================================================


package rouleau

import "fmt"

type Rouleaux struct {
	A     []int
	B     []int
	Coups int
}

func (r *Rouleaux) Sa() {
	if len(r.A) >= 2 {
		r.A[0], r.A[1] = r.A[1], r.A[0]
	}
	r.Coups++
}

func (r *Rouleaux) Sb() {
	if len(r.B) >= 2 {
		r.B[0], r.B[1] = r.B[1], r.B[0]
	}
	r.Coups++
}
func (r *Rouleaux) Ss() {
	if len(r.A) >= 2 {
		r.A[0], r.A[1] = r.A[1], r.A[0]
	}
	if len(r.B) >= 2 {
		r.B[0], r.B[1] = r.B[1], r.B[0]
	}
	r.Coups++
}

func (r *Rouleaux) Pa() {
	if len(r.B) >= 1 {
		valeur := r.B[0]
		r.B = r.B[1:]
		r.A = append([]int{valeur}, r.A...)
	}
	r.Coups++
}

func (r *Rouleaux) Pb() {
	if len(r.A) >= 1 {
		valeur := r.A[0]
		r.A = r.A[1:]
		r.B = append([]int{valeur}, r.B...)
	}
	r.Coups++
}

func (r *Rouleaux) Ra() {
	if len(r.A) >= 2 {
		haut := r.A[0]
		r.A = append(r.A[1:], haut)
	}
	r.Coups++
}

func (r *Rouleaux) Rb() {
	if len(r.B) >= 2 {
		haut := r.B[0]
		r.B = append(r.B[1:], haut)
	}
	r.Coups++
}

func (r *Rouleaux) Rr() {
	if len(r.A) >= 2 {
		haut := r.A[0]
		r.A = append(r.A[1:], haut)
	}
	if len(r.B) >= 2 {
		haut := r.B[0]
		r.B = append(r.B[1:], haut)
	}
	r.Coups++
}

func (r *Rouleaux) Rra() {
	if len(r.A) >= 2 {
		dernier := r.A[len(r.A)-1]
		r.A = append([]int{dernier}, r.A[:len(r.A)-1]...)
	}
	r.Coups++
}

func (r *Rouleaux) Rrb() {
	if len(r.B) >= 2 {
		dernier := r.B[len(r.B)-1]
		r.B = append([]int{dernier}, r.B[:len(r.B)-1]...)
	}
	r.Coups++
}

func (r *Rouleaux) Rrr() {
	if len(r.A) >= 2 {
		dernier := r.A[len(r.A)-1]
		r.A = append([]int{dernier}, r.A[:len(r.A)-1]...)
	}
	if len(r.B) >= 2 {
		dernier := r.B[len(r.B)-1]
		r.B = append([]int{dernier}, r.B[:len(r.B)-1]...)
	}
	r.Coups++
}


func (r *Rouleaux) Appliquer(nom string) error {
	switch nom {
	case "sa":
		r.Sa()
	case "sb":
		r.Sb()
	case "ss":
		r.Ss()
	case "pa":
		r.Pa()
	case "pb":
		r.Pb()
	case "ra":
		r.Ra()
	case "rb":
		r.Rb()
	case "rr":
		r.Rr()
	case "rra":
		r.Rra()
	case "rrb":
		r.Rrb()
	case "rrr":
		r.Rrr()
	default:
		return fmt.Errorf("jutsu inconnu : %s", nom)
	}

	return nil
}

func (r *Rouleaux) EstTrie() bool {
	if len(r.B) != 0 {
		return false
	}

	for i := 0; i < len(r.A)-1; i++ {
		if r.A[i] > r.A[i+1] {
			return false
		}
	}

	return true
}
