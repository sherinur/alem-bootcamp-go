package bootcamp

import "github.com/alem-platform/ap"

func WrapperPrintStr(fn func(str *string)) func(str *string) {
 return func(s *string) {
  for id := range *s {
   ap.PutRune(rune((*s)[id]))
  }
  ap.PutRune('\n')
 }
}

func WrapperRot1(fn func(str *string)) func(str *string) {
 return func(s *string) {
  runes := []rune(*s)
  for id := range runes {
   symbol := runes[id]
   if symbol >= 'a' && symbol <= 'z' {
    symbol = (symbol-'a'+1)%26 + 'a'
   } else if symbol >= 'A' && symbol <= 'Z' {
    symbol = (symbol-'A'+1)%26 + 'A'
   }
   runes[id] = symbol
  }
  *s = string(runes)
 }
}

func WrapperRot13(fn func(str *string)) func(str *string) {
 return func(s *string) {
  runes := []rune(*s)
  for id := range runes {
   symbol := runes[id]
   if symbol >= 'a' && symbol <= 'z' {
    symbol = (symbol-'a'+13)%26 + 'a'
   } else if symbol >= 'A' && symbol <= 'Z' {
    symbol = (symbol-'A'+13)%26 + 'A'
   }
   runes[id] = symbol
  }
  *s = string(runes)
 }
}

func WrapperReverseStr(fn func(str *string)) func(str *string) {
 return func(s *string) {
  runes := []rune(*s)
  for i := 0; i < (len(runes))/2; i++ {
   runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
  }
  *s = string(runes)
 }
}

func WrapFunctions(decs []func(fn func(str *string)) func(str *string)) func(str *string) {
 return func(str *string) {
  for _, dec := range decs {
   dec(func(s *string) {})(str)
  }
 }
}