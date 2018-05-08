package decompose

import (
	"github.com/intdxdt/sset"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/deque"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/lnr"
	"github.com/TopoSimplify/opts"
)

//hull geom
func hullGeom(coords []*geom.Point) geom.Geometry {
	var g geom.Geometry

	if len(coords) > 2 {
		g = geom.NewPolygon(coords)
	} else if len(coords) == 2 {
		g = geom.NewLineString(coords)
	} else {
		g = coords[0].Clone()
	}
	return g
}

//Type DP
type dpTest struct {
	id        string
	Hulls     *deque.Deque
	Pln       *pln.Polyline
	Meta      map[string]interface{}
	Opts      *opts.Opts
	ScoreFn   lnr.ScoreFn
	SimpleSet *sset.SSet
}

func (self *dpTest) Id() string {
	return self.id
}

func (self *dpTest) Options() *opts.Opts {
	return self.Opts
}

func (self *dpTest) NodeQueue() *deque.Deque {
	return self.Hulls
}

func (self *dpTest) Simple() []int {
	return []int{}
}

func (self *dpTest) Coordinates() []*geom.Point {
	return self.Pln.Coordinates
}

func (self *dpTest) Polyline() *pln.Polyline {
	return self.Pln
}
