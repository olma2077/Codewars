package kata

func max(arr []int) (int, int) {
  i, max := 0, 0
  for j, x := range arr {
    if x > max {
      i, max = j, x
    }
  }
  return i, max
}

func MakeValley(arr []int) []int {
    length := len(arr)
    arr2 := make([]int, length)
    for i := 0; i < length; i++ {
      j, x := max(arr)
      arr[j] = arr[len(arr)-1]
      arr = arr[:len(arr)-1]
      if i % 2 == 0 {
        arr2[i/2] = x
      } else {
        arr2[length-i/2-1] = x
      }
    }
    return arr2
}
