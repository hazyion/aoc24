package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

func checker(ilevels []int) int {
  diff := ilevels[0] - ilevels[1]
  if diff < -3 || diff == 0 || diff > 3{
    return 0
  }
  ouch := -1
  for i := 1; i < len(ilevels) - 1; i++{
    curDiff := ilevels[i] - ilevels[i+1]
    if curDiff * diff < 0{
      ouch = i
      break
    }
    if curDiff <  -3 || curDiff == 0 || curDiff > 3{
      ouch = i
      break
    }
  }
  return ouch
}

func main() {
  data, err := os.ReadFile("inputs/d2.txt")
  if err != nil{
    fmt.Printf("Error reading file")
    return
  }
  
  recs := strings.Split(string(data), "\n")
  total := 0
  totalB := 0
  for _, val := range recs{
    levels := strings.Fields(val)
    var ilevels []int
    if len(levels) == 0{
      break
    }
    for _, num := range levels{
      inum, err := strconv.Atoi(num)
      if err != nil{
	fmt.Printf("Error")
      }
      ilevels = append(ilevels, inum)
    }

    // Part 1
    ouch := checker(ilevels)
    if ouch == -1{
      total += 1
      totalB += 1
    } else{
      // Part 2
      for _, ind := range []int{ouch-1, ouch, ouch+1}{
	tempL := make([]int, len(ilevels))
	copy(tempL, ilevels)
	if ind < 0 || ind >= len(ilevels){
	  continue
	}
	newList := append(tempL[:ind], tempL[ind+1:]...)
	res := checker(newList)
	if res == -1{
	  totalB += 1
	  break
	}
      }
    }
  }
  fmt.Println(total)
  fmt.Println(totalB)
}
