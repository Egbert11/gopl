package main

import (
	"fmt"
	"os"
	"strconv"
)

type Meter float64 // 米
type Foot float64  // 英尺

func MToF(m Meter) Foot {
	return Foot(m * 3.2808)
}

func FToM(f Foot) Meter {
	return Meter(f/3.2808)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}

func (f Foot) String() string {
	return fmt.Sprintf("%g ft", f)
}

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		m := Meter(t)
		f := Foot(t)
		fmt.Printf("%s=%s, %s=%s\n", m, MToF(m), f, FToM(f))
	}
}
