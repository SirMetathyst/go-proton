package modelbuilder

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAliasBuilderNewAlias(t *testing.T) {
	Convey("Given a new Alias using AliasBuilder with `` and ``", t, func() {
		alias, error := AliasBuilder.NewAlias("", "")
		Convey("NewAlias should return an error", func() {
			So(error, ShouldNotBeNil)
			Convey("error should equal `Alias: `ID` Undefined.`", func() {
				So(error.Error(), ShouldEqual, "Alias: `ID` Undefined.")
			})
			Convey("NewAlias should return nil", func() {
				So(alias, ShouldBeNil)
			})
		})
	})
	Convey("Given a new Alias using AliasBuilder with `my_id` and `my_value`", t, func() {
		alias, error := AliasBuilder.NewAlias("my_id", "my_value")
		Convey("NewAlias should not return an error", func() {
			So(error, ShouldBeNil)
			Convey("Alias should not be nil", func() {
				So(alias, ShouldNotBeNil)
				Convey("Alias id should be `my_id`", func() {
					So(alias.GetID(), ShouldEqual, "my_id")
				})
				Convey("Alias value should be `my_value`", func() {
					So(alias.GetValue(), ShouldEqual, "my_value")
				})
			})
		})
	})
}

func TestAliasBuilderSetID(t *testing.T) {
	Convey("Given a new AliasBuilder", t, func() {
		ab := NewAliasBuilder()
		Convey("AliasBuilder should not be nil", func() {
			So(ab, ShouldNotBeNil)
			Convey("when AliasBuilder id is `my_id`", func() {
				ab.SetID("my_id")
				Convey("when Build is called on AliasBuilder", func() {
					alias, error := ab.Build()
					Convey("Build should return error", func() {
						So(error.Error(), ShouldEqual, "Alias: `Value` Undefined.")
						Convey("Alias should be nil", func() {
							So(alias, ShouldBeNil)
						})
					})
				})
			})
		})
	})
}

func TestAliasBuilderSetValue(t *testing.T) {
	Convey("Given a new AliasBuilder", t, func() {
		ab := NewAliasBuilder()
		Convey("AliasBuilder should not be nil", func() {
			So(ab, ShouldNotBeNil)
			Convey("when AliasBuilder value is `my_value`", func() {
				ab.SetValue("my_value")
				Convey("when Build is called on AliasBuilder", func() {
					alias, error := ab.Build()
					Convey("Build should return error", func() {
						So(error.Error(), ShouldEqual, "Alias: `ID` Undefined.")
						Convey("Alias should be nil", func() {
							So(alias, ShouldBeNil)
						})
					})
				})
			})
		})
	})
}

func TestAliasBuilderReset(t *testing.T) {
	Convey("Given a new AliasBuilder", t, func() {
		ab := NewAliasBuilder()
		Convey("AliasBuilder should not be nil", func() {
			So(ab, ShouldNotBeNil)
			Convey("when AliasBuilder id is `my_id` and value is `my_value`", func() {
				ab.SetID("my_id")
				ab.SetValue("my_value")
				Convey("AliasBuilder id should equal `my_id`", func() {
					So(ab.id, ShouldEqual, "my_id")
					Convey("AliasBuilder value should equal `my_value`", func() {
						So(ab.value, ShouldEqual, "my_value")
						Convey("when Reset is called on AliasBuilder", func() {
							ab.Reset()
							Convey("AliasBuilder id should equal ``", func() {
								So(ab.id, ShouldBeEmpty)
							})
							Convey("AliasBuilder value should equal ``", func() {
								So(ab.value, ShouldBeEmpty)
							})
						})
					})
				})
			})
		})
	})
}
