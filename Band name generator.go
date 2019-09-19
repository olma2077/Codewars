package kata

import "strings"

func bandNameGenerator(word string) string {
  if word[0] == word[len(word)-1] {
    return strings.Title(word + word[1:])
  } else {
    return "The " + strings.Title(word)
  }
}
