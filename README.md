# ter-lri

Travail d'Étude et de Recherche sur l'exploration de semigroupes numériques avec mise en place d'un algorithme de vol de tâches.

Il se déroule au LRI (Laboratoire de Recherche en Informatique), rattaché à l'Université Paris-Sud et au CNRS, dans l'équipe GALaC (Graphes, Algorithmes et Combinatoire).

Tuteur : [Florent Hivert](https://www.lri.fr/~hivert/)

## Déroulement

- [x] Étude du problème d'exploration de semigroupes auquel on va s'intéresser
- [x] Implémentation basique en Go
- [x] Parallélisation et optimisation basique
- [x] Étude comparative avec la version optimisée en C++
- [ ] Profiler le code Go avec perf et pprof
- [ ] Documentation détaillée du fonctionnement des algorithmes (naïf et optimisé)
- [ ] Optimisation avancée du Go ou utilisation de cgo avec appel au code de `NumericMonoid` généré par GCC
- [ ] Test du compilateur Intel et de LLVM/Clang
- [ ] Implémentation du système distribué
- [ ] Mise en place du cluster : image, SSH/DSH, etc.
- [ ] Finalisation du système distribué et lancement des calculs

## Problèmes complémentaires

* Distribution du problème des N reines
* [Mise en place du vol de tâche](https://github.com/Didayolo/spark) dans [Spark](http://spark.apache.org/)
* Essais avec CUDA ?
* Utiliser des instructions AVX1 voire AVX2

## Ressources

* [NumericMonoid](https://github.com/hivert/NumericMonoid)
* [Machinery](https://github.com/RichardKnop/machinery) : système de distribution de tâches
* [Glow](https://github.com/chrislusf/glow) : un framework de distribution MapReduce similaire à Hadoop, Spark, etc.
* [Gleam](https://github.com/chrislusf/gleam) : pareil que Glow mais plus performant et utilise LuaJIT
* [Un livre sur la programmation réseau en Go](https://www.gitbook.com/book/jannewmarch/network-programming-with-go-golang-/details)
* [Un livre sur la concurrence en Go](http://shop.oreilly.com/product/9781783983483.do)
* [Functory](http://functory.lri.fr/About.html) : une bibliothèque robuste de calcul distribué en OCaml
* [Exploring the tree of numerical semigroups](https://hal.archives-ouvertes.fr/hal-00823339/document)
* [HPC in Combinatorics: Application of Work-Stealing](https://github.com/OpenDreamKit/OpenDreamKit/raw/master/WP5/T5.6/HPC-Combi.pdf)
* [Profiling en Go](https://blog.golang.org/profiling-go-programs)
