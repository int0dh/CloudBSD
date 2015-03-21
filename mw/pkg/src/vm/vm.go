package vm

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os/exec"
	"strings"
)

// Yes, there may be more than just one interface
// or one disk per VM
// but let`s keep thing simply at this stage
type VmConfig struct {
	Vmname   string
	Storage  string
	Netif    string
	Memsize  int
	Numcores int
}

type VmState struct {
	config *VmConfig
	// pid of the user part of hypervisor (bhyve)
	pid int
	// state of VM, will be encoded further, not used for now
	state int
	// path to the loader (for bhyve it is either bhyveload or grub-bhyve)
	Loader string
	// path to the user-part of hypervisor
	Hypervisor string
}

type CustomHypervisor interface {
	PrepareLoaderArgs([]string)
	PrepareHypervisorArgs([]string)
}

type Hypervisor interface {
	InitVM(string) error
	CreateVM() error
	DestroyVM() error
	CustomHypervisor
}

func (r *VmState) CreateVM() error {

	var args []string

	if r.Loader != "" {
		err, args := r.PrepareLoaderArgs(args)
		if err != nil {
			return err
		}
		cmd := exec.Command(r.Loader, args...)

		cmd.Stdin = strings.NewReader("\n\n")
		if err = cmd.Start(); err != nil {
			return err
		}
		if err = cmd.Wait(); err != nil {
			return err
		}
	}
	if r.Hypervisor != "" {
		err, args := r.PrepareHypervisorArgs(args)
		cmd := exec.Command(r.Hypervisor, args...)
		if err = cmd.Start(); err != nil {
			return err
		}
	}
	return nil
}

func (r *VmState) InitVM(fileName string) error {
	var result VmConfig

	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fileContent, &result)
	if err != nil {
		return err
	}
	r.config = &result

	return nil
}

func (r *VmState) DestroyVM() error {
	return errors.New("Destroy VM not implemented")
}
