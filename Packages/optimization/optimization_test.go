package optimization

import (
    "testing"
    "fmt"
    "math"
)

// TestGoldmin calls optimization.Goldmin with a function, x lower, x upper, error limit and max iterations, checking 
// for a valid return value.
func TestGoldmin(t *testing.T) {
    xl := 0.0
    xu := 4.0
    es := 1e-4
    maxit := 50
    f := func(x float64) float64 {
        return (x*x)/10.0 - 2.0*math.Sin(x)
    }
    xwant := 1.4276
    fxwant := -1.7757
    x, fx, ea, iter, err := Goldmin(f, xl, xu, es, maxit)
    msg := fmt.Sprintf("%f, %f, %f, %d", x, fx, ea, iter)
    want := fmt.Sprintf("%f, %f, %f, %d < %d", xwant, fxwant, es, iter, maxit)
    xwithininterval := (x <= xwant + es) && (x >= xwant - es)
    fxwithininterval := (fx <= fxwant + es) && (fx >= fxwant - es)
    if !xwithininterval || !fxwithininterval || ea > es || iter > maxit || err != nil {
        t.Fatalf(`Goldmin(f, 0, 4, 1e-4, 50) = %q, %v, want match for %#v, nil`, msg, err, want)
    }
}

// TestParabolic calls optimization.Parabolic with a function, x lower, x upper, error limit and max iterations, checking 
// for a valid return value.
func TestParabolic(t *testing.T) {
    xl := 0.0
    xm := 1.0
    xu := 4.0
    es := 1e-4
    maxit := 50
    f := func(x float64) float64 {
        return (x*x)/10.0 - 2.0*math.Sin(x)
    }
    xwant := 1.4276
    fxwant := -1.7757
    x, fx, ea, iter, err := Parabolic(f, xl, xm, xu, es, maxit)
    msg := fmt.Sprintf("%f, %f, %f, %d", x, fx, ea, iter)
    want := fmt.Sprintf("%f, %f, %f, %d < %d", xwant, fxwant, es, iter, maxit)
    xwithininterval := (x <= xwant + es) && (x >= xwant - es)
    fxwithininterval := (fx <= fxwant + es) && (fx >= fxwant - es)
    if !xwithininterval || !fxwithininterval || ea > es || iter > maxit || err != nil {
        t.Fatalf(`Parabolic(f, 0, 1, 4, 1e-4, 50) = %q, %v, want match for %#v, nil`, msg, err, want)
    }
}
