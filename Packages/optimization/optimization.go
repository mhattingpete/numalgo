package optimization

import (
    "errors"
    "math"
)

// Goldmin (Golden-Section search) 
// input:
// the function to find the minimum for (f), lower limit (xl), upper limit (xu), error deviation (es), maximum iterations (iter)
// output: 
// the estimated x (x), function value (fx), error estimate (ea), iterations done (iter)
func Goldmin(f func(float64) float64, xl float64, xu float64, es float64, maxit int) (x float64, fx float64, ea float64, iter int, err error) {
    if es < 0.0 {
        return 0.0, 0.0, 0.0, 0, errors.New("es must be greater than 0")
    }
    phi := (1.0+math.Sqrt(5.0))/2.0
    var d, x1, x2 float64
    iter = 0
    for ; iter < maxit; iter ++ {
        d = (phi - 1.0)*(xu - xl)
        x1 = xl + d
        x2 = xu - d
        if f(x1) < f(x2) {
            x = x1
            xl = x2
        } else {
            x = x2
            xu = x1
        }
        if x != 0 {
            ea = (2.0 - phi)*math.Abs((xu - xl)/x)*100.0
        }
        if ea <= es {
            break;
        }
    }
    fx = f(x)
    return x, fx, ea, iter, nil
}

// Parabolic (Parabolic-Interpolation search) 
// input:
// the function to find the minimum for (f), lower limit (xl), intermediate point (xm), upper limit (xu), error deviation (es), maximum iterations (iter)
// output: 
// the estimated x (x), function value (fx), error estimate (ea), iterations done (iter)
func Parabolic(f func(float64) float64, xl float64, xm float64, xu float64, es float64, maxit int) (x float64, fx float64, ea float64, iter int, err error) {
    if es < 0.0 {
        return 0.0, 0.0, 0.0, 0, errors.New("es must be greater than 0")
    }
    if xm < xl || xm > xu {
        return 0.0, 0.0, 0.0, 0, errors.New("The following condition is not met: xl < xm < xu")
    }
    var x1, x2, x3, x4, f1, f2, f3, f4, xold float64
    x1, x2, x3 = xl, xm, xu
    f1, f2, f3 = f(x1), f(x2), f(x3)
    if f2 > f1 || f2 > f3 {
        return 0.0, 0.0, 0.0, 0, errors.New("The following condition is not met: f(xl) > f(xm) < f(xu)")
    }
    iter = 0
    for ; iter < maxit; iter ++ {
        xold = x4
        x4 = x2 - 0.5*(math.Pow(x2-x1,2.0)*(f2-f3)-math.Pow(x2-x3,2)*(f2-f1))/((x2-x1)*(f2-f3)-(x2-x3)*(f2-f1))
        f4 = f(x4)
        if f4 < f2 {
            if x4 < x2 { // swap x2 with x4 and x3 with x2, no need to update x4 as it is overwritten
                x3 = x2
                f3 = f2
                x2 = x4
                f2 = f4
            } else { // swap x1 with x2 and x2 with x4, no need to update x4 as it is overwritten
                x1 = x2
                f1 = f2
                x2 = x4
                f2 = f4
            }
        } else {
            x4 = x2
            f4 = f2
            break;
        }
        if x4 != 0 {
            ea = math.Abs((x4 - xold) / x4) * 100.0
        }
        if ea <= es {
            break;
        }
    }
    x = x4 
    fx = f(x)
    return x, fx, ea, iter, nil
}