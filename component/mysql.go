//go:build mysql || full || mini
// +build mysql full mini

package build

import (
	_ "github.com/nzai/trojan-go/statistic/mysql"
)
