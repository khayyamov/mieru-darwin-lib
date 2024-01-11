package Mieru

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"runtime"
)

func redirectStd(path string) error {
	stdin := path + "/stdin.log"
	if stats, err := os.Stat(stdin); err == nil && stats.Size() > 0 {
		_ = os.Rename(stdin, stdin+".old")
	}
	stdinFile, err := os.Create(stdin)
	if err != nil {
		return err
	}
	if runtime.GOOS != "android" {
		if os.Getuid() > 0 {
			err = stdinFile.Chown(os.Getuid(), os.Getgid())
			if err != nil {
				stdinFile.Close()
				os.Remove(stdinFile.Name())
				return err
			}
		}
	}
	err = unix.Dup2(int(stdinFile.Fd()), int(os.Stdin.Fd()))
	if err != nil {
		stdinFile.Close()
		os.Remove(stdinFile.Name())
		return err
	}

	stdout := path + "/stdout.log"
	if stats, err := os.Stat(stdout); err == nil && stats.Size() > 0 {
		_ = os.Rename(stdout, stdout+".old")
	}
	stdoutFile, err := os.Create(stdout)
	if err != nil {
		return err
	}
	if runtime.GOOS != "android" {
		if os.Getuid() > 0 {
			err = stdoutFile.Chown(os.Getuid(), os.Getgid())
			if err != nil {
				stdoutFile.Close()
				os.Remove(stdoutFile.Name())
				return err
			}
		}
	}
	err = unix.Dup2(int(stdoutFile.Fd()), int(os.Stdout.Fd()))
	if err != nil {
		stdoutFile.Close()
		os.Remove(stdoutFile.Name())
		return err
	}

	stderr := path + "/stderr.log"
	if stats, err := os.Stat(stderr); err == nil && stats.Size() > 0 {
		_ = os.Rename(stderr, stderr+".old")
	}
	stderrFile, err := os.Create(stderr)
	if err != nil {
		return err
	}
	if runtime.GOOS != "android" {
		if os.Getuid() > 0 {
			err = stderrFile.Chown(os.Getuid(), os.Getgid())
			if err != nil {
				stderrFile.Close()
				os.Remove(stderrFile.Name())
				return err
			}
		}
	}
	err = unix.Dup2(int(stderrFile.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		stderrFile.Close()
		os.Remove(stderrFile.Name())
		return err
	}
	return nil
}

//	func Run(config string, stdPath string) error {
//		print("kilo print")
//		fmt.Print("kilo fmt.Print")
//		err := redirectStd(stdPath)
//		if err != nil {
//			return err
//		}
//		err = cli.HandleRunFunc([]string{"", "", config, stdPath})
//		if err != nil {
//			return err
//		}
//		return nil
//	}
func Run() string {
	print("kilo print")
	fmt.Print("kilo fmt.Print")
	return "Hello from go"
}
