package random

import (
	"fmt"

	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/evolve"
	"github.com/alcortesm/conway/grid"
)

// Random is a Conway universe that starts with random alives.
type Random struct {
	width  int
	height int
	status conway.Grid
}

// New returns a new Random universe.
func New(w, h, n int) (*Random, error) {
	if w < grid.MinWidth {
		return nil, fmt.Errorf("width < %d, was %d", grid.MinWidth, w)
	}
	if h < grid.MinHeight {
		return nil, fmt.Errorf("height < %d, was %d", grid.MinHeight, h)
	}
	if n < 0 {
		return nil, fmt.Errorf("number of initial alives must be > 0, was %d", n)
	}
	return &Random{
		width:  w,
		height: h,
		status: grid.NewRandom(w, h, n),
	}, nil
}

// Status implements conway.Universe.
func (r *Random) Status() conway.Grid {
	return r.status
}

// Tick implements conway.Universe.
func (r *Random) Tick() error {
	alives, err := evolve.Evolve(r.status)
	if err != nil {
		return err
	}
	r.status, err = grid.New(r.status.Width(), r.status.Height(), alives)
	if err != nil {
		return err
	}
	return nil
}
