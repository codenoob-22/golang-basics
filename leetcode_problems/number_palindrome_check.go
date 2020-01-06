func isPalindrome(x int) bool {
    a := 0
    y := x
    for y > 0 {
        a *= 10
        a += y % 10
        y /= 10
    }
    return a == x 
}
