#include "resizeman.h"

VALUE rb_mResizeman;

void
Init_resizeman(void)
{
  rb_mResizeman = rb_define_module("Resizeman");
}
