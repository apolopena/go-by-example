package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func benchmark(start time.Time, name string) string {
	elapsed := time.Since(start)
	return fmt.Sprintf("%s took %s", name, elapsed)
}

type Primes struct {
	Prime     []int
	NotPrime  []int
	Benchmark string
}

func (rp Primes) data() []int {
	return append(rp.Prime, rp.NotPrime...)
}

func (rp Primes) len() int {
	return len(rp.data())
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func isPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func isPrime(val int) bool {
	if !(isPrimeSqrt((val))) {
		return false
	}
	return true
}

func primes(min int, max int, total int, verbose bool) (Primes, error) {
	start := time.Now()
	report := Primes{}

	if min >= max {
		return report, fmt.Errorf("min: %d out of range for max: %d", min, max)
	}

	for i := min; i <= total; i++ {
		if isPrime(i) {
			if verbose {
				fmt.Printf("%d is prime\n", i)
			}
			report.Prime = append(report.Prime, i)
		} else {
			if verbose {
				fmt.Printf("%d is not prime\n", i)
			}
			report.NotPrime = append(report.NotPrime, i)
		}
	}

	report.Benchmark = fmt.Sprintf("%s", time.Since(start))
	return report, nil
}

func randPrimes(min int, max int, total int, verbose bool) (Primes, error) {
	start := time.Now()
	report := Primes{}

	if min >= max {
		return report, fmt.Errorf("min: %d out of range for max: %d", min, max)
	}

	for i := min; i <= total; i++ {
		num := randInt(min, max)
		if isPrime(num) {
			if verbose {
				fmt.Printf("%d is prime\n", num)
			}
			report.Prime = append(report.Prime, num)
		} else {
			if verbose {
				fmt.Printf("%d is not prime\n", num)
			}
			report.NotPrime = append(report.NotPrime, num)
		}
	}

	report.Benchmark = fmt.Sprintf("%s", time.Since(start))
	return report, nil
}

func main() {
	min, max, total := 1, 10000, 100000
	if r, e := primes(min, max, total, false); e != nil {
		fmt.Println("Error: ", e)
	} else {
		fmt.Println("---------primes Report---------")
		//fmt.Printf("Data: %v\n", r.data())
		fmt.Printf("Iterations: %v\n", r.len())
		fmt.Println("Total prime numbers found: ", len(r.Prime))
		fmt.Println("Total unprime numbers found: ", len(r.NotPrime))
		fmt.Println("Execution time: ", r.Benchmark)
	}

	if r, e := randPrimes(min, max, total, false); e != nil {
		fmt.Println("Error: ", e)
	} else {
		fmt.Println("---------randPrimes Report---------")
		//fmt.Printf("Data: %v\n", r.data())
		fmt.Printf("Iterations: %v\n", r.len())
		fmt.Println("Total prime numbers found: ", len(r.Prime))
		fmt.Println("Total unprime numbers found: ", len(r.NotPrime))
		fmt.Println("Execution time: ", r.Benchmark)
	}
}
