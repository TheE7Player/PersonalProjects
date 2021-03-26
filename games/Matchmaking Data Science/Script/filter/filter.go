package filter

import (
	"fmt"
	"strconv"

	f "CSGO_STATISTICS/sort"
)

var table map[string][]string

func Run(data map[string][]string) {
	fmt.Println("Running statistics...")
	table = data
	fmt.Println("Map Frequency", get_map_freq())
	fmt.Println("Result Count", get_result_count())
	fmt.Println("Waiting Time per Hour", get_avg_time_by_hour())
	fmt.Println("Most Streaks of X", get_streaks())
}

// greatest common divisor (GCD) via Euclidean algorithm (Go Playground Example)
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func get_map_freq() map[string]float32 {
	maps := table["Map"]

	result := make(map[string]float32)

	for _, level := range maps {
		if len(level) > 2 {
			result[level]++
		}
	}

	totalMaps := len(maps)
	// Then we get the percentage
	for k, v := range result {
		result[k] = (v / float32(totalMaps)) * 100
	}

	// Sort the map from value
	result = f.SortByFloat(result)

	return result
}

func get_result_count() map[string]string {
	results := table["Result"]
	returnVal := make((map[string]string))

	win := 0
	loss := 0
	tie := 0

	for _, k := range results {
		if len(k) > 2 {
			switch k {
			case "Won":
				win++
			case "Tie":
				tie++
			case "Lost":
				loss++
			}
		}
	}

	returnVal["Won"] = strconv.Itoa(win)
	returnVal["Tie"] = strconv.Itoa(tie)
	returnVal["Loss"] = strconv.Itoa(loss)

	gcd_div := gcd(win, loss)

	leftRatio := win / gcd_div
	rightRatio := loss / gcd_div

	returnVal["W/L Ratio"] = fmt.Sprintf("%d:%d", leftRatio, rightRatio)
	return returnVal
}

func get_avg_time_by_hour() map[string]string {
	// Time, Waiting Time

	hour := table["Time"]
	wait := table["Waiting Time"]
	size := len(hour)

	appendMap := make(map[string][]int64, size)
	returnBack := make(map[string]string, size)

	for i := 0; i < size; i++ {

		time := string([]rune(hour[i])[:2])
		waitTime := []rune(wait[i])

		if len(waitTime) != 5 {
			continue
		}

		minutes := string(waitTime[:2])
		seconds := string(waitTime[3:])

		min, errMin := strconv.ParseInt(minutes, 10, 64)
		total, errTotal := strconv.ParseInt(seconds, 10, 64)

		if errMin != nil {
			fmt.Println("Error, please look at 'errMin'")
		}

		if errTotal != nil {
			fmt.Println("Error, please look at 'errTotal'")
		}

		for min > 0 {
			min--
			total += 60
		}

		appendMap[time] = append(appendMap[time], total)
	}

	// Now average each hour
	avg := int64(0)
	avgT := 0
	for k, v := range appendMap {
		for _, amt := range v {
			avg += amt
		}

		// fmt.Sprintf("%d", (int(avg) / len(appendMap[k])))
		avgT = int(avg) / len(appendMap[k])

		returnBack[k] = fmt.Sprintf("%d", avgT)

		avg = 0
		avgT = 0
	}

	return returnBack
}

func get_streaks() map[string]int {
	results := table["Result"]
	returnBack := make(map[string]int, 3)

	winStreak := 0
	lossStreak := 0
	tieStreak := 0

	lastResult := results[0]
	currentResult := ""
	size := len(results)

	switch lastResult {
	case "Won":
		winStreak++
	case "Lost":
		lossStreak++
	case "Tie":
		tieStreak++
	}

	streakCounter := 1
	for i := 1; i < size; i++ {
		currentResult = results[i]
		lastResult = results[i-1]

		if currentResult != lastResult {
			switch lastResult {
			case "Won":
				if streakCounter > winStreak {
					winStreak = streakCounter
				}
			case "Lost":
				if streakCounter > lossStreak {
					lossStreak = streakCounter
				}
			case "Tie":
				if streakCounter > tieStreak {
					tieStreak = streakCounter
				}
			}

			streakCounter = 1

			continue
		}
		streakCounter++
	}

	returnBack["Wins"] = winStreak
	returnBack["Ties"] = tieStreak
	returnBack["Losses"] = lossStreak
	return returnBack
}
