package main

import "fmt"

func buy(bitcoin, altcoin bool) (bool, bool, bool) {
	buyBitcoin := bitcoin != altcoin
	buyCrypto := bitcoin || altcoin
	buyAnything := bitcoin && altcoin
	return buyBitcoin, buyCrypto, buyAnything
}

func main() {
	buyBitcoin, buyCrypto, buyAnything := buy(true, true)
	fmt.Printf("Buy bitcoin: %t Buy crypto: %t Buy altcoin: %t", buyBitcoin, buyCrypto, buyAnything)
}
