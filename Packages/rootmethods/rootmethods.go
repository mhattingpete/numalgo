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

// Bisection 
// input:
// the function to find the root for (f), lower limit (xl), upper limit (xu), error deviation (es), maximum iterations (iter)
// output: 
// the estimated root (root), function value (fx), error estimate (ea), iterations done (iter)
func Bisection(f func(float64) float64, xl float64, xu float64, es float64, maxit int) (root float64, fx float64, ea float64, iter int, err error) {
    if es < 0.0 {
        return 0.0, 0.0, 0.0, 0, errors.New("es must be greater than 0")
    }
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

// Newtraph (Newton-Raphson)
// input:
// the function to find the root for (f), derivative function of f (df), initial guess (xu), error deviation (es), maximum iterations (iter)
// output: 
// the estimated root (root), function value (fx), error estimate (ea), iterations done (iter)
func Newtraph(f func(float64) float64, df func(float64) float64, xr float64, es float64, maxit int) (root float64, fx float64, ea float64, iter int, err error) {
    if es < 0.0 {
        return 0.0, 0.0, 0.0, 0, errors.New("es must be greater than 0")
    }
    iter = 0
    var xrold float64
    ea = 100
    for ; iter < maxit; iter ++ {
        xrold = xr
        xr -= f(xr)/df(xr)
        if xr != 0 {
            ea = math.Abs((xr - xrold) / xr) * 100.0
        }
        if ea <= es {
            break;
        }
    }
    root = xr
    fx = f(xr)
    return root, fx, ea, iter, nil
}

// Secant (Variation of Newton-Raphson)
// input:
// the function to find the root for (f), pertubation fraction (p), initial guess (xr), error deviation (es), maximum iterations (iter)
// output: 
// the estimated root (root), function value (fx), error estimate (ea), iterations done (iter)
func Secant(f func(float64) float64, p float64, xr float64, es float64, maxit int) (root float64, fx float64, ea float64, iter int, err error) {
    if es < 0.0 {
        return 0.0, 0.0, 0.0, 0, errors.New("es must be greater than 0")
    }
    iter = 0
    var xrold float64
    ea = 100
    for ; iter < maxit; iter ++ {
        xrold = xr
        xr -= (p*xr*f(xr))/(f(xr+p*xr)-f(xr))
        if xr != 0 {
            ea = math.Abs((xr - xrold) / xr) * 100.0
        }
        if ea <= es {
            break;
        }
    }
    root = xr
    fx = f(xr)
    return root, fx, ea, iter, nil
}

// InverseQuadracticInterpolation
// input:
// the function to find the root for (f), pertubation fraction (p), initial guess (xr), error deviation (es), maximum iterations (iter)
// output: 
// the estimated root (root), function value (fx), error estimate (ea), iterations done (iter)
func InverseQuadracticInterpolation(f func(float64) float64, p float64, xr float64, es float64, maxit int) (root float64, fx float64, ea float64, iter int, err error) {
    if es < 0.0 {
        return 0.0, 0.0, 0.0, 0, errors.New("es must be greater than 0")
    }
    iter = 0
    var xrold, x1, x2, y1, yr, y2 float64
    ea = 100
    for ; iter < maxit; iter ++ {
        xrold = xr
        x1 = xr - p*xr
        x2 = xr + p*xr
        y1 = f(x1)
        yr = f(xr)
        y2 = f(x2)
        xr = ((y1*yr)/((y2-y1)*(y2-yr)))*x2 + ((y2*yr)/((y1-y2)*(y1-yr)))*x1 + ((y2*y1)/((yr-y2)*(yr-y1)))*xr
        if xr != 0 {
            ea = math.Abs((xr - xrold) / xr) * 100.0
        }
        if ea <= es {
            break;
        }
    }
    root = xr
    fx = f(xr)
    return root, fx, ea, iter, nil
}

// BrentsMethod
// A method that combines the Secant method, Bisection and Inverse Quadractic Interpolation
// input:
// the function to find the root for (f), lower limit (xl), upper limit (xu), error deviation (es), maximum iterations (iter)
// output: 
// the estimated root (root), function value (fx), error estimate (ea), iterations done (iter)
func BrentsMethod(f func(float64) float64, xl float64, xu float64, es float64, maxit int) (root float64, fx float64, ea float64, iter int, err error) {
    if es < 0.0 {
        return 0.0, 0.0, 0.0, 0, errors.New("es must be greater than 0")
    }
    a := xl
    b := xu
    fa := f(a)
    fb := f(b)
    c := a
    fc := fa
    d := b - c
    e := d
    iter = 0
    var m, tol, s, p, q, r float64
    ea = 100
    for ; iter < maxit; iter ++ {
        if fb == 0.0 {
            break;
        }
        if math.Signbit(fa) == math.Signbit(fb) { // if needed rearrange points
            a = c
            fa = fc
            d = b - c
            e = d
        }
        if math.Abs(fa) < math.Abs(fb) {
            c = b
            b = a
            a = c
            fc = fb
            fb = fa
            fa = fc
        }
        m = 0.5*(a - b) // Termination test and possible exit
        tol = 2.0*es*math.Max(math.Abs(b), 1.0)
        if math.Abs(m) <= tol || fb == 0.0 {
            break;
        }
        // Choose open methods or bisection
        if math.Abs(e) >= tol && math.Abs(fc) > math.Abs(fb) {
            s = fb/fc
            if a == c { // Secant method
                p = 2.0*m*s
                q = 1.0 - s
            } else { // Inverse quadractic interpolation
                q = fc/fa
                r = fb/fa
                p = s*(2.0*m*q*(q - r) - (b - c)*(r - 1.0))
                q = (q - 1.0)*(r - 1.0)*(s - 1.0)
            }
            if p > 0.0 {
                q = -q
            } else {
                p = -p
            }
            if 2.0*p < 3.0*m*q - math.Abs(tol*q) && p < math.Abs(0.5*e*q) {
                e = d
                d = p/q
            } else {
                d = m
                e = m
            }
        } else { // Bisection
            d = m
            e = m
        }
        c = b
        fc = fb
        if math.Abs(d) > tol {
            b += d
        } else {
            if b-a >= 0.0 {
                b -= tol
            } else {
                b += tol
            }
        }
        fb = f(b)
    }
    root = b
    fx = fb
    ea = math.Abs(m)
    return root, fx, ea, iter, nil
}
