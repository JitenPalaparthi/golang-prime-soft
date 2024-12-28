package square_test

import (
	"shapes/shape/square"
	"testing"
)

func TestWhat(t *testing.T) {
	sq := square.New(10)
	expect := "Square"
	actual := sq.What()
	if expect != actual {
		t.Fail()
	}
}

func TestArea(t *testing.T) {
	sq := square.New(10)
	if sq.Area() != 100 {
		t.Fail()
	}
}

func TestPerimer(t *testing.T) {
	sq := square.New(10)
	if sq.Perimeter() != 40 {
		t.Fail()
	}
}
