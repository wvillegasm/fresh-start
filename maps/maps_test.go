package maps_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/wvillegasm/fresh-start/maps"
)

var _ = Describe("Maps", func() {

	const expectedHomePage string = "/"

	var currentHomePage string

	BeforeEach(func() {
		currentHomePage = GetUrl(HOME_PAGE)
	})

	Describe("Pages", func() {
		It("should get home page /", func() {
			Expect(currentHomePage).To(Equal(expectedHomePage))
		})

		It("should get contact page /contact", func() {
			Expect(GetUrl(CONTACTS)).To(Equal("/contacts"))
		})
	})
})
