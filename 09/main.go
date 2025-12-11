package main

import (
	"bufio"
	"cmp"
	"os"
	"slices"
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
	x, y int
}

func PtFromLine(line string) Pt {
	parts := strings.Split(line, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return Pt{x, y}
}

type Interval struct {
	start, end int
}

func (iv Interval) Intersects(other Interval) bool {
	return min(iv.end, other.end) >= max(iv.start, other.start)
}

func (iv Interval) Sub(other Interval) []Interval {
	if !iv.Intersects(other) {
		return []Interval{iv}
	}

	common := iv.Common(other)
	if common.start <= iv.start && common.end >= iv.end {
		return nil
	}

	ret := make([]Interval, 0, 2)
	if common.start > iv.start {
		ret = append(ret, Interval{iv.start, common.start - 1})
	}
	if common.end < iv.end {
		ret = append(ret, Interval{common.end + 1, iv.end})
	}
	return ret
}

func (iv Interval) Common(other Interval) Interval {
	return Interval{max(iv.start, other.start), min(iv.end, other.end)}
}

type Line struct {
	interval Interval
	fixed    int
	dir      Dir
}

func LineFromPoints(p1, p2 Pt) Line {
	var interval Interval
	var fixed int
	var dir Dir

	if p1.x == p2.x {
		interval = Interval{min(p1.y, p2.y), max(p1.y, p2.y)}
		fixed = p1.x
		if p2.y > p1.y {
			dir = UP
		} else {
			dir = DOWN
		}
	} else {
		interval = Interval{min(p1.x, p2.x), max(p1.x, p2.x)}
		fixed = p1.y
		if p2.x > p1.x {
			dir = RIGHT
		} else {
			dir = LEFT
		}
	}
	return Line{interval, fixed, dir}
}

type Rect struct {
	orig, end Pt
	hIval     Interval
	vIval     Interval
}

func RectFromCorners(c1, c2 Pt) Rect {
	return Rect{
		orig:  Pt{min(c1.x, c2.x), min(c1.y, c2.y)},
		end:   Pt{max(c1.x, c2.x), max(c1.y, c2.y)},
		hIval: Interval{min(c1.x, c2.x), max(c1.x, c2.x)},
		vIval: Interval{min(c1.y, c2.y), max(c1.y, c2.y)},
	}
}

func (r Rect) Intersects(other Rect) bool {
	return r.hIval.Intersects(other.hIval) && r.vIval.Intersects(other.vIval)
}

func (r Rect) Area() int {
	return (r.end.x - r.orig.x + 1) * (r.end.y - r.orig.y + 1)
}

func main() {
	file, _ := os.Open("inputs/09.in")

	points := make([]Pt, 0, 100)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		points = append(points, PtFromLine(scanner.Text()))
	}

	// Pre-process: normalize to 0 using the minimum coordinates
	minX, minY := points[0].x, points[0].y
	for _, p := range points {
		minX, minY = min(minX, p.x), min(minY, p.y)
	}
	maxX, maxY := 0, 0
	for i := range points {
		points[i].x -= minX
		points[i].y -= minY
		maxX, maxY = max(maxX, points[i].x), max(maxY, points[i].y)
	}

	lines := make([]Line, 0, len(points)-1)
	for i := 0; i < len(points)-1; i++ {
		lines = append(lines, LineFromPoints(points[i], points[i+1]))
	}

	// Pre-index horizontal lines by Y coordinate for fast lookup
	linesByY := make(map[int][]Line)
	for _, line := range filteredLines(lines, map[Dir]bool{LEFT: true, RIGHT: true}) {
		y := line.fixed
		linesByY[y] = append(linesByY[y], line)
	}

	outsideIvals := make([][]Interval, maxY+1)
	for i := range outsideIvals {
		outsideIvals[i] = []Interval{{0, maxX}}
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
			queue := []Interval{line.interval}
			oppositeIdx := 0

			for len(queue) > 0 {
				currentIval := queue[0]

				oppLine := oppositeLines[oppositeIdx]
				oppIval := oppLine.interval
				if oppLine.fixed <= line.fixed || !currentIval.Intersects(oppIval) {
					oppositeIdx++
					continue
				}

				queue = append(queue[1:], currentIval.Sub(oppIval)...)
				overlap := currentIval.Common(oppIval)
				toRemove := Interval{line.fixed, oppLine.fixed}

				for y := overlap.start; y <= overlap.end; y++ {
					updated := make([]Interval, 0, len(outsideIvals[y])*2)
					for _, outer := range outsideIvals[y] {
						updated = append(updated, outer.Sub(toRemove)...)
					}
					outsideIvals[y] = updated
				}
			}
		}

		// Subtract horizontal lines from the outside intervals. If these lines (BC) form a corner with
		// the _opposite_ line we're casting to (AB), they'll be shadowed, but points along these lines
		// should still be considered "in" (so the whole AC interval is "in").
		//
		//                │
		//      │ ◀────── │
		//      ▼ ◀────── ▲
		//      │ ◀────── │
		// ◀────┘ ◀────── │
		// C     B      A │
		//
		for y := 0; y <= maxY; y++ {
			if lines, ok := linesByY[y]; ok {
				for _, line := range lines {
					updated := make([]Interval, 0, len(outsideIvals[y])*2)
					for _, outer := range outsideIvals[y] {
						updated = append(updated, outer.Sub(line.interval)...)
					}
					outsideIvals[y] = updated
				}
			}
		}

		// Consolidate outside intervals into rectangles - this reduces our predicates from ~90k to ~2k
		outsideRects := make([]Rect, 0, maxY)
		for startY, endY := 0, 0; startY <= maxY && endY <= maxY; {
			baseIvals := outsideIvals[startY]
			nextIvals := outsideIvals[endY]
			if equalIntervalSlices(nextIvals, baseIvals) {
				endY++
				continue
			}

			for _, ival := range baseIvals {
				outsideRects = append(outsideRects, RectFromCorners(Pt{ival.start, startY}, Pt{ival.end, endY - 1}))
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

func equalIntervalSlices(a, b []Interval) bool {
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
	return (d + 2) % 4
}

func filteredLines(lines []Line, dirs map[Dir]bool) []Line {
	filtered := make([]Line, 0, len(lines))

	for _, l := range lines {
		if dirs[l.dir] {
			filtered = append(filtered, l)
		}
	}
	return filtered
}

func sortedLines(lines []Line, dirs map[Dir]bool) []Line {
	filtered := filteredLines(lines, dirs)

	slices.SortFunc(filtered, func(l1, l2 Line) int {
		return cmp.Compare(l1.fixed, l2.fixed)
	})
	return filtered
}
