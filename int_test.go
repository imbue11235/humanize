package humanize

import (
	"fmt"
	"math/big"
	"testing"
)

func TestInt(t *testing.T) {
	fmt.Println(Int(1250000000))
	fmt.Println(BigInt(big.NewInt(1250000000)))
}
