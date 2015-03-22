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
	args = append(args, r.config.Storage)
	args = append(args, "-m")
	args = append(args, strconv.Itoa(r.config.Memsize))
	args = append(args, r.config.Vmname)

	return nil, args
}

func (r *VmState) PrepareHypervisorArgs(args []string) (error, []string) {
	if r.config == nil {
		return errors.New("VM is not configured!"), nil
	}
	args = args[0:0]
	args = append(args, "-AI -H -P -s ")
	args = append(args, "-H")
	args = append(args, "-P")
	args = append(args, "-s 0:0,hostbridge")
	args = append(args, "-s 1:0,lpc")
	args = append(args, "-s 2:0,virtio-net,"+r.config.Netif)
	args = append(args, "-s 3:0,virtio-blk,"+r.config.Storage)
	args = append(args, "-l "+r.config.Console)
	args = append(args, "-c "+strconv.Itoa(r.config.Numcores))
	args = append(args, "-m "+strconv.Itoa(r.config.Memsize))
	args = append(args, r.config.Vmname)

	return nil, args
}
