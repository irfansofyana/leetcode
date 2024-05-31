func checkPowersOfThree(n int) bool {
  for n > 0 {
    dig := n % 3
    if dig != 0 && dig != 1 {
        return false
    }
    n /= 3
  }

  return n == 0 || n == 1
}
