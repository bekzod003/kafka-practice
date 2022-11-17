package smth

import (
	"testing"
)

func TestKafka(t *testing.T) {
	if err := NewKafkaClient("localhost", "29092").Consumer([]string{
		"lol", "bekzod",
	}); err != nil {
		t.Fatal(err)
	}

	t.Log("Successfully done")
}
