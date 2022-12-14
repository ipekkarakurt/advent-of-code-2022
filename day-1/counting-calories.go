package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "sort"
)

func splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + 2, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

func main() {
  calories_per_elf := make([]int, 0)

  f, _ := os.Open("input.txt")

  scanner := bufio.NewScanner(f)
  scanner.Split(splitFunc)

  for scanner.Scan() {
    var calories_sum int

    for _, row := range strings.Split(scanner.Text(), "\n") {
			num, _ := strconv.Atoi(row)
      calories_sum += num
		}

    calories_per_elf = append(calories_per_elf, calories_sum)
  }

  sort.Ints(calories_per_elf)
  length := len(calories_per_elf)

  fmt.Println(calories_per_elf[length-1])
  max_3 := calories_per_elf[length-3:]
  fmt.Println(max_3[0] + max_3[1] + max_3[2])
}
