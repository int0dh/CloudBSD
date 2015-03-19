# CloudBSD

The goal for this project is provide the bhyve and jails
virtualization combined with ZFS atop of minimal FreeBSD 
system. The system is going to consist of kernel, minimal
set of system libraries and service utilities plus middleware. 
The middleware is to

1) Create VM instance
2) Destroy VM instance
3) Move VM instance to another computing node.
4) Monitor the VM instances and take pre-configured actions
in depend on certain conditions, like VM faults or hangs.
5) Backup and replicate the VM storage.
6) Provide REST api for all above.

The ZFS ZVOLs and iSCSI targets will be used as a VM storage.
The VALE will be used as an inter-VM ethernet switch.
