package proton

import (
	"testing"

	goconvey "github.com/smartystreets/goconvey/convey"
)

func TestContextList(t *testing.T) {
	goconvey.Convey("when creating a new context list", t, func() {
		contextList := NewContextList()
		goconvey.Convey("context list should not be nil", func() {
			goconvey.So(contextList, goconvey.ShouldNotBeNil)
		})
	})

	goconvey.Convey("when creating a new context list", t, func() {
		contextList := NewContextList()
		goconvey.Convey("and adding a nil context", func() {
			err := contextList.AddContext(nil)
			goconvey.Convey("error should be `"+ErrContextListTriedToAddNilContext.Error()+"`", func() {
				goconvey.So(err, goconvey.ShouldEqual, ErrContextListTriedToAddNilContext)
			})
		})

		goconvey.Convey("and adding a valid context", func() {
			context, _ := NewContext("test")
			err := contextList.AddContext(context)
			goconvey.Convey("error should be nil", func() {
				goconvey.So(err, goconvey.ShouldBeNil)
			})

			goconvey.Convey("and retrieving a list of contexts", func() {
				slice := contextList.ContextSlice()
				goconvey.Convey("slice length should equal `1`", func() {
					goconvey.So(len(slice), goconvey.ShouldEqual, 1)
				})
			})

			goconvey.Convey("and adding an context with the same id as before", func() {
				context, _ := NewContext("test")
				err := contextList.AddContext(context)
				goconvey.Convey("error should equal `"+ErrContextListTriedToAddDuplicateContextID.Error()+"`", func() {
					goconvey.So(err, goconvey.ShouldEqual, ErrContextListTriedToAddDuplicateContextID)
				})
			})
		})
	})
}
