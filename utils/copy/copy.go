package copy

import (
	"io"
	"os"

	"github.com/jsorrell/www.jacksorrell.com/data"
)

// WriteAssetToDisk copies a file from assets to os filesystem
func WriteAssetToDisk(src, dst string) error {
	in, err := data.Assets.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}
