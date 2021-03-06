package intest

import (
	"fmt"
	"io"
	"testing"
)

func TestExchangeThese(t *testing.T) {
	ExchangeThese()
}

func TestRead(t *testing.T) {
	const size = 18
	robert := &StringPair{"Robert L.", "Stevenson"}
	david := StringPair{"Davide ", "Zhang"}
	//n, err := david.Read(ToBytes(david))
	fmt.Printf("Stevenson length first, second: %d, %d\n", len(robert.first), len(robert.second))

	for _, reader := range []io.Reader{robert, &david} {
		raw, err := ToBytes(reader, size)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Raw %q\n", raw)
	}
}
func TestDataBlock(t *testing.T) {
	DataBlock()
}
