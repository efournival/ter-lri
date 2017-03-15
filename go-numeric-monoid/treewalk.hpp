#include <cilk/cilk.h>
#include <cilk/reducer.h>
#include <cilk/reducer_list.h>
#include "monoid.hpp"

typedef unsigned long int results_type[MAX_GENUS];
void walk_children_stack(monoid m, results_type &res);
void walk_children_stack(monoid m, ind_t bound, results_type &res);

struct Results
{
  results_type values;
  inline Results() {reset();};
  inline void reset() {for(int i=0; i<MAX_GENUS; i++) values[i] = 0;};
};

class ResultsReducer
{
  struct Monoid: cilk::monoid_base<Results>
  {
    inline static void reduce (Results *left, Results *right){
      for(auto i=0; i<MAX_GENUS; i++) left->values[i] += right->values[i];
    }
  };
private:
  cilk::reducer<Monoid> imp_;
public:
  ResultsReducer() : imp_() {};
  inline unsigned long int & operator[](ind_t i) {
    return imp_.view().values[i];
  };
  inline results_type &get_array() {return imp_.view().values;}
  inline void reset() {imp_.view().reset();}
};

extern ResultsReducer cilk_results;

#ifdef TBB
#include <tbb/scalable_allocator.h>
extern cilk::reducer_list_append<monoid, tbb::scalable_allocator<monoid>> cilk_list_results;
#else
extern cilk::reducer_list_append<monoid> cilk_list_results;
#endif


void walk_children(monoid m);
void walk_children(monoid m, ind_t bound);
void list_children(monoid m, ind_t bound);
