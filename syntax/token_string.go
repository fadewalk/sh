// Code generated by "stringer -type Token"; DO NOT EDIT

package syntax

import "fmt"

const _Token_name = "illegalTokEOFLitLitWord'\"`&&&||||&$$'$\"${$[$($(([(((}])));;;;&;;&!++--***==!=<=>=+=-=*=/=%=&=|=^=<<=>>=>>><<><&>&>|<<<<-<<<&>&>><(>(+:+-:-?:?=:=%%%###^^^,,,///:!-e-f-d-c-b-p-S-L-g-u-r-w-x-s-t-z-n-o-v-R=~-nt-ot-ef-eq-ne-le-ge-lt-gt?(*(+(@(!("

var _Token_index = [...]uint8{0, 10, 13, 16, 23, 24, 25, 26, 27, 29, 31, 32, 34, 35, 37, 39, 41, 43, 45, 48, 49, 50, 52, 53, 54, 55, 57, 58, 60, 62, 65, 66, 68, 70, 71, 73, 75, 77, 79, 81, 83, 85, 87, 89, 91, 93, 95, 97, 100, 103, 104, 106, 107, 109, 111, 113, 115, 117, 120, 123, 125, 128, 130, 132, 133, 135, 136, 138, 139, 141, 142, 144, 145, 147, 148, 150, 151, 153, 154, 156, 157, 159, 160, 161, 163, 165, 167, 169, 171, 173, 175, 177, 179, 181, 183, 185, 187, 189, 191, 193, 195, 197, 199, 201, 203, 206, 209, 212, 215, 218, 221, 224, 227, 230, 232, 234, 236, 238, 240}

func (i Token) String() string {
	if i < 0 || i >= Token(len(_Token_index)-1) {
		return fmt.Sprintf("Token(%d)", i)
	}
	return _Token_name[_Token_index[i]:_Token_index[i+1]]
}
