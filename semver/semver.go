package semver

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type SemVer struct {
	Value     string
	Major     int
	Minor     int
	Increment int
}

func (v *SemVer) String() string {

	return fmt.Sprintf("%v.%v.%v", v.Major, v.Minor, v.Increment)
}

func New(value string) (*SemVer, error) {
	splits := strings.Split(value, ".")
	if len(splits) != 3 {
		return nil, errors.New("Not in format 'n.n.n'")
	}
	major, err := strconv.Atoi(splits[0])
	if err != nil {
		return nil, errors.New("Major version not an integer.")
	}
	minor, err := strconv.Atoi(splits[1])
	if err != nil {
		return nil, errors.New("Minor version not an integer.")
	}
	increment, err := strconv.Atoi(splits[2])
	if err != nil {
		return nil, errors.New("Increment version not an integer.")
	}
	return &SemVer{Value: value, Major: major, Minor: minor, Increment: increment}, nil
}
