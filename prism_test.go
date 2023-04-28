package prism

import (
	"fmt"
	fuzzyPid "github.com/ivan97/fuzzy-pid"
	"testing"
)

func TestPrism(t *testing.T) {
	pid := fuzzyPid.NewFuzzyPid(1200, 650, 0.3, 1.0, 0.6, 0.01, 0.02, 0.01).SetDefaultConfig()
	target := 500.0
	actual := 0.0

	for i := 0; i < 20; i++ {
		signal := pid.Realize(target, actual)
		actual += signal
		fmt.Printf("===============signal: %f actual: %f======================\n", signal, actual)
	}
}
