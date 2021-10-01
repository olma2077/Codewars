var fibonacci = function(n) {
    let  cache = [0, 1]
    
    let fib_rec = (x) => {
      if(typeof(cache[x]) === 'undefined') {
        cache[x] = fib_rec(x-1) + fib_rec(x-2)
      } 
      return cache[x]
    }

    return fib_rec(n)
  }