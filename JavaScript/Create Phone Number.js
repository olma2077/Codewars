function createPhoneNumber(numbers){
    let res = '('
    for (i = 0; i < 3; i++) {
      res += numbers[i].toString()
    }
    res += ') '
    for (i = 3; i < 6; i++) {
      res += numbers[i].toString()
    }
    res += '-'
    for (i = 6; i < 10; i++) {
      res += numbers[i].toString()
    }
    return res
}
