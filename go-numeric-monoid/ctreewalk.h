#ifndef _GO_TREEWALK_H_
#define _GO_TREEWALK_H_

#ifdef __cplusplus
extern "C" {
#endif

typedef void* Monoid;
int InitCilk();
unsigned long* WalkChildren(Monoid);
Monoid InitFullN(void);
Monoid RemoveGenerator(Monoid, unsigned int);
void WalkChildrenStack(Monoid, long unsigned int*);
unsigned int Genus(Monoid);
void FreeMonoid(Monoid);

typedef void* GeneratorIterator;
GeneratorIterator NewGeneratorIterator(Monoid);
int MoveNext(GeneratorIterator);
unsigned int GetGen(GeneratorIterator);
void FreeGeneratorIterator(GeneratorIterator);

#ifdef __cplusplus
}
#endif

#endif
