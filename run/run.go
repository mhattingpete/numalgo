package main

import (
	"fmt"
	"example.com/rootmethods"
)

func main() {
	start := 0.0
    stop := 1.0
    numsteps := 10
    msg, err := rootmethods.Linspace(start, stop, numsteps)
    fmt.Println(msg,err)
    fmt.Println("Bisection")
    // bisection
    xl := -10.0
    xu := 10.0
    es := 0.0001
    maxit := 50
    f := func(x float64) float64 {
        return 2.0*x - 3.0
    }
    root, fx, ea, iter, err := rootmethods.Bisection(f, xl, xu, es, maxit)
    fmt.Println(root, fx, ea, iter, err)
}

/*
2x - 3 = 0
2x = 3
x = 1.5
*/