function anagrams(word, words) {
    return words.filter(item => item.split('').sort().join('') ===
                                word.split('').sort().join(''))
  }