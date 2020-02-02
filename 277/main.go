package _277

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
    return func(n int) int {
        famous := 0
        for i := 1; i < n; i++ {
            if famous != i && knows(famous, i) {
                famous = i
            }
        }
        for i := 0; i < n; i++ {
            if famous != i {
                if knows(famous, i) || !knows(i, famous) {
                    return -1
                }
            }
        }
        return famous
    }
}
