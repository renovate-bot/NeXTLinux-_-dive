package viewmodel

import (
	"github.com/nextlinux/dive/dive/image"
)

type LayerSelection struct {
	Layer                                                      *image.Layer
	BottomTreeStart, BottomTreeStop, TopTreeStart, TopTreeStop int
}
