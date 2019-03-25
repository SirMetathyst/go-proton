package proton

import (
	"testing"

	goconvey "github.com/smartystreets/goconvey/convey"
)

func TestAliasList(t *testing.T) {
	goconvey.Convey("when creating a new alias list", t, func() {
		aliasList := NewAliasList()
		goconvey.Convey("alias list should not be nil", func() {
			goconvey.So(aliasList, goconvey.ShouldNotBeNil)
		})
	})

	goconvey.Convey("when creating a new alias list", t, func() {
		aliasList := NewAliasList()
		goconvey.Convey("and adding a nil alias", func() {
			err := aliasList.AddAlias(nil)
			goconvey.Convey("error should be `"+ErrAliasListTriedToAddNilAlias.Error()+"`", func() {
				goconvey.So(err, goconvey.ShouldEqual, ErrAliasListTriedToAddNilAlias)
			})
		})

		goconvey.Convey("and adding a valid alias", func() {
			alias, _ := NewAlias("test", "value")
			err := aliasList.AddAlias(alias)
			goconvey.Convey("error should be nil", func() {
				goconvey.So(err, goconvey.ShouldBeNil)
			})

			goconvey.Convey("and retrieving a list of aliases", func() {
				slice := aliasList.AliasSlice()
				goconvey.Convey("slice length should equal `1`", func() {
					goconvey.So(len(slice), goconvey.ShouldEqual, 1)
				})
			})

			goconvey.Convey("and adding an alias with the same id as before", func() {
				alias, _ := NewAlias("test", "value2")
				err := aliasList.AddAlias(alias)
				goconvey.Convey("error should equal `"+ErrAliasListTriedToAddDuplicateAliasID.Error()+"`", func() {
					goconvey.So(err, goconvey.ShouldEqual, ErrAliasListTriedToAddDuplicateAliasID)
				})
			})
		})
	})
}
