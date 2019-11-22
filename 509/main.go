package _509

var mem []int = make([]int, 31)

func fib(N int) int {
    if N <= 1 {
        return N
    } else if mem[N] != 0 {
        return mem[N]
    } else {
        mem[N] = fib(N-1) + fib(N-2)
        return mem[N]
    }
}
