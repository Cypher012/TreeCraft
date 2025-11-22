package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadLine(prompt string) string {
	fmt.Print(prompt)
	raw, _ := reader.ReadString('\n')
	return strings.TrimSpace(raw)
}

func Get[T any](prompt string) T {
	for {
		raw := ReadLine(prompt)

		var zero T
		switch any(zero).(type) {

		case int:
			if raw == "" {
				return any(1).(T)
			}
			n, err := strconv.Atoi(raw)
			if err == nil {
				return any(n).(T)
			}

		case string:
			return any(raw).(T)
		}

		fmt.Println("Invalid input. Try again.")
	}
}

func Confirm(msg string) bool {
	res := ReadLine(msg + " (y/n): ")
	return strings.ToLower(res) == "y"
}
