package main

import "fmt"

func main() {

	var input struct {
		JamKerja    int
		TarifPerJam float64
	}

	var gaji float64

	fmt.Print("masukan jumlah jam kerja: ")
	fmt.Scan(&input.JamKerja)
	fmt.Print("masukan tarif jam kerja: ")
	fmt.Scan(&input.TarifPerJam)

	if input.JamKerja > 40 {
		lembur := float64(input.JamKerja-40) * input.TarifPerJam * 1.5
		gaji = (40 * input.TarifPerJam) + lembur
	} else {
		gaji = float64(input.JamKerja) * input.TarifPerJam
	}

	fmt.Printf("Gaji karyawan adalah: rp%.2f\n", gaji)
}
