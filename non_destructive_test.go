package o

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	tests := []struct {
		name     string
		ra       Ring
		cycles   uint
		expected []uint
	}{
		{"basic5/13", NewRing(5), 13, []uint{3, 4, 0, 1, 2}},
		{"basic5/6", NewRing(5), 6, []uint{1, 2, 3, 4, 0}},
		{"mask4/13", NewRing(4), 13, []uint{1, 2, 3, 0}},
		{"mask4/6", NewRing(4), 6, []uint{2, 3, 0, 1}},
	}
	for _, elt := range tests {
		test := elt
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			var i uint
			for i = 0; i < test.cycles; i++ {
				ForcePush(test.ra)
			}
			assert.Equal(t, test.expected, All(test.ra))
		})
	}
}

func TestRev(t *testing.T) {
	tests := []struct {
		name     string
		ra       Ring
		cycles   uint
		expected []uint
	}{
		{"basic5/13", NewRing(5), 13, []uint{2, 1, 0, 4, 3}},
		{"basic5/6", NewRing(5), 6, []uint{0, 4, 3, 2, 1}},
		{"mask4/13", NewRing(4), 13, []uint{0, 3, 2, 1}},
		{"mask4/6", NewRing(4), 6, []uint{1, 0, 3, 2}},
	}
	for _, elt := range tests {
		test := elt
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			var i uint
			for i = 0; i < test.cycles; i++ {
				ForcePush(test.ra)
			}
			assert.Equal(t, test.expected, Rev(test.ra))
		})
	}
}

func TestStartEnd(t *testing.T) {
	tests := []struct {
		name               string
		ra                 Ring
		cycles             uint
		start1, end1, end2 uint
	}{
		// filled beyond their capacity:
		{"basic5/13", NewRing(5), 13, 3, 5, 3},
		{"basic5/6", NewRing(5), 6, 1, 5, 1},
		{"mask4/13", NewRing(4), 13, 1, 4, 1},
		{"mask4/6", NewRing(4), 6, 2, 4, 2},
		// Filled to less than capacity:
		{"mask4/2", NewRing(4), 2, 0, 2, 0},
		{"basic5/2", NewRing(5), 2, 0, 2, 0},
	}
	for _, elt := range tests {
		test := elt
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			var i uint
			for i = 0; i < test.cycles; i++ {
				ForcePush(test.ra)
			}
			t.Logf("e: %d:%d, 0:%d", test.start1, test.end1, test.end2)
			t.Logf("g: %d:%d, 0:%d - %d:%d", Start1(test.ra), End1(test.ra), End2(test.ra),
				test.ra.start(), test.ra.capacity())
			assert.Equal(t, test.start1, Start1(test.ra))
			assert.Equal(t, test.end1, End1(test.ra))
			assert.Equal(t, test.end2, End2(test.ra))
		})
	}
}
