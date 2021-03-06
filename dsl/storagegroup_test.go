package dsl_test

import (
	"github.com/goadesign/gorma"
	gdsl "github.com/goadesign/gorma/dsl"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/dsl"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StorageGroup", func() {
	var name string
	var dsl func()

	BeforeEach(func() {
		Design = nil
		Errors = nil
		name = "mysql"
		dsl = nil
		gorma.GormaDesign = nil
		InitDesign()

	})

	JustBeforeEach(func() {

		gdsl.StorageGroup(name, dsl)

		RunDSL()

	})

	Context("with no DSL", func() {
		BeforeEach(func() {
			name = "mysql"
		})

		It("produces a valid Storage Group definition", func() {
			Ω(Design.Validate()).ShouldNot(HaveOccurred())
			Ω(gorma.GormaDesign.Name).Should(Equal(name))
		})
	})

	Context("with an already defined Storage Group with the same name", func() {
		BeforeEach(func() {
			name = "mysql"
		})

		It("produces an error", func() {
			gdsl.StorageGroup(name, dsl)
			Ω(Errors).Should(HaveOccurred())
		})
	})

	Context("with an already defined Storage Group with a different name", func() {
		BeforeEach(func() {
			name = "mysql"
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
				dsl = func() {
					gdsl.Description(description)
				}
			})

			It("sets the storage group description", func() {
				Ω(gorma.GormaDesign.Description).Should(Equal(description))
			})
		})

	})
})
