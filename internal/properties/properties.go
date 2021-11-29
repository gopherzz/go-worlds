package properties

import (
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

var (
	BasicAtlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)
)
