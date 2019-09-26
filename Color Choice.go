ackage kata

import "math"

// Factorial won't fit even uint64, so we should work around.

func Choose(n, x int) int {
  // Edge cases, avoid calculations
  if x == 1 {
    return n
  } else if n == x || x == 0 {
    return 1
  }
  
  // We'll reduct the biggest factorial in denominator and part of factorial in numerator. 
  bottom := x
  if bottom < n - x {
    bottom = n - x
  } else {
    x = n - x
  }
    
  result := 1.0
  
  // Multiply-divide in turns to avoid variable overflow until we
  // deplete both numerator factorial leftover and denominator factorial. 
  for n - bottom > 0 || x > 0 {
    if bottom < n {
      bottom++
      result *= float64(bottom)
    }
    if x > 0 {
      result /= float64(x)
      x--
    }
  }
  
  return int(math.Round(result))
}

func CheckChoose(m, n int) int {
  for x := 0; x <= n; x++ {
    if m == Choose(n, x) {
      return x
    }
  }
  return -1
}
