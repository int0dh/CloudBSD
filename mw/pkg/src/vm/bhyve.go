package vm

import (
	"errors"
	"strconv"
)

type BhyveHypervisor interface {
	PrepareLoaderArgs([]string) (error, []string)
	PrepareHypervisorArgs([]string) (error, []string)
}

func (r *VmState) PrepareLoaderArgs(args []string) (error, []string) {

	if r.config == nil {
		return errors.New("VM is not configured"), nil
	}
	args = args[0:0]
	args = append(args, "-d")
	args = append(args, r.config.storage)
	args = append(args, "-m")
	args = append(args, strconv.Itoa(r.config.memsize))
	args = append(args, r.config.vmname)

	return nil, args
}

func (r *VmState) PrepareHypervisorArgs(args []string) (error, []string) {
	if r.config == nil {
		return errors.New("VM is not configured!"), nil
	}
	return errors.New("PrepareHypervisorArgs is not implemented yet"), nil
}
