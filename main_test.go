package main

import (
	"testing"
)

func TestDevideByThree(t *testing.T) {
	for _, c := range []struct {
		in   int
		want bool
	}{
		{1, false},
		{2, false},
		{3, true},
		{4, false},
	} {
		got := DevideBy(c.in, 3)
		if got != c.want {
			t.Errorf("DevideByThree(%d,3) == %t, want %t", c.in, got, c.want)
		}
	}
}
func TestDevideByFive(t *testing.T) {
	for _, c := range []struct {
		in   int
		want bool
	}{
		{1, false},
		{2, false},
		{3, false},
		{4, false},
		{5, true},
	} {
		got := DevideBy(c.in, 5)
		if got != c.want {
			t.Errorf("DevideByFive(%d,5) == %t, want %t", c.in, got, c.want)
		}
	}
}
func TestFizzBuzz(t *testing.T) {
	for _, c := range []struct {
		in   int
		want string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
	} {
		got := FizzBuzz(c.in)
		if got != c.want {
			t.Errorf("FizzBuzz(%d) == %s, want %s", c.in, got, c.want)
		}
	}
}
func TestContainThree(t *testing.T) {
	for _, c := range []struct {
		in   int
		want bool
	}{
		{1, false},
		{2, false},
		{3, true},
		{4, false},
		{5, false},
	} {
		got := Contain(c.in, 3)
		if got != c.want {
			t.Errorf("Contain(%d,3) == %t, want %t", c.in, got, c.want)
		}
	}
}
func TestContainFive(t *testing.T) {
	for _, c := range []struct {
		in   int
		want bool
	}{
		{1, false},
		{2, false},
		{3, false},
		{4, false},
		{5, true},
	} {
		got := Contain(c.in, 5)
		if got != c.want {
			t.Errorf("Contain(%d,5) == %t, want %t", c.in, got, c.want)
		}
	}
}
func TestMain(t *testing.T) {
	main()
}
