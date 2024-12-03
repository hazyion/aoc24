package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func multer(mult string) int {
  cindex := strings.Index(mult, `,`)
  oindex := strings.Index(mult, `(`)
  pindex := strings.Index(mult, `)`)
  fnum, err := strconv.Atoi(mult[oindex+1:cindex])
  if err != nil{
    fmt.Println("error extracting number")
  }
  snum, err := strconv.Atoi(mult[cindex+1:pindex])
  if err != nil{
    fmt.Println("error extracting number")
  }
  return fnum * snum
}

func main(){
  data, err := os.ReadFile("inputs/d3.txt")
  if err != nil{
    fmt.Printf("Error reading file")
    return
  }
  str := string(data)

  // Part 1
  regex, err := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
  if err != nil{
    fmt.Println("error regex")
  }
  res := regex.FindAllString(str, -1)
  total := 0
  for _, mult := range res{
    total += multer(mult)
  }
  fmt.Println(total)

  // Part 2
  newRegex := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
  res = newRegex.FindAllString(str, -1)
  valid := true
  total = 0
  for _, mult := range res{
    if mult == "don't()"{
      valid = false
      continue
    }
    if mult == "do()"{
      valid = true
      continue
    }
    if(valid){
      total += multer(mult)
    }
  }
  fmt.Println(total)
}
