// Copyright 2018 Sergio Correia <sergio@correia.cc>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/klauspost/cpuid"
	"github.com/qrwteyrutiyoup/ryzen-stabilizator/boosting"
	"github.com/qrwteyrutiyoup/ryzen-stabilizator/c6"
)

const (
	program   = "Ryzen Stabilizator Tabajara"
	copyright = "Copyright (C) 2018 Sergio Correia <sergio@correia.cc>"

	// The family number for Zen processors.
	amdZenFamily = 0x17
)

var (
	version = "unspecified/git version"
)

// rsProfile represents a single profile, in which we may have definitions for
// both C6 C-state and processor boosting. Both of these parameters are string
// and accept values `enabled' and `disabled'.
type rsProfile struct {
	C6       string `toml:"c6"`
	Boosting string `toml:"boosting"`
}

// rsConfig represents the config file for Ryzen Stabilizator, with its two
// profiles, `Boot' and `Resume'.
type rsConfig struct {
	Boot   rsProfile `toml:"boot"`
	Resume rsProfile `toml:"resume"`
}

// sanityCheck performs a few checks to be sure we should be running this
// program.
func sanityCheck() error {
	switch {
	// Check if we are running on an AMD processor.
	case cpuid.CPU.VendorID != cpuid.AMD:
		return fmt.Errorf("this is not an AMD processor")
	// Check if it is the right family, 17h (Zen).
	case cpuid.CPU.Family != amdZenFamily:
		return fmt.Errorf("wrong family of AMD processors; expected 23 (17h), got %d", cpuid.CPU.Family)
	// Check if we are running as root.
	case os.Geteuid() != 0:
		return fmt.Errorf("you need to be root to use this program")
	}
	return nil
}

// disableC6 disables C6 C-state.
func disableC6() {
	fmt.Printf("Disabling C6 C-state:   ")
	err := c6.Disable()
	if err != nil {
		fmt.Printf("oops: %v\n", err)
		return
	}
	fmt.Println("SUCCESS")
}

// enableC6 enables C6 C-state.
func enableC6() {
	fmt.Printf("Enabling C6 C-state:   ")
	err := c6.Enable()
	if err != nil {
		fmt.Printf("oops: %v\n", err)
		return
	}
	fmt.Println("SUCCESS")
}

// disableBoosting disables processor boosting.
func disableBoosting() {
	fmt.Printf("Disabling processor boosting:   ")
	err := boosting.Disable()
	if err != nil {
		fmt.Printf("oops: %v\n", err)
		return
	}
	fmt.Println("SUCCESS")
}

// enableBoosting enables processor boosting.
func enableBoosting() {
	fmt.Printf("Enabling processor boosting:   ")
	err := boosting.Enable()
	if err != nil {
		fmt.Printf("oops: %v\n", err)
		return
	}
	fmt.Println("SUCCESS")
}

// showStatus displays the current status of both C6 C-state and processor
// boosting.
func showStatus() {
	c6Status := "C6 C-state is DISABLED."
	c6Enabled, err := c6.Enabled()
	if err == nil {
		if c6Enabled {
			c6Status = "C6 C-state is ENABLED."
		}
	} else {
		c6Status = fmt.Sprintf("Error while obtaining status of C6 C-state: %v", err)
	}
	fmt.Printf("\n%s\n", c6Status)

	boostingEnabled, err := boosting.Enabled()
	boostingStatus := "Processor boosting is DISABLED."
	if err == nil {
		if boostingEnabled {
			boostingStatus = "Processor boosting is ENABLED."
		}
	} else {
		boostingStatus = fmt.Sprintf("Error while obtaining status of processor boosting: %v", err)
	}
	fmt.Println(boostingStatus)
}

func handleConfigurationFile(configFile, profileName string) {
	config := rsConfig{}
	selectedProfile := &rsProfile{}

	switch strings.ToLower(profileName) {
	case "":
		fmt.Printf("Error: you need to specify a profile available in the provided configuration file.\nE.g.: %s --config=%q --profile=%q\n\n", os.Args[0], configFile, "boot")
		return
	case "boot":
		selectedProfile = &config.Boot
		break
	case "resume":
		selectedProfile = &config.Resume
		break
	default:
		fmt.Printf("Error: invalid profile %q; valid profiles are %q and %q\n", profileName, "boot", "resume")
		return
	}

	// Reading and parsing the configuration file provided.
	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error: unable to read contents of config file %q: %v.\n", configFile, err)
		return
	}

	if _, err = toml.Decode(string(buf), &config); err != nil {
		fmt.Printf("Error: problem parsing config file %q: %v.\n\n", configFile, err)
		return
	}

	// Now we perform the actions indicated by the specified profile.
	fmt.Printf("Config file: %q; profile: %q\n", configFile, profileName)
	switch strings.ToLower(selectedProfile.Boosting) {
	case "enable":
		enableBoosting()
	case "disable":
		disableBoosting()
	}
	switch strings.ToLower(selectedProfile.C6) {
	case "enable":
		enableC6()
	case "disable":
		disableC6()
	}
	// Current status of both C6 C-state and processor boosting.
	showStatus()
}

func main() {
	fmt.Printf("%s %s\n%s\n\n", program, version, copyright)

	err := sanityCheck()
	if err != nil {
		fmt.Printf("Error: %v.\n", err)
		return
	}

	configFilePtr := flag.String("config", "", "ryzen-stabilizator config file; requires a profile to be specified")
	profileNamePtr := flag.String("profile", "", fmt.Sprintf("profile from the provided config file; either %q or %q", "boot", "resume"))
	enableC6Ptr := flag.Bool("enable-c6", false, "Enable C6 C-state")
	disableC6Ptr := flag.Bool("disable-c6", false, "Disable C6 C-state")
	enableBoostingPtr := flag.Bool("enable-boosting", false, "Enable processor boosting")
	disableBoostingPtr := flag.Bool("disable-boosting", false, "Disable processor boosting")
	flag.Parse()

	// Handle config file with associated profile.
	if *configFilePtr != "" {
		handleConfigurationFile(*configFilePtr, *profileNamePtr)
		return
	}

	// Regular handling of command-line arguments, if we are not using config
	// file with predefined profiles.
	// C6.
	switch {
	case *disableC6Ptr:
		disableC6()
	case *enableC6Ptr:
		enableC6()
	}

	// Boosting.
	switch {
	case *disableBoostingPtr:
		enableBoosting()
	case *enableBoostingPtr:
		enableBoosting()
	}

	// Current status of both C6 C-state and processor boosting.
	showStatus()
}
