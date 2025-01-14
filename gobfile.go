package pogreb

import (
	"encoding/gob"

	"github.com/maitai/pogreb/fs"
)

func readGobFile(fsys fs.FileSystem, name string, v interface{}, readOnly bool) error {
	f, err := openFile(fsys, name, false, readOnly)
	if err != nil {
		return err
	}
	defer f.Close()
	dec := gob.NewDecoder(f)
	return dec.Decode(v)
}

func writeGobFile(fsys fs.FileSystem, name string, v interface{}, readOnly bool) error {
	f, err := openFile(fsys, name, true, readOnly)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	return enc.Encode(v)
}
