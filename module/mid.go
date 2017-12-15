package module

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"go_spider/errors"
)

var DefaultSNGen = NewSNGenertor(1, 0)

type MID string

func GenMID() error {
	mtype := "string"
	errMsg := fmt.Sprintf("illegal module type: %s", mtype)
	return errors.NewIllegalParameterError(errMsg)
}

func SplitMID(mid MID) ([]string, error) {

}
