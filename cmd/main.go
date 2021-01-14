package main

import (
	"fmt"
	"github.com/smartwalle/tsv"
)

func main() {
	t, err := tsv.OpenTSV("./tsv.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range t.Rows() {
		fmt.Println(row.Value(0), row.Value(1), row.Value(2))
	}

	//c , err := tsv.OpenCSV("./csv.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(c.RowAtIndex(0).Int64(0))
}
