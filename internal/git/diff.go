package git

import (
	"context"
	"strconv"
	"strings"
)

// Diff returns the uncommitted diff (staged or unstaged) bounded by
// maxLines of output.
