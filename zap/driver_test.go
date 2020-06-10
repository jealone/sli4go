package zap

import (
	"testing"

	"github.com/jealone/sli4go"
)

func TestZap(t *testing.T) {
	sli4go.GetLogger().Info("test zap log")
}
