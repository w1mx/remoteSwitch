// Code generated by "stringer -type=rbModel"; DO NOT EDIT.

package remotebox

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[rbUnknown-0]
	_ = x[rb1x6-1]
	_ = x[rb2x6-2]
	_ = x[rb1x8-3]
	_ = x[rb2x8-4]
	_ = x[rb2x12-5]
	_ = x[rb4sq-6]
	_ = x[rb4sqplus-7]
}

const _rbModel_name = "rbUnknownrb1x6rb2x6rb1x8rb2x8rb2x12rb4sqrb4sqplus"

var _rbModel_index = [...]uint8{0, 9, 14, 19, 24, 29, 35, 40, 49}

func (i rbModel) String() string {
	if i < 0 || i >= rbModel(len(_rbModel_index)-1) {
		return "rbModel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _rbModel_name[_rbModel_index[i]:_rbModel_index[i+1]]
}
