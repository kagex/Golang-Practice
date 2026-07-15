package main

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
	"testing"
)

func runProgram(input string) (string, error) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin = bytes.NewBufferString(input)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	output := stdout.String()
	if stderr.Len() > 0 {
		return "", errors.New(stderr.String())
	}

	return strings.TrimSpace(output), nil
}

type testCase struct {
	name     string
	input    string
	expected string
}

func TestTimeOfDay(t *testing.T) {
	tests := []testCase{
		{"утро 6ч", "6\n", "Сейчас 6ч. - утро."},
		{"утро 7ч", "7\n", "Сейчас 7ч. - утро."},
		{"утро 11ч", "11\n", "Сейчас 11ч. - утро."},
		{"день 12ч", "12\n", "Сейчас 12ч. - день."},
		{"день 13ч", "13\n", "Сейчас 13ч. - день."},
		{"день 17ч", "17\n", "Сейчас 17ч. - день."},
		{"вечер 18ч", "18\n", "Сейчас 18ч. - вечер."},
		{"вечер 21ч", "21\n", "Сейчас 21ч. - вечер."},
		{"вечер 22ч", "22\n", "Сейчас 22ч. - вечер."},
		{"ночь 23ч", "23\n", "Сейчас 23ч. - ночь."},
		{"ночь 0ч", "0\n", "Сейчас 0ч. - ночь."},
		{"ночь 5ч", "5\n", "Сейчас 5ч. - ночь."},
		{"отрицательное число", "-5\n", "Неверно задано время"},
		{"больше 24", "25\n", "Неверно задано время"},
		{"текст", "abc\n", "Неверно задано время"},
		{"пустая строка", "\n", "Неверно задано время"},
		{"граничное 5ч", "5\n", "Сейчас 5ч. - ночь."},
		{"граничное 11ч", "11\n", "Сейчас 11ч. - утро."},
		{"граничное 17ч", "17\n", "Сейчас 17ч. - день."},
		{"граничное 22ч", "22\n", "Сейчас 22ч. - вечер."},
		{"граничное 24ч", "24\n", "Сейчас 24ч. - ночь."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := runProgram(tt.input)
			if err != nil {
				t.Fatalf("Error running program: %v", err)
			}
			if !strings.Contains(output, tt.expected) {
				t.Errorf("Expected to contain %q, got %q", tt.expected, output)
			}
		})
	}
}
