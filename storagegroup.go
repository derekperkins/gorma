package gorma

import (
	"fmt"
	"sort"

	"github.com/goadesign/goa/dslengine"
)

// IterateStores runs an iterator function once per Relational Store in the
// StorageGroup's Store list.
func (sd *StorageGroupDefinition) IterateStores(it StoreIterator) error {
	if sd.RelationalStores != nil {
		names := make([]string, len(sd.RelationalStores))
		i := 0
		for n := range sd.RelationalStores {
			names[i] = n
			i++
		}
		sort.Strings(names)
		for _, n := range names {
			if err := it(sd.RelationalStores[n]); err != nil {
				return err
			}
		}
	}
	return nil
}

// Context returns the generic definition name used in error messages.
func (sd StorageGroupDefinition) Context() string {
	if sd.Name != "" {
		return fmt.Sprintf("StorageGroup %#v", sd.Name)
	}
	return "unnamed Storage Group"
}

// DSL returns this object's DSL.
func (sd StorageGroupDefinition) DSL() func() {
	return sd.DefinitionDSL
}

// Children returns a slice of this objects children.
func (sd StorageGroupDefinition) Children() []dslengine.Definition {
	var stores []dslengine.Definition
	for _, s := range sd.RelationalStores {
		stores = append(stores, s)
	}
	return stores
}

// IterateSets goes over all the definition sets of the StorageGroup: the
// StorageGroup definition itself, each store definition, models and fields.
func (sd *StorageGroupDefinition) IterateSets(iterator dslengine.SetIterator) {
	// First run the top level StorageGroup

	iterator([]dslengine.Definition{sd})
	sd.IterateStores(func(store *RelationalStoreDefinition) error {
		iterator([]dslengine.Definition{store})
		store.IterateModels(func(model *RelationalModelDefinition) error {
			iterator([]dslengine.Definition{model})
			model.IterateFields(func(field *RelationalFieldDefinition) error {
				iterator([]dslengine.Definition{field})
				return nil
			})
			return nil
		})
		return nil
	})
}
