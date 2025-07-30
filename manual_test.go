package main

import (
	"fmt"
	"testing"
)

func TestManualAmount(t *testing.T) {
	testCases := []string{
		"ALICI: Test User\nAÇIKLAMA: Payment\nİŞLEM TUTARI: 1,234.56 TL",
		"ALICI: Test User\nAÇIKLAMA: Payment\nTUTARI: 500.00 TL",
		"ALICI: Test User\nAÇIKLAMA: Payment\nHAVALE TUTARI: 750.25 TL",
		"ALICI: Test User\nAÇIKLAMA: Payment\nSome text 250.75 TL more text",
		"ALICI: Test User\nAÇIKLAMA: Payment\n1.500,00 TL",
	}

	for i, input := range testCases {
		fmt.Printf("Test %d:\n", i+1)
		fmt.Printf("Input: %s\n", input)
		result := extractFields(input)
		fmt.Printf("Output: %s\n\n", result)
	}
}
