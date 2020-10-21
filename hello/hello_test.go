package hello

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, world."
	if actual := Hello(); actual != expected {
		t.Errorf("Hello() = %q, expected %q", actual, expected)
	}
}

func TestProverb(t *testing.T) {
	expected := "Concurrency is not parallelism."
	if actual := Proverb(); expected != actual {
		t.Errorf("Proverb() =  %q, want %q", actual, expected)
	}
}
