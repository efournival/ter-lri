#include "treewalk.hpp"
#include "monoid.hpp"
#include "ctreewalk.h"
#include <cilk/cilk_api.h>
#include <iostream>

int InitCilk()
{
	std::string nproc = "0";
	if (__cilkrts_set_param("nworkers", nproc.c_str() ) != __CILKRTS_SET_PARAM_SUCCESS)
		return 0;
	return __cilkrts_get_nworkers();
}

Monoid InitFullN()
{
	monoid *m = new monoid;
	init_full_N(*m);
	return (void*)m;
}

unsigned long* WalkChildren(Monoid nm)
{
	walk_children(*(monoid*)nm);
	static unsigned long res[MAX_GENUS];
	for (unsigned int i=0; i<MAX_GENUS; i++)
		res[i] = cilk_results[i];
	return res;
}

Monoid RemoveGenerator(Monoid nm, unsigned int generator)
{
	monoid *m = new monoid;
	remove_generator(*m, *(monoid*)nm, generator);
	return (void*)m;
}

void WalkChildrenStack(Monoid nm, unsigned long int* results)
{
	walk_children_stack(*(monoid*)nm, results);
}

unsigned int Genus(Monoid nm)
{
	return ((monoid*)nm)->genus;
}

void FreeMonoid(Monoid nm)
{
	delete (monoid*)nm;
}

GeneratorIterator NewGeneratorIterator(Monoid nm)
{
	auto *it = new generator_iter<CHILDREN>(*(monoid*)nm);
	return (void*)it;
}

int MoveNext(GeneratorIterator gi)
{
	return ((generator_iter<CHILDREN>*)gi)->move_next();
}

unsigned int GetGen(GeneratorIterator gi)
{
	return ((generator_iter<CHILDREN>*)gi)->get_gen();
}

void FreeGeneratorIterator(GeneratorIterator gi)
{
	delete (generator_iter<CHILDREN>*)gi;
}
