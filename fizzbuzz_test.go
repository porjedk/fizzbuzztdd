package fizzbuzz

import "testing"

func TestFizzBuzz1Return1(t *testing.T) {
	got := FizzBuzz(1)
	want := "1"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestFizzBuzz2Return2(t *testing.T) {
	got := FizzBuzz(2)
	want := "2"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestFizzBuzz3ReturnFizz(t *testing.T) {
	got := FizzBuzz(3)
	want := "Fizz"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}	
}

func TestFizzBuzz5ReturnBuzz(t *testing.T) {
	got := FizzBuzz(5)
	want := "Buzz"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}	
}

func TestFizzBuzz15ReturnFizzBuzz(t *testing.T) {
	got := FizzBuzz(15)
	want := "FizzBuzz"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}	
}