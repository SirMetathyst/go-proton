package proton

import (
	"testing"

	goconvey "github.com/smartystreets/goconvey/convey"
)

func TestAlias(t *testing.T) {
	goconvey.Convey("when creating a new alias", t, func() {

		goconvey.Convey("without an id", func() {
			alias, err := NewAlias("", "")
			goconvey.Convey("alias is nil", func() {
				goconvey.So(alias, goconvey.ShouldBeNil)
			})
			goconvey.Convey("error should be `"+ErrAliasIDUndefined.Error()+"`", func() {
				goconvey.So(err, goconvey.ShouldEqual, ErrAliasIDUndefined)
			})
		})

		goconvey.Convey("with an id of `test` and value of ``", func() {
			alias, err := NewAlias("test", "")
			goconvey.Convey("alias is nil", func() {
				goconvey.So(alias, goconvey.ShouldBeNil)
			})
			goconvey.Convey("error should be `"+ErrAliasValueUndefined.Error()+"`", func() {
				goconvey.So(err, goconvey.ShouldEqual, ErrAliasValueUndefined)
			})
		})

		goconvey.Convey("with id that contains spaces", func() {
			alias, err := NewAlias("test  ", "")
			goconvey.Convey("alias is nil", func() {
				goconvey.So(alias, goconvey.ShouldBeNil)
			})
			goconvey.Convey("error should be `"+ErrAliasIDContainsWhitespace.Error()+"`", func() {
				goconvey.So(err, goconvey.ShouldEqual, ErrAliasIDContainsWhitespace)
			})
		})

		goconvey.Convey("with id that contains tabs", func() {
			alias, err := NewAlias("test	", "")
			goconvey.Convey("alias is nil", func() {
				goconvey.So(alias, goconvey.ShouldBeNil)
			})
			goconvey.Convey("error should be `"+ErrAliasIDContainsWhitespace.Error()+"`", func() {
				goconvey.So(err, goconvey.ShouldEqual, ErrAliasIDContainsWhitespace)
			})
		})

		goconvey.Convey("with an id of `test` and value contains spaces", func() {
			alias, err := NewAlias("test", "value  ")
			goconvey.Convey("alias is nil", func() {
				goconvey.So(alias, goconvey.ShouldBeNil)
			})
			goconvey.Convey("error should be `"+ErrAliasValueContainsWhitespace.Error()+"`", func() {
				goconvey.So(err, goconvey.ShouldEqual, ErrAliasValueContainsWhitespace)
			})
		})

		goconvey.Convey("with an id of `test` and value contains tabs", func() {
			alias, err := NewAlias("test", "value	")
			goconvey.Convey("alias is nil", func() {
				goconvey.So(alias, goconvey.ShouldBeNil)
			})
			goconvey.Convey("error should be `"+ErrAliasValueContainsWhitespace.Error()+"`", func() {
				goconvey.So(err, goconvey.ShouldEqual, ErrAliasValueContainsWhitespace)
			})
		})
	})

	goconvey.Convey("when given a valid alias", t, func() {
		alias, err := NewAlias("test", "value")
		goconvey.Convey("alias should not be nil", func() {
			goconvey.So(alias, goconvey.ShouldNotBeNil)
		})
		goconvey.Convey("error should be nil", func() {
			goconvey.So(err, goconvey.ShouldBeNil)
		})
		goconvey.Convey("retrieving the id", func() {
			id := alias.ID()
			goconvey.Convey("should equal `test`", func() {
				goconvey.So(id, goconvey.ShouldEqual, String("test"))
			})
		})
		goconvey.Convey("retrieving the `value`", func() {
			value := alias.Value()
			goconvey.Convey("should equal `value`", func() {
				goconvey.So(value, goconvey.ShouldEqual, String("value"))
			})
		})
	})
}
