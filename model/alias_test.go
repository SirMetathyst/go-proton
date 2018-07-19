package model

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestNewAlias(t *testing.T) {
	c.Convey("Given a new Alias when id is `` and value is ``", t, func() {
		alias, err := NewAlias("", "")
		c.Convey("NewAlias should return an error", func() {
			c.So(err, c.ShouldNotBeNil)
			c.Convey("Error should equal `Alias: `ID` Undefined.`", func() {
				c.So(err.Error(), c.ShouldEqual, "Alias: `ID` Undefined.")
			})
			c.Convey("Alias should be nil", func() {
				c.So(alias, c.ShouldBeNil)
			})
		})
	})
	c.Convey("Given a new Alias when id is `my_id` and value is ``", t, func() {
		alias, err := NewAlias("my_id", "")
		c.Convey("NewAlias should return an error", func() {
			c.So(err, c.ShouldNotBeNil)
			c.Convey("Error should equal `Alias: `Value` Undefined.`", func() {
				c.So(err.Error(), c.ShouldEqual, "Alias: `Value` Undefined.")
			})
			c.Convey("Alias should be nil", func() {
				c.So(alias, c.ShouldBeNil)
			})
		})
	})
	c.Convey("Given a new Alias when id is `my_id` and value is `my_value`", t, func() {
		alias, err := NewAlias("my_id", "my_value")
		c.Convey("NewAlias should not return an error", func() {
			c.So(err, c.ShouldBeNil)
			c.Convey("Alias should not be nil", func() {
				c.So(alias, c.ShouldNotBeNil)
				c.Convey("Alias id should equal `my_id`", func() {
					c.So(alias.ID(), c.ShouldEqual, "my_id")
				})
				c.Convey("Alias value should equal `my_value`", func() {
					c.So(alias.Value(), c.ShouldEqual, "my_value")
				})
			})
		})
	})
}
