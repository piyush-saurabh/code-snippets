// Very simple implementation of a container
// Run: /usr/bin/sudo /usr/local/go/bin/go run main.go run /bin/bash

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

func main() {

	switch os.Args[1] {
	case "run":
		// Create a new namespace
		run()
	case "child":
		// Run the commands inside the newly created namespace
		child()
	default:
		panic("help")
	}

}

// Create & setup a new namespace
func run() {

	fmt.Printf("Running %v \n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2])...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Adding containerization
	// Create the new namespaces specified in this struct
	// Specify the flags for unshare system call ref. http://man7.org/linux/man-pages/man2/unshare.2.html
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	// fork and exec syscall
	must(cmd.Run())

}

func child() {

	fmt.Printf("Running %v \n", os.Args[2:])

	// Create a control group
	cg()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// [UTS namespace demo] Set the hostname inside the container
	must(syscall.Sethostname([]byte("container")))

	// // [PID namespace demo] Mount a file system inside the container and make it as the root using chroot
	// must(syscall.Chroot("/path/to/some/filesystem"))
	// must(syscall.Chdir("/"))

	// // Mount the proc so that we can us ps to see the running processes. Also unmount this at the end
	// must(syscall.Mount("proc", "proc", "proc", 0, ""))

	// [Mount namespace demo] Mount /dev/shm
	must(syscall.Mount("/dev/shm", "/tmp", "tmpfs", 0, ""))

	// fork and exec syscall
	must(cmd.Run())

	// //  [PID namespace demo] Unmount proc filesystem
	// must(syscall.Unmount("proc", 0))

	// [Mount namespace demo] Unmount /dev/shm
	must(syscall.Unmount("/dev/shm", 0))

}

// Setting up a control group
func cg() {
	cgroups := "/sys/fs/cgroup"

	// Assume that pids file system (of type cgroup) is mounted at /sys/fs/cgroup
	pids := filepath.Join(cgroups, "pids")

	// Create a new directory inside pids cgroup (this directory is kind of a docker container)
	os.Mkdir(filepath.Join(pids, "piyush"), 0755)

	// Limit the max number of process running inside cgroup to 20
	must(ioutil.WriteFile(filepath.Join(pids, "piyush/pid.max"), []byte("20"), 0700))

	// Remove the new cgroup after the container exits
	must(ioutil.WriteFile(filepath.Join(pids, "piyush/notify_on_release"), []byte("1"), 0700))

	// Write the PID of currently running process into cgroup.procs
	// This will attach a pid cgroup to our process
	// limit the number of process's spawned by our current process to 20
	must(ioutil.WriteFile(filepath.Join(pids, "piyush/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

// Wrapper function to catch the error messages during syscall
func must(err error) {
	if err != nil {
		panic(err)
	}
}
