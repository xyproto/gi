// Code generated by "stringer -type=ViewBoxMeetOrSlice"; DO NOT EDIT.

package svg

import (
	"fmt"
	"strconv"
)

const _ViewBoxMeetOrSlice_name = "MeetSlice"

var _ViewBoxMeetOrSlice_index = [...]uint8{0, 4, 9}

func (i ViewBoxMeetOrSlice) String() string {
	if i < 0 || i >= ViewBoxMeetOrSlice(len(_ViewBoxMeetOrSlice_index)-1) {
		return "ViewBoxMeetOrSlice(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ViewBoxMeetOrSlice_name[_ViewBoxMeetOrSlice_index[i]:_ViewBoxMeetOrSlice_index[i+1]]
}

func (i *ViewBoxMeetOrSlice) FromString(s string) error {
	for j := 0; j < len(_ViewBoxMeetOrSlice_index)-1; j++ {
		if s == _ViewBoxMeetOrSlice_name[_ViewBoxMeetOrSlice_index[j]:_ViewBoxMeetOrSlice_index[j+1]] {
			*i = ViewBoxMeetOrSlice(j)
			return nil
		}
	}
	return fmt.Errorf("String %v is not a valid option for type ViewBoxMeetOrSlice", s)
}
