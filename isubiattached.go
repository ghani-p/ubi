// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

import (
	"fmt"
	"os"
	"strconv"
)

// IsUbiAttached() is used to determine whether the specified UBI device
// is attached.
func IsUbiAttached(ubiNum int32) (bool, error) {
	dev := "/sys/devices/virtual/ubi/ubi" + strconv.Itoa(int(ubiNum))
	_, err := os.Stat(dev)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("Unable to stat %s: %s", dev, err)
	}
	return true, nil
}
