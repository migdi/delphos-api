package models_test

import (
	. "github.com/migdi/delphos-api/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Advice", func() {
	Context("when creating advices", func() {
		It("must return advice model attrs", func() {
			user := NewUser("Foo", "bar@bar.com", "foobar123")
			advice := NewAdvice("id", "Content of Advice", *user)
			Expect(advice.ID()).To(Equal("id"))
			Expect(advice.Content()).To(Equal("Content of Advice"))
		})
	})
})
