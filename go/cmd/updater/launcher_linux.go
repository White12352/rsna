package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var local_qt_theme bool

func Launcher() {
	log.Println("Running as launcher")
	wd, _ := filepath.Abs(".")
	nd := filepath.Join(wd, "./nekoray") // nekoray's dictory, not your:)

	_debug := flag.Bool("debug", false, "Debug mode")
	flag.Parse()

	cmd := exec.Command(nd, flag.Args()...)

	ld_env := "LD_LIBRARY_PATH=" + filepath.Join(wd, "./usr/lib")
	qt_plugin_env := "QT_PLUGIN_PATH=" + filepath.Join(wd, "./usr/plugins")

	// Check AppImage
	// _, ok := os.LookupEnv("APPIMAGE")
	// if ok {
	// 	enabledTUN := os.Getenv("NKR_APPIMAGE_TUN")
	// 	if enabledTUN == "1" || enabledTUN == "true" {
	// 		// cmd.Args = append(cmd.Args, "-flag_restart_tun_on")
	// 		cmd = exec.Command("sudo", append(append([]string{nd}, flag.Args()...), "-flag_restart_tun_on")...)
	// 	}
	// }

	// Qt 5.12 abi is usually compatible with system Qt 5.15
	// But use package Qt 5.12 by default.
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "NKR_FROM_LAUNCHER=1")
	cmd.Env = append(cmd.Env, ld_env, qt_plugin_env)

	log.Println(ld_env, qt_plugin_env, cmd)

	if *_debug {
		cmd.Env = append(cmd.Env, "QT_DEBUG_PLUGINS=1")
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd.Start()
	}
}
