// Code generated by "stringer -type=BaselineShifts"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

const _BaselineShifts_name = "ShiftBaselineShiftSuperShiftSubBaselineShiftsN"

var _BaselineShifts_index = [...]uint8{0, 13, 23, 31, 46}

func (i BaselineShifts) String() string {
	if i < 0 || i >= BaselineShifts(len(_BaselineShifts_index)-1) {
		return "BaselineShifts(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BaselineShifts_name[_BaselineShifts_index[i]:_BaselineShifts_index[i+1]]
}

func (i *BaselineShifts) FromString(s string) error {
	for j := 0; j < len(_BaselineShifts_index)-1; j++ {
		if s == _BaselineShifts_name[_BaselineShifts_index[j]:_BaselineShifts_index[j+1]] {
			*i = BaselineShifts(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: BaselineShifts")
}