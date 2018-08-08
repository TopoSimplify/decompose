package decompose

import (
	"github.com/TopoSimplify/rng"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/lnr"
	"github.com/TopoSimplify/pln"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/iter"
)

type scoreRelationFn func(float64) bool

//Douglas-Peucker decomposition at a given threshold
func DouglasPeucker(
	id *iter.Igen, pln pln.Polyline, scoreFn lnr.ScoreFn,
	scoreRelation scoreRelationFn, geomFn func(geom.Coords)geom.Geometry,
) []node.Node {
	var k, n int
	var val float64
	var coordinates geom.Coords
	var hque []node.Node

	if pln.LineString == nil {
		return hque
	}

	var r = pln.Range()
	var stack = make([]rng.Rng, 0, (r.J-r.I)+1)
	stack = append(stack, r)

	for !(len(stack) == 0) {
		n = len(stack) - 1
		r = stack[n]
		stack = stack[:n]

		coordinates = pln.SubCoordinates(r)
		k, val = scoreFn(coordinates)
		k = r.I + k //offset

		if scoreRelation(val) {
			hque = append(hque, node.CreateNode(id, coordinates, r, geomFn))
		} else {
			stack = append(stack, rng.Range(k, r.J)) // right
			stack = append(stack, rng.Range(r.I, k)) // left
		}
	}
	return hque
}
