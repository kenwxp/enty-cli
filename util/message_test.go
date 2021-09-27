package util

import "testing"

func TestSendCheckCodeMessage(t *testing.T) {
	SendCheckCodeMessage("15361445990", "156893")
}

func TestEmail(t *testing.T) {
	SingleMail("terilscaub@gmail.com", "184684")
}
