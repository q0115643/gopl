// go test -bench=. gopl/ch2/2-3
package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount4(0x1234567890ABCDEF)
	}
}

func TestPopCount(t *testing.T) {
	t.Log("testing PopCount function...")
	result := PopCount(1)
	if result != 1 {
		t.Errorf("expected 1 bit set for (int 1) but got %d", result)
	} else {
		t.Log("1 bit set for (int 1)")
	}

	result = PopCount(0xffef)
	if result != 15 {
		t.Errorf("expected 15 bits set for 0xffef but got %d", result)
	} else {
		t.Log("15 bits set for 0xffef")
	}
}

func TestPopCount2(t *testing.T) {
	t.Log("testing PopCount function...")
	result := PopCount2(1)
	if result != 1 {
		t.Errorf("expected 1 bit set for (int 1) but got %d", result)
	} else {
		t.Log("1 bit set for (int 1)")
	}

	result = PopCount2(0xffef)
	if result != 15 {
		t.Errorf("expected 15 bits set for 0xffef but got %d", result)
	} else {
		t.Log("15 bits set for 0xffef")
	}
}

func TestPopCount3(t *testing.T) {
	t.Log("testing PopCount function...")
	result := PopCount3(1)
	if result != 1 {
		t.Errorf("expected 1 bit set for (int 1) but got %d", result)
	} else {
		t.Log("1 bit set for (int 1)")
	}

	result = PopCount3(0xffef)
	if result != 15 {
		t.Errorf("expected 15 bits set for 0xffef but got %d", result)
	} else {
		t.Log("15 bits set for 0xffef")
	}
}

func TestPopCount4(t *testing.T) {
	t.Log("testing PopCount function...")
	result := PopCount4(1)
	if result != 1 {
		t.Errorf("expected 1 bit set for (int 1) but got %d", result)
	} else {
		t.Log("1 bit set for (int 1)")
	}

	result = PopCount4(0x2)
	if result != 1 {
		t.Errorf("expected 1 bits set for 0x2 but got %d", result)
	} else {
		t.Log("1 bits set for 0x2")
	}

	result = PopCount4(0xffef)
	if result != 15 {
		t.Errorf("expected 15 bits set for 0xffef but got %d", result)
	} else {
		t.Log("15 bits set for 0xffef")
	}
}
