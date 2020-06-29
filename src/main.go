package main


import (
	"fmt"
	"./set1"
)


func main() {
	encrypted := set1.RepeatingKeyXOR("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal", "ICE")
	fmt.Println(encrypted)
}