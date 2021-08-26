function tribonacci(signature,n){
    res = [ ...signature ]
    if (n < 4) {
      res.length = n
    } else {
      for (i = n - 3; i > 0; i--) {
        l = res.length
        res.push(res[l-1] + res[l-2] + res[l-3])
      }
    }
    return res
}
