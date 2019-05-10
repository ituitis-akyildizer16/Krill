package git

import (
	"context"
	"fmt"
	"strings"
)

// BlameLine is a single parsed line of `git blame --porcelain`.
type BlameLine struct {
	Commit    string
	Author    string
	Timestamp string
	LineNo    int
	Content   string
}

// Blame returns up to maxLines lines of porcelain blame for a path,
