// go package bs_fsbndm
//
// The MIT License (MIT)
// Copyright (c) 2016 Andreas Briese, eduToolbox@Bri-C GmbH, Sarstedt

// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

/*
 * 'nos esse quasi nanos gigantum umeris insidentes' (Bernhard von Chartres, 1120)
 * The giants in this respect:
 * This is a modification of the Forward Semplified BNDM algorithm published by
 * S. Faro and T. Lecroq (2008):
 * Efficient Variants of the Backward-Oracle-Matching Algorithm.
 * Proceedings of the Prague Stringology Conference 2008, pp.146--160, Czech Technical University in Prague, Czech Republic, (2008).
 * Lizence of the authors C-implementation: GNU General Public License V.3 as published by the Free Software Foundation
 *
 * Modifications: the longpat variant uses the prefix instead of the suffix; 64bit-implementation
 */

package bs_fsbndm

import (
	"errors"
)

// Alphabet & Errors
var (
	ALPHABET    = 256
	NEEDLESHORT = errors.New("Length needle is < 2")
	NEEDLELONG  = errors.New("Length needle > length haystack")
)

func Index(haystack, needle *[]byte) (int, error) {

	// check length needle
	if len(*haystack) < len(*needle) {
		return -1, NEEDLELONG
	}
	if len(*needle) < 2 {
		return -1, NEEDLESHORT
	}

	return findFI(haystack, needle), nil
}

func Count(haystack, needle *[]byte) (int, error) {

	// check length needle
	if len(*haystack) < len(*needle) {
		return -1, NEEDLELONG
	}
	if len(*needle) < 2 {
		return -1, NEEDLESHORT
	}

	return count(haystack, needle), nil
}

func FindAll(haystack, needle *[]byte) (found []int, e error) {

	// check length needle
	if len(*haystack) < len(*needle) {
		return found, NEEDLELONG
	}
	if len(*needle) < 2 {
		return found, NEEDLESHORT
	}

	return findALL(haystack, needle), nil
}
