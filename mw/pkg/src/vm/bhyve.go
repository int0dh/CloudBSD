package vm

import (
	"errors"
	"strconv"
)

type BhyveHypervisor interface {
	PrepareLoaderArgs([]string) error
	PrepareHypervisorArgs([]string) error
}

func (r *VmState) PrepareLoaderArgs(args []string) error {

	if r.config == nil {
		return errors.New("VM is not configured")
	}
	args = append(args, "-d")
	args = append(args, r.config.storage)
	args = append(args, "-m")
	args = append(args, strconv.Itoa(r.config.memsize))
	args = append(args, r.config.vmname)

	return nil
}

func (r *VmState) PrepareHypervisorArgs(args []string) error {
	if r.config == nil {
		return errors.New("VM is not configured!")
	}
	return errors.New("PrepareHypervisorArgs is not implemented yet")
}
