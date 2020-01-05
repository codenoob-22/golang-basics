func max( a, b int) int{
    if a >= b {
        return a
    } else {
        return b
    }
}
func abs( a int) int{
    if( a < 0) {
        return -1 * a
    } else {
        return a
    }
}
func helper(a, b, c, d int) int {
    return max(abs(a - c), abs(b - d))
}

func minTimeToVisitAllPoints(points [][]int) int {
    sum := 0
    
    for i :=0; i <len(points) -1; i++ {
        sum += helper(points[i][0], points[i][1], points[i+1][0], points[i+1][1])
    }
    
    return sum
}