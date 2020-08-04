package Lesson3

import "testing"

func TestEncodeFromFileDecodeToFileCountSumAllGlobalIDs(t *testing.T) {
	fileName := "../data-test.json"
	result := EncodeFromFileDecodeToFileCountSumAllGlobalIDs(fileName)
	if result != 6 {
		t.Error("Wrong answer! Expected 6, got ", result)
	}
}
