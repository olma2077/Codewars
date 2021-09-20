function findMissingLetter(array) {
  for (i = 0; i < array.length; i++) {
    if (array[i].charCodeAt() - array[0].charCodeAt() !== i) {
      return String.fromCharCode(array[i - 1].charCodeAt() + 1)
    } 
  }
}
