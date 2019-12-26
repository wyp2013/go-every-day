package main

import (
	"fmt"
	"math"
)

// hashrate*10^15*600/(difficulty*2^32)
func getLucky(hashrate float64, difficulty float64) float64 {
	hash := hashrate * math.Pow(10, 15)*600
	diff := difficulty * math.Pow(2, 32)
	lucky := hash/diff
	fmt.Println(lucky)
	return lucky*24*6*3
}

/*
 $latest_block = $this->BTCapi->getLatestBlock();
        $btc_number_in_one_block = 1 / pow(2, floor($latest_block['height'] / 210000)) * 50;
        $income = pow(10, 12) * 86400 / (pow(2, 32) * $latest_block['difficulty']) * $btc_number_in_one_block;
        $this->logger->error("getHashrateIncome4BTC ", $latest_block, [$btc_number_in_one_block, $income]);
        return $income;
 */

func getCoinIncome(height int, difficulty float64) {
	btc_number_in_one_block := 50.0 / float64(math.Pow(2, math.Floor(float64(height/210000.0))))
	income := math.Pow(10, 12) * 86400 / (math.Pow(2, 32) *difficulty) * btc_number_in_one_block
	fmt.Println(income)
}

func main() {
	// fmt.Println(55/getLucky(12530, 12973235968799.777))
	// fmt.Println(62/getLucky(17709, 12973235968799.777))

	getCoinIncome(607618, 12876842089682.479)
}


