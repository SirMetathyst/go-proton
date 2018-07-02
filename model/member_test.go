package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewMember(t *testing.T) {
	Convey("Given a new member", t, func() {
		member := NewMember()
		Convey("should not return nil", func() {
			So(member, ShouldNotBeNil)
		})
	})
}

func TestNewMemberWithValue(t *testing.T) {
	Convey("Given a new member with value `my_value`", t, func() {
		member := NewMemberWithValue("", "my_value", 0)
		Convey("should not return nil", func() {
			So(member, ShouldNotBeNil)
			Convey("member value should be `my_value`", func() {
				So(member.value, ShouldEqual, "my_value")
			})
		})
	})
}

func TestNewMemberWithID(t *testing.T) {
	Convey("Given a new member with id `my_id`", t, func() {
		member := NewMemberWithID("my_id")
		Convey("should not return nil", func() {
			So(member, ShouldNotBeNil)
			Convey("member id should be `my_id`", func() {
				So(member.id, ShouldEqual, "my_id")
			})
		})
	})
}
