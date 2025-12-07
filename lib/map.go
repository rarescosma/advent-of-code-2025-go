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
	numCols := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		lineLen := len(line)
		numCols = max(numCols, lineLen)
		row := make([]byte, lineLen)
		copy(row, line)
		buf = append(buf, row)
	}
	return Map[byte]{Buf: buf, NumRows: len(buf), NumCols: numCols}
}

func (m *Map[C]) Get(r, c int) C {
	return m.Buf[r][c]
}

func (m *Map[C]) Set(r, c int, b C) {
	m.Buf[r][c] = b
}

func (m *Map[C]) EqualizeRows(filler C) {
	for r := range m.NumRows {
		for range m.NumCols - len(m.Buf[r]) {
			m.Buf[r] = append(m.Buf[r], filler)
		}
	}
}

func (m *Map[C]) GetRow(r int) []C {
	return m.Buf[r]
}

func (m *Map[C]) Transpose() {
	var buf [][]C
	for c := range m.NumCols {
		buf = append(buf, make([]C, 0))
		for r := range m.NumRows {
			buf[c] = append(buf[c], m.Get(r, c))
		}
	}
	*m = Map[C]{Buf: buf, NumRows: m.NumCols, NumCols: m.NumRows}
}
