package main

import (
	"fmt"

	"github.com/dwisiswant0/unch/pkg/unch"
)

func main() {
	message := "1f y0u c4n r34d th15 s3cr3t m3ss4g3, c0ngr4tul4t10ns!"

	plain := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	encoded, _ := unch.Encode(message, plain)
	decoded, _ := unch.Decode(encoded)

	fmt.Println("message:", message)
	fmt.Println("plain:", plain)
	fmt.Println("plain len:", len(plain))
	println()
	fmt.Println("encoded:", encoded)
	fmt.Println("encoded len:", len(encoded))
	fmt.Println("encoded == plain:", encoded == plain)
	println()
	fmt.Println("decoded:", decoded)
	fmt.Println("decoded == message:", decoded == message)
}
