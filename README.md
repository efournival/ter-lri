# ter-lri

Travail d'Étude et de Recherche sur l'exploration de semi-groupes numériques avec mise en place d'un algorithme de vol de tâches

Tuteur : [Florent Hivert](https://www.lri.fr/~hivert/)

## Déroulement

* Étude du problème d'exploration de semi-groupes auquel on va s'intéresser
* Implémentation basique en Go avec parallélisation
* Étude comparative avec la version optimisée en C++
* Implémentation du système distribué
* Mise en place du cluster : image, SSH/DSH, etc.
* Optimisation avancée du Go ou utilisation de cgo avec appel au code de `NumericMonoid` généré par GCC
* Finalisation du système distribué et lancement des calculs

## Problèmes complémentaires

* Distribution du problème des N reines
* Mise en place du vol de tâche dans [Spark](http://spark.apache.org/)

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
