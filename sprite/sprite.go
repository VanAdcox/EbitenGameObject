package sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/png"
	"log"
	"os"
)

type Sprite struct {
	gameObject GameObject
	X          float64
	Y          float64
	Dx         float64
	Dy         float64
	XScale     float64
	YScale     float64

	Img            *ebiten.Image
	OriginalWidth  int
	OriginalHeight int
}

func createBounds(imgPath string) (int, int) {
	file, err := os.Open(imgPath)
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Panicln(err)
	}

	return img.Bounds().Max.X, img.Bounds().Max.Y
}
func (s *Sprite) Init(x, y float64, imgPath string) {
	ebImg, _, err := ebitenutil.NewImageFromFile(imgPath)
	if err != nil {
		log.Panicln(err)
	}
	s.Img = ebImg

	s.X = x
	s.Y = y

	s.XScale = 1
	s.YScale = 1

	s.OriginalWidth, s.OriginalHeight = createBounds(imgPath)
}
func (s *Sprite) GoTo(x, y int) {
	s.X = float64(x)
	s.Y = float64(y)
}
func (s *Sprite) Resize(width, height int) {
	s.XScale = float64(width) / float64(s.OriginalWidth)
	s.YScale = float64(height) / float64(s.OriginalHeight)
}
func (s *Sprite) GenerateImgOptions() *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(s.XScale, s.YScale)
	op.GeoM.Translate(s.X-float64(s.PixelWidth())/2, s.Y-float64(s.PixelHeight())/2)
	return op
}
func (s Sprite) PixelHeight() int {
	return int(float64(s.OriginalHeight) * s.YScale)
}
func (s Sprite) PixelWidth() int {
	return int(float64(s.OriginalWidth) * s.XScale)
}
