package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// processStack processes the data that is in a string, separate it, and put it into the stack with appropriate location.
func processStack(stack [][]uint8, data string) {
	// based on the sample data
	// [A] [B] [C]
	// this means we need to read every 3 characters + 1 space.

	for i := 0; i < len(data); i += 4 {
		n := data[i : i+3]
		if n[1] != ' ' {
			stack[i/4] = append(stack[i/4], n[1])
		}
	}
}

// moveElement1 moves count elements from the index from to index to within the given stack. This function is for part 1
// of the puzzle.
func moveElement1(stack [][]uint8, count int, from int, to int) {
	for i := 0; i < count; i++ {
		lfrom := len(stack[from-1])
		if lfrom == 0 {
			// nothing to move anymore.
			break
		}

		stack[to-1] = append(stack[to-1], stack[from-1][lfrom-1])
		stack[from-1] = stack[from-1][:lfrom-1]
	}
}

// moveElement1 moves count elements from the index from to index to within the given stack. This function is for part 2
// of the puzzle.
func moveElement2(stack [][]uint8, count int, from int, to int) {
	l := len(stack[from-1])
	ub := l
	lb := ub - count
	if lb < 0 {
		lb = 0
	}

	for i := lb; i < ub; i++ {
		stack[to-1] = append(stack[to-1], stack[from-1][i])
	}

	stack[from-1] = stack[from-1][:lb]
}

func main() {
	// Read the input file.
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %s", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	temp := make([]string, 0)
	stack1 := make([][]uint8, 0)
	stack2 := make([][]uint8, 0)
	for {
		l, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatalf("failed to read input file: %s", err)
		}

		if err == io.EOF {
			break
		}
		l = strings.TrimSuffix(l, "\r\n")

		// First of all, we need to read n number of lines, until we find the list of numbers. This is the indicator
		// that we need to process the stack.
		{
			tmp := strings.Split(strings.TrimSpace(l), " ")
			if strings.TrimSpace(tmp[0]) == "1" {
				stackCount := 0
				for i := 0; i < len(tmp); i++ {
					if tmp[i] != "" {
						stackCount++
					}
				}

				stack1 = make([][]uint8, stackCount)
				stack2 = make([][]uint8, stackCount)
				for i := 0; i < stackCount; i++ {
					stack1[i] = make([]uint8, 0)
					stack2[i] = make([]uint8, 0)
				}

				// Process it.
				for i := len(temp) - 1; i >= 0; i-- {
					processStack(stack1, temp[i])
					processStack(stack2, temp[i])
				}

				break
			} else {
				temp = append(temp, l)
			}
		}
	}

	r.ReadString('\n') // Remove one empty line

	for {
		// Continue reading the rest of the input file. Now, all is about the movement.
		var c, f, t int
		n, _ := fmt.Fscanf(r, "move %d from %d to %d\r\n", &c, &f, &t)
		if n == 0 {
			break
		}

		moveElement1(stack1, c, f, t)
		moveElement2(stack2, c, f, t)
	}

	// Print the result of part 1.
	for i := 0; i < len(stack1); i++ {
		l := len(stack1[i])

		if l > 0 {
			fmt.Printf("%c", stack1[i][l-1])
		}
	}

	fmt.Println()

	// Print the result of part 2.
	for i := 0; i < len(stack2); i++ {
		l := len(stack2[i])

		if l > 0 {
			fmt.Printf("%c", stack2[i][l-1])
		}
	}
}
