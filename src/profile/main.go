package main

import (
	"math"
	"math/rand/v2"
	"os"
	"runtime/pprof"
	"strconv"
)

func main() {
	profileFtoaRandomBits(1000000, 64, "random_bits64.pprof")
	profileFtoaRandomDigits(1000000, 64, "random_digits64.pprof")
	profileFtoaRandomBits(1000000, 32, "random_bits32.pprof")
	profileFtoaRandomDigits(1000000, 32, "random_digits32.pprof")
}

type testdata struct {
	val     float64
	bitSize int
}

func profileFtoaRandomBits(samples, bitSize int, profFile string) {
	var tests []testdata
	for range samples {
		var val float64
		switch bitSize {
		case 32:
			val = float64(math.Float32frombits(rand.Uint32()))
		case 64:
			val = math.Float64frombits(rand.Uint64())
		}

		tests = append(tests, testdata{val, bitSize})
	}

	f, err := os.Create(profFile)
	if err != nil {
		panic(err)
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for _, tt := range tests {
		strconv.ProfileDragonboxFtoa(tt.val, tt.bitSize)
		strconv.ProfileRyuFtoaShortest(tt.val, tt.bitSize)
	}
}

func profileFtoaRandomDigits(samples, bitSize int, profFile string) {
	var tests []testdata
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

			tests = append(tests, testdata{val, bitSize})
		}
	}

	f, err := os.Create(profFile)
	if err != nil {
		panic(err)
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for _, tt := range tests {
		if !(tt.bitSize == 32 || tt.bitSize == 64) {
			panic("bitsize is illegal" + strconv.FormatInt(int64(tt.bitSize), 10))
		}
		strconv.ProfileDragonboxFtoa(tt.val, tt.bitSize)
		strconv.ProfileRyuFtoaShortest(tt.val, tt.bitSize)
	}
}
