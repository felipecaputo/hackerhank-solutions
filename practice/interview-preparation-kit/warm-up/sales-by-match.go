package warmup

import (
	"math"
)

/*
 * Complete the 'sockMerchant' function below.
 *
 * There is a large pile of socks that must be paired by color. Given an array
 * of integers representing the color of each sock, determine how many pairs
 * of socks with matching colors there are.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
	// create a map to put socks of the same color together
	socksPile := make(map[int32]float64)

	for _, color := range ar {
		v, ok := socksPile[color] // if there are socks of the same color, get the count
		if !ok {                  // if there aren't, start the count at 0
			v = 0
		}

		v++                  // increment the number of socks
		socksPile[color] = v // and update the pule count
	}

	numOfPairs := int32(0) // starts the all colors sock's pairs count
	for _, v := range socksPile {
		numOfPairs += int32(math.Trunc(v / 2)) // for each color get the number of pairs by dividing by 2 and truncating decimals
	}

	return numOfPairs
}
