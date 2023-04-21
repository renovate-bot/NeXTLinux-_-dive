package docker

import (
	"fmt"
	"github.com/nextlinux/dive/dive/image"
	"os"
)

type archiveResolver struct{}

func NewResolverFromArchive() *archiveResolver {
	return &archiveResolver{}
}

func (r *archiveResolver) Fetch(path string) (*image.Image, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	img, err := NewImageArchive(reader)
	if err != nil {
		return nil, err
	}
	return img.ToImage()
}

func (r *archiveResolver) Build(args []string) (*image.Image, error) {
	return nil, fmt.Errorf("build option not supported for docker archive resolver")
}
