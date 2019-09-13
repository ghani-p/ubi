// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isMounted looks to see if the specified device is mounted anywhere
func isMounted(dev string) (bool, error) {
	f, err := os.Open("/proc/mounts")
	if err != nil {
		return false, fmt.Errorf("Unable to open /proc/mounts: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if fields[0] == dev {
			return true, nil
		}
	}
	return false, scanner.Err()
}

// IsUbiMounted() is used to determine whether the specified UBI device
// is mounted.
func IsUbiMounted(ubiNum int32, ubiVol int32) (bool, error) {
	return isMounted("/dev/ubi" + strconv.Itoa(int(ubiNum)) + "_" +
		strconv.Itoa(int(ubiVol)))
}
