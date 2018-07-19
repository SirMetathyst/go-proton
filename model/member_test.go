package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewMember(t *testing.T) {
	Convey("Given a new Member", t, func() {
		member, err := NewMember("", "", 0)
		Convey("NewMember should return error", func() {
			So(err, ShouldNotBeNil)
			Convey("NewMember should return nil", func() {
				So(member, ShouldBeNil)
			})
		})
	})
}
