function findOutlier(integers){
    let parity
    
    if ((integers[0] + integers[1]) % 2) {
      if ((integers[0] + integers[2]) % 2) {
        return integers[0]
      } else {
        return integers[1]
      }
    } else {
      parity = integers[0] % 2
    }
     
    for (i = 2; i < integers.length; i++) {
      if (integers[i] % 2 && !parity) {
          return integers[i]
      }
      if (!(integers[i] % 2) && parity) {
          return integers[i]
      }
    }
}
