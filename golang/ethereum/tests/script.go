package main 

import(
	"fmt"
	"math/big"
)

func main() {
	amount := new(big.Int)
	amount.SetString("1e2", 10)
	fmt.Println(amount)
}