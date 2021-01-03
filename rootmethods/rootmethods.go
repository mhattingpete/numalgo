package rootmethods

import (
    "errors"
    "math"
)

// Linspace returns an array of floats in the range [start, stop) with numsteps numbers
func Linspace(start float64, stop float64, numsteps int) ([]float64, error) {
    if numsteps <= 0 {
        return nil, errors.New("numsteps must be greater than 0")
    }
    stepsize := float64((stop-start)/float64(numsteps))
    x := make([]float64, numsteps, numsteps)
    for i := 0; i < numsteps; i++ {
        x[i] = start + stepsize*float64(i)
    }
    return x, nil
}

func Bisection(f func(float64) float64, xl float64, xu float64, es float64, maxit int) (root float64, fx float64, ea float64, iter int, err error) {
    if test := f(xl)*f(xu); test > 0.0 {
        return 0.0, 0.0, 0.0, 0, errors.New("No sign change in interval")
    }
    iter = 0
    xr := xl
    var xrold float64
    ea = 100
    for ; iter < maxit; iter ++ {
        xrold = xr
        xr = (xl + xu) / 2.0
        if xr != 0 {
            ea = math.Abs((xr - xrold) / xr) * 100.0
        }
        if test := f(xl)*f(xr); test < 0.0 {
            xu = xr
        } else if test > 0.0 {
            xl = xr
        } else {
            ea = 0
        }
        if ea <= es {
            break;
        }
    }
    root = xr
    fx = f(xr)
    return root, fx, ea, iter, nil
}
