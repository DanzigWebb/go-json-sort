package flags

import "flag"

var (
	// Path - flag to folder path
	Path string
)

// Init call flag parser
func Init() {
	flag.StringVar(&Path, "path", "i18n", "path to folder")
	flag.Parse()
}
