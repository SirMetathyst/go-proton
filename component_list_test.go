package proton

import (
	"testing"

	goconvey "github.com/smartystreets/goconvey/convey"
)

func TestComponentList(t *testing.T) {
	goconvey.Convey("when creating a new component list", t, func() {
		componentList := NewComponentList()
		goconvey.Convey("component list should not be nil", func() {
			goconvey.So(componentList, goconvey.ShouldNotBeNil)
		})
	})

	goconvey.Convey("when creating a new component list", t, func() {
		componentList := NewComponentList()

		contextList := NewContextList()
		context, _ := NewContext("test")
		contextList.AddContext(context)


		componentMemberList := NewComponentMemberList()
		componentMember, _ := NewComponentMember("id", "value", SingleEntityIndex)
		componentMemberList.AddMember(componentMember)

		goconvey.Convey("and adding a nil component", func() {
			err := componentList.AddComponent(nil)
			goconvey.Convey("error should be `"+ErrComponentListTriedToAddNilComponent.Error()+"`", func() {
				goconvey.So(err, goconvey.ShouldEqual, ErrComponentListTriedToAddNilComponent)
			})
		})

		goconvey.Convey("and adding a valid component", func() {
			component, _ := NewComponent("test", "", false, 0, 0, 0, 0, contextList, componentMemberList)
			err := componentList.AddComponent(component)
			goconvey.Convey("error should be nil", func() {
				goconvey.So(err, goconvey.ShouldBeNil)
			})

			goconvey.Convey("and retrieving a list of components", func() {
				slice := componentList.ComponentSlice()
				goconvey.Convey("slice length should equal `1`", func() {
					goconvey.So(len(slice), goconvey.ShouldEqual, 1)
				})
			})
			

			goconvey.Convey("and retrieving components with context id", func() {
				slice := componentList.ComponentsWithContextID("test")
				goconvey.Convey("slice length should equal `1`", func() {
					goconvey.So(len(slice), goconvey.ShouldEqual, 1)
					goconvey.Convey("component id is `test`", func() {
						id := slice[0].ID()
						goconvey.So(id, goconvey.ShouldEqual, "test")
					})
				})
			})

			goconvey.Convey("and retrieving components with entity index", func() {
				slice := componentList.ComponentsWithEntityIndex()
				goconvey.Convey("slice length should equal `1`", func() {
					goconvey.So(len(slice), goconvey.ShouldEqual, 1)
					goconvey.Convey("component id is `test`", func() {
						id := slice[0].ID()
						goconvey.So(id, goconvey.ShouldEqual, "test")
					})
				})
			})

			goconvey.Convey("and adding an component with the same id as before", func() {
				component, _ := NewComponent("test", "", false, 0, 0, 0, 0, contextList, componentMemberList)
				err := componentList.AddComponent(component)
				goconvey.Convey("error should equal `"+ErrComponentListTriedToAddDuplicateComponentID.Error()+"`", func() {
					goconvey.So(err, goconvey.ShouldEqual, ErrComponentListTriedToAddDuplicateComponentID)
				})
			})
		})
	})
}
