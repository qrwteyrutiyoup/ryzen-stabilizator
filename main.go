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
	"os"

	"github.com/qrwteyrutiyoup/ryzen-stabilizator/boosting"
	"github.com/qrwteyrutiyoup/ryzen-stabilizator/c6"
)

const (
	program   = "Ryzen Stabilizator Tabajara"
	copyright = "Copyright (C) 2018 Sergio Correia <sergio@correia.cc>"
)

var (
	version = "unspecified/git version"
)

func main() {
	fmt.Printf("%s %s\n%s\n\n", program, version, copyright)
	// Check if we are running as root.
	if os.Geteuid() != 0 {
		fmt.Printf("You need to be root to use this program.\n")
		return
	}

	enablePtr := flag.Bool("enable-c6", false, "Enable C6 C-state")
	disablePtr := flag.Bool("disable-c6", false, "Disable C6 C-state")
	enableBoosting := flag.Bool("enable-boosting", false, "Enable processor boosting")
	disableBoosting := flag.Bool("disable-boosting", false, "Disable processor boosting")
	flag.Parse()

	// C6.
	switch {
	case *disablePtr:
		fmt.Printf("Disabling C6 C-state:   ")
		err := c6.Disable()
		if err != nil {
			fmt.Printf("oops: %v\n\n", err)
			break
		}
		fmt.Printf("SUCCESS\n\n")
	case *enablePtr:
		fmt.Printf("Enabling C6 C-state:   ")
		err := c6.Enable()
		if err != nil {
			fmt.Printf("oops: %v\n\n", err)
			break
		}
		fmt.Printf("SUCCESS\n\n")
	}

	// Boosting.
	switch {
	case *disableBoosting:
		fmt.Printf("Disabling processor boosting:   ")
		err := boosting.Disable()
		if err != nil {
			fmt.Printf("oops: %v\n\n", err)
			break
		}
		fmt.Printf("SUCCESS\n\n")
	case *enableBoosting:
		fmt.Printf("Enabling processor boosting:   ")
		err := boosting.Enable()
		if err != nil {
			fmt.Printf("oops: %v\n\n", err)
			break
		}
		fmt.Printf("SUCCESS\n\n")
	}

	// Current status of both C6 C-state and processor boosting.
	c6Status := "C6 C-state is DISABLED."
	c6Enabled, err := c6.Enabled()
	if err == nil {
		if c6Enabled {
			c6Status = "C6 C-state is ENABLED."
		}
	} else {
		c6Status = fmt.Sprintf("Error while obtaining status of C6 C-state: %v", err)
	}
	fmt.Println(c6Status)

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
