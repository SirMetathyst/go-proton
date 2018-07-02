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

func TestComponentSetID(t *testing.T) {
	Convey("Given a new Component with id `my_test_id`", t, func() {
		component := NewComponentWithID("my_test_id")
		Convey("should not return nil", func() {
			So(component, ShouldNotBeNil)
			Convey("when attempt to change component id to `new_id`", func() {
				Convey("component id should be `my_test_id`", func() {
					component.SetID("new_id")
					So(component.id, ShouldEqual, "my_test_id")
				})
			})
		})
	})
	Convey("Given a new Component with id ``", t, func() {
		component := NewComponentWithID("")
		Convey("should not return nil", func() {
			So(component, ShouldNotBeNil)
			Convey("when attempt to change component id to `new_id`", func() {
				Convey("component id should be `new_id`", func() {
					component.SetID("new_id")
					So(component.id, ShouldEqual, "new_id")
				})
			})
		})
	})
}

func TestComponentGetID(t *testing.T) {
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

func TestComponentHasPrefix(t *testing.T) {
	Convey("Given a new Component", t, func() {
		component := NewComponent()
		Convey("should not return nil", func() {
			So(component, ShouldNotBeNil)
			Convey("component should have no prefix", func() {
				So(component.HasPrefix(), ShouldBeFalse)
			})
		})
	})
}

func TestComponentGetPrefix(t *testing.T) {
	Convey("Given a new Component", t, func() {
		component := NewComponent()
		Convey("should not return nil", func() {
			So(component, ShouldNotBeNil)
			Convey("component prefix should be ``", func() {
				So(component.GetPrefix(), ShouldBeEmpty)
			})
		})
	})
}

func TestComponentSetPrefix(t *testing.T) {
	Convey("Given a new Component", t, func() {
		component := NewComponent()
		Convey("should not return nil", func() {
			So(component, ShouldNotBeNil)
			Convey("component prefix should be ``", func() {
				So(component.GetPrefix(), ShouldBeEmpty)
			})
			Convey("when component prefix is set to `my_prefix`", func() {
				component.SetPrefix("my_prefix")
				Convey("component prefix should be `my_prefix`", func() {
					So(component.GetPrefix(), ShouldEqual, "my_prefix")
				})
			})
		})
	})
}

func TestComponentGetPrefixOrDefault(t *testing.T) {
	Convey("Given a new Component", t, func() {
		component := NewComponent()
		Convey("should not return nil", func() {
			So(component, ShouldNotBeNil)
			Convey("component default prefix should be `is`", func() {
				So(component.GetPrefixOrDefault(), ShouldEqual, "is")
			})
			Convey("when one or more members are added to the component", func() {
				component.member["member"] = &Member{}
				Convey("component default prefix should be `has`", func() {
					So(component.GetPrefixOrDefault(), ShouldEqual, "has")
				})
			})
			Convey("unless a prefix is set, in which case the user set prefix is returned", func() {
				Convey("when component prefix is `my_prefix`. GetPrefixOrDefault() should return `my_prefix`", func() {
					component.SetPrefix("my_prefix")
					So(component.GetPrefixOrDefault(), ShouldEqual, "my_prefix")
				})
			})
		})
	})
}

func TestComponentIsUnique(t *testing.T) {
	Convey("Given a new Component", t, func() {
		component := NewComponent()
		Convey("should not return nil", func() {
			So(component, ShouldNotBeNil)
			Convey("component should not be unique", func() {
				So(component.IsUnique(), ShouldBeFalse)
			})
		})
	})
}
