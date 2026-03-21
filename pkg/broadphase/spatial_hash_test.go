package broadphase

import (
	"testing"
	"github.com/rudransh61/Physix-go/pkg/vector"
)

type mockEnt struct {
	ID int
}

func TestSpatialHash(t *testing.T) {
	// Create Hash Map with 50x50 cells
	cellSize := 50.0
	sh := NewSpatialHash(cellSize, 800, 800)

	obj1 := &mockEnt{ID: 1}
	obj2 := &mockEnt{ID: 2} // Near obj1
	obj3 := &mockEnt{ID: 3} // Far away

	// Cell indices: (X/50, Y/50)
	// (25,25) -> Cell (0,0)
	sh.Add(obj1, vector.Vector{X: 25, Y: 25})
	// (60,60) -> Cell (1,1) -> which is adjacent to (0,0)
	sh.Add(obj2, vector.Vector{X: 60, Y: 60})
	// (500,500) -> Cell (10,10) -> very far
	sh.Add(obj3, vector.Vector{X: 500, Y: 500})

	// Query region around obj1
	queryPos := vector.Vector{X: 30, Y: 30}
	nearby := sh.Query(queryPos)

	if len(nearby) != 2 {
		t.Errorf("SpatialHash.Query() expected 2 nearby objects, got %d", len(nearby))
	}

	foundObj1, foundObj2, foundObj3 := false, false, false

	for _, obj := range nearby {
		m, ok := obj.(*mockEnt)
		if ok {
			if m.ID == 1 { foundObj1 = true }
			if m.ID == 2 { foundObj2 = true }
			if m.ID == 3 { foundObj3 = true }
		}
	}

	if !foundObj1 || !foundObj2 {
		t.Errorf("SpatialHash.Query() failed to find expected adjacent cell objects")
	}
	if foundObj3 {
		t.Errorf("SpatialHash.Query() returned mathematically distant objects incorrectly")
	}

	// Test Clearing implementation
	sh.Clear()
	nearbyClear := sh.Query(queryPos)
	if len(nearbyClear) != 0 {
		t.Errorf("SpatialHash.Clear() expected 0 nearby objects, got %d", len(nearbyClear))
	}
}
