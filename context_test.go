package proton

import (
	"testing"

	goconvey "github.com/smartystreets/goconvey/convey"
)

func TestContext(t *testing.T) {
	goconvey.Convey("when creating a new context", t, func() {
		goconvey.Convey("without an id", func() {
			context, err := NewContext("")
			goconvey.Convey("context is nil", func() {
				goconvey.So(context, goconvey.ShouldBeNil)
			})
			goconvey.Convey("error should be `"+ErrContextIDUndefined.Error()+"`", func() {
				goconvey.So(err, goconvey.ShouldEqual, ErrContextIDUndefined)
			})
		})

		goconvey.Convey("with an id of `test`", func() {
			context, err := NewContext("test")
			goconvey.Convey("context is not nil", func() {
				goconvey.So(context, goconvey.ShouldNotBeNil)
			})
			goconvey.Convey("error be nil", func() {
				goconvey.So(err, goconvey.ShouldBeNil)
			})
		})

		goconvey.Convey("with id that contains spaces", func() {
			context, err := NewContext("test  ")
			goconvey.Convey("context is nil", func() {
				goconvey.So(context, goconvey.ShouldBeNil)
			})
			goconvey.Convey("error should be `"+ErrContextIDContainsWhitespace.Error()+"`", func() {
				goconvey.So(err, goconvey.ShouldEqual, ErrContextIDContainsWhitespace)
			})
		})

		goconvey.Convey("with id that contains tabs", func() {
			context, err := NewContext("test	")
			goconvey.Convey("context is nil", func() {
				goconvey.So(context, goconvey.ShouldBeNil)
			})
			goconvey.Convey("error should be `"+ErrContextIDContainsWhitespace.Error()+"`", func() {
				goconvey.So(err, goconvey.ShouldEqual, ErrContextIDContainsWhitespace)
			})
		})
	})

	goconvey.Convey("when given a valid context", t, func() {
		context, err := NewContext("test")
		goconvey.Convey("context should not be nil", func() {
			goconvey.So(context, goconvey.ShouldNotBeNil)
		})
		goconvey.Convey("error should be nil", func() {
			goconvey.So(err, goconvey.ShouldBeNil)
		})
		goconvey.Convey("retrieving the id", func() {
			id := context.ID()
			goconvey.Convey("should equal `test`", func() {
				goconvey.So(id, goconvey.ShouldEqual, String("test"))
			})
		})
	})
}
