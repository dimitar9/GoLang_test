package main

import (
    "fmt"
    "time"
    "math/rand"
    "math"
    "math/cmplx"
    "runtime"
)

func main() {
	fmt.Println("Welcome here!",1)
    fmt.Println("The time is",time.Now())
    fmt.Println("My favorite number is", rand.Intn(10))
    fmt.Printf("Now you have %g problems.",math.Nextafter(2,3))
    fmt.Println(math.Pi)
    fmt.Println(add(1,2))

    a, b := swap("hello", "world")
    fmt.Println(a,b)

    var z  complex128 = cmplx.Sqrt(-5 + 12i)
    fmt.Printf("z is %T %v ", z ,z)

    sum := 0
    for i:= 0; i< 10; i++ {
    	sum += i
    }
    fmt.Println(sum)

    fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	var arr[2] string
	arr[0] = "hello"
	fmt.Println(arr[0],arr)
}

func add(x int, y int) int {
	return x + y
}

func swap(x,y string) (string, string) {
	return y, x
}