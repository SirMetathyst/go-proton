package proton

import (
	"testing"

	goconvey "github.com/smartystreets/goconvey/convey"
)

func TestComponentMemberList(t *testing.T) {
	goconvey.Convey("when creating a new component member list", t, func() {
		componentMemberList := NewComponentMemberList()
		goconvey.Convey("component member list should not be nil", func() {
			goconvey.So(componentMemberList, goconvey.ShouldNotBeNil)
		})
	})

	goconvey.Convey("when creating a new component member list", t, func() {
		componentMemberList := NewComponentMemberList()
		goconvey.Convey("and adding a nil component member", func() {
			err := componentMemberList.AddMember(nil)
			goconvey.Convey("error should be `"+ErrComponentMemberListTriedToAddNilMember.Error()+"`", func() {
				goconvey.So(err, goconvey.ShouldEqual, ErrComponentMemberListTriedToAddNilMember)
			})
		})

		goconvey.Convey("and adding a valid component member", func() {
			componentMember, _ := NewComponentMember("test", "value", SingleEntityIndex)
			err := componentMemberList.AddMember(componentMember)
			goconvey.Convey("error should be nil", func() {
				goconvey.So(err, goconvey.ShouldBeNil)
			})

			goconvey.Convey("and retrieving a list of component members", func() {
				slice := componentMemberList.MemberSlice()
				goconvey.Convey("slice length should equal `1`", func() {
					goconvey.So(len(slice), goconvey.ShouldEqual, 1)
				})
			})

			goconvey.Convey("and retrieving component members with entity index", func() {
				slice := componentMemberList.MembersWithEntityIndex()
				goconvey.Convey("slice length should equal `1`", func() {
					goconvey.So(len(slice), goconvey.ShouldEqual, 1)
					goconvey.Convey("component member id should equal `test`", func() {
						id := slice[0].ID()
						goconvey.So(id, goconvey.ShouldEqual, "test")
					})
				})
			})

			goconvey.Convey("and retrieving component member with id `test`", func() {
				componentMember := componentMemberList.MemberWithID("test")
				goconvey.Convey("member should not be nil", func() {
					goconvey.So(componentMember, goconvey.ShouldNotBeNil)
					goconvey.Convey("component member id should equal `test`", func() {
						id := componentMember.ID()
						goconvey.So(id, goconvey.ShouldEqual, "test")
					})
				})
			})

			goconvey.Convey("and adding an component with the same id as before", func() {
				componentMember, _ := NewComponentMember("test", "value", SingleEntityIndex)
				err := componentMemberList.AddMember(componentMember)
				goconvey.Convey("error should equal `"+ErrComponentMemberListTriedToAddDuplicateMemberID.Error()+"`", func() {
					goconvey.So(err, goconvey.ShouldEqual, ErrComponentMemberListTriedToAddDuplicateMemberID)
				})
			})
		})
	})
}
