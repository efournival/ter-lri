# ter-lri

Travail d'Étude et de Recherche sur l'exploration de semigroupes numériques avec mise en place d'un algorithme de vol de tâches.

Il se déroule au LRI (Laboratoire de Recherche en Informatique), rattaché à l'Université Paris-Sud et au CNRS, dans l'équipe GALaC (Graphes, Algorithmes et Combinatoire).

Tuteur : [Florent Hivert](https://www.lri.fr/~hivert/)

## Déroulement

- [x] Étude du problème d'exploration de semigroupes auquel on va s'intéresser
- [x] Implémentation basique en Go
- [x] Parallélisation et optimisation basique
- [x] Étude comparative avec la version optimisée en C++
- [x] Profiler le code Go avec `perf` et `pprof`
- [x] Utilisation de `cgo` avec appel au code de `NumericMonoid` généré par GCC
- [x] Traduction de la boucle principale en Go natif utilisant les Goroutines à la place de Cilk du code original
- [x] Implémentation d'un système de type fork-join distribué
- [x] Mise en place du vol de tâche au sein du modèle fork-join
- [ ] Finalisation et lancement des calculs

## Problèmes complémentaires

* Distribution du problème des N reines
* [Un problème similaire](https://github.com/Didayolo/spark) avec [Spark](http://spark.apache.org/)
* Utiliser des instructions AVX1 voire AVX2
* Test du compilateur Intel et de LLVM/Clang
* [Bug Cilk dans GCC](https://gcc.gnu.org/bugzilla/show_bug.cgi?id=80038)

## Ressources

* [NumericMonoid](https://github.com/hivert/NumericMonoid)
* [Exploring the tree of numerical semigroups](https://hal.archives-ouvertes.fr/hal-00823339/document)
* [HPC in Combinatorics: Application of Work-Stealing](https://github.com/OpenDreamKit/OpenDreamKit/raw/master/WP5/T5.6/HPC-Combi.pdf)
* [Le code source de l'implémentation du vol de tâches dans Sage](https://github.com/sagemath/sage/blob/master/src/sage/parallel/map_reduce.py)
* [Profiling en Go](https://blog.golang.org/profiling-go-programs)
* [Un livre sur la programmation réseau en Go](https://www.gitbook.com/book/jannewmarch/network-programming-with-go-golang-/details)
* [Un livre sur la concurrence en Go](http://shop.oreilly.com/product/9781783983483.do)
* [Les modalités d'évaluation des TER](https://perso.limsi.fr/allauzen/webpages/pmwiki.php?n=Cours.M1-TER)
