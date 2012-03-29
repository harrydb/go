// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package grayscale

import (
	"image"
	"image/color"
	"image/draw"
)

const (
	NEIGHBOR4 = 0
	NEIGHBOR8 = 1
)

// A list of all points in one connected component.
type CoCo []image.Point

type cocoData struct {
	label         []int
	width, height int
	neighborhood  int
	roots         []int
}

// CoCos returns all connected components of the provided color.
//
// 4-connected or 8-connected neighborhoods can be specified using
// grayscale.NEIGHBOR4 and grayscale.NEIGHBOR8. The algorithm is implemented
// using a two-pass union/find approach.
func CoCos(m *image.Gray, color uint8, neighborhood int) []CoCo {
	var data cocoData
	data.width = m.Rect.Max.X - m.Rect.Min.X
	data.height = m.Rect.Max.Y - m.Rect.Min.Y
	data.label = make([]int, data.width*data.height)
	data.neighborhood = neighborhood
	data.passOne(m, color)
	return data.passTwo()
}

// CocoRemove
func CoCoRemove(m draw.Image, coco CoCo, c color.Color) {
	for _, p := range coco {
		m.Set(p.X, p.Y, c)
	}
}

func (data *cocoData) passTwo() []CoCo {
	numPixels := len(data.label)
	count := 0
	// renumber
	for _, a := range data.roots {
		root := data.label[a]
		if root == a {
			data.label[a] = count + numPixels
			count++
		} else {
			data.label[a] = data.label[root]
		}
	}
	cocos := make([]CoCo, count)
	for i := 0; i < numPixels; i++ {
		if data.label[i] == -1 {
			continue
		}
		root := data.find(i) - numPixels
		x := i % data.width
		y := i / data.width
		cocos[root] = append(cocos[root], image.Point{x, y})
	}
	return cocos
}

func (data *cocoData) passOne(m *image.Gray, color uint8) {
	neighbor := make([]int, 0, 4)
	for y := 0; y < data.height; y++ {
		offset := y * m.Stride
		row := m.Pix[offset : offset+data.width]

		for x := 0; x < data.width; x++ {
			pos := y*data.width + x
			if row[x] != color {
				data.label[pos] = -1
				continue
			}
			neighbor = neighbor[:0]
			neighbor = data.addNeighbor(x-1, y, neighbor)
			neighbor = data.addNeighbor(x, y-1, neighbor)
			if data.neighborhood == NEIGHBOR8 {
				neighbor = data.addNeighbor(x-1, y-1, neighbor)
				neighbor = data.addNeighbor(x+1, y-1, neighbor)
			}

			if len(neighbor) == 0 {
				data.label[pos] = pos
				data.roots = append(data.roots, pos)
				continue
			}

			minLabel := neighbor[0]
			for _, label := range neighbor {
				if label < minLabel {
					minLabel = label
				}
			}
			data.label[pos] = minLabel

			for _, label := range neighbor {
				if label == minLabel {
					continue
				}
				data.union(minLabel, label)
			}
		}
	}
}

func (data *cocoData) find(a int) int {
	if a >= len(data.label) || data.label[a] == a {
		return a
	}
	data.label[a] = data.find(data.label[a])
	return data.label[a]
}

func (data *cocoData) union(a, b int) {
	aRoot := data.find(a)
	bRoot := data.find(b)
	if aRoot == bRoot {
		return
	}
	if aRoot < bRoot {
		data.label[bRoot] = aRoot
	} else {
		data.label[aRoot] = bRoot
	}
}

func (data *cocoData) addNeighbor(x, y int, neighbor []int) []int {
	if x < 0 || y < 0 || x >= data.width || y >= data.height {
		return neighbor
	}
	label := data.label[y*data.width+x]
	if label == -1 {
		return neighbor
	}
	return append(neighbor, label)
}
