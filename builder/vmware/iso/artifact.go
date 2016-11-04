package iso

import (
	"fmt"
	"log"
)

// Artifact is the result of running the VMware builder, namely a set
// of files associated with the resulting machine.
type Artifact struct {
	builderId string
	dir       OutputDir
	f         []string
}

func (a *Artifact) BuilderId() string {
	return a.builderId
}

func (a *Artifact) Files() []string {
	return a.f
}

func (*Artifact) Id() string {
	return "VM"
}

func (a *Artifact) String() string {
	return fmt.Sprintf("VM files in directory: %s", a.dir)
}

func (a *Artifact) State(name string) interface{} {
	return nil
}

func (a *Artifact) Destroy() error {
	log.Printf("XXX Berne: builder/vmware/iso/artifact.go Artifact::Destroy Remove %s", a.dir)
	return a.dir.RemoveAll()
}
