#ifndef _GO_CTREEWALK_H_
#define _GO_CTREEWALK_H_

#include <stdint.h>

extern const int SIZEOF_MONOID;

#ifdef __cplusplus
extern "C" {
#endif

typedef void* Monoid;
int InitCilk();
unsigned long* WalkChildren(Monoid);
Monoid InitFullN();
Monoid RemoveGenerator(Monoid, uint_fast64_t);
void WalkChildrenStack(Monoid, unsigned long int*);
unsigned int Genus(Monoid);
void Print(Monoid);

typedef void* GeneratorIterator;
GeneratorIterator NewGeneratorIterator(Monoid);
int MoveNext(GeneratorIterator);
uint_fast64_t GetGen(GeneratorIterator);
uint8_t Count(GeneratorIterator);

#ifdef __cplusplus
}
#endif

#endif
