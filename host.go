package main

import (
	"fmt"
	"os"
	"os/exec"
	"polydawn.net/siphon"
)

type hostOpts struct {
	Address string `short:"L" long:"addr" optional:"true" default:"defaults to unix://siphon.sock" description:"Address to bind to and await client attachings, of the form unix://path/to/socket" `
	Command string `short:"c" long:"command" optional:"true" default:"defaults to /bin/sh" description:"Command to execute inside the new psuedoterminal" `
}

func init() {
	parser.AddCommand("host", "host", "Host a process", &hostOpts{
		Address: "unix://siphon.sock",
		Command: "/bin/sh",
	})
}

func (opts *hostOpts) Execute(args []string) error {
	addr, err := ParseNewAddr(opts.Address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "siphon: %s\n", err)
		os.Exit(1)
	}
	cmd := exec.Command(opts.Command)

	fmt.Printf("Hosting %s at %s\n", opts.Command, addr.Label())

	host := siphon.NewHost(cmd, addr)

	host.Serve(); defer host.UnServe()
	host.Start()
	exitCode := host.Wait()
	fmt.Printf("siphon: %s exited %d\r\n", opts.Command, exitCode)

	return nil
}
