package main

import "github.com/kumadee/adhoc/pkg/multipersist"

func main() {
	multipersist.MinMaxHighestPersistNumbers(0, multipersist.MaxUint64, multipersist.NaiveMultiPersitence)
}
