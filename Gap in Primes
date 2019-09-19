// This solution includes 2 varians of looking up for a requested gap.
// First one is faster but dirty: modifies loop counter within loop.

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

func Gap(g, m, n int) []int {    
    // no need to look up in case gap is wider than provided range
    if n - m < g {
      return nil
    }

// First variant, looking up for expected gap, than checking there are no other primes within
    for i := m; i <= n-g; i++ {
      // checking pairs with distance g to be both primes
      if isPrime(i) && isPrime(i+g) {
        var nexti int
        // none between them should be prime
        for j := i+1; j < i+g; j++ {  
          if isPrime(j) {
            nexti = j
            break
          }
        }
        if nexti == 0 {
          return []int{i, i+g}
        } else {
          // if prime was found, we can skip to it for prime searching
          i = nexti-1
        }
      }
    }
// Second variant, looking for a prime gap, then checking if it of requested size
/*      if isPrime(i) {
        for j := i+1; j <= i+g; j++ {
          if isPrime(j) {
            if j - i == g {
              return []int{i, i+g}
            }
            break
          }
        }
      }
    }
*/    return nil
}
