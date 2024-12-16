package helloworld

import (
	"bytes"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	ebiten_res "github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type Gopher struct {
	Image *ebiten.Image
}

type Test struct {
	Image *ebiten.Image
}

var GophersImg = ReadEbImgByte(ebiten_res.Ebiten_png)
var TestImg = ReadEbImgByte(ebiten_res.Runner_png)

func ReadEbImgByte(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

// Canvas resource for drawing.
type Canvas struct {
	// The screen image.
	Image *ebiten.Image
	// Current screen width.
	Width int
	// Current screen height.
	Height int
}

type Render struct {
	canvas  generic.Resource[Canvas]
	gophers generic.Resource[Gopher]
	test    generic.Resource[Test]
}

func (s *Render) Load(world *ecs.World) {
	s.canvas = generic.NewResource[Canvas](world)
	s.gophers = generic.NewResource[Gopher](world)
	s.test = generic.NewResource[Test](world)
}

func (s *Render) Update(world *ecs.World) {
}

func (s *Render) Draw(world *ecs.World) {
	canvas := s.canvas.Get()
	gophers := s.gophers.Get()
	test := s.test.Get()

	canvas.Image.Fill(color.White)

	{
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(100, 100)
		canvas.Image.DrawImage(gophers.Image, &op)
	}
	{
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(100, 200)
		canvas.Image.DrawImage(test.Image, &op)
	}
}
