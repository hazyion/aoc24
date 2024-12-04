package main

import (
	"fmt"
	"os"
	"strings"
)

func xmas(x int, y int, n int, l int, mx *[]string) int {
  mxx := *mx
  lk := map[int]string{
    1: "M",
    2: "A",
    3: "S",
  }
  total := 0

  // up
  oops := false
  for i := 1; i < 4; i++{
    if x-i < 0 || string(mxx[x-i][y]) != lk[i]{
      oops = true
      break
    }
  }
  if !oops{
    total += 1
  }

  // down
  oops = false
  for i := 1; i < 4; i++{
    if x+i >= n || string(mxx[x+i][y]) != lk[i]{
      oops = true
      break
    }
  }
  if !oops{
    total += 1
  }

  // left
  oops = false
  for i := 1; i < 4; i++{
    if y-i < 0 || string(mxx[x][y-i]) != lk[i]{
      oops = true
      break
    }
  }
  if !oops{
    total += 1
  }

  // right 
  oops = false
  for i := 1; i < 4; i++{
    if y+i >= l || string(mxx[x][y+i]) != lk[i]{
      oops = true
      break
    }
  }
  if !oops{
    total += 1
  }

  // NW (northwest diagonal)
  oops = false
  for i := 1; i < 4; i++{
    ix, iy := x-i, y-i
    if ix < 0 || iy < 0 || string(mxx[ix][iy]) != lk[i]{
      oops = true
      break
    }
  }
  if !oops{
    total += 1
  }

  // NE
  oops = false
  for i := 1; i < 4; i++{
    ix, iy := x-i, y+i
    if ix < 0 || iy >= l || string(mxx[ix][iy]) != lk[i]{
      oops = true
      break
    }
  }
  if !oops{
    total += 1
  }

  // SE
  oops = false
  for i := 1; i < 4; i++{
    ix, iy := x+i, y+i
    if ix >= n || iy >= l || string(mxx[ix][iy]) != lk[i]{
      oops = true
      break
    }
  }
  if !oops{
    total += 1
  }

  // SW
  oops = false
  for i := 1; i < 4; i++{
    ix, iy := x+i, y-i
    if ix >= n || iy < 0 || string(mxx[ix][iy]) != lk[i]{
      oops = true
      break
    }
  }
  if !oops{
    total += 1
  }

  return total
}

func xmas2(x int, y int, n int, l int, mx *[]string) bool {
  mxx := *mx
  // NW Diagonal
  ix, iy, ix2, iy2 := x-1, y-1, x+1, y+1
  if ix < 0 || iy < 0 || ix2 >= n || iy2 >= l{
    return false
  }
  if !(string(mxx[ix][iy]) == "M" && string(mxx[ix2][iy2]) == "S") && !(string(mxx[ix][iy]) == "S" && string(mxx[ix2][iy2]) == "M"){
    return false
  }
  // NE Diagonal
  ix, iy, ix2, iy2 = x+1, y-1, x-1, y+1
  if ix >= n || iy < 0 || ix2 < 0 || iy2 >= l{
    return false
  }
  if !(string(mxx[ix][iy]) == "M" && string(mxx[ix2][iy2]) == "S") && !(string(mxx[ix][iy]) == "S" && string(mxx[ix2][iy2]) == "M"){
    return false
  }
  return true 
}

func main(){
  data, err := os.ReadFile("inputs/d4.txt")
  if err != nil{
    fmt.Printf("Error reading file")
    return
  }
  mx := strings.Fields(string(data))
  l, n := len(mx[0]), len(mx)

  total, total2 := 0, 0
  for i := range n{
    for j := range l{
      // Part 1
      if string(mx[i][j]) == "X"{
	total += xmas(i, j, n, l, &mx)
      }

      // Part 2
      if string(mx[i][j]) == "A"{
	if xmas2(i, j, n, l, &mx){
	  total2 += 1
	}
      }
    }
  }
  fmt.Println(total)
  fmt.Println(total2)
}
