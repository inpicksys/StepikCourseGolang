package Lesson3

import "testing"

func TestChannelWork(t *testing.T) {
	num := 4
	channel := make(chan int, 1)
	go ChannelWork(channel, num)
	var result int
	result = <-channel
	close(channel)
	if result != 5 {
		t.Error("Wrong answer! Expected 5, got ", result)
	}
}

func TestSendModifiedStringToChannel(t *testing.T) {
	str, result := "hop", ""
	channel := make(chan string, 5)
	defer close(channel)
	go SendModifiedStringToChannel(channel, str)
	for i := 0; i < 5; i++ {
		result += <-channel
	}
	close(channel)
	if result != "hop hop hop hop hop " {
		t.Error("Wrong answer! Expected hop hop hop hop hop , got ", result)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	str1, str2, str3, str4, result := "str1", "str2", "str2", "str3", ""
	in := make(chan string, 4)
	out := make(chan string, 4)
	go RemoveDuplicates(in, out)
	in <- str1
	in <- str2
	in <- str3
	in <- str4
	close(in)
	for i := 0; i < 4; i++ {
		result += <-out
	}
	if result != "str1str2str3" {
		t.Error("Wrong answer! Expected str1str2str3, got ", result)
	}
}
