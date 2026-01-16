// LeetCode L134

package main

func main() {
	// fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	// fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
	// fmt.Println(canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}))
	// fmt.Println(canCompleteCircuit([]int{5, 8, 2, 8}, []int{6, 5, 6, 6}))
	// fmt.Println(canCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2}))
}

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	total := 0
	current := 0

	for i := range gas {
		total += gas[i] - cost[i]
		current += gas[i] - cost[i]
		if current < 0 {
			current = 0
			start++
		}
	}

	if total < 0 {
		return -1
	} else {
		return start
	}
}

func canCompleteCircuitWhat(gas []int, cost []int) int {
	i := 0
	for i < len(gas) {
		if gas[i] > 0 && gas[i] >= cost[i] {
			j := i + 1
			if j >= len(cost) {
				j = 0
			}
			for j != i {
				if gas[j] >= cost[j] {
					j++
					if j >= len(cost) {
						j = 0
					}
				} else {
					break
				}
			}
			if j == i {
				return i
			}

			// fmt.Println("REached hear!")

			j = i + 1
			gasLeft := gas[i] - cost[i]
			reachable := true
			if j >= len(cost) {
				j = 0
			}
			for j != i {

				gasLeft = gasLeft + gas[j] - cost[j]

				if gasLeft >= 0 {
					j++
					if j >= len(cost) {
						j = 0
					}
				} else {
					reachable = false
					break
				}
				// fmt.Println(i, j)
			}
			// fmt.Println(i)
			if reachable {
				return i
			}
		}
		i++
	}
	return -1
}
func canCompleteCircuitFailed(gas []int, cost []int) int {
	// find min cost with max gas
	minCostIdx := 0
	maxGasInMinCosts := gas[0]
	for i := 1; i < len(cost); i++ {
		if cost[i] < cost[minCostIdx] {
			minCostIdx = i
			maxGasInMinCosts = gas[i]
		} else if cost[i] == cost[minCostIdx] && gas[i] > maxGasInMinCosts {
			minCostIdx = i
			maxGasInMinCosts = gas[i]
		}
	}
	// fmt.Println(minCostIdx) ///

	i := minCostIdx
	gasLeft := gas[i] - cost[i]
	if gasLeft >= 0 {
		i++
		for i != minCostIdx {
			// simple circular
			if i >= len(cost) {
				i = 0
			}

			gasLeft = gasLeft + gas[i] - cost[i]
			if gasLeft >= 0 {
				i++
			} else {
				break
			}
		}
	} else {
		return -1
	}

	if i == minCostIdx {
		return minCostIdx
	} else {
		return -1
	}

}
