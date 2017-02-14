#ifndef _GO_TREEWALK_H_
#define _GO_TREEWALK_H_

#ifdef __cplusplus
extern "C" {
#endif

	typedef void* Monoid;
	int InitCilk();
	unsigned long* WalkChildren(Monoid);
	Monoid InitFullN(void);
	void FreeMonoid(Monoid);

#ifdef __cplusplus
}
#endif

#endif
