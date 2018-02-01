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

package aslr

import (
	"io/ioutil"
	"strings"
)

const (
	aslrControlFile = "/proc/sys/kernel/randomize_va_space"
)

// changeASLR receives a parameter indicating whether it should enable or
// disable address space layout randomization (ASLR).
func changeASLR(enable bool) error {
	// Following info from https://askubuntu.com/a/318476.
	// 0 - no randomization. Everything is static.
	// 1 - conservative randomization. Shared libraries, stack, mmap(), VDSO and
	//     heap are randomized.
	// 2 - full randomization. In addition to elements listed in the previous
	//     point, memory managed through brk() is also randomized.
	// We enable by setting full randomization (2), and disable with no
	// randomization (0).
	value := []byte("0")
	if enable {
		value = []byte("2")
	}
	return ioutil.WriteFile(aslrControlFile, value, 0644)
}

// Enabled returns a boolean indicating whether ASLR is enabled or not.
func Enabled() (bool, error) {
	value, err := ioutil.ReadFile(aslrControlFile)
	if err != nil {
		return false, err
	}

	enabled := true
	if strings.Trim(string(value), "\n") == "0" {
		enabled = false
	}
	return enabled, nil
}

// Disabled returns a boolean indicating whether ASLR is disabled.
func Disabled() (bool, error) {
	enabled, err := Enabled()
	if err != nil {
		return false, err
	}
	return !enabled, nil
}

// Enable enables ASLR.
func Enable() error {
	// We pass `true' to enable ASLR.
	return changeASLR(true)
}

// Disable disables ASLR.
func Disable() error {
	// We pass `false' to disable ASLR.
	return changeASLR(false)
}
