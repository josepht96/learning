func climbStairs(n int) int {
	//how is this useful in programming
    a, b := 1, 1
    for i := 0; i < n; i++ {
        a, b = b, a + b
    }
    return a
}