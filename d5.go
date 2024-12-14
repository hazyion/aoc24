package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func update_check(prev_key *map[int][]int, update *[]int) (bool, []int) {
	// Logic
	// Check if update[i] | update[i+1] exists (enough since transitive condition ensures update[j] | update[i+1] for all j < i+1)
	// If not, check if update[i+1] | update[j] exists for all j < i+1 (check for contradiction)
	upd := *update
	fine := true
	// Part 1
	for i := 0; i < len(upd)-1; i++ {
		arr, ok := (*prev_key)[upd[i]]
		if ok {
			if slices.Contains(arr, upd[i+1]) {
				continue
			} else {
				arr1, ok1 := (*prev_key)[upd[i+1]]
				if !ok1 {
					continue
				}
				for j := i; j >= 0; j-- {
					if slices.Contains(arr1, upd[j]) {
						fine = false
						break
					}
				}
				if !fine {
					break
				}
			}
		} else {
			arr, ok := (*prev_key)[upd[i+1]]
			if !ok {
				continue
			}
			for j := i; j >= 0; j-- {
				if slices.Contains(arr, upd[j]) {
					fine = false
					break
				}
			}
			if !fine {
				break
			}
		}
	}
	if fine {
		return true, []int{}
	} else {
		// Part 2
		for i := 0; i < len(upd)-1; i++ {
			arr, ok := (*prev_key)[upd[i]]
			if ok {
				if slices.Contains(arr, upd[i+1]) {
					continue
				} else {
					arr1, ok1 := (*prev_key)[upd[i+1]]
					if ok1 {
						for j := 0; j < i+1; j++ {
							if slices.Contains(arr1, upd[j]) {
								last := upd[i+1]
								for k := i + 1; k > j; k-- {
									upd[k] = upd[k-1]
								}
								upd[j] = last
								break
							}
						}
					}
				}
			} else {
				arr1, ok1 := (*prev_key)[upd[i+1]]
				if ok1 {
					for j := 0; j < i+1; j++ {
						if slices.Contains(arr1, upd[j]) {
							last := upd[i+1]
							for k := i + 1; k > j; k-- {
								upd[k] = upd[k-1]
							}
							upd[j] = last
							break
						}
					}
				}

			}
		}
		return false, upd
	}
}

func main() {
	data, err := os.ReadFile("inputs/d5.txt")
	if err != nil {
		fmt.Printf("Error reading file")
		return
	}

	inp := strings.Fields(string(data))
	prev_key := make(map[int][]int)

	total, total_new := 0, 0
	for _, i := range inp {
		if strings.Contains(i, "|") {
			data := strings.Split(i, "|")
			fnum, err := strconv.Atoi(data[0])
			if err != nil {
				fmt.Println("error converting to int")
			}
			snum, err := strconv.Atoi(data[1])
			if err != nil {
				fmt.Println("error converting to int")
			}
			val, ok := prev_key[fnum]
			if !ok {
				prev_key[fnum] = []int{snum}
			} else {
				prev_key[fnum] = append(val, snum)
			}
		} else {
			data := strings.Split(i, ",")
			var update []int
			for _, val := range data {
				num, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println("error converting to int")
				}
				update = append(update, num)
			}
			yea, arr := update_check(&prev_key, &update)
			if yea {
				total += update[len(update)/2]
			} else {
				total_new += arr[len(arr)/2]
			}
		}
	}

	fmt.Println(total)
	fmt.Println(total_new)
}
