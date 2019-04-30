package proton_test

import (
	proton "github.com/SirMetathyst/go-proton"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

// Describe Component Member List ...
var _ = Describe("Component Member List", func() {

	// Adding a nil component member ...
	It("will panic when a nil component member is added", func() {
		componentMemberList := proton.NewComponentMemberList()
		Expect(func() { componentMemberList.AddComponentMember(nil) }).To(Panic())
	})

	// Adding a valid component member ...
	DescribeTable("adding a valid component member",
		func(componentMemberData ComponentMemberData) {
			componentMemberList := proton.NewComponentMemberList()
			componentMember, _ := proton.NewComponentMember(componentMemberData.id, componentMemberData.value, componentMemberData.entityIndexType)
			Expect(componentMemberList.AddComponentMember(componentMember)).To(BeNil())
		}, ValidComponentMemberEntrySlice...,
	)

	// Adding a valid component member alias ...
	DescribeTable("adding a valid component member alias",
		func(componentMemberAliasData ComponentMemberAliasData) {
			componentMemberList := proton.NewComponentMemberList()
			componentMember, _ := proton.NewComponentMemberAlias(componentMemberAliasData.id, componentMemberAliasData.alias, componentMemberAliasData.entityIndexType)
			Expect(componentMemberList.AddComponentMember(componentMember)).To(BeNil())
		}, ValidComponentMemberAliasEntrySlice...,
	)

	// Adding a duplicate component member ...
	DescribeTable("adding a duplicate component member",
		func(componentMemberData ComponentMemberData) {
			componentMemberList := proton.NewComponentMemberList()
			componentMember, _ := proton.NewComponentMember(componentMemberData.id, componentMemberData.value, componentMemberData.entityIndexType)
			componentMemberList.AddComponentMember(componentMember)
			Expect(componentMemberList.AddComponentMember(componentMember)).To(Equal(proton.ErrComponentMemberListTriedToAddDuplicateMemberID))
		}, ValidComponentMemberEntrySlice...,
	)

	// Adding a duplicate component member alias ...
	DescribeTable("adding a duplicate component member alias",
		func(componentMemberAliasData ComponentMemberAliasData) {
			componentMemberList := proton.NewComponentMemberList()
			componentMember, _ := proton.NewComponentMemberAlias(componentMemberAliasData.id, componentMemberAliasData.alias, componentMemberAliasData.entityIndexType)
			componentMemberList.AddComponentMember(componentMember)
			Expect(componentMemberList.AddComponentMember(componentMember)).To(Equal(proton.ErrComponentMemberListTriedToAddDuplicateMemberID))
		}, ValidComponentMemberAliasEntrySlice...,
	)

	// Add component member entries to ComponentMemberList
	addComponentMemberEntries := func(componentMemberList *proton.ComponentMemberList, entries []TableEntry) []*proton.ComponentMember {
		componentMemberSlice := make([]*proton.ComponentMember, 0)
		for _, entry := range entries {
			componentMemberData := entry.Parameters[0].(ComponentMemberData)
			componentMember, _ := proton.NewComponentMember(componentMemberData.id, componentMemberData.value, componentMemberData.entityIndexType)
			componentMemberList.AddComponentMember(componentMember)
			componentMemberSlice = append(componentMemberSlice, componentMember)
		}
		return componentMemberSlice
	}

	// Returning a slice of component member's
	It("will return a slice of component member's", func() {
		componentMemberList := proton.NewComponentMemberList()
		componentMemberSlice := addComponentMemberEntries(componentMemberList, ValidComponentMemberEntrySlice)
		Expect(componentMemberList.ComponentMemberSlice()).To(HaveLen(len(ValidComponentMemberEntrySlice)))
		for _, componentMember := range componentMemberSlice {
			Expect(componentMemberList.ComponentMemberSlice()).To(ContainElement(componentMember))
		}
	})

	// Returning a slice of component member's with entity index's
	It("will return a slice of component member's with entity index's", func() {
		componentMemberList := proton.NewComponentMemberList()
		componentMemberSlice := addComponentMemberEntries(componentMemberList, ValidComponentMemberEntryHalfEntityIndexSlice)
		Expect(componentMemberList.ComponentMembersWithEntityIndex()).To(HaveLen(2))
		for _, componentMember := range componentMemberSlice {
			Expect(componentMemberList.ComponentMemberSlice()).To(ContainElement(componentMember))
		}
	})

})
