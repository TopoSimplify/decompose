package decompose

import (
	"simplex/pln"
	"simplex/rng"
	"simplex/node"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/deque"
	"simplex/opts"
	"simplex/lnr"
	"github.com/intdxdt/sset"
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

func (self *dpTest) Score(coordinates []*geom.Point, rg *rng.Range) (int, float64) {
	return self.ScoreFn(coordinates, rg)
}

func linear_coords(wkt string) []*geom.Point{
	return geom.NewLineStringFromWKT(wkt).Coordinates()
}

func create_hulls(indxs [][]int, coords []*geom.Point) []*node.Node {
	poly := pln.New(coords)
	hulls := make([]*node.Node, 0)
	for _, o := range indxs {
		r := rng.NewRange(o[0], o[1])
		hulls = append(hulls, node.New(poly.SubCoordinates(r), r, hullGeom))
	}
	return hulls
}
