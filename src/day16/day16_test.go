package day16

import (
	"testing"
)

func TestSampleOne(t *testing.T) {
	input := `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`
	result := Solve(input)
	if result != 7036 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestSampleTwo(t *testing.T) {
	input := `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`
	result := Solve(input)
	if result != 11048 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestPart1(t *testing.T) {
	input := `#############################################################################################################################################
#.....#.......#.............#.........#.......................#.......#...#...............#.....#.#.................................#.#....E#
#.#####.###.###.#######.###.#.#.#######.#.###.#####.###.#######.#.###.###.#.#####.#####.###.#.#.#.#.#####.###.#.###.#######.###.#.#.#.#.#.#.#
#.......#...#...#.....#.#.#.#.#.#...#...#.....#...#.#...#.......#...#.#...#...#...#.#...#...............#...#...#.#.......#.#...#.#.#...#.#.#
#########.#.#.###.#####.#.#.#.#.#.#.#.#########.#.#.###.#.#########.#.#.#.#.###.#.#.#.###.###.#####.###.###.#.#.#.###.###.#.#.###.#.#.###.#.#
#.......#.#.#.#...#...#.#.#.#.#...#.....#.......#.#...#.#.#...#...#.#...#.#.#...#.....#...#.#...#...#.....#.#.....#.....#.#...#...#.#.#...#.#
#.#####.#.###.#.#.#.#.#.#.#.###########.#######.#.###.#.#.###.#.#.#.#.#####.#.###.###.#.###.###.#.###.#.###.#.#.#.#.#.#.#.#.#.#.###.###.#####
#...#...#.......#.#.#...#.#...........................#.#.#...#.#...#.#.....#.#.....#.#...#...#.....................#.#.#...................#
#.#.#.#######.#####.#####.###########.#######.#####.#.###.#.###.#######.#####.###.#.#.###.#.#.#.#.#######.###.#.#####.#.#.#.#.###.###.#.###.#
#.#...#.....#.#...#...#.......#.......#.....#...#...#.....#.#...#.....#...#...#.....#...#...#.#...#.......#...#.#...#.#.#.#.#...#.#...#...#.#
#.#.###.###.###.#.#.#.###.###.#.#.#####.###.###.#.#########.#.###.###.###.#.###.#.###.#######.#####.#######.#.#.#.#.#.###.###.#.#.#.#######.#
#.#.....#.......#.#.#...#.#...#.#...#...#.#.#.#.#.......#...#...#...#.......................#...#.#.#.....#.#.....#.#.........#...#.#.....#.#
#.###.###########.#####.#.#####.#.#.#.#.#.#.#.#.###.###.#.#.###.###.#########.###.#.#.#.###.#.#.#.#.#####.#.#.#.#####.#######.#####.#.###.#.#
#...#...................#.....#.#.#.#.#.#.....#.....#.#...#.#.#...#.....#...#.....#.#.#.#.#.....#.#.....#.#...#.#.........#...........#.....#
#.#.#.#.###.#.###############.#.#.#.#.#.#############.#####.#.###.#####.#.#.#####.#.#.#.#.#######.#####.#.###.#.#.###.###.#.#.###.###########
#.#.#.#.#.#.#.#.....#.......#.#...#...#...............#...#...#.#...#...#.#.......#.#.#.#.......#...#...#.....#.#.#.......#.#...#.#.........#
#.#.###.#.#.###.###.###.#.#.#.###.#.#####.###.###.###.#.#.###.#.#.###.###.###.###.#.#.#.###.###.###.#.###.#####.#.#.#.###.#.#.#.###.#######.#
#.#.....#.#...#.#.#...#.#.#.#...#.#.#...#...#...#...#...#...#...#.....#...#...#.......#.....#.#.....#...#.......#...#.#...#...#...#.#.#.....#
#.#######.###.#.#.###.#.#.#.#.#.#.###.#.#####.#.#.#.#########.#.#######.#.#.#.#.#############.#####.###.#########.#####.###.#####.#.#.#.#.###
#.#.......#.#.....#...#.#.#...#.#.#...#.....#.#.#.#.........#.#.........#.#.#.#.#.................#.#.#...#.......#.........#...#...#.#.....#
#.#.#####.#.#######.#####.#######.#.#####.#.###.#.#########.#.#.#########.#.#.#.###.#######.#.#####.#.###.#.#######.#.###.###.#.#####.###.#.#
#.#.....#.......#...#...#...#.....#.#...........#.#.......#...#.#...#.....#.#.#...#.#.......#.#...#...#...#.....#...#...#.#.................#
#.#####.#######.#.###.#.###.#.#####.#######.#####.#.#.###.#####.#.#.###.###.#.###.###.#####.###.#.###.#.#######.#.#####.#.#####.#.#.###.#.#.#
#.#...#.#.....#...#...#...#.#.......#.....#.#.....#.#...#.#...#.#.#...#.#...#.#.#...#.#...#.#...#...#.#.........#.....#.#.......#...#...#.#.#
#.#.#.#.#.###.#.###.#####.#.#.#####.#.###.###.#########.###.#.#.#.###.#.#.###.#.###.#.#.#.###.###.###.###############.#.###########.#.###.#.#
#.#.#.#.#.#.....#...#...#...#.....#...#.#...#.#...#...#.#...#.#...#...#.#.#.#.#.........#.#...#.#.....#.......#.......#.........#...#.#.....#
#.#.#.#.#.#######.###.#.#.#.#####.###.#.###.#.#.#.#.#.#.#.###.#####.#####.#.#.#########.#.#.###.#######.#######.###############.#.###.#.#.###
#.#.#...#.........#...#.#.#.....#...#...#.#.....#...#...#.#.#...#...#.....#.#.....#.....#...#.....#.....#.....#...#...#...#.....#...#.#.#...#
#.#.#########.#######.#.#.#.###.###.###.#.###########.###.#.###.#.###.#.###.#####.#.#########.#.#.###.#.#.###.#.#.#.#.#.#.#.#######.#.#.#.#.#
#...#...#...#.......#.#.#.#.#.....#.......#...#.....#.#...#...#.#...#...#.......#...#.....#...#.#.....#.#...#...#...#...#...#...............#
###.#.#.#.#.#######.#.#.#.#.#.#######.#####.#.#.###.###.###.###.#.#.###.#.#.#.#.#####.#.###.###.###.#######.#####.###########.#.###.#.#######
#...#.#...#...#...#...#.#.#.#.......#.....#.#.....#.....#...#...#.#...#...#.#...#.....#...#...#.#.........#.....#.........#...#.....#...#...#
#.#.#.#######.###.#####.#.#.#######.#####.#.#####.#######.#.#.#######.#.###.#.###.#.###.#.###.#.#.#######.#####.#.#######.#.#####.#####.#.#.#
#.#...#.....#...#.....#...#...#...#.....#...#...#.#...#...#...........#.....#.#...#...#.#.....#...#.#...#.#...#...#.....#.#.#.....#...#...#.#
#.#.###.###.###.#.###.#.#.###.#.###.#########.#.#.#.#.#.#####.#.#########.#.#.###.###.###########.#.#.#.#.#.#.###.###.#.#.#.#.###.#.#.#####.#
#.#.#...#.#...#.....#...#.#.#...#...#.........#.#...#.#...#...#.....#...#...#...#...#...........#...#.#.#...#...#...#.#...#.......#.#.......#
#.#.#.###.###.#####.#.###.#.###.#.###.#####.###.###.#.###.#.#.###.#.#.#.#######.#.#.###########.#####.#.#####.#.###.#.###########.#.#########
#.#.#.#.....#.#...#.#...#.#.#...#.....#...#.#...#.....#...#.#.....#...#...#...#.#.#.....#...#...#.....#.........#...#.......#.....#.........#
#.#.#.#####.#.###.#.#####.#.#.#########.###.#.#.#.###.#.###.###########.#.#.###.#.#####.#.#.#.###.#########.#####.#########.#####.#########.#
#.#.#.#.....#...#.#.....#.............#...#.#...#.#...#...#.......#.#...#.#.#...#.....#.#...#.#...#.........#.....#.............#...#.......#
###.#.#.#######.#.#####.#.#.#####.###.###.#.#.###.#.#####.#######.#.#.###.#.#.#########.###.#.#.###.#####.#.#.#####.#.#########.#.###.#.###.#
#...#.#.......#...#.....#.#.#.......#.....#.#.....#.#.......#.......#.#.#...#.........#...#.#.#.....#.....#.#...#...#.........#.#.#...#...#.#
#.#.#.#.#####.###.#.#####.###.###.#.###.###.#######.#.#####.#.#####.#.#.###.#########.###.#.#.#.#####.#####.###.#####.#######.#.#.#.###.#.#.#
#.#...#.....#...#.#.#.........#...#.#...#...#.....#.#.#.....#...#...#.#.....#.......#.....#.#.#.....#...#...#.#.....#.#.#.........#.#.....#.#
#.#.#.###.#####.#.#.###.#######.###.#.###.#.#.###.###.###.#####.#.###.#.#####.###.#.#######.#.#####.#.#.#.###.#####.#.#.#.#####.#.#.#.#.#.#.#
#.#.#.#...#...#...#...........#.#.....#...#.#...#.....#...#.....#...#.#.....#.#.#.#.#.......#.......#.#.#...#...#...#...#...#...#...#...#.#.#
#.###.###.#.#.#.#.###.#######.#.#######.###.###.#######.###.#######.#.#####.#.#.#.###.###.#############.###.#.#.#.#####.###.#.###.###.#.###.#
#...#...#.#.#...#.....#...#.....#.......#.....#...#.....#...#.....#.#...#...#...#.......#...#...........#.#...#.#.....#...#.#...#.#...#.....#
#.#.#.#.###.###.#.#####.#.#######.###############.###.#.#.###.#.#.#.###.#.#####.###.###.###.#.#######.#.#.#####.#####.#.###.###.#.#####.###.#
#.#.#.#.....#...#.#.....#...#...#.#.............#.....#.#...#.#.#.....#.#.#...#...#.......#...#.......#.....#...#...#...#...#...#...........#
###.#########.#.#.#.#######.#.#.#.#.#########.#########.###.###.###.###.###.#.###.###.#.#######.###########.#.###.#.#####.###.###.#.#.#.#.###
#...#.....#...#.#.........#...#.#...#...#.#...#...........#...#.......#.....#...#...#.#...#.......#.....#...#...#.#.....#.#.#.#.....#...#...#
#.###.###.#.#####.#######.#####.#.###.#.#.#.###.#############.#####.#.#########.###.#####.#.#.###.#####.#.#####.#.#######.#.#.#########.###.#
#.#...#.....#.....#...#...#...#.#.#...#...#.....#.............#...#.#...#.....#...#...#...#.....#.....#.......#.#.#.......#.#.#...........#.#
#.#.#######.#.#####.#.#.###.#.#.#.#.#.###########.#############.#.###.#.#.###.###.###.#.#.#.###.#####.#########.#.#.#######.#.#.#######.#.#.#
#.........#.#...#.#.#.#.#.#.#...#...#.............#.....#.......#.#...#.#.#.#...#.....#.#.#...#.......#.......#...#.#.......#.#.......#.#...#
#.#######.#.###.#.#.#.#.#.#.###.#####.#############.#.###.#.#####.#.###.#.#.###.#######.#####.#######.#.###.#.#.###.#.#####.#.#.#####.#.#.###
#...#.....#.#...#.#.#...#.#...#.#...#.#.......#.....#.....#...#...#.#...#.#...#.#...............#.....#...#.#...#...#.....#.#.#...#...#.....#
#.#.#.#######.###.#.#####.###.###.#.#.#.#.#.#.#.#############.#.###.#.###.###.#.#.#####.#.#.#.#.###.#####.#.#####.###.###.###.###.#.###.#.###
#.#.#.......#...#.#.#.......#.#...#...#...#.#.#...#.....#.....#.#...#...#...#.....#.#...#.#.#.#.......#...#.#.....#.#.#...#...#...#.#.......#
###.#.#####.###.#.#.#.#.#####.#.#########.#.#.###.#.#.###.#####.#.#####.###.#.#####.#.###.#.#####.#.#.#.###.#.#####.#.#.#.#.#######.#.#.#.#.#
#...#.....#...#...#.#.#...#...#.#.......#.#.#...#.....#...#...#.#...#.#.....#.#...#.....#.#.#.....#...#.....#...#...#.#.#.#.#.......#...#.#.#
#.#.#########.###.#.#.###.#.#.#.#.###.###.#.###.#####.#.#####.#.#.#.#.###.#####.#.#######.#.#.#####.#######.###.###.#.#.#.#.#.#########.#.#.#
#.#.........#...#.#.#...#.#.....#...#.....#...#.....#.#.#.....#.#.#.#.....#.....#...#.....#...#...#.#.......#.......#.#.....#.....#.......#.#
#.#.#######.###.#.#.#####.#######.#.#########.#####.###.###.###.#.#.###.###.#######.#.#.#######.#.#.#.#.#.###.###.###.#######.###.###.#.#.#.#
#.#...#...#...#...#.......#.......#.....#.........#...#...#...#.#.#...#.#.#...#...#...#.......#.#.#...#.#.#.#.#.....#.#.......#.............#
#.###.#.#.###.###########.#.###########.#.#.#####.###.#.#.#.#.#.#.###.#.#.###.#.#.#.###.#.###.###.#.###.#.#.#.#####.#.###.#####.###.#.#.#####
#...#...#.#.#.#...#.....#.#...#...........#.......#...#.#.#.#...#...#.#...#...#.#...#.....#.....#.......#...#...........#.#.......#.#...#...#
#.#.#####.#.#.#.#.#.###.#.###.#####.#######.#.#####.###.#.###.#####.#.#####.#########.#####.###.#######.###.###.#.#.###.#.#.#####.#.#.###.#.#
#.#.#.......#.#.#...#...#...#...#...#.....#.#.......#...#...#...#...#.#...#...#...#...#.......#...#.....#...#...#.#...#.#...#...#.#.#.....#.#
#.#.#.#######.#.#####.#####.###.#.#.#.###.#.#############.#.###.#.###.#.#.#.#.#.#.#.#####.###.###.#######.###.###.#####.#####.#.#.#.#.#####.#
#.#.#.#.#.....#.#...#.....#...#.#...#...#.#...#...........#.#...#...#.#.#.#.#...#...#...#.#.#...#...........#...#...#...#.....#.#.#.#.#...#.#
#.#.#.#.#.#####.#.#.#####.###.#.#.#.###.#.###.#.#.#####.###.#.#######.###.#.#####.#.#.#.#.#.###.###########.###.###.#.###.#####.#.#.#.###.#.#
#.#.....#...#...#.#...#...#...#.#.#...#.#.....#...#...#.#...#.........#...#.#.#...#...#.#.#.......#.......#...#...#.#.#.#.#...#...#.#.....#.#
###########.#.###.#.###.#.#.###.#.###.#.#######.#.###.#.#.###.#########.#.#.#.#.#######.#.#.#.#####.#####.###.###.#.#.#.#.#.#.#####.#######.#
#.........#.......#.#...#.#.....#...#.#...#...#.#.#...#.#.#...#.......#.#...#.#.#.........#.....#...#.....#.....#.#...#...#.#.#.....#.......#
#.#######.###########.#.#.#######.###.###.#.#.###.#.###.#.###.#.#####.#.###.#.#.#.#########.#.###.###.#########.#.###.#####.#.#.#####.#.#####
#.................#...#.#...#.....#...#...#.#.....#.#...#...#.#.#.....#...#...#...#...#.#.....#...#...#.......#.#...#.......#...#.#...#.#...#
#.#.#.#.###.###.#.#.###.#.###.#####.#.#.###.#######.#.#####.###.#.###.###.###.#####.#.#.#.#####.###.#.#.#.###.###.#.#############.#.#.###.#.#
#...#.#...#...#.#...#...............#.#.#.....#.....#.....#.#...#...#.#.#.#...#...#.#.#...#.....#.....#.#.#.#...#.#.#.....#.......#.#.....#.#
#.###.###.###.#.#.#.###.###.#####.#.#.#.###.#.#.#.#.#####.#.#.#####.#.#.#.#####.#.#.#.#####.#####.#######.#.###.###.#.#####.###.#.#.###.#.#.#
#...#.#.#.#...#...#...#...#.#.....#.#.#...#.#.#.#.#.....#.#...#...#.#...#.......#...#...#...#...#.#.......#...#...#.#.#...#...#.#.#.#...#.#.#
###.#.#.#.#.#.#####.#.#.###.#.###.#.#####.###.###.#######.#.###.#.#.###############.###.#.###.###.#####.###.#####.#.#.#.#.#.#.#.###.#.#.#.#.#
#...#.#.#...#.....#.#.#.#...#.#...#.....#.....#...#.......#.#...#...#.........#.......#...#.......#.....#.........#.#...#...#.#.....#.#...#.#
#.###.#.###.###.#.#.#.###.###.#######.#########.#.#.#######.#.#######.#######.#.###.#######.#######.#####.#########.#.#####.#.#######.#.###.#
#.#...#.....#...#...#.#...#.#.......#...........#.#.#.......#.#.......#...#...#...#.....#.#...#.#.....#...#.........#.......#.#.......#.#...#
###.#.#######.#######.#.###.#.#####.#.###.#.#####.#.#######.#.#.#####.###.#.#####.###.#.#.###.#.#.###.#.#######.#.###.#.###.#.#.#######.#.###
#...#.#.....#...#...#...#...#.#.#.....#.#.#.#...#.#...#...#.#.#.#.......#.#.....#...#.#.....#.#...#...#.......#.#.#...#.#...#...#.....#.#...#
#.#.###.###.###.#.#.#####.#.#.#.#.###.#.#.###.#.#####.#.#.#.#.#.#.#####.#.#####.###.#.#.#####.#.###.#########.###.#.###.#########.#####.###.#
#.#.#...#.....#.#.#.#...#.#.#...#.#.....#.....#.....#...#.#.#.#...#.#...#...#...#...#.#.#.....#.#...#...#.......#.#...#...........#.....#...#
#.###.###.###.#.#.#.#.#.###.###.#.#########.#####.#####.#.#.#.#####.#.#####.#.#.#.###.###.#######.###.###.#.###.#.#########.#######.#####.###
#.......#...#.#...#...#.#.....#.#.........#.....#.......#.#.#.......#.#.....#.#.#...#...#.#.......#...#...#...#.#.#...#.....#...........#...#
#.#####.###.###########.#.#####.#####.#.#.#.#.#.###.#####.###.#.###.#.#.#####.#####.###.#.#.#.#####.#.#.#.#.###.#.#.#.#.#####.#########.#.#.#
#.#...#...#.......#.....#...#.......#...#...#.#...#.#.........#.......#.....#.......#.#.....#.#.....#.....#.#...#.....#.#.......#.....#.#...#
###.#.###.#.###.###.#######.#.#############.#.#####.#.###############.#.###.#########.#.###.#.#############.#.#######.#.#.###.###.###.###.#.#
#...#...#.#.#...#...#.....#.#.#...........#.#.....#...#.#...#.......#.#...#.............#...#.......#.....#.#.........#.#...#.....#.....#.#.#
#.#####.#.#.#.###.###.###.#.#.#.#########.###.###.#####.#.#.#.#####.#.###.#######.#######.#########.#.###.#.###########.###.#.#.#######.#.#.#
#.....#.#.#.#...#.#.#...#.#...#.....#...#...#.#.........#.#...#...#.#.#.#.....................#.....#...#...#.....#.........#.#.#...#.....#.#
#.###.#.#.#.###.#.#.#.#.#.#########.#.#.###.###.#####.###.#####.#.#.#.#.#.#####.#############.#.###.#.#.###.#.#####.#########.###.#.#####.#.#
#.#...#.#...#...#.#...#.#...........#.#.........#.....#...#.....#.#.#.#.......#.#.......#...#...#.....#...#.#.#.........#...#.#...#...#...#.#
#.#.###.###.#.###.#.###.#############.###########.#.###.#####.###.#.#.#####.#.#.#.###.###.#.#.###.#####.#.#.#.#.#######.#.#.#.#.#####.###.#.#
#.#...#.......#...#...#.#.......#.....#.....#.....#.....#.....#...#...#...#.#.#.#.#...#...#.#.#...#...#.#...#...#...#...#.#.#.#.....#...#...#
#.###.#####.#.#.#######.#######.#.#####.###.#.#############.###.#######.#.###.#.#.###.#.###.#.#.###.#.#.#.#######.#.#.#####.#.#.#######.#.###
#...#...#...#.#.#.......#.......#...#.....#.#...#.........#.#...#.......#...#.....#.....#.#...#.#...#...#.......#.#.#.......#...#.......#...#
#######.###.#.#.###.#####.#####.###.#####.#####.#####.###.#.###.###.#.#####.#######.#####.#######.#############.#.#.#####.#.#####.#######.#.#
#.......#...#.#...#...........#.#...#.....#.....#...#.#...#...#...#.#...#...........#.....#.....#.#.....#.....#.#.#.#.......#.....#.....#...#
#.#######.#.#.###.#.#.#######.###.###.#####.#.###.#.#.#.#####.###.#####.#.###.#########.#.###.#.#.#.#####.#.###.#.#.#########.#####.###.#.###
#.....#.#.......#...#...#...#.....#...#.....#.....#...#.....#...#.......#.....#.........#.....#...#.......#...#.#.#.........#...#.....#.#.#.#
#.###.#.#.#.#########.#.#.#.#######.###.#######.###########.#.#.#################.#################.#########.#.#.#.#######.###.#.###.###.#.#
#.#.#.#.....#.........#.#.#.....#...#...#.............#...#...#.#.#.......#...............#.......#...#.#...#...#.#.#...#.#.#...#...#...#.#.#
#.#.#.###.###.###.#.#.#.#.###.#.#.#.#.#####.#######.#.#.#.###.#.#.#.###.#.#.#.#######.###.#.###.#####.#.#.#.#####.###.#.#.#.#.#######.#.#.#.#
#...#...#.#...#.#...#.#.#...#.....#.#.#.....#.....#.#...#...#.#.#.....#.#.#.#.#.......#.....#.#.....#.#...#.......#.....#...#.#.....#.#.#.#.#
#.#.###.#.#.###.#####.#.###########.#.#.#####.#.#.#.###.#.###.#.#####.#.#.###.#.###.#########.#.#.#.#.#.#######.#.#.###.#.###.#.#.#.###.#.#.#
#.#.....#.#.#.......#.......#.......#...#.....#.#...#.....#...#...#...#.#...#.#...#.#.......#...#.#.......#...#.#.#.#...#.#...#.#.......#...#
#.#.#.#####.#.#.#####.#####.#.#####.###.###.###.#####.#####.#####.#.###.###.#.#####.#.#####.#####.###.###.#.#.#.#.#.#.###.#.###.###########.#
#.#...#...#.#.#...#.......#...#...#.#.....#.#...#.....#...#.....#.#...#.#.#.#.....#.....#...#.....#...#.....#.#.#.#.#...#...#...#...#.....#.#
###.#.#.#.#.#####.#.###.#########.#.###.#.###.#.#######.#.###.#.#.#.###.#.#.#.###.#######.#.#.#####.#.#######.#.#.#.###.#######.#.#.#.###.#.#
#.....#.#...#...#.#...#.#.........#.....#...#.#.#.....#.#.....#...#...#.#.#.#.#.#.......#.#.........#.....#...#.......#.#.....#...#...#...#.#
#.#.###.#####.#.#.###.###.#####.#.#########.#.###.###.#.###.#.#####.#.#.#.#.#.#.#######.#.#.###.###.#####.###.#########.#.###.###.#####.###.#
#.#...#.#.....#...#.#...#...#...#.#.......#.#.#...#...#...#.#.....#.#...#.#.#...#.......#.#...#...#.#...#...#.#.......#...#.#.....#.....#...#
#.#.#.#.#.#######.#.###.###.#.###.###.#.###.#.#.###.###.###.#####.#.#####.#.###.#.#.#.###.###.###.#.#.#####.###.#####.#.###.#####.#.#####.###
#...#...#.....#.#.....#...#.#.#.#...#.#...#.#...#.#.....#...#...#...#.....#.#.....#.#.......#...#...#.....#.....#.....#...#.....#.#.........#
#.#.#####.###.#.#####.###.#.#.#.###.#####.#.#.#.#.#######.#####.#####.#.#.#.#.###.#.#.#####.###.#########.#######.#######.#.###.#.###########
#.#.....#...#.#.#.....#...#.#.#...........#.#...#.#.......#...........#.#.#.#...#...#.....#...#.#.......#...........#...#...#...#...........#
#.#.###.###.#.#.#.#####.###.#.#####.#######.#.###.#.#######.#########.###.#.#.###.#######.###.#.#.#.#.#.###.###.###.#.#######.###.#########.#
#...#...#.#.#.#.#.....#...#.#.#...#.#.....#.#.....#...#.............#.....#.#.....#.....#...#.#.#.#.#.........#.#...#.#.....#.#.....#.....#.#
###.#.###.#.#.#.#.#.#.###.#.#.#.#.#.#.###.#.#####.#.#.#.#####.#.###.###.###.#.#.###.#.#####.#.#.###.#########.#.#.#.#.#.#.#.#.###.#.#.#.###.#
#...#.#.................#.#.#...#.#.#...#.#.#.....#...#.#...#.#...#...#.#...#.#.#...#.....#.#.#...#.....#...#.#.......................#...#.#
#.#.#.#.#####.#.#.#.###.#.#.#.###.###.#.#.#.#.#####.###.###.#.#.#.###.###.#####.#.#.#.###.#.#.###.#####.#.#.#.#.#.#.#####.#.#####.#.#.#.#.#.#
#.#...#.#.....#...#.....#.#.#...#...#...#...#.#...#...#.....#.....#.......................#.#...#.....#.#.#...#.#.#.....#.#.#...#.....#.#.#.#
#.#######.#####.#.#####.#.#.#.#.###.#.#######.#.#.#.#.#########.###.###.#####.###.###.#.###.###.#####.#.#.#####.#.#.###.#.#.#.###.#####.#.#.#
#.......#.#.............#.#...#...#.#.#.#...#...#...#.........#...#.#...#.....#.#.#.#...#...............#.....#.....................#.......#
#######.#.#.#####.#######.###.###.#.#.#.#.#.#####.#.###.#####.#.###.#.#####.###.#.#.#.###.###.#####.#########.#.#####.###.#####.#.#.#.#.#####
#.#...#...#.#.......#.......#.#...#.#.#...#.......................#.#.......#.#...#.#...#...#.....#...#.........#.....#.......#.#...#.......#
#.#.#.#####.#.#.#####.#####.#.#.#.#.#.#####.#.#####.#.###.#.###.#.#.#########.#.###.#.#.###.#.###.###.#########.#.#########.#.#.#.#.###.###.#
#...#.....#.#.#...#...#.....#...#.#.#.#...#...#.....#.#.#...#...#.#.#.........#.......#...#.#...#...#.......#...#.#.......#.#.........#...#.#
#.#######.#.#.###.#.#.#######.###.#.#.#.#.#####.###.#.#.###.#.###.#.#.#######.#######.###.#.#######.#######.#####.#.#####.#######.###.#.#.#.#
#.#.....#...#...#.#.#.#...........#...#.#.........................#.#.....#.#.#.....#.#...#.......#...#.....#...#.....#.#...#.....#.......#.#
#.###.#.#########.#.###.#############.#.#####.###.#.#.#########.###.#####.#.#.#.#.###.#.#########.###.#.#####.#.#.###.#.###.#.###.#####.###.#
#S....#...........#.....................#.........#.............#...........#...#.....#...............#.......#.......#.......#.............#
#############################################################################################################################################`
	result := Solve(input)
	if result != 134588 {
		t.Errorf("Calculated solution was not expected")
	}
}
