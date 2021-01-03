package main

import (
	"fmt"
	"math"
	"example.com/rootmethods"
	"example.com/optimization"
)

func main() {
	start := 0.0
    stop := 1.0
    numsteps := 10
    msg, err := rootmethods.Linspace(start, stop, numsteps)
    fmt.Println(msg,err)
    fmt.Println("\nRootmethods")
    fmt.Println("\nBisection")
    // Bisection
    xl := -10.0
    xu := 10.0
    es := 0.0001
    maxit := 50
    f := func(x float64) float64 {
        return 2.0*x - 3.0
    }
    root, fx, ea, iter, err := rootmethods.Bisection(f, xl, xu, es, maxit)
    fmt.Println(root, fx, ea, iter, err)
    fmt.Println("\nNewton-Raphson")
    // Newtraph
    xr := -10.0
    df := func(x float64) float64 {
        return 2.0
    }
    root, fx, ea, iter, err = rootmethods.Newtraph(f, df, xr, es, maxit)
    fmt.Println(root, fx, ea, iter, err)
    fmt.Println("\nSecant")
    // Secant
    p := 1e-6
    root, fx, ea, iter, err = rootmethods.Secant(f, p, xr, es, maxit)
    fmt.Println(root, fx, ea, iter, err)
    fmt.Println("\nInverseQuadracticInterpolation")
    // InverseQuadracticInterpolation
    p = 1e-1
    root, fx, ea, iter, err = rootmethods.InverseQuadracticInterpolation(f, p, xr, es, maxit)
    fmt.Println(root, fx, ea, iter, err)
	fmt.Println("\nBrentsMethod")
    // BrentsMethod
    es = 1e-6
    root, fx, ea, iter, err = rootmethods.BrentsMethod(f, xl, xu, es, maxit)
    fmt.Println(root, fx, ea, iter, err)
    fmt.Println("\n\nOptimizations")
    fmt.Println("\nGoldmin")
    // Goldmin
    xl = 0.0
    xu = 4.0
    es = 1e-4
    maxit = 50
    f = func(x float64) float64 {
        return (x*x)/10.0 - 2.0*math.Sin(x)
    }
    root, fx, ea, iter, err = optimization.Goldmin(f, xl, xu, es, maxit)
    fmt.Println(root, fx, ea, iter, err)
    fmt.Println("\nParabolic")
    // Parabolic
    xm := 1.0
    root, fx, ea, iter, err = optimization.Parabolic(f, xl, xm, xu, es, maxit)
    fmt.Println(root, fx, ea, iter, err)
}
