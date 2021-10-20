package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	Convey("First stage", t, func() {
		So(5, ShouldEqual, 5)
	})

	Convey("Second stage", t, func() {
		assert.Equal(t, 5, 5)
		assert.Equal(t, 5, 6)
	})
}
