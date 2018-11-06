package main

import "github.com/SirMetathyst/go-blackboard"

const (
	KeyFile            string = "File"
	KeyOutputDirectory string = "OutputDirectory"
	KeyWatchFile       string = "WatchFile"
	KeyWatchFileEnable string = "WatchFileEnable"
)

// SetupBlackboard ...
func SetupBlackboard(bb *blackboard.BB) error {
	SetFile(bb, "proton.proton")
	SetOutputDirectory(bb, "Generated")
	SetWatchFile(bb, []string{})
	SetWatchFileEnable(bb, false)
	return nil
}

// SetFile ...
func SetFile(bb *blackboard.BB, v string) {
	bb.SetString(KeyFile, v)
}

// File ...
func File(bb *blackboard.BB) string {
	return *bb.StringP(KeyFile)
}

// SetOutputDirectory ...
func SetOutputDirectory(bb *blackboard.BB, v string) {
	bb.SetString(KeyOutputDirectory, v)
}

// OutputDirectory ...
func OutputDirectory(bb *blackboard.BB) string {
	return *bb.StringP(KeyOutputDirectory)
}

// SetWatchFile ...
func SetWatchFile(bb *blackboard.BB, v []string) {
	bb.SetStringSlice(KeyWatchFile, v)
}

// WatchFile ...
func WatchFile(bb *blackboard.BB) []string {
	return *bb.StringSliceP(KeyWatchFile)
}

// SetWatchFileEnable ...
func SetWatchFileEnable(bb *blackboard.BB, v bool) {
	bb.SetBool(KeyWatchFileEnable, v)
}

// WatchFileEnable ...
func WatchFileEnable(bb *blackboard.BB) bool {
	return *bb.BoolP(KeyWatchFileEnable)
}
