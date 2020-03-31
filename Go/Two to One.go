package kata

import "strings"

func TwoToOne(s1 string, s2 string) string {
  var result string
  s := s1 + s2
  for i := 'a'; i <= 'z'; i++ {
    if strings.Index(s, string(i)) != -1 {
      result += string(i)
    }
  }
  return result
}
