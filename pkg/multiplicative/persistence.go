package multiplicative

import (
	"errors"
	"log"
	"strconv"
)

// MaxUint64 maximum value of uint64
const MaxUint64 = ^uint64(0)

var (
	// ErrorNotInRange0To9 error when rune is not in the equivalent 0 to 9
	ErrorNotInRange0To9 = errors.New("rune is not in the range of 0 to 9")
)

// Persistence type of the Persistence function
type Persistence func(number uint64) (uint64, error)

// NaivePersistence takes an integer and converts
// it into string digits and keeps on multiplying them
// until the result is of 1 digit.
func NaivePersistence(number uint64) (uint64, error) {
	if number < 10 {
		return 0, nil
	}
	var result uint64 = 1
	for _, digit := range strconv.FormatUint(number, 10) {
		value, err := aToi(digit)
		if err != nil {
			return 0, err
		}
		result *= value
	}
	per, err := NaivePersistence(result)
	return 1 + per, err
}

func aToi(n rune) (uint64, error) {
	switch n {
	case '0':
		return 0, nil
	case '1':
		return 1, nil
	case '2':
		return 2, nil
	case '3':
		return 3, nil
	case '4':
		return 4, nil
	case '5':
		return 5, nil
	case '6':
		return 6, nil
	case '7':
		return 7, nil
	case '8':
		return 8, nil
	case '9':
		return 9, nil
	default:
		return 0, ErrorNotInRange0To9
	}
}

// MinMaxHighestPersistNumbers finds the min & max number
// with the highest persistence for uint64
func MinMaxHighestPersistNumbers(start uint64, end uint64, f Persistence) (uint64, uint64, uint64) {
	min, max, per := uint64(0), uint64(0), uint64(0)
	for i := start; i <= end; i++ {
		v, err := f(i)
		if err != nil {
			return min, max, per
		}
		log.Printf("F(%d)=%d\n", i, v)
		if v > per {
			per = v
			min = i
		}
		if v == per && i > max {
			max = i
		}
	}
	log.Printf("min=%d, max=%d, per=%d\n", min, max, per)
	return min, max, per
}
