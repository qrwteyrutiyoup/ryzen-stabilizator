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

package c6

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"runtime"
)

// ryzenC6MSR stores the offset and target bit of a given feature. MSR stands
// for model-specific register.
type ryzenC6MSR struct {
	offset int64
	bit    uint64
}

var (
	// msr has info for core and package C6, a C-state (idle power saving
	// state). Magic numbers for the MSR obtained from ZenStates-Linux project
	// available at https://github.com/r4m0n/ZenStates-Linux.
	msr = []ryzenC6MSR{
		// C6 package.
		{0xC0010292, 1 << 32},
		// C6 core.
		{0xC0010296, (1 << 22) | (1 << 14) | (1 << 6)},
	}
)

// readMSR reads the MSR of a given CPU at a given offset.
func readMSR(offset int64, cpu int) (uint64, error) {
	fname := fmt.Sprintf("/dev/cpu/%d/msr", cpu)
	f, err := os.OpenFile(fname, os.O_RDONLY, 0666)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	f.Seek(offset, io.SeekStart)
	data := make([]byte, 8)
	if _, err = f.Read(data); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(data), nil
}

// writeMSR writes a value to a specific CPU MSR at a given offset.
func writeMSR(offset int64, cpu int, value uint64) error {
	fname := fmt.Sprintf("/dev/cpu/%d/msr", cpu)
	f, err := os.OpenFile(fname, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Seek(offset, io.SeekStart)
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, value)
	_, err = f.Write(data)
	return err
}

// changeC6 eithers enable or disable the C6 (both core and package) C-state,
// depending on whether the provided parameter is true or false, respectively.
func changeC6(enable bool) error {
	cpus := runtime.NumCPU()
	for _, m := range msr {
		value := m.bit
		if !enable {
			value = ^(m.bit)
		}
		for c := 0; c < cpus; c++ {
			if err := writeMSR(m.offset, c, value); err != nil {
				return err
			}
		}
	}
	return nil
}

// c6Enable returns true or false depending on whether C6 C-state is enabled or
// disabled, respectively.
func c6Enabled() (bool, error) {
	cpus := runtime.NumCPU()
	for c := 0; c < cpus; c++ {
		for _, m := range msr {
			data, err := readMSR(m.offset, c)
			if err != nil {
				return false, err
			}
			if data&(m.bit) == m.bit {
				return true, nil
			}
		}
	}
	return false, nil
}

// Enable enables C6 C-state.
func Enable() error {
	// Passing true to indicate we want C6 enabled.
	return changeC6(true)
}

// Disable disables C6 C-state.
func Disable() error {
	// Passing false to indicate we want C6 disabled.
	return changeC6(false)
}

// Enabled returns true if C6 C-state is enabled.
func Enabled() (bool, error) {
	return c6Enabled()
}

// Disabled returns true if C6 C-state is disabled.
func Disabled() (bool, error) {
	enabled, err := c6Enabled()
	if err != nil {
		return false, err
	}
	return !enabled, nil
}
