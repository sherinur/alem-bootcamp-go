package bootcamp

func ChainFunctions(funcs []func(*string)) func(*string) {
 return func(s *string) {
  for _, f := range funcs {
   f(s)
  }
 }
}