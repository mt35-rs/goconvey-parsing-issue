package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	Convey("First stage", t, func() {
		So(5, ShouldEqual, 5)
		So(5, ShouldEqual, 6)
	})
}
