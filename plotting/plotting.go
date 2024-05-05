package plotting

import (
	"fmt"
	"io"
	"os"
	difur "lab3/lab3/dufur"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func BuildPlot(u [][]float64, Nt, Nx int, title string) {
	page := components.NewPage()
	page.AddCharts(
		surface3DBase(u, Nt, Nx, title),
	)
	f, err := os.Create(title + ".html")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
	fmt.Println("PLOT: \"", f.Name(), "\" was created")
}

func surface3DBase(u [][]float64, Nt, Nx int, title string) *charts.Surface3D {
	surface3d := charts.NewSurface3D()
	surface3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange:    &opts.VisualMapInRange{Color: surfaceRangeColor},
			Max:        6,
			Min:        -1,
		}),
	)

	surface3d.AddSeries(title, surface3dData1(u, Nt, Nx))
	return surface3d
}

func ComparePlots(u1, u2 difur.Difur, title1, title2 string) {
	page := components.NewPage()
	page.AddCharts(
		surface3DBase2(u1, u2, title1, title2),
	)
	f, err := os.Create(title1 + title2 +"plots Comparision.html")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
	fmt.Println("PLOT: \"", f.Name(), "\" was created")
}

func surface3DBase2(u1, u2 difur.Difur, title1, title2 string) *charts.Surface3D {
	surface3d := charts.NewSurface3D()
	surface3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title1 + title2}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange:    &opts.VisualMapInRange{Color: surfaceRangeColor},
			Max:        35,
			Min:        0,
		}),
	)

	surface3d.AddSeries(title1, surface3dData1(u1.U, u1.Nt, u1.Nx))
	surface3d.AddSeries(title2, surface3dData1(resample(u1.U, u2.U), u1.Nt, u1.Nx))
	return surface3d
}

func resample(u1, u2 [][]float64) [][]float64 {
    r := make([][]float64, len(u1))
    for i := range r {
        r[i] = make([]float64, len(u1[0]))
    }

    dt := (int)(len(u2) / len(u1))
    h := (int)(len(u2[0]) / len(u1[0]))
	
    r[0][0] = u2[0][0]

    for n := 0; n < len(u1); n++ {
        for j := 0; j < len(u1[0]); j++ {
            r[n][j] = u2[n*dt][j*h]
        }
    }

    return r
}

func surface3dData1(u [][]float64, Nt, Nx int) []opts.Chart3DData {
	data := make([][3]interface{}, 0)
	for i := 0; i < Nt; i++ {
		y := i
		for j := 0; j < Nx; j++ {
			x := j
			z := u[i][j]
			data = append(data, [3]interface{}{x, y, z})
		}
	}

	ret := make([]opts.Chart3DData, 0, len(data))
	for _, d := range data {
		ret = append(ret, opts.Chart3DData{
			Value: []interface{}{d[0], d[1], d[2]},
		})
	}
	return ret
}

var surfaceRangeColor = []string{
	"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
	"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}
