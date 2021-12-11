package bilderbergclub

import (
	_ "embed" // to embed members file
	"runtime"
	"strings"
)

func MustApproveThisMember() {
	_, caller, _, ok := runtime.Caller(1)
	if !ok {
		panic("bilderbergclub: cannot verify participant")
	}
	if _, ok := members[caller]; !ok {
		panic("bilderbergclub: caller is not a club participant")
	}
	// Welcome back, the mighty of this world.
}

//go:embed members.txt
var membersFile string

var members = func() map[string]struct{} {
	members := make(map[string]struct{})
	for _, m := range strings.Split(membersFile, "\n") {
		members[m] = struct{}{}
	}
	return members
}()
