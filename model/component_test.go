package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewComponent(t *testing.T) {
	Convey("Given a new Component", t, func() {
		component := NewComponent()
		Convey("should not return nil", func() {
			So(component, ShouldNotBeNil)
		})
	})
}

func TestNewComponentWithID(t *testing.T) {
	Convey("Given a new Component with id `my_test_id`", t, func() {
		component := NewComponentWithID("my_test_id")
		Convey("should not return nil", func() {
			So(component, ShouldNotBeNil)
			Convey("component id should be `my_test_id`", func() {
				So(component.GetID(), ShouldEqual, "my_test_id")
			})
		})
	})
}
