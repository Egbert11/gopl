package popcount

import "fmt"


var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountRewrite(x uint64) int {
	count := 0
	for x != 0 {
		x &= x-1
		count += 1
	}
	return count
}

func main() {
	var i uint64 = 0
	for ; i < 100; i++ {
		fmt.Printf("%d:%d\n", i, PopCountRewrite(i))
	}
}