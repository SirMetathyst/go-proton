package proton_test

import (
	proton "github.com/SirMetathyst/go-proton"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

// ComponentMemberData ...
type ComponentMemberData struct {
	id              string
	value           string
	entityIndexType proton.EntityIndexType
	expectedErr     error
}

// ComponentMemberAliasData ...
type ComponentMemberAliasData struct {
	id              string
	alias           *proton.Alias
	entityIndexType proton.EntityIndexType
	expectedErr     error
}

var (

	// InvalidComponentMemberEntrySlice ...
	InvalidComponentMemberEntrySlice = []TableEntry{
		Entry("", ComponentMemberData{"", "", proton.EntityIndexTypeNone, proton.ErrComponentMemberIDUndefined}),                // 0
		Entry("", ComponentMemberData{"", "", proton.EntityIndexTypeSingle, proton.ErrComponentMemberIDUndefined}),              // 1
		Entry("", ComponentMemberData{"", "", proton.EntityIndexTypeMultiple, proton.ErrComponentMemberIDUndefined}),            // 2
		Entry("", ComponentMemberData{" ", "", proton.EntityIndexTypeNone, proton.ErrComponentMemberIDContainsWhitespace}),      // 3
		Entry("", ComponentMemberData{"id", "", proton.EntityIndexTypeNone, proton.ErrComponentMemberValueUndefined}),           // 4
		Entry("", ComponentMemberData{"id", " ", proton.EntityIndexTypeNone, proton.ErrComponentMemberValueContainsWhitespace}), // 5
		Entry("", ComponentMemberData{"id", "value", 3, proton.ErrComponentMemberEntityIndexInvalid}),                           // 6
		Entry("", ComponentMemberData{"id", "value", 4, proton.ErrComponentMemberEntityIndexInvalid}),                           // 7
		Entry("", ComponentMemberData{"id", "value", -1, proton.ErrComponentMemberEntityIndexInvalid}),                          // 8
	}

	// InvalidComponentMemberAliasEntrySlice ...
	InvalidComponentMemberAliasEntrySlice = []TableEntry{
		Entry("", ComponentMemberAliasData{"", nil, proton.EntityIndexTypeNone, proton.ErrComponentMemberIDUndefined}),           // 0
		Entry("", ComponentMemberAliasData{" ", nil, proton.EntityIndexTypeNone, proton.ErrComponentMemberIDContainsWhitespace}), // 1
		Entry("", ComponentMemberAliasData{"id", nil, proton.EntityIndexTypeNone, proton.ErrComponentMemberAliasShouldNotBeNil}), // 2
		Entry("", ComponentMemberAliasData{"id", ValidAlias, 3, proton.ErrComponentMemberEntityIndexInvalid}),                    // 3
		Entry("", ComponentMemberAliasData{"id", ValidAlias, -1, proton.ErrComponentMemberEntityIndexInvalid}),                   // 4
	}

	// ValidComponentMemberEntrySlice ...
	ValidComponentMemberEntrySlice = []TableEntry{
		Entry("", ComponentMemberData{"id", "value", proton.EntityIndexTypeNone, nil}),               // 0
		Entry("", ComponentMemberData{"id2", "value", proton.EntityIndexTypeSingle, nil}),            // 1
		Entry("", ComponentMemberData{"idjd242*£^()", "value", proton.EntityIndexTypeMultiple, nil}), // 2
		Entry("", ComponentMemberData{"the_id", "the_value", proton.EntityIndexTypeNone, nil}),       // 3
	}

	// ValidComponentMemberAliasEntrySlice ...
	ValidComponentMemberAliasEntrySlice = []TableEntry{
		Entry("", ComponentMemberAliasData{"id", ValidAlias, proton.EntityIndexTypeNone, nil}),               // 0
		Entry("", ComponentMemberAliasData{"id2", ValidAlias, proton.EntityIndexTypeSingle, nil}),            // 1
		Entry("", ComponentMemberAliasData{"idjd242*£^()", ValidAlias, proton.EntityIndexTypeMultiple, nil}), // 2
		Entry("", ComponentMemberAliasData{"the_id", ValidAlias, proton.EntityIndexTypeNone, nil}),           // 3
	}

	// ValidComponentMemberEntryHalfEntityIndexSlice ...
	ValidComponentMemberEntryHalfEntityIndexSlice = []TableEntry{
		Entry("", ComponentMemberData{"id", "value", proton.EntityIndexTypeNone, nil}),               // 0
		Entry("", ComponentMemberData{"id2", "value", proton.EntityIndexTypeSingle, nil}),            // 1
		Entry("", ComponentMemberData{"idjd242*£^()", "value", proton.EntityIndexTypeMultiple, nil}), // 2
		Entry("", ComponentMemberData{"the_id", "the_value", proton.EntityIndexTypeNone, nil}),       // 3
	}
)

// Describe Component Member ...
var _ = Describe("Component Member", func() {

	// Creating an invalid component member ...
	DescribeTable("creating an invalid component member without an alias",
		func(componentMemberData ComponentMemberData) {
			componentMember, err := proton.NewComponentMember(componentMemberData.id, componentMemberData.value, componentMemberData.entityIndexType)
			Expect(err).To(Equal(componentMemberData.expectedErr))
			Expect(componentMember).To(BeNil())
		}, InvalidComponentMemberEntrySlice...,
	)

	// Creating an invalid component member ...
	DescribeTable("creating an invalid component member with an alias",
		func(componentMemberAliasData ComponentMemberAliasData) {
			componentMember, err := proton.NewComponentMemberAlias(componentMemberAliasData.id, componentMemberAliasData.alias, componentMemberAliasData.entityIndexType)
			Expect(err).To(Equal(componentMemberAliasData.expectedErr))
			Expect(componentMember).To(BeNil())
		}, InvalidComponentMemberAliasEntrySlice...,
	)

	// Creating a valid component member ...
	DescribeTable("creating a valid component member without an alias",
		func(componentMemberData ComponentMemberData) {
			componentMember, err := proton.NewComponentMember(componentMemberData.id, componentMemberData.value, componentMemberData.entityIndexType)
			Expect(err).To(BeNil())
			Expect(componentMember).ToNot(BeNil())
			Expect(componentMember.ID()).To(Equal(proton.String(componentMemberData.id)))
			Expect(componentMember.Value()).To(Equal(proton.String(componentMemberData.value)))
			Expect(componentMember.EntityIndexType()).To(Equal(componentMemberData.entityIndexType))
		}, ValidComponentMemberEntrySlice...,
	)

	// Creating a valid component member ...
	DescribeTable("creating a valid component member with an alias",
		func(componentMemberAliasData ComponentMemberAliasData) {
			componentMember, err := proton.NewComponentMemberAlias(componentMemberAliasData.id, componentMemberAliasData.alias, componentMemberAliasData.entityIndexType)
			Expect(err).To(BeNil())
			Expect(componentMember).ToNot(BeNil())
			Expect(componentMember.ID()).To(Equal(proton.String(componentMemberAliasData.id)))
			Expect(componentMember.Value()).To(Equal(proton.String(componentMemberAliasData.alias.Value())))
			Expect(componentMember.Alias()).To(Equal(componentMemberAliasData.alias))
			Expect(componentMember.EntityIndexType()).To(Equal(componentMemberAliasData.entityIndexType))
		}, ValidComponentMemberAliasEntrySlice...,
	)

})
