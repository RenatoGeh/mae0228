package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"math/rand"
)

const (
	SEED = 10101
	W    = 1920
	H    = 1080
)

func newPlot(title, xl, yl string) *plot.Plot {
	P, e := plot.New()
	if e != nil {
		panic(e)
	}
	P.Title.Text = title
	P.X.Label.Text = xl
	P.Y.Label.Text = yl
	return P
}

func throwCoins(n int) []int {
	T := make([]int, n)
	for i := 0; i < n; i++ {
		k := rand.Int() % 2
		if k == 0 {
			T[i] = -1
		} else {
			T[i] = 1
		}
	}
	return T
}

func plotSums(n int) (plotter.XYs, []int) {
	T := throwCoins(n)
	P := make(plotter.XYs, n)
	P[0].X, P[0].Y = 0, 0
	var S int
	for i := 1; i <= n; i++ {
		S += T[i-1]
		P[i-1].X = float64(i)
		P[i-1].Y = float64(S)
	}
	return P, T
}

func plotA(n int) ([]int, []int, []int, int, float64, float64) {
	P := newPlot("Item (a)", "n", "S")

	A, Ta := plotSums(n)
	B, Tb := plotSums(n)
	C, Tc := plotSums(n)
	plotutil.AddLinePoints(P,
		"A", A,
		"B", B,
		"C", C)

	P.Save(W, H, "plot_a.png")
	return Ta, Tb, Tc, n, P.Y.Min, P.Y.Max
}

func plotNSums(n int, T []int) plotter.XYs {
	P := make(plotter.XYs, n)
	var S int
	for i := 0; i < n; i++ {
		S += T[i]
		P[i].X = float64(i + 1)
		P[i].Y = float64(S) / float64(i+1)
	}
	return P
}

func plotC(Ta, Tb, Tc []int, n int, mX, mY float64) ([]int, []int, []int, int, float64, float64) {
	P := newPlot("Item (c)", "n", "Y")

	P.Y.Min, P.Y.Max = mX, mY
	plotutil.AddLinePoints(P,
		"Y_A", plotNSums(n, Ta),
		"Y_B", plotNSums(n, Tb),
		"Y_C", plotNSums(n, Tc))

	P.Save(W, H, "plot_c.png")
	return Ta, Tb, Tc, n, mX, mY
}

func plot2Sums(n int, T []int) plotter.XYs {
	P := make(plotter.XYs, n)
	P[0].X, P[0].Y = 0, float64(T[0])
	X := T[0]
	for i := 1; i < n; i++ {
		Z := X + T[i]
		P[i].X = float64(i)
		P[i].Y = float64(Z) / 2.0
		X = T[i]
	}
	return P
}

func plotD(Ta, Tb, Tc []int, n int, mX, mY float64) {
	P := newPlot("Item (d)", "n", "Z")

	P.Y.Min, P.Y.Max = mX, mY
	plotutil.AddLinePoints(P,
		"Z_A", plot2Sums(n, Ta),
		"Z_B", plot2Sums(n, Tb),
		"Z_C", plot2Sums(n, Tc))

	P.Save(W, H, "plot_d.png")
}

func main() {
	rand.Seed(SEED)
	plotD(plotC(plotA(500)))
}
