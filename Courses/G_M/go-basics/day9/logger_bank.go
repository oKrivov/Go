package main

type BankLogger struct {
	Logs []string
}

func (bl *BankLogger) Log(s string) {
	bl.Logs = append(bl.Logs, s)
}
