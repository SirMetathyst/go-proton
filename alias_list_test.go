package proton_test

import (
	proton "github.com/SirMetathyst/go-proton"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

// Describe Alias List ...
var _ = Describe("Alias List", func() {

	// Adding nil alias ...
	It("will panic when a nil alias is added", func() {
		aliasList := proton.NewAliasList()
		Expect(func() { aliasList.AddAlias(nil) }).To(Panic())
	})

	// Adding valid alias ...
	DescribeTable("adding a valid alias",
		func(aliasData AliasData) {
			aliasList := proton.NewAliasList()
			alias, _ := proton.NewAlias(aliasData.id, aliasData.value)
			Expect(aliasList.AddAlias(alias)).To(BeNil())
		}, ValidAliasEntrySlice...,
	)

	// Adding duplicate alias ...
	DescribeTable("adding a duplicate alias",
		func(aliasData AliasData) {
			aliasList := proton.NewAliasList()
			alias, _ := proton.NewAlias(aliasData.id, aliasData.value)
			aliasList.AddAlias(alias)
			Expect(aliasList.AddAlias(alias)).To(Equal(proton.ErrAliasListTriedToAddDuplicateAliasID))
		}, ValidAliasEntrySlice...,
	)

	// Return slice of alias
	It("will return a slice of alias", func() {

		// Add alias entries to AliasList
		addAliasEntries := func(aliasList *proton.AliasList, entries []TableEntry) []*proton.Alias {
			aliasSlice := make([]*proton.Alias, 0)
			for _, entry := range entries {
				aliasData := entry.Parameters[0].(AliasData)
				alias, _ := proton.NewAlias(aliasData.id, aliasData.value)
				aliasList.AddAlias(alias)
				aliasSlice = append(aliasSlice, alias)
			}
			return aliasSlice
		}

		aliasList := proton.NewAliasList()
		aliasSlice := addAliasEntries(aliasList, ValidAliasEntrySlice)
		Expect(aliasList.AliasSlice()).To(HaveLen(len(ValidAliasEntrySlice)))

		for _, alias := range aliasSlice {
			Expect(aliasList.AliasSlice()).To(ContainElement(alias))
		}
	})

})
