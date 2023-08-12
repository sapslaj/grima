package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/mitchellh/go-ps"
	"github.com/prometheus/procfs"
)

func despair(err any) {
	log.Print(err)
	if os.Getenv("PLEASE_BE_GENTLE") == "true" {
		panic("i've fucked up but you asked me to be gentle so i'm just gonna crash ok kthnxbye")
	}
	log.Print("I'VE FUCKED UP AND YOU'RE COMING DOWN WITH ME")
	cmds := [][]string{}
	cmds = append(cmds, []string{"systemctl", "poweroff"})
	cmds = append(cmds, []string{"shutdown", "-h", "now"})
	cmds = append(cmds, []string{"sudo", "shutdown", "-h", "now"})
	cmds = append(cmds, []string{"poweroff", "-f"})
	cmds = append(cmds, []string{"sudo", "poweroff", "-f"})
	cmds = append(cmds, []string{"shutdown", "/s", "/t", "0"})
	cmds = append(cmds, []string{"shutdown", "/s", "/f"})
	for _, cmd := range cmds {
		log.Print(cmd)
		exec.Command(cmd[0], cmd[1:]...).Run()
	}
	p, err := os.FindProcess(1)
	if err == nil {
		p.Kill()
		p.Signal(syscall.SIGTERM)
		p.Signal(syscall.SIGKILL)
	}
	panic("FUCK")
}

func kill(pid int) {
	log.Printf("CHROME DETECTED ON PID %d; EXTERMINATING OR DIE TRYING", pid)
	process, err := os.FindProcess(pid)
	if err != nil {
		despair(err)
	}
	if err = process.Kill(); err != nil {
		if err = process.Signal(syscall.SIGTERM); err != nil {
			if err = process.Signal(syscall.SIGKILL); err != nil {
				despair(err)
			}
		}
	}
}

func huntWindows() {
	// yeah i know this isn't as comprehensive as the *nix version i don't run
	// Windows sorry
	processes, err := ps.Processes()
	if err != nil {
		despair(err)
	}
	possibleExes := []string{
		"chrome.exe",
		"google-chrome.exe",
		"chromium.exe",
		"Discord.exe",
		"Code.exe",
	}
	for _, process := range processes {
		for _, possibility := range possibleExes {
			if strings.Contains(process.Executable(), possibility) {
				kill(process.Pid())
			}
		}
	}
}

func huntNix() {
	fs, err := procfs.NewDefaultFS()
	if err != nil {
		despair(err)
	}
	procs, err := fs.AllProcs()
	if err != nil {
		despair(err)
	}
	for _, proc := range procs {
		cmdLine, err := proc.CmdLine()
		if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
			despair(err)
		}
		for _, arg := range cmdLine {
			bannedStrings := []string{
				"chrome",
				"chromium",
				"electron",
			}
			for _, banned := range bannedStrings {
				if strings.Contains(arg, banned) {
					kill(proc.PID)
				}
			}
		}
	}
}

func main() {
	log.Print("starting the hunt")
	for {
		if runtime.GOOS == "windows" {
			huntWindows()
		} else {
			huntNix()
		}
		time.Sleep(5 * time.Second)
	}
}
