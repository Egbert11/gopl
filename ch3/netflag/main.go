package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func TurnDown(v *Flags)  {
	*v &^= FlagUp
}

func SetBroadcast(v *Flags)  {
	*v |= FlagBroadcast
}

func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}


const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func main()  {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))

	fmt.Printf("%d\n", KiB)
	fmt.Printf("%d\n", MiB)
	fmt.Printf("%d\n", GiB)
	fmt.Printf("%d\n", TiB)
	fmt.Printf("%d\n", PiB)
	fmt.Printf("%d\n", EiB)
	fmt.Printf("%d\n", ZiB)
	fmt.Printf("%d\n", YiB)
}
