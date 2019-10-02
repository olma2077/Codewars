package kata

func MaximumSubarraySum(numbers []int) int {
  max_local := 0
  max := 0
  for _, x := range(numbers) {
    max_local += x
    if max_local < 0 {
      max_local = 0
    } else {
      if max_local > max {
        max = max_local
      }
    }
  }
  return max
}
