// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// IsUbi is used to determine whether the specified MTD device contains
// a UBI volume. We read the first 4 bytes of the specified MTD device
// and check for the magic header UBI#.
func IsUbi(mtdNum int32) (u bool, err error) {
	dev := "/dev/mtd" + strconv.Itoa(int(mtdNum))
	m, err := os.OpenFile(dev, os.O_RDWR, 0666)
	if err != nil {
		return false, fmt.Errorf("Unable to open %s: %s", dev, err)
	}
	defer m.Close()
	buf := make([]byte, 4)
	_, err = io.ReadAtLeast(m, buf, 4)
	if err != nil {
		return false, fmt.Errorf("Error reading %s: %s", dev, err)
	}
	return buf[0] == 'U' && buf[1] == 'B' &&
		buf[2] == 'I' && buf[3] == '#', nil
}
