package difur

import "math"

const(
	Nx_end int = 1
  	Nt_end int = 1
)

type Difur struct {
	H float64
	Dt float64
	Nx int
	Nt int
	U [][]float64
}

func Cut(interval int, delta float64) int {
	return (int)(math.Ceil(float64(interval) / delta))
}

func (difur Difur) SetN() Difur {
	difur.Nx = Cut(Nx_end, difur.H)
	difur.Nt = Cut(Nt_end, difur.Dt)
	return difur
}

func (difur Difur) LeftImplicitScheme() Difur{
	H, Dt, Nx, Nt := difur.H , difur.Dt, difur.Nx, difur.Nt
  
	u := make([][]float64, Nt)
	for i := range u {
	  u[i] = make([]float64, Nx)
	}
  
	for j := 1; j < Nx+1; j++ {
		  u[0][j-1] = (3 * math.Pow(H*float64(j-1), 2))
	  }
  
	for n := 0; n < Nt - 1; n++ {
  
	  u[n+1][0] = 3 * float64(n+1) * Dt
  
	  for j := 1; j < Nx; j++ {
		u[n+1][j] = ( u[n][j] + (Dt*u[n+1][j-1])/(2*H) + 3*Dt*(H*float64(j)+1) ) / ( 1 + Dt/(2*H) )
	  }
	}
	difur.U = u
	return difur
  }