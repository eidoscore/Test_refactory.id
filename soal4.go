package main

import (
	"fmt"
	"math/rand"
	"time"
)

func xrand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	nomortebakan := xrand(1, 100)
	attempts := 0
	var tebakan int

	fmt.Println("Pilih Angka 1 - 100.")

	for attempts < 10 {
		fmt.Println("Masukan Angka :")
		fmt.Scanf("%d", &tebakan)
		attempts++
		if tebakan < nomortebakan {
			fmt.Println("Terlalu Rendah.")
		}
		if tebakan > nomortebakan {
			fmt.Println("Terlalu Tinggi.")
		}
		if tebakan == nomortebakan {
			break
		}
	}
	if tebakan == nomortebakan {
		fmt.Printf("Tebakan Kamu ada %d \n", attempts)
	} else {
		fmt.Printf("Nomor Tebakannya adalah %d\n", nomortebakan)
	}
}
