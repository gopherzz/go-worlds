package gameobject

import "github.com/faiface/pixel/pixelgl"

type GameObject interface {
	Draw(*pixelgl.Window)
	UseController(*pixelgl.Window)
}
