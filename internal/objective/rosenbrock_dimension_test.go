package objective

import "testing"

func TestRosenbrock2DDimension(t *testing.T) {
	fn := NewRosenbrock2D()
	if fn.Dimension() != 2 {
		t.Fatalf("Dimension() = %d, want 2", fn.Dimension())
	}
}

func TestRosenbrockNDimension(t *testing.T) {
	fn, err := NewRosenbrock(4)
	if err != nil {
		t.Fatalf("NewRosenbrock() returned error: %v", err)
	}

	if fn.Dimension() != 4 {
		t.Fatalf("Dimension() = %d, want 4", fn.Dimension())
	}
}
