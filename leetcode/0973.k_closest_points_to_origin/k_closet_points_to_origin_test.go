package leetcode

import "testing"

func Test_kClosest(t *testing.T) {
	var tests = []struct {
		points   [][]int
		k        int
		expected [][]int
	}{
		{[][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
		{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2, [][]int{{3, 3}, {-2, 4}}},
		{
			[][]int{{-2, -6}, {-7, -2}, {-9, 6}, {10, 3}, {-8, 1}, {2, 8}},
			5,
			[][]int{{-2, -6}, {-7, -2}, {10, 3}, {-8, 1}, {2, 8}},
		},
	}

	for _, e := range tests {
		result := kClosest(e.points, e.k)
		if !sliceOfSliceEqual(result, e.expected) {
			t.Errorf(
				"expected k closet point are %v but got %v",
				e.expected,
				result,
			)
		}
	}
}

func sliceOfSliceEqual(matrixa, matrixb [][]int) bool {
	if len(matrixa) != len(matrixb) {
		return false
	}

	for _, rowa := range matrixa {
		found := false
		for _, rowb := range matrixb {
			if rowa[0] == rowb[0] && rowa[1] == rowb[1] {
				found = true
			}
		}
		if found == false {
			return false
		}
	}
	return true
}
