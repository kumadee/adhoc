package main

import "github.com/kumadee/adhoc/pkg/multiplicative"

func main() {
	multiplicative.MinMaxHighestPersistNumbers(0, multiplicative.MaxUint64, multiplicative.NaivePersistence)
}
