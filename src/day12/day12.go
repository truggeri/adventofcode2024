package day12

import (
	"strings"
)

type point struct {
	x, y int
}

type fence struct {
	x, y      int
	direction rune
}

func (p point) surroundingFences() []fence {
	top := fence{x: p.x, y: p.y, direction: '-'}
	bottom := fence{x: p.x, y: p.y + 1, direction: '-'}
	left := fence{x: p.x, y: p.y, direction: '|'}
	right := fence{x: p.x + 1, y: p.y, direction: '|'}
	return []fence{top, bottom, left, right}
}

type patch struct {
	area      uint
	perimeter uint
}

func Solve(input string) uint {
	crops := parseInput(input)
	var cost uint = 0
	for _, crop := range crops {
		uniquePatches, fencesUsed := findUniquePatches(crop)
		details := patchDetails(uniquePatches, fencesUsed)
		for _, d := range details {
			cost += d.area * d.perimeter
		}
	}
	return cost
}

func parseInput(input string) map[rune][]point {
	crops := make(map[rune][]point)
	for j, line := range strings.Split(input, "\n") {
		for i, plant := range line {
			crops[plant] = append(crops[plant], point{x: i, y: j})
		}
	}
	return crops
}

func findUniquePatches(crop []point) (map[uint][]point, map[fence]uint) {
	var nextPatchId uint = 1
	patches := make(map[uint][]point)
	fencesByPatchId := make(map[fence]uint)
	for _, p := range crop {
		usedNewPatchId := true
		currentPatchId := nextPatchId
		newlyUsedFences := make([]fence, 0)

		for _, f := range p.surroundingFences() {
			if fencesByPatchId[f] == 0 {
				newlyUsedFences = append(newlyUsedFences, f)
				continue
			}

			if fencesByPatchId[f] != currentPatchId {
				patches[fencesByPatchId[f]] = append(patches[fencesByPatchId[f]], patches[currentPatchId]...)
				delete(patches, currentPatchId)
				currentPatchId = fencesByPatchId[f]
				usedNewPatchId = false
			}
			delete(fencesByPatchId, f)
		}

		for _, f := range newlyUsedFences {
			fencesByPatchId[f] = currentPatchId
		}

		patches[currentPatchId] = append(patches[currentPatchId], p)

		if usedNewPatchId {
			nextPatchId++
		}
	}

	return patches, fencesByPatchId
}

func patchDetails(patches map[uint][]point, fencesUsed map[fence]uint) []patch {
	details := make([]patch, len(patches))
	i := 0
	for _, points := range patches {
		details[i].area = uint(len(points))
		for _, p := range points {
			for _, f := range p.surroundingFences() {
				if fencesUsed[f] > 0 {
					details[i].perimeter++
				}
			}
		}
		i++
	}
	return details
}
