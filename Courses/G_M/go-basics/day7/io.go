package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// readLine читает одну строку из reader и возвращает её trimmed.
// Если EOF или другая ошибка — возвращает её.
func readLine(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

// parseParams парсит строку "a op b" и возвращает a, b, op или ошибку.
// Не использует panic — всегда возвращает ошибку при некорректном вводе.
func parseParams(line string) (int, int, string, error) {
	fields := strings.Fields(line)

	if len(fields) != 3 {
		return 0, 0, "", fmt.Errorf("invalid input format: expected  `a op b`")
	}
	aStr, op, bStr := fields[0], fields[1], fields[2]

	a, err := strconv.Atoi(aStr)
	if err != nil {
		return 0, 0, "", fmt.Errorf("invalid first numder: %w", err)
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		return 0, 0, "", fmt.Errorf("invalid second numder: %w", err)
	}
	if (op != "+") && op != "-" && op != "*" && op != "/" {
		return 0, 0, "", fmt.Errorf("invalid operator: %s", op)
	}

	return a, b, op, nil
}

// getCalcParams интерактивно запрашивает у пользователя параметры операции.
// Возвращает a,b,op или ошибку (например io.EOF).
func getCalcParams(r *bufio.Reader) (int, int, string, error) {
	for {
		fmt.Println("Enter parmetrs for calculation (example: 2 * 2). Press Ctrl+D to exit.")
		fmt.Print(">")

		line, err := readLine(r)
		if err != nil {
			// распространяется на io.EOF или другие I/O ошибки
			return 0, 0, "", err
		}
		if line == "" {
			fmt.Println("Empty input? try again.")
			continue
		}

		a, b, op, err := parseParams(line)
		if err != nil {
			fmt.Println("Input error:", err)
			continue // попросим ввести снова
		}
		return a, b, op, nil
	}
}
