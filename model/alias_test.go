package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewAlias(t *testing.T) {
	Convey("Given a new Alias", t, func() {
		alias := NewAlias()
		Convey("should not return nil", func() {
			So(alias, ShouldNotBeNil)
		})
	})
}

func TestNewAliasWithID(t *testing.T) {
	Convey("Given a new Alias with id `my_test_id`", t, func() {
		alias := NewAliasWith("my_test_id", "")
		Convey("should not return nil", func() {
			So(alias, ShouldNotBeNil)
			Convey("alias id should be `my_test_id`", func() {
				So(alias.GetID(), ShouldEqual, "my_test_id")
			})
			Convey("when attempt to change alias id to `new_id`", func() {
				Convey("alias id should be `my_test_id`", func() {
					alias.SetID("new_id")
					So(alias.GetID(), ShouldEqual, "my_test_id")
				})
			})
		})
	})
}

func TestNewAliasWithValue(t *testing.T) {
	Convey("Given a new Alias with id `my_test_id`", t, func() {
		alias := NewAliasWith("", "my_value")
		Convey("should not return nil", func() {
			So(alias, ShouldNotBeNil)
			Convey("alias value should be `my_value`", func() {
				So(alias.GetValue(), ShouldEqual, "my_value")
			})
			Convey("when attempt to change alias value to `new_value`", func() {
				Convey("alias value should be `new_value`", func() {
					alias.SetValue("new_value")
					So(alias.GetValue(), ShouldEqual, "new_value")
				})
			})
		})
	})
}
