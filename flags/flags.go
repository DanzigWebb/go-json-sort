package flags

import "flag"

var (
	// Path - flag to folder path
	Path string
)

// Parse call flag parser
func Parse() {
	flag.StringVar(&Path, "path", "i18n", "path to folder")
	flag.Parse()
}
