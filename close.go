// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

func (u Ubictrl) Close() (err error) {
	return u.f.Close()
}
