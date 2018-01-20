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
	flag.Parse()

	switch {
	case *disablePtr:
		fmt.Printf("Disabling C6 C-state:   ")
		err := c6.Disable()
		if err != nil {
			fmt.Printf("oops: %v\n", err)
			break
		}
		fmt.Printf("SUCCESS\n")
	case *enablePtr:
		fmt.Printf("Enabling C6 C-state:   ")
		err := c6.Enable()
		if err != nil {
			fmt.Printf("oops: %v\n", err)
			break
		}
		fmt.Printf("SUCCESS\n")
	default:
		enabled, err := c6.Enabled()
		if err != nil {
			fmt.Printf("Error while obtaining status of C6 C-state: %v\n", err)
		}
		status := "C6 C-state is DISABLED."
		if enabled {
			status = "C6 C-state is ENABLED."
		}
		fmt.Println(status)
	}
}
