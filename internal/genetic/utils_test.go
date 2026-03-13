package genetic

import (
	"reflect"
	"testing"
)

func TestClamp(t *testing.T) {
	if got := Clamp(5.0, 0.0, 10.0); got != 5.0 {
		t.Errorf("Clamp(5.0, 0, 10) = %f; want 5.0", got)
	}
	if got := Clamp(-5.0, 0.0, 10.0); got != 0.0 {
		t.Errorf("Clamp(-5.0, 0, 10) = %f; want 0.0", got)
	}
	if got := Clamp(15.0, 0.0, 10.0); got != 10.0 {
		t.Errorf("Clamp(15.0, 0, 10) = %f; want 10.0", got)
	}

	if got := Clamp(15, 0, 10); got != 10 {
		t.Errorf("Clamp(15, 0, 10) = %d; want 10", got)
	}
}

func TestCloneGenes(t *testing.T) {
	original := []float64{1.1, 2.2, 3.3}
	cloned := CloneGenes(original)

	if !reflect.DeepEqual(original, cloned) {
		t.Errorf("CloneGenes() = %v; want %v", cloned, original)
	}

	cloned[0] = 99.9
	if original[0] == 99.9 {
		t.Error("The original sclice was editted")
	}
}
