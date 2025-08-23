package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	fuzzFtoaRandomBits(100000, 64)
	fuzzFtoaRandomDigits(100000, 64)
	fuzzFtoaRandomBits(100000, 32)
	fuzzFtoaRandomDigits(100000, 32)
	fmt.Println("success")
}

func fuzzFtoaRandomBits(samples, bitSize int) {
	for range samples {
		var val float64
		switch bitSize {
		case 32:
			val = float64(math.Float32frombits(rand.Uint32()))
		case 64:
			val = math.Float64frombits(rand.Uint64())
		}

		output1, _ := strconv.RunDragonboxFtoa(val, bitSize)
		output2, _ := strconv.RunRyuFtoaShortest(val, bitSize)

		if output1 != output2 {
			panic(fmt.Sprintf("Mismatch:\nInput: %f\nDragonbox output: %s\nRyu output: %s", val, output1, output2))
		}
	}
}

func fuzzFtoaRandomDigits(samples, bitSize int) {
	for n := 1; n <= 17; n++ {
		for range samples {
			var val float64
			switch bitSize {
			case 32:
				val = float64(math.Float32frombits(rand.Uint32()))
			case 64:
				val = math.Float64frombits(rand.Uint64())
			}

			// truncate to n digits
			s := strconv.FormatFloat(val, 'e', n, bitSize)
			val, _ = strconv.ParseFloat(s, bitSize)

			output1, _ := strconv.RunDragonboxFtoa(val, bitSize)
			output2, _ := strconv.RunRyuFtoaShortest(val, bitSize)

			if output1 != output2 {
				panic(fmt.Sprintf("Mismatch:\nInput: %f\nDragonbox output: %s\nRyu output: %s", val, output1, output2))
			}
		}
	}
}
