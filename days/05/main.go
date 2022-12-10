package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Stack struct {
	data []string
}

func NewStack() *Stack {
	return &Stack{
		data: []string{},
	}
}

func (s *Stack) Push(v string) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() string {
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return item
}

func (s *Stack) Peek() string {
	if len(s.data) == 0 {
		return "0"
	}

	item := s.data[len(s.data)-1]
	return item
}

type Crate struct {
	stacks map[int]*Stack
}

func NewCrate() *Crate {
	return &Crate{
		stacks: map[int]*Stack{},
	}
}

func (c *Crate) PushIntoStack(id int, value string) {
	_, ok := c.stacks[id]
	if !ok {
		c.stacks[id] = NewStack()
	}

	c.stacks[id].Push(value)
}

func (c *Crate) PopFromStack(id int) string {
	stack := c.stacks[id]
	return stack.Pop()
}

func (c *Crate) PeekStacks() []string {
	var (
		result []string
	)

	for i := 0; i < len(c.stacks); i++ {
		stack := c.stacks[i]
		item := stack.Peek()
		result = append(result, item)
	}

	return result
}

type Movement struct {
	Quantity int
	From     int
	To       int
}

func NewMovement(spec string) *Movement {
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(spec, -1)

	quantity, _ := strconv.Atoi(numbers[0])
	from, _ := strconv.Atoi(numbers[1])
	to, _ := strconv.Atoi(numbers[2])

	return &Movement{
		Quantity: quantity,
		From:     from,
		To:       to,
	}
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = input.Close()
	}()

	fileScanner := bufio.NewScanner(input)

	var (
		crate = NewCrate()
		lines []string
	)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) == 0 {
			break
		}

		lines = append(lines, line)
	}

	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]

		for i := 0; i < len(line); i++ {
			char := string(line[i])
			if char == "[" {
				v := string(line[i+1])
				stackId := i / 4
				crate.PushIntoStack(stackId, v)
			}
		}
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		movement := NewMovement(line)

		from := movement.From - 1
		to := movement.To - 1

		auxStack := NewStack()

		for i := 0; i < movement.Quantity; i++ {
			v := crate.PopFromStack(from)
			auxStack.Push(v)
		}

		for i := 0; i < movement.Quantity; i++ {
			v := auxStack.Pop()
			crate.PushIntoStack(to, v)
		}
	}

	fmt.Println(crate.PeekStacks())
}
