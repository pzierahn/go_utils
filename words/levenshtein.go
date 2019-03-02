package words

//
// Matchr: an approximate string matching library for the Go programming language
//
// Copyright (C) 2013-2014 Ant Zucaro
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301  USA
//
// You can contact Ant Zucaro at azucaro at gmail dot com.
//

//
// Levenshtein computes the Levenshtein distance between two
// strings. The returned value - distance - is the number of insertions,
// deletions, and substitutions it takes to transform one
// string (s1) into another (s2). Each step in the transformation "costs"
// one distance point.
//
func Levenshtein(s1 string, s2 string) (distance int) {

	r1 := []rune(s1)
	r2 := []rune(s2)

	rows := len(r1) + 1
	cols := len(r2) + 1

	var d1 int
	var d2 int
	var d3 int
	var i int
	var j int

	dist := make([]int, rows*cols)

	for i = 0; i < rows; i++ {
		dist[i*cols] = i
	}

	for j = 0; j < cols; j++ {
		dist[j] = j
	}

	for j = 1; j < cols; j++ {

		for i = 1; i < rows; i++ {

			if r1[i-1] == r2[j-1] {

				dist[(i*cols)+j] = dist[((i-1)*cols)+(j-1)]

			} else {

				d1 = dist[((i-1)*cols)+j] + 1
				d2 = dist[(i*cols)+(j-1)] + 1
				d3 = dist[((i-1)*cols)+(j-1)] + 1

				dist[(i*cols)+j] = min(d1, min(d2, d3))
			}
		}
	}

	distance = dist[(cols*rows)-1]

	return
}

func min(a int, b int) int {

	if a < b {
		return a
	} else {
		return b
	}
}
