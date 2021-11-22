package main

func main() {
	_ = int(1.23)     // 1.23不能被表示为int类型值。
	_ = uint8(-1)     // -1不能被表示为uint8类型值。
	_ = float64(1+2i) // 1+2i不能被表示为float64类型值。

	// -1e+1000不能被表示为float64类型值。不允许溢出。
	_ = float64(-1e1000)
	// 0x10000000000000000做为int值将溢出。
	_ = int(0x10000000000000000)

	// 字面量65.0的默认类型是float64（不是一个整数类型）。
	_ = string(65.0)
	// 66+0i的默认类型是complex128（不是一个整数类型）。
	_ = string(66+0i)
}