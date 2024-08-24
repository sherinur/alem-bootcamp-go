package bootcamp

func findLongestChain(cities []string, currentChain []string, used map[string]bool) []string {
	if len(currentChain) == 0 {
		longestChain := []string{}
		for _, city := range cities {
			chain := findLongestChain(cities, []string{city}, map[string]bool{city: true})
			if len(chain) > len(longestChain) {
				longestChain = chain
			}
		}
		return longestChain
	} else {
		lastCity := currentChain[len(currentChain)-1]
		lastChar := lastCity[len(lastCity)-1]
		longestChain := make([]string, len(currentChain))
		copy(longestChain, currentChain)
		for _, city := range cities {
			if !used[city] && (city[0] == lastChar || city[0]+32 == lastChar || city[0] == lastChar+32) {
				used[city] = true
				chain := findLongestChain(cities, append(currentChain, city), used)
				if len(chain) > len(longestChain) {
					longestChain = chain
				}
				used[city] = false
			}
		}
		return longestChain
	}
}

func GameCities(cities []string) []string {
	return findLongestChain(cities, []string{}, map[string]bool{})
}
