// Code generated by "stringer -type=AlignVert"; DO NOT EDIT.

package gi

import (
	"fmt"
	"strconv"
)

const _AlignVert_name = "AlignTopAlignVCenterAlignBottomAlignVJustifyAlignBaselineAlignVertN"

var _AlignVert_index = [...]uint8{0, 8, 20, 31, 44, 57, 67}

func (i AlignVert) String() string {
	if i < 0 || i >= AlignVert(len(_AlignVert_index)-1) {
		return "AlignVert(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AlignVert_name[_AlignVert_index[i]:_AlignVert_index[i+1]]
}

func StringToAlignVert(s string) (AlignVert, error) {
	for i := 0; i < len(_AlignVert_index)-1; i++ {
		if s == _AlignVert_name[_AlignVert_index[i]:_AlignVert_index[i+1]] {
			return AlignVert(i), nil
		}
	}
	return 0, fmt.Errorf("String %v is not a valid option for type AlignVert", s)
}
