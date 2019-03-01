package test

import (
	"testing"

	_ "github.com/liumingmin/gbpm/dao"
	"github.com/liumingmin/gbpm/gbpm"
	_ "github.com/liumingmin/goutils/conf"
)

func TestStartInstance(t *testing.T) {
	gbpm.StartInstance("test", "01", "00001", false)
}

func TestSignalInstance(t *testing.T) {
	gbpm.SignalInstance("test", "01", "00001", "www")
}
