package rootmethods

import (
    "testing"
    "reflect"
    "fmt"
)

// TestLinspace calls rootmethods.Linspace with a start, stop and numsteps, checking 
// for a valid return value.
func TestLinspace(t *testing.T) {
    start := 0.0
    stop := 1.0
    numsteps := 10
    want := []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}
    msg, err := Linspace(start, stop, numsteps)
    if reflect.DeepEqual(want, msg) || err != nil {
        t.Fatalf(`Linspace(0.0, 1.0, 11) = %v, %v, want match for %#v, nil`, msg, err, want)
    }
}

// TestLinspaceZero calls rootmethods.Linspace with zero numsteps, 
// checking for an error.
func TestLinspaceZero(t *testing.T) {
    start := 0.0
    stop := 1.0
    numsteps := 0
    msg, err := Linspace(start, stop, numsteps)
    if msg != nil || err == nil {
        t.Fatalf(`Linspace(start, stop, 0) = %v, %v, want empty, error`, msg, err)
    }
}

// TestBisection calls rootmethods.Bisection with a start, stop and numsteps, checking 
// for a valid return value.
func TestBisection(t *testing.T) {
    xl := -10.0
    xu := 10.0
    es := 0.0001
    maxit := 50
    f := func(x float64) float64 {
        return 2.0*x - 3.0
    }
    rootwant := 1.5
    fxwant := 0.0
    root, fx, ea, iter, err := Bisection(f, xl, xu, es, maxit)
    msg := fmt.Sprintf("%f, %f, %f, %d", root, fx, ea, iter)
    want := fmt.Sprintf("%f, %f, %f, %d < %d", rootwant, fxwant, 0.0, iter, maxit)
    rootwithininterval := (root <= rootwant + es) && (root >= rootwant - es)
    fxwithininterval := (fx <= fxwant + es) && (fx >= fxwant - es)
    if !rootwithininterval || !fxwithininterval || ea > es || iter > maxit || err != nil {
        t.Fatalf(`Bisection(f: x->2x-3, -10, 10, 0.0001, 50) = %q, %v, want match for %#v, nil`, msg, err, want)
    }
}

// TestBisectionNoSign calls rootmethods.Bisection with xl and xu, 
// checking for an error when no sign change is found.
func TestBisectionNoSign(t *testing.T) {
     xl := 0.0
    xu := 1.0
    es := 0.0001
    maxit := 50
    f := func(x float64) float64 {
        return 2.0*x - 3.0
    }
    root, fx, ea, iter, err := Bisection(f, xl, xu, es, maxit)
    msg := fmt.Sprintf("%f, %f, %f, %d", root, fx, ea, iter)
    want := fmt.Sprintf("%f, %f, %f, %d", 0.0, 0.0, 0.0, 0)
    if root != 0.0 || fx != 0.0 || ea != 0.0 || iter != 0 || err == nil {
        t.Fatalf(`Bisection(f: x->2x-3, 0, 1, 0.0001, 50) = %q, %v, want %q, error`, msg, err, want)
    }
}