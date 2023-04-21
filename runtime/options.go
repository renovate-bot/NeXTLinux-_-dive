package runtime

import (
	"github.com/nextlinux/dive/dive"
	"github.com/spf13/viper"
)

type Options struct {
	Ci           bool
	Image        string
	Source       dive.ImageSource
	IgnoreErrors bool
	ExportFile   string
	CiConfig     *viper.Viper
	BuildArgs    []string
}
