package kata

import "strconv"

func Min(arr []int64) (int, int64) {
  i, min := 0, arr[0]
  for j, x := range arr {
    if x < min {
      i, min = j, x
    }
  }
  return i, min
}

func IsCoprime(arr []int64) bool {
    // if one is factor of another
    for _, i := range arr {
        for _, j := range arr {
            if i == j {
                continue
            }
            if i % j == 0 {
                return false
            }
        }  
    }
    
    // Looking for common factors up to sqrt of the minimal moduli.
    _, min := Min(arr)
    for i := 2; int64(i * i) <= min; i++ {
        count := 0
        for _, x := range arr {
            if x % int64(i) == 0 {
                count++
            }
        }
        if count > 1 {
            return false
        }
    }
    return true
}

func FromNbToStr(n int64, sys []int64) string {
    // product check
    prod := int64(1)
    for _, i := range sys {
        prod *= i
    }
    if prod < n {
        return "Not applicable"
    }    
    
    if !IsCoprime(sys) {
      return "Not applicable"
    }
    
    result := ""
    for _, x := range sys {
        result = result + "-" + strconv.Itoa(int(n % x)) + "-"
    }
    return result
}
