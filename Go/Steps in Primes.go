package kata

import "math"

// smartish search for divisors
func isPrime(x int) bool {
// if it is not devised by 2, we can skip all even divisors
  if x % 2 == 0 {
      return false
  }
// try all odd divisors up to square root candidate
  for i := 3; i <= int(math.Floor(math.Sqrt(float64(x)))); i += 2 {
    if x % i == 0 {
      return false
    }
  }
  return true
}

func Step(g, m, n int) []int {    
    // no need to look up in case gap is wider than provided range
    if n - m < g {
      return nil
    }

    for i := m; i <= n-g; i++ {
      // checking pairs with distance g to be both primes
      if isPrime(i) && isPrime(i+g) {
          return []int{i, i+g}
      }
    }
    return nil
}
