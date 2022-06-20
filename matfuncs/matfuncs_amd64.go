// Copyright 2022 The NLP Odyssey Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 && gc && !purego

package matfuncs

import "github.com/nlpodyssey/spago-mat-asm/matfuncs/cpu"

var hasAVX = cpu.X86.HasAVX
