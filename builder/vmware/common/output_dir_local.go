package common

import (
	"log"
	"os"
	"path/filepath"
)

// LocalOutputDir is an OutputDir implementation where the directory
// is on the local machine.
type LocalOutputDir struct {
	dir string
}

func (d *LocalOutputDir) DirExists() (bool, error) {
	_, err := os.Stat(d.dir)
	return err == nil, nil
}

func (d *LocalOutputDir) ListFiles() ([]string, error) {
	files := make([]string, 0, 10)
	log.Printf("XXX Berne: builder/vmware/common/output_dir_local.go ListFiles path: %s", d.dir)

	visit := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("XXX Berne: builder/vmware/common/output_dir_local.go ListFiles path: %s, visit err", path)
			return err
		}
		if !info.IsDir() {
			log.Printf("XXX Berne: builder/vmware/common/output_dir_local.go ListFiles path: %s, is not dir, appending to files", path)
			files = append(files, path)
		}
		return nil
	}

	return files, filepath.Walk(d.dir, visit)
}

func (d *LocalOutputDir) MkdirAll() error {
	return os.MkdirAll(d.dir, 0755)
}

func (d *LocalOutputDir) Remove(path string) error {
	log.Printf("XXX Berne: builder/vmware/common/output_dir_local.go Remove path: %s", d.dir)
	return os.Remove(path)
}

func (d *LocalOutputDir) RemoveAll() error {
	log.Printf("XXX Berne: builder/vmware/common/output_dir_local.go RemoveAll path: %s", d.dir)
	return os.RemoveAll(d.dir)
}

func (d *LocalOutputDir) SetOutputDir(path string) {
	log.Printf("XXX Berne: builder/vmware/common/output_dir_local.go SetOutputDir path: %s", path)
	d.dir = path
}

func (d *LocalOutputDir) String() string {
	return d.dir
}
