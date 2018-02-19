package image

import (
	"io/ioutil"
	"sync"

	"github.com/DMarby/picsum-photos/vips"
)

// Processor is an image processor
type Processor struct {
}

var instance *Processor
var once sync.Once

// GetInstance returns the processor instance, and creates it if neccesary.
func GetInstance() (*Processor, error) {
	var err error

	once.Do(func() {
		err = vips.Initialize()
		instance = &Processor{}
	})

	return instance, err
}

type Image struct {
	data []byte
}

// Shutdown shuts down the image processor and deinitialises vips
func (p *Processor) Shutdown() {
	vips.Shutdown()
}

// TODO: What should we expose? Just resize, crop, etc?
func (p *Processor) LoadImage(path string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	image, err := vips.ResizeImage(buf, 500, 500)
	if err != nil {
		return err
	}

	image, err = vips.Grayscale(image)
	if err != nil {
		return err
	}

	imageBuffer, err := vips.SaveToBuffer(image)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./fixtures/result.jpg", imageBuffer, 0644)
	if err != nil {
		return err
	}

	return nil
}
