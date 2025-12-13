package main

import (
	"errors"
	"fmt"
)

func ProcessTransaction(a *Account, amount int) error {
	err := a.Withdraw(amount)
	if err == nil {
		return nil
	}

	// 1. Бизнес-ошибки (оборачиваем)
	if errors.Is(err, ErrInsufficientFunds) {
		return fmt.Errorf("transaction declined: %w", err)
	}

	// 2. Ошибки валидации
	var neg NegativeAmountError
	if errors.As(err, &neg) {
		return fmt.Errorf("invalid transaction amount: %w", err)
	}

	// 3. Системные состояния
	if errors.Is(err, ErrAccountFrozen) {
		return err
	}

	// 4. Фолбэк
	return err
}
