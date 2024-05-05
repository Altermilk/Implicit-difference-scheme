package main

import (
	"encoding/csv"
	"fmt"
	difur "lab3/lab3/dufur"
	"lab3/lab3/plotting"
	"log"
	"os"
	"strconv"
)

func main() {

  u01 := difur.Difur{H: 0.1, Dt: 0.1}.SetN().LeftImplicitScheme()
  u001 := difur.Difur{H: 0.01, Dt: 0.01}.SetN().LeftImplicitScheme()
  makeCSV(u01.U, "results-h01")
  makeCSV(u001.U, "results-h001")

  plotting.BuildPlot(u01.U, u01.Nt, u01.Nx, "lab3_h01")
  plotting.BuildPlot(u001.U, u001.Nt, u001.Nx, "lab3_h001")

  plotting.ComparePlots(u01, u001, "НРС h=0,1", "НРС h=0,01")

}

func makeCSV(u [][]float64, filename string) {
  f, err := os.Create(filename + ".csv")

  if err != nil {
    log.Fatalln("failed to open file", err)
  }

  w := csv.NewWriter(f)

  for i := 0; i < len(u); i++ {
    tmp := make([]string, len(u[0]))
    for j := 0; j < len(u[0]); j++ {
      tmp[j] = strconv.FormatFloat(u[i][j], 'f', 2, 64)
    }
    if err := w.Write(tmp); err != nil {
      log.Fatal(err)
    }
  }
  w.Flush()
  f.Close()
  fmt.Println("CSV: \"", f.Name(), "\" was created")
}

