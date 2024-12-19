package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
	UNTUK TESTING REST API

curl -X POST http://localhost:8080/gaji \
-H "Content-Type: application/json" \
-d '{"jam_kerja": 45, "tarif_per_jam": 20000}'
*/
func main() {

	http.HandleFunc("POST /gaji", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			JamKerja    int     `json:"jam_kerja"`
			TarifPerJam float64 `json:"tarif_per_jam"`
		}

		var resp struct {
			Gaji float64 `json:"gaji"`
		}

		byte, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		err = json.Unmarshal(byte, &input)
		if err != nil {
			log.Println(err)
			return
		}

		if input.JamKerja > 40 {
			lembur := float64(input.JamKerja-40) * input.TarifPerJam * 1.5
			resp.Gaji = (40 * input.TarifPerJam) + lembur
		} else {
			resp.Gaji = float64(input.JamKerja) * input.TarifPerJam
		}

		fmt.Fprintln(w, resp)
	})

	server := http.Server{
		Addr: ":8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}
}
