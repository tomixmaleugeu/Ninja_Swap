// =====================================================================
//  internal/strategie/shikamaru.go — LA STRATÉGIE DE SHIKAMARU
//  Mission 3 (Jônin · Rang B)
// =====================================================================
//
//  Objectif : passer sous 1 100 jutsus pour 100 archives
//             et sous 8 500 pour 500.
//
//  À FAIRE :
//
//   ÉTAPE 1 — Normaliser
//   Remplacer chaque archive par son RANG : 0 pour la plus petite,
//   1 pour la suivante… jusqu'à n-1.
//   Exemple :  42 -7 100 3   →   2 0 3 1
//   (Le paquet sort est autorisé ici, et seulement ici.)
//
//   ÉTAPE 2 — Choisir AU MOINS une stratégie :
//
//   - Les escouades (chunks)
//     Sceller dans B d'abord tous les rangs 0 à 19, puis 20 à 39, etc.
//     Ensuite, invoquer vers A en prenant toujours la plus grande archive de B.
//     → Testez PLUSIEURS tailles d'escouade et gardez la meilleure.
//
//   - Le tri par bits (radix)
//     Pour chaque bit des rangs (du plus faible au plus fort) :
//     sceller dans B les archives dont le bit vaut 0, garder celles à 1,
//     puis tout invoquer. Ordre de grandeur : ≈ 1 080 pour 100.
//
//   ÉTAPE 3 — Mesurer et comparer avec la patrouille (mission 2).
//
//  QUESTIONS DE RÉFLEXION (réponses dans le README) :
//   - Pourquoi la normalisation simplifie-t-elle autant le problème ?
//   - Comment avez-vous choisi la taille de vos escouades ? Mesures à l'appui.
//
// =====================================================================

package strategie

import "ninja-swap/internal/rouleau"

func Normaliser(valeurs []int) []int {
	rangs := make([]int, len(valeurs))

	for i := 0; i < len(valeurs); i++ {
		for j := 0; j < len(valeurs); j++ {
			if valeurs[j] < valeurs[i] {
				rangs[i]++
			}
		}
	}

	return rangs
}

func NormaliserRouleaux(r *rouleau.Rouleaux) {
	r.A = Normaliser(r.A)
}

func Radix(r *rouleau.Rouleaux) []string {
	suite := []string{}
  NormaliserRouleaux(r)

  plusGrand := len(r.A) - 1
  nbBits := 0
  for (plusGrand >> nbBits) != 0 {
	  nbBits++
  }

  n := len(r.A)

for bit := 0; bit < nbBits; bit++ {

	for i := 0; i < n; i++ {
		if (r.A[0] >> bit) & 1 == 0 {
	    r.Pb()
      suite = append(suite, "pb")
    } else {
	    r.Ra()
      suite = append(suite, "ra")
}
	}

	for len(r.B) > 0 {
	    r.Pa()
      suite = append(suite, "pa")
  }
}

	return suite
}
