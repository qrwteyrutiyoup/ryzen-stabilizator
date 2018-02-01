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

package boosting

import (
	"io/ioutil"
	"os"
	"strings"
)

const (
	boostingControlFile = "/sys/devices/system/cpu/cpufreq/boost"
)

// changeProcessorBoosting receives a parameter indicating whether it should
// enable or disable processor boosting.
func changeProcessorBoosting(enable bool) error {
	value := []byte("0")
	if enable {
		value = []byte("1")
	}
	return ioutil.WriteFile(boostingControlFile, value, 0644)
}

// Available returns a boolean indicating whether we have boosting control
// available or not. Disabling AMD Cool'n'Quiet, for instance, prevents cpufreq
// module from loading, which in turn, makes boosting control unavailable.
func Available() bool {
	if _, err := os.Stat(boostingControlFile); err == nil {
		return true
	}
	return false
}

// Enabled returns a boolean indicating whether processor boosting is enabled
// or not.
func Enabled() (bool, error) {
	value, err := ioutil.ReadFile(boostingControlFile)
	if err != nil {
		return false, err
	}

	enabled := true
	if strings.Trim(string(value), "\n") == "0" {
		enabled = false
	}
	return enabled, nil
}

// Disabled returns a boolean indicating whether processor boosting is
// disabled.
func Disabled() (bool, error) {
	enabled, err := Enabled()
	if err != nil {
		return false, err
	}
	return !enabled, nil
}

// Enable enables processor boosting.
func Enable() error {
	// We pass `true' to enable boosting.
	return changeProcessorBoosting(true)
}

// Disable disables processor boosting.
func Disable() error {
	// We pass `false' to disable boosting.
	return changeProcessorBoosting(false)
}
