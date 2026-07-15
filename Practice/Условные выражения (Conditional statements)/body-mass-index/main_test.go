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

func TestBMI(t *testing.T) {
	tests := []testCase{
		{"недостаточный вес", "50\n170\n", "Ваш ИМТ: 17.30\nКатегория: Недостаточный вес"},
		{"нормальный вес среднее", "70\n175\n", "Ваш ИМТ: 22.86\nКатегория: Нормальный вес"},
		{"избыточный вес", "80\n170\n", "Ваш ИМТ: 27.68\nКатегория: Избыточный вес"},
		{"ожирение", "95\n170\n", "Ваш ИМТ: 32.87\nКатегория: Ожирение"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := runProgram(tt.input)
			if err != nil {
				t.Fatalf("Error running program: %v", err)
			}
			expectedWithPrompt := "Введите ваш вес (кг): Введите ваш рост (см): \n" + tt.expected
			if output != expectedWithPrompt {
				t.Errorf("Expected %q, got %q", expectedWithPrompt, output)
			}
		})
	}
}

func TestBMIErrors(t *testing.T) {
	errorTests := []struct {
		name  string
		input string
	}{
		{"отрицательный вес", "-5\n170\n"},
		{"отрицательный рост", "70\n-170\n"},
		{"ноль вес", "0\n170\n"},
		{"ноль рост", "70\n0\n"},
		{"текст вес", "abc\n170\n"},
		{"текст рост", "70\nabc\n"},
		{"пустая строка вес", "\n170\n"},
		{"пустая строка рост", "70\n\n"},
	}

	for _, tt := range errorTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := runProgram(tt.input)
			if err == nil {
				t.Errorf("Expected error for invalid input %q", tt.input)
			}
		})
	}
}
