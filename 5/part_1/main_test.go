package main

import "testing"

func TestSeedExtractor(t *testing.T) {

	got := ExtractSeeds("Seeds: 1 2 3 4 5 6 7 8 9 10")
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if len(got) != 10 {
		t.Errorf("Expected 10 seeds, got %d", len(got))
	}

	for i, v := range got {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], v)
		}
	}

}
