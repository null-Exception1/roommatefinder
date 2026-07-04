package handlers

import (
	"encoding/json"
	"fmt"
	"golang/db"
	"golang/globals"
	"golang/structs"
	"net/http"
)

func Blocks(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/text")

	Occupancies := make(map[string]map[string]int, 0)
	rows := db.Query("SELECT * FROM people", globals.Globaldb)

	//fmt.Println(rows)
	for _, row := range rows {
		if _, ok := Occupancies[row.Blockno]; ok {
			Occupancies[row.Blockno][row.Roomno] = Occupancies[row.Blockno][row.Roomno] + 1
		} else {
			Occupancies[row.Blockno] = make(map[string]int, 0)
			Occupancies[row.Blockno][row.Roomno] = 1
		}
	}
	Block := make(map[string]*structs.Block, 0)

	for key := range Occupancies {
		Block[key] = &structs.Block{Partial: 0, Full: 0}
		for _, occupancy := range Occupancies[key] {
			if occupancy >= 2 { // remind me to check the occupancy max limit later
				Block[key].Full = Block[key].Full + 1
			} else {
				Block[key].Partial = Block[key].Partial + 1
			}
		}
	}
	//fmt.Println(Block)
	str, _ := json.Marshal(Block)

	fmt.Fprintf(w, "%s", string(str))
}
