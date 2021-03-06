package dsl

import "github.com/goadesign/gorma"

// Store represents a database.  Gorma lets you specify
// a database type, but it's currently not used for any generation
// logic.
func Store(name string, storeType gorma.RelationalStorageType, dsl func()) {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.

	checkInit()
	if s, ok := storageGroupDefinition(true); ok {
		if s.RelationalStores == nil {
			s.RelationalStores = make(map[string]*gorma.RelationalStoreDefinition)
		}
		store, ok := s.RelationalStores[name]
		if !ok {
			store = &gorma.RelationalStoreDefinition{
				Name:             name,
				DefinitionDSL:    dsl,
				Parent:           s,
				Type:             storeType,
				RelationalModels: make(map[string]*gorma.RelationalModelDefinition),
			}
		}
		s.RelationalStores[name] = store
	}

}

// NoAutomaticIDFields applies to a `Store` type.  It allows you
// to turn off the default behavior that will automatically create
// an ID/int Primary Key for each model.
func NoAutomaticIDFields() {
	if s, ok := relationalStoreDefinition(false); ok {
		s.NoAutoIDFields = true
	}
}

// NoAutomaticTimestamps applies to a `Store` type.  It allows you
// to turn off the default behavior that will automatically create
// an `CreatedAt` and `UpdatedAt` fields for each model.
func NoAutomaticTimestamps() {
	if s, ok := relationalStoreDefinition(false); ok {
		s.NoAutoTimestamps = true
	}
}

// NoAutomaticSoftDelete applies to a `Store` type.  It allows
// you to turn off the default behavior that will automatically
// create a `DeletedAt` field (*time.Time) that acts as a
// soft-delete filter for your models.
func NoAutomaticSoftDelete() {
	if s, ok := relationalStoreDefinition(false); ok {
		s.NoAutoSoftDelete = true
	}
}
