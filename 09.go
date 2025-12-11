package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Dir int

const (
	UP Dir = iota
	RIGHT
	DOWN
	LEFT
)

type Pt struct {
	X, Y int
}

func PtFromLine(line string) Pt {
	parts := strings.Split(line, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return Pt{X: x, Y: y}
}

type Ival struct {
	Start, End int
}

func (iv Ival) Intersects(other Ival) bool {
	return min(iv.End, other.End) >= max(iv.Start, other.Start)
}

func (iv Ival) Sub(other Ival) []Ival {
	if !iv.Intersects(other) {
		return []Ival{iv}
	}

	common := iv.Common(other)
	if common.Start <= iv.Start && common.End >= iv.End {
		return nil
	}

	ret := make([]Ival, 0, 2)
	if common.Start > iv.Start {
		ret = append(ret, Ival{iv.Start, common.Start - 1})
	}
	if common.End < iv.End {
		ret = append(ret, Ival{common.End + 1, iv.End})
	}
	return ret
}

func (iv Ival) Common(other Ival) Ival {
	return Ival{max(iv.Start, other.Start), min(iv.End, other.End)}
}

type Line struct {
	Start, End Pt
	Dir        Dir
	interval   Ival
}

func LineFromPoints(p1, p2 Pt) Line {
	var dir Dir
	var iv Ival
	if p1.X == p2.X {
		iv = Ival{min(p1.Y, p2.Y), max(p1.Y, p2.Y)}
		if p2.Y > p1.Y {
			dir = UP
		} else {
			dir = DOWN
		}
	} else {
		iv = Ival{min(p1.X, p2.X), max(p1.X, p2.X)}
		if p2.X > p1.X {
			dir = RIGHT
		} else {
			dir = LEFT
		}
	}
	return Line{Start: p1, End: p2, Dir: dir, interval: iv}
}

type Rect struct {
	Orig, End Pt
	hIval     Ival
	vIval     Ival
}

func RectFromCorners(c1, c2 Pt) Rect {
	return Rect{
		Orig:  Pt{min(c1.X, c2.X), min(c1.Y, c2.Y)},
		End:   Pt{max(c1.X, c2.X), max(c1.Y, c2.Y)},
		hIval: Ival{min(c1.X, c2.X), max(c1.X, c2.X)},
		vIval: Ival{min(c1.Y, c2.Y), max(c1.Y, c2.Y)},
	}
}

func (r Rect) Intersects(other Rect) bool {
	return r.hIval.Intersects(other.hIval) && r.vIval.Intersects(other.vIval)
}

func (r Rect) Area() int {
	return (r.End.X - r.Orig.X + 1) * (r.End.Y - r.Orig.Y + 1)
}

func main() {
	file, _ := os.Open("inputs/09.in")

	points := make([]Pt, 0, 100)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		points = append(points, PtFromLine(scanner.Text()))
	}

	// Pre-process: normalize to 0 using the minimum coordinates
	minX, minY := points[0].X, points[0].Y
	for _, p := range points {
		minX, minY = min(minX, p.X), min(minY, p.Y)
	}
	maxX, maxY := 0, 0
	for i := range points {
		points[i].X -= minX
		points[i].Y -= minY
		maxX, maxY = max(maxX, points[i].X), max(maxY, points[i].Y)
	}

	lines := make([]Line, 0, len(points)-1)
	for i := 0; i < len(points)-1; i++ {
		lines = append(lines, LineFromPoints(points[i], points[i+1]))
	}

	// Pre-index horizontal lines by Y coordinate for fast lookup
	linesByY := make(map[int][]Line)
	for _, line := range filteredLines(lines, map[Dir]bool{LEFT: true, RIGHT: true}) {
		y := line.Start.Y
		linesByY[y] = append(linesByY[y], line)
	}

	outsideIvals := make([][]Ival, maxY+1)
	for i := range outsideIvals {
		outsideIvals[i] = []Ival{{0, maxX}}
	}

	inner := func(startDir Dir) {
		startLines := sortedLines(lines, map[Dir]bool{startDir: true})
		oppositeLines := sortedLines(lines, map[Dir]bool{opposite(startDir): true})

		// Process vertical lines going in the `startDir` direction:
		//
		// For each line going in the `startDir` direction find the closest (wrt the non-changing dimension)
		// set of lines that go in the _opposite_ direction and "cover" the full the interval of the line.
		//
		// The intervals defined by casting rays from the processed line to these opposite lines are "inside",
		// so they can be subtracted from the "outside" intervals.
		for idx := range len(startLines) {
			line := startLines[idx]
			queue := []Ival{line.interval}
			oppositeIdx := 0

			for len(queue) > 0 {
				currentIval := queue[0]
				queue = queue[1:]

				oppLine := oppositeLines[oppositeIdx]
				oppIval := oppLine.interval
				if oppLine.Start.X >= line.Start.X || !currentIval.Intersects(oppIval) {
					oppositeIdx++
					queue = append(queue, currentIval)
					continue
				}

				remaining := currentIval.Sub(oppIval)
				queue = append(queue, remaining...)
				overlap := currentIval.Common(oppIval)
				toRemove := Ival{oppLine.Start.X, line.Start.X}

				for y := overlap.Start; y <= overlap.End; y++ {
					updated := make([]Ival, 0, len(outsideIvals[y])*2)
					for _, outer := range outsideIvals[y] {
						updated = append(updated, outer.Sub(toRemove)...)
					}
					outsideIvals[y] = updated
				}
			}
		}

		// Subtract horizontal lines from the outside intervals. If these lines form a corner with
		// the _opposite_ line we're casting to, they'll be shadowed, but points along these lines
		// should still be considered "in".
		for y := 0; y <= maxY; y++ {
			if lines, ok := linesByY[y]; ok {
				for _, line := range lines {
					updated := make([]Ival, 0, len(outsideIvals[y])*2)
					for _, outer := range outsideIvals[y] {
						updated = append(updated, outer.Sub(line.interval)...)
					}
					outsideIvals[y] = updated
				}
			}
		}

		// Consolidate outside intervals into rectangles - this reduces our predicates from ~90k to ~2k
		outsideRects := make([]Rect, 0, maxY)
		startY, endY := 0, 0
		for startY <= maxY && endY <= maxY {
			baseIvals := outsideIvals[startY]
			nextIvals := outsideIvals[endY]
			if equalIntervalSlices(nextIvals, baseIvals) {
				endY++
				continue
			}

			for _, ival := range baseIvals {
				outsideRects = append(outsideRects, RectFromCorners(Pt{ival.Start, startY}, Pt{ival.End, endY - 1}))
			}
			startY = endY
		}

		p1, p2 := 0, 0

		for i := 0; i < len(points)-1; i++ {
		out:
			for j := len(points) - 1; j > i; j-- {
				testRect := RectFromCorners(points[i], points[j])
				area := testRect.Area()
				if area > p1 {
					p1 = area
				}

				// Early exit if we already found a larger p2
				if area <= p2 {
					continue
				}

				for k := range outsideRects {
					if testRect.Intersects(outsideRects[k]) {
						continue out
					}
				}
				p2 = area
			}
		}

		println("p1:", p1)
		println("p2:", p2)
	}

	defer func() {
		if recover() != nil {
			inner(DOWN)
		}
	}()

	inner(UP)
}

func equalIntervalSlices(a, b []Ival) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func opposite(d Dir) Dir {
	return Dir((int(d) + 2) % 4)
}

func filteredLines(lines []Line, dirs map[Dir]bool) []Line {
	filtered := make([]Line, 0, len(lines))
	for _, l := range lines {
		if dirs[l.Dir] {
			filtered = append(filtered, l)
		}
	}
	return filtered
}

func sortedLines(lines []Line, dirs map[Dir]bool) []Line {
	filtered := filteredLines(lines, dirs)

	upDown := dirs[UP] || dirs[DOWN]
	sort.Slice(filtered, func(i, j int) bool {
		if upDown {
			return filtered[i].Start.X < filtered[j].Start.X
		}
		return filtered[i].Start.Y < filtered[j].Start.Y
	})
	return filtered
}
