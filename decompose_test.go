package decompose

import (
	"testing"
	"simplex/opts"
	"simplex/node"
	"simplex/offset"
	"github.com/intdxdt/geom"
	"github.com/franela/goblin"
	"simplex/pln"
)

func TestDecompose(t *testing.T) {
	var g = goblin.Goblin(t)

	g.Describe("hull decomposition", func() {
		g.It("should test decomposition of a line", func() {
			options := &opts.Opts{
				Threshold:              50.0,
				MinDist:                20.0,
				RelaxDist:              30.0,
				KeepSelfIntersects:     true,
				AvoidNewSelfIntersects: true,
				GeomRelation:           true,
				DistRelation:           false,
				DirRelation:            false,
			}

			var scoreRelation = func(val float64) bool {
				return val <= options.Threshold
			}
			// self.relates = relations(self)
			var wkt = "LINESTRING ( 470 480, 470 450, 490 430, 520 420, 540 440, 560 430, 580 420, 590 410, 630 400, 630 430, 640 460, 630 490, 630 520, 640 540, 660 560, 690 580, 700 600, 730 600, 750 570, 780 560, 790 550, 800 520, 830 500, 840 480, 850 460, 900 440, 920 440, 950 480, 990 480, 1000 520, 1000 570, 990 600, 1010 620, 1060 600 )"
			var coords = geom.NewLineStringFromWKT(wkt).Coordinates()
			var inst   = &dpTest{Pln: pln.New(coords), Opts: options, ScoreFn: offset.MaxOffset}

			inst.Opts.Threshold = 120
			hulls := DouglasPeucker(inst, scoreRelation, hullGeom)

			g.Assert(hulls.Len()).Equal(4)

			inst.Opts.Threshold = 150
			hulls = DouglasPeucker(inst, scoreRelation, hullGeom)

			g.Assert(hulls.Len()).Equal(1)
			h := hulls.Get(0).(*node.Node)
			g.Assert(h.Range.AsSlice()).Equal([]int{0, len(coords) - 1})
		})
	})
}
