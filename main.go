package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	version = "0.0.1"
)

var (
	fp  = kingpin.Arg("file", "Path of JSON Schema").File()
	op  = kingpin.Flag("output", "Path of Go struct file").Short('o').String()
	pkg = kingpin.Flag("package", "Package name for Go struct file").Default("main").Short('p').String()
)

func exec() error {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Version(version)
	kingpin.Parse()
	defer func() {
		(*fp).Close()
	}()
	var err error
	var con *Convertor
	if *fp != nil {
		con, err = Read(*fp)
	} else {
		info, err := os.Stdin.Stat()
		if err != nil {
			return fmt.Errorf("File does not exists: %s", err)
		}
		if info.Size() == 0 {
			kingpin.Usage()
			return err
		}
		con, err = Read(os.Stdin)
	}
	if err != nil {
		return fmt.Errorf("Cannot read: %s", err)
	}
	err = con.Extract()
	if err != nil {
		return fmt.Errorf("Cannot extract: %s", err)
	}
	con.SetPackage(*pkg)
	output := os.Stdout
	if *op != "" {
		output, err = os.Create(*op)
		if err != nil {
			return fmt.Errorf("Cannot create: %s", err)
		}
	}
	if err := con.Write(output); err != nil {
		return fmt.Errorf("Cannot write: %s", err)
	}
	return nil
}

func main() {
	if err := exec(); err != nil {
		kingpin.Errorf("%s", err)
		os.Exit(1)
	}
}
