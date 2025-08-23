package main

import (
	"encoding/csv"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	benchFtoaRandomBits(1000000, 1000, 64, "random_bits64.csv")
	benchFtoaRandomDigits(100000, 1000, 64, "random_digits64.csv")
	benchFtoaRandomBits(1000000, 1000, 32, "random_bits32.csv")
	benchFtoaRandomDigits(100000, 1000, 32, "random_digits32.csv")
}

func benchFtoaRandomBits(samples, iter, bitSize int, csvFile string) error {
	f, err := os.Create(csvFile)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{"bits", "dragonbox_ns", "ryu_ns"})

	for range samples {
		var bits uint64
		var val float64
		switch bitSize {
		case 32:
			bits = uint64(rand.Uint32())
			val = float64(math.Float32frombits(uint32(bits)))
		case 64:
			bits = rand.Uint64()
			val = math.Float64frombits(bits)
		}

		var total1, total2 time.Duration
		for range iter {
			_, elapsed1 := strconv.RunDragonboxFtoa(val, bitSize)
			_, elapsed2 := strconv.RunRyuFtoaShortest(val, bitSize)
			total1 += elapsed1
			total2 += elapsed2
		}
		mean1 := float64(total1.Nanoseconds()) / float64(iter)
		mean2 := float64(total2.Nanoseconds()) / float64(iter)

		w.Write([]string{
			strconv.FormatUint(bits, 10),
			strconv.FormatFloat(mean1, 'e', -1, 64),
			strconv.FormatFloat(mean2, 'e', -1, 64),
		})
	}

	return nil
}

func benchFtoaRandomDigits(samples, iter int, bitSize int, csvFile string) error {
	f, err := os.Create(csvFile)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{"digits", "dragonbox_ns", "ryu_ns"})

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

			var total1, total2 time.Duration
			for range iter {
				_, elapsed1 := strconv.RunDragonboxFtoa(val, bitSize)
				_, elapsed2 := strconv.RunRyuFtoaShortest(val, bitSize)
				total1 += elapsed1
				total2 += elapsed2
			}
			mean1 := float64(total1.Nanoseconds()) / float64(iter)
			mean2 := float64(total2.Nanoseconds()) / float64(iter)

			w.Write([]string{
				strconv.FormatInt(int64(n), 10),
				strconv.FormatFloat(mean1, 'e', -1, 64),
				strconv.FormatFloat(mean2, 'e', -1, 64),
			})
		}
	}

	return nil
}
