package dijkstra

import "testing"

func pathsEqual(a, b []Point) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestRun_Simple(t *testing.T) {
	input := [][]int{
		{1, 1, 1},
		{1, 0, 0},
		{1, 1, 1},
	}
	start := Point{X: 0, Y: 0}
	end := Point{X: 2, Y: 2}

	expectedDistance := 5
	expectedPath := []Point{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
	}

	distance, path := Run(input, start, end)

	if distance != expectedDistance {
		t.Errorf("Expected distance %d, got %d", expectedDistance, distance)
	}

	if !pathsEqual(path, expectedPath) {
		t.Errorf("Expected path %v, got %v", expectedPath, path)
	}
}

func TestRun_Obstacles(t *testing.T) {
	input := [][]int{
		{1, 0, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	start := Point{X: 0, Y: 0}
	end := Point{X: 0, Y: 2}

	expectedDistance := 7
	expectedPath := []Point{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
		{X: 1, Y: 2},
		{X: 0, Y: 2},
	}

	distance, path := Run(input, start, end)

	if distance != expectedDistance {
		t.Errorf("Expected distance %d, got %d", expectedDistance, distance)
	}

	if !pathsEqual(path, expectedPath) {
		t.Errorf("Expected path %v, got %v", expectedPath, path)
	}
}

func TestRun_Complex(t *testing.T) {
	input := [][]int{
		{1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 0, 1},
		{1, 0, 1, 0, 0, 1},
		{1, 1, 1, 1, 0, 1},
	}
	start := Point{X: 0, Y: 0}
	end := Point{X: 2, Y: 3}

	expectedDistance := 10
	expectedPath := []Point{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
		{X: 3, Y: 0},
		{X: 4, Y: 0},
		{X: 4, Y: 1},
		{X: 4, Y: 2},
		{X: 3, Y: 2},
		{X: 2, Y: 2},
		{X: 2, Y: 3},
	}

	distance, path := Run(input, start, end)

	if distance != expectedDistance {
		t.Errorf("Expected distance %d, got %d", expectedDistance, distance)
	}

	if !pathsEqual(path, expectedPath) {
		t.Errorf("Expected path %v, got %v", expectedPath, path)
	}
}
