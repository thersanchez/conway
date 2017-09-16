package grid_test

import (
	"testing"

	"github.com/alcortesm/conway/grid"
)

func TestSizeOK(t *testing.T) {
	testSizeOK(t, 3, 3)
	testSizeOK(t, 3, 4)
	testSizeOK(t, 4, 3)
	testSizeOK(t, 4, 4)
	testSizeOK(t, 10, 10)
	testSizeOK(t, 20, 10)
	testSizeOK(t, 10, 20)
	testSizeOK(t, 10000, 20000)
}

func testSizeOK(t *testing.T, width, height int) {
	t.Helper()
	g, err := grid.New(width, height)
	if err != nil {
		t.Fatalf("cannot create grid: %v", err)
	}
	if w := g.Width(); w != width {
		t.Errorf("wrong width: expected %d, got %d", width, w)
	}
	if h := g.Height(); h != height {
		t.Errorf("wrong height: expected %d, got %d", height, h)
	}
}

func TestSizeError(t *testing.T) {
	testSizeError(t, 2, 3)
	testSizeError(t, 3, 2)
	testSizeError(t, 2, 2)
	testSizeError(t, 0, 0)
	testSizeError(t, -6, 3)
	testSizeError(t, 3, -6)
	testSizeError(t, 10000, -2)
	testSizeError(t, -2, 10000)
}

func testSizeError(t *testing.T, width, height int) {
	t.Helper()
	_, err := grid.New(width, height)
	if err == nil {
		t.Errorf("new grid was supposed to fail and it did not")
	}
}
