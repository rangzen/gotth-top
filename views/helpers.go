package views

import (
	"fmt"
	"math"
)

// PrettyByteSize returns a human-readable byte size string.
// e.g. 1337 -> 1.3KiB
// Reference: https://stackoverflow.com/questions/1094841/get-a-human-readable-version-of-a-file-size/1094933#1094933
// Reference: https://gist.github.com/anikitenko/b41206a49727b83a530142c76b1cb82d?permalink_comment_id=4467913#gistcomment-4467913
func PrettyByteSize(b int) string {
	bf := float64(b)
	for _, unit := range []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"} {
		if math.Abs(bf) < 1024.0 {
			return fmt.Sprintf("%3.1f%sB", bf, unit)
		}
		bf /= 1024.0
	}
	return fmt.Sprintf("%.1fYiB", bf)
}
