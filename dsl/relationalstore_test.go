package dsl_test

import (
	"github.com/goadesign/gorma"
	gdsl "github.com/goadesign/gorma/dsl"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RelationalStore", func() {
	var sgname, name string
	var dsl func()

	BeforeEach(func() {
		Design = nil
		Errors = nil
		sgname = "production"
		dsl = nil
		name = ""
		gorma.GormaDesign = nil
		InitDesign()

	})

	JustBeforeEach(func() {

		gdsl.StorageGroup(sgname, func() {
			gdsl.Store(name, gorma.MySQL, dsl)
		})

		Run()

	})

	Context("with no DSL", func() {
		BeforeEach(func() {
			name = "mysql"
		})

		It("produces a valid Relational Store definition", func() {
			Ω(Design.Validate()).ShouldNot(HaveOccurred())
			sg := gorma.GormaDesign
			Ω(sg.RelationalStores[name].Name).Should(Equal(name))
		})
	})

	Context("with an already defined Relational Store with the same name", func() {
		BeforeEach(func() {
			name = "mysql"
		})

		It("produces an error", func() {
			gdsl.StorageGroup(sgname, func() {
				gdsl.Store(name, gorma.MySQL, dsl)
			})
			Ω(Errors).Should(HaveOccurred())
		})
	})

	Context("with an already defined Relational Store with a different name", func() {
		BeforeEach(func() {
			sgname = "mysql"
		})

		It("returns an error", func() {
			gdsl.StorageGroup("news", dsl)
			Ω(Errors).Should(HaveOccurred())
		})
	})

	Context("with valid DSL", func() {
		JustBeforeEach(func() {
			Ω(Errors).ShouldNot(HaveOccurred())
			Ω(Design.Validate()).ShouldNot(HaveOccurred())
		})

		Context("with a description", func() {
			const description = "description"

			BeforeEach(func() {
				name = "mysql"
				dsl = func() {
					gdsl.Description(description)
				}
			})

			It("sets the relational store description", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].Description).Should(Equal(description))
			})
			It("auto id generation defaults to true", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoIDFields).Should(Equal(false))
			})
			It("auto timestamps defaults to true", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoTimestamps).Should(Equal(false))
			})
			It("auto soft delete defaults to true", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoSoftDelete).Should(Equal(false))
			})
		})
		Context("with NoAutomaticIDFields", func() {
			BeforeEach(func() {
				name = "mysql"
				dsl = func() {
					gdsl.NoAutomaticIDFields()
				}
			})

			It("auto id generation should be off", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoIDFields).Should(Equal(true))
			})
			It("auto timestamps defaults to true", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoTimestamps).Should(Equal(false))
			})
			It("auto soft delete defaults to true", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoSoftDelete).Should(Equal(false))
			})
		})
		Context("with NoAutomaticTimestamps", func() {
			BeforeEach(func() {
				name = "mysql"
				dsl = func() {
					gdsl.NoAutomaticTimestamps()
				}
			})

			It("auto id generation should be on", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoIDFields).Should(Equal(false))
			})
			It("auto timestamps should be off", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoTimestamps).Should(Equal(true))
			})
			It("auto soft delete should be on", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoSoftDelete).Should(Equal(false))
			})
		})
		Context("with NoAutomaticSoftDelete", func() {
			BeforeEach(func() {
				name = "mysql"
				dsl = func() {
					gdsl.NoAutomaticSoftDelete()
				}
			})

			It("auto id generation should be on", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoIDFields).Should(Equal(false))
			})
			It("auto timestamps should be on", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoTimestamps).Should(Equal(false))
			})
			It("auto soft delete should be off", func() {
				sg := gorma.GormaDesign
				Ω(sg.RelationalStores[name].NoAutoSoftDelete).Should(Equal(true))
			})
		})

	})
})
