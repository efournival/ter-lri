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
	monoid* m = new monoid;
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

void FreeMonoid(Monoid nm)
{
	monoid* m = (monoid*)nm;
	delete m;
}
