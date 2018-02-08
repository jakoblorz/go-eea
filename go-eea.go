package main

import (
	"fmt"

	"github.com/caarlos0/env"
	eea "github.com/jakoblorz/go-eea/lib"
)

func main() {

	cfg := eea.ExtendedEuclidianParameters{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	for i, r := range cfg.Calculate() {
		fmt.Printf("%d: A=%d B=%d Q=%d S=%d T=%d\n", i, r.A, r.B, r.Q, r.S, r.T)
	}
}
