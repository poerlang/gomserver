package handle

import (
	"quadtree"
	. "twof"
)

const (
	QuadtreeInitSize = 6400
)

type GameMap struct {
	Name string
	Tree *quadtree.Quadtree
}

var (
	// All players and monsters are stored in an Quadtree, to make it quick to find distances.
	lowerLeftNearCorner = TwoF{-QuadtreeInitSize, -QuadtreeInitSize}
	upperLeftFarCorner  = TwoF{QuadtreeInitSize, QuadtreeInitSize}
	playerQuadtree      = quadtree.MakeQuadtree(lowerLeftNearCorner, upperLeftFarCorner, 1)
	MapA                = GameMap{"Map A", playerQuadtree} //地图A
)
