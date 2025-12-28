package main

import (
	"fmt"
	"math/big"
	"time"
)

// Fungsi untuk menghitung digit persistence secara iteratif
func persistenceIteratif(n *big.Int) int {
	count := 0
	temp := new(big.Int).Set(n)
	zero := big.NewInt(0)
	ten := big.NewInt(10)
	prod := new(big.Int)

	for {
		if temp.Cmp(ten) < 0 {
			break
		}

		prod.SetInt64(1)
		t := new(big.Int).Set(temp)

		for t.Cmp(zero) > 0 {
			mod := new(big.Int)
			mod.Mod(t, ten)
			prod.Mul(prod, mod)
			t.Div(t, ten)
		}

		temp.Set(prod)
		count++
	}

	return count
}

// Fungsi untuk menghitung digit persistence secara rekursif
func persistenceRekursif(n *big.Int) int {
	ten := big.NewInt(10)
	if n.Cmp(ten) < 0 {
		return 0
	}

	prod := big.NewInt(1)
	t := new(big.Int).Set(n)
	mod := new(big.Int)

	for t.Cmp(big.NewInt(0)) > 0 {
		mod.Mod(t, ten)
		prod.Mul(prod, mod)
		t.Div(t, ten)
	}

	return 1 + persistenceRekursif(prod)
}

func main() {
	var input string
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&input)

	// Konversi input menjadi big.Int
	n := new(big.Int)
	n, ok := n.SetString(input, 10)
	if !ok {
		fmt.Println("Input tidak valid!")
		return
	}

	// Jumlah pengulangan untuk mengukur waktu
	const repetitions = 100000

	var iterCount, rekCount int

	// Iteratif
	startIter := time.Now()
	for i := 0; i < repetitions; i++ {
		iterCount = persistenceIteratif(n)
	}
	elapsedIter := time.Since(startIter)

	// Rekursif
	startRek := time.Now()
	for i := 0; i < repetitions; i++ {
		rekCount = persistenceRekursif(n)
	}
	elapsedRek := time.Since(startRek)

	// Tampilkan hasil
	fmt.Println("\n=== HASIL PERHITUNGAN DIGIT PERSISTENCE ===")
	fmt.Printf("Versi Iteratif : %d\n", iterCount)
	fmt.Printf("Iteratif total waktu : %v\n", elapsedIter)
	fmt.Printf("Iteratif rata-rata (ns) : %d\n", elapsedIter.Nanoseconds()/repetitions)
	fmt.Printf("Iteratif rata-rata (µs) : %d\n", elapsedIter.Microseconds()/repetitions)

	fmt.Printf("\nVersi Rekursif : %d\n", rekCount)
	fmt.Printf("Rekursif total waktu : %v\n", elapsedRek)
	fmt.Printf("Rekursif rata-rata (ns) : %d\n", elapsedRek.Nanoseconds()/repetitions)
	fmt.Printf("Rekursif rata-rata (µs) : %d\n", elapsedRek.Microseconds()/repetitions)
}
