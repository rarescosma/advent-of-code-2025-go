package lib

import "bufio"

type Pos struct {
	Row int
	Col int
}

type Cell interface {
	byte | int | uint64
}

type Map[c Cell] struct {
	Buf     [][]c
	NumRows int
	NumCols int
}

func NewByteMap(scanner *bufio.Scanner) Map[byte] {
	var buf [][]byte
	numRows := 0
	numCols := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		buf = append(buf, make([]byte, 0))
		buf[numRows] = append(buf[numRows], line...)
		numCols = max(numCols, len(buf[numRows]))
		numRows++
	}
	return Map[byte]{Buf: buf, NumRows: numRows, NumCols: numCols}
}

func (m *Map[C]) Get(p Pos) C {
	return m.Buf[p.Row][p.Col]
}

func (m *Map[C]) Set(p Pos, b C) {
	m.Buf[p.Row][p.Col] = b
}

func (m *Map[C]) Append(row int, b C) {
	m.Buf[row] = append(m.Buf[row], b)
}

func (m *Map[C]) GetRow(r int) []C {
	return m.Buf[r]
}

func (m *Map[C]) Transpose() Map[C] {
	var buf [][]C
	for c := range m.NumCols {
		buf = append(buf, make([]C, 0))
		for r := range m.NumRows {
			buf[c] = append(buf[c], m.Get(Pos{Row: r, Col: c}))
		}
	}
	return Map[C]{Buf: buf, NumRows: m.NumCols, NumCols: m.NumRows}
}
