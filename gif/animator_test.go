package gif

import (
	"io/ioutil"
	"testing"

	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/grid"
)

func Test(t *testing.T) {
	f, err := ioutil.TempFile("", "test_animator_")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	t.Logf("file written at %s", f.Name())

	a := NewAnimator(100)
	frames := 5
	for i := 0; i < frames; i++ {
		g, err := grid.New(100, 50, []conway.Coord{})
		if err != nil {
			t.Fatal(err)
		}
		a.Add(g)
	}
	if err := a.Encode(f); err != nil {
		t.Error(err)
	}
}
