package touchbar

import (
	"github.com/LouisBrunner/go-touchbar/pkg/internal/darwin"
)

// New allows to create a new Touch Bar for this application
// Note: only one Touch Bar can be active at a given time.
var New = darwin.NewTouchBar
