// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !linux
// +build !darwin
// +build !dragonfly
// +build !freebsd
// +build !netbsd
// +build !solaris

package runtime

func sysargs(argc int32, argv **byte) {
}
