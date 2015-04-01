package main

import (
	"log"
	"time"
	"vm"
)

const config = "vmdesc.json"
const loader = "/usr/sbin/bhyveload"
const hypervisor = "/usr/sbin/bhyve"

func main() {

	myVm := &vm.VmState{
		Loader:     loader,
		Hypervisor: hypervisor,
	}
	if err := myVm.InitVM(config); err != nil {
		log.Fatal(err)
	}
	if err := myVm.CreateVM(); err != nil {
		log.Fatal(err)
	}
	// In real code there will be console redirect goroutine here
	for {
		time.Sleep(10)
	}
}
