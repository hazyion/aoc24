package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
  data, err := os.ReadFile("inputs/d1.txt")
  if err != nil{
    fmt.Printf("Error reading file")
    return
  }

  var lnums, rnums []int
  pairs := strings.Split(string(data), "\n")
  for _, val := range pairs{
    nums := strings.Fields(val)
    if len(nums) == 0{
      break
    }
    lnum, err := strconv.Atoi(nums[0])
    if err != nil {
      fmt.Println("Error:", err)
      return
    }
    lnums = append(lnums, lnum)
    rnum, err := strconv.Atoi(nums[1])
    if err != nil {
      fmt.Println("Error:", err)
      return
    }
    rnums = append(rnums, rnum)
  }

  // Part a
  sort.Ints(lnums)
  sort.Ints(rnums)

  diff := 0
  for ind := range len(lnums){
    x := lnums[ind] - rnums[ind] 
    if x < 0{
      x *= -1
    }
    diff += x
  }

  fmt.Println(diff)

  // Part b
	myMap := make(map[string]int)
  for _, val := range rnums{
    strval := strconv.Itoa(val)
    f, ex := myMap[strval]
    if !ex{
      myMap[strval] = 1
    } else{
      myMap[strval] = f + 1
    }
  }

  total := 0
  for _, val := range lnums{
    strval := strconv.Itoa(val)
    f, ex := myMap[strval]
    if ex{
      total += f * val
    }
  }
  fmt.Println(total)
}
