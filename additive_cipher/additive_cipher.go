package main

import (
	"fmt"
)

var (
	v2p map[int]string
	v2c map[int]string
	p2v map[string]int
	c2v map[string]int
	i   int
)

func p2c(plain string, k int) string {
	var cipher string
	for _, p := range plain {
		c := (p2v[string(p)] + k) % 26
		cipher += v2c[c]
	}
	fmt.Printf("\nPlaintext: %s", plain)
	fmt.Printf("\nKey: %d", k)
	return cipher
}

func c2p(cipher string, k int) string {
	var plain string
	for _, c := range cipher {
		c1 := c2v[string(c)] - k
		if c1 <= 0 {
			c1 += 26
		}
		p := (c1) % 26
		plain += v2p[p]
	}
	fmt.Printf("\nCiphertext: %s", cipher)
	fmt.Printf("\nKey: %d", k)
	return plain
}

func main() {

	v2p = make(map[int]string)
	v2c = make(map[int]string)
	p2v = make(map[string]int)
	c2v = make(map[string]int)

	// Plaintext Map v --> p
	for i = 97; i <= 122; i++ {
		v2p[i-97] = string(rune(i))
	}

	// Plaintext Map p --> v
	for i = 97; i <= 122; i++ {
		p2v[string(rune(i))] = i - 97
	}

	// Ciphertext Map v --> c
	for i = 65; i <= 90; i++ {
		v2c[i-65] = string(rune(i))
	}

	// Ciphertext Map c --> v
	for i = 65; i <= 90; i++ {
		c2v[string(rune(i))] = i - 65
	}

	plain := "hello"
	k := 15
	fmt.Printf("\nENCRYPTION:")
	fmt.Printf("\nCiphertext: %s\n", p2c(plain, k))
	fmt.Printf("\nDECRYPTION:")
	fmt.Printf("\nPlaintext: %s\n", c2p("WTAAD", k))
}
