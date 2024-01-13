package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 4 {
		fmt.Println("Not enough arguments")
		return
	}

	q, err1 := strconv.Atoi(args[0])
	m, err2 := strconv.Atoi(args[1])
	n, err3 := strconv.Atoi(args[2])
	p, err4 := strconv.ParseFloat(args[3], 64)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Println("Error parsing arguments")
		return
	}

	switch q {
	case 1:
		theoretical, support := theoreticalDistribution(n, p)
		fmt.Println("Support:", sliceToString(support))
		fmt.Println("Theoretical:", sliceToString(theoretical))
	case 2:
		outcomes := inverseCDFBinomial(m, n, p)
		fmt.Println(sliceToString(outcomes))
	case 3:
		bernOutcomes := bernBinomial(m, n, p)
		fmt.Println(sliceToString(bernOutcomes))
	default:
		fmt.Println("Invalid option")
	}
}

func sliceToString(slice []float64) string {
	strSlice := make([]string, len(slice))
	for i, v := range slice {
		if v == math.Trunc(v) {
			strSlice[i] = fmt.Sprintf("%.0f", v)
		} else {
			strSlice[i] = fmt.Sprintf("%g", v)
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(strSlice, ", "))
}

func theoreticalDistribution(n int, p float64) ([]float64, []float64) {
	support := make([]float64, n+1)
	theoretical := make([]float64, n+1)
	binom := distuv.Binomial{
		N: float64(n),
		P: p,
	}

	for i := range support {
		support[i] = float64(i)
		theoretical[i] = binom.Prob(float64(i))
	}

	return theoretical, support
}

func inverseCDFBinomial(m, n int, p float64) []float64 {
	outcomes := make([]float64, n+1)
	binomialDist := distuv.Binomial{
		N: float64(n),
		P: p,
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < m; i++ {
		u := rand.Float64()
		cdf := 0.0
		for j := 0; j <= n; j++ {
			cdf += binomialDist.Prob(float64(j))
			if u <= cdf {
				outcomes[j]++
				break
			}
		}
	}

	return outcomes
}

func bernBinomial(m, n int, p float64) []float64 {
	outcomes := make([]float64, n+1)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < m; i++ {
		successCount := 0
		for j := 0; j < n; j++ {
			if rand.Float64() <= p {
				successCount++
			}
		}
		outcomes[successCount]++
	}

	return outcomes
}
