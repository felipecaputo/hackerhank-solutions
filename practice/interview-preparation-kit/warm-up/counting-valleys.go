package warmup

import (
	"strings"
)

/*
 * Complete the 'countingValleys' function below.
 *
 * Given the sequence of up and down steps during a hike, find and print the number of valleys walked through.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	currentAltitude := 0 //sea level
	isInValley := false
	valleyCount := int32(0)

	for _, s := range strings.Split(path, "") {
		// First we need to calculate current altitude
		switch s {
		case "U":
			currentAltitude += 1
		case "D":
			currentAltitude -= 1
		}

		// if we are in a valley and we reached sea level, we are not in a valley anymore
		if isInValley && currentAltitude >= 0 {
			isInValley = false
		}

		// if we reached an altitude bellow sea level and we are not yet in a valley, now we are
		if currentAltitude < 0 && !isInValley {
			isInValley = true
			valleyCount++
		}
	}

	return valleyCount
}
