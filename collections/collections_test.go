package collections_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/wvillegasm/fresh-start/collections"
)

var _ = Describe("Collections", func() {
	Context("Slice of Names ", func() {
		It("Should get an array of three elements", func() {
			actual := Names()
			expected := 3
			Expect(len(actual)).To(Equal(expected))
		})

		It("Should get only two names", func() {
			actual := Names()[0:2]
			expected := 2
			Expect(len(actual)).To(Equal(expected))
		})
	})

	Context("Conditionals", func() {
		It("Should return 10", func() {
			actual := Sum(1, 9)
			expected := 10
			Expect(actual).To(Equal(expected))
		})
	})

	Context("Ages", func() {
		It("Should return category when age is provided", func() {
			dataAges := []struct {
				age      uint8
				expected string
			}{
				{5, "Child"},
				{12, "Child"},
				{13, "Teenager"},
				{18, "Teenager"},
				{20, "Adult"},
				{68, "Old man"},
			}

			for _, ageData := range dataAges {
				actual, err := Ages(ageData.age)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(actual).To(Equal(ageData.expected))

			}
		})

		It("Should return error when negative age is passed", func() {
			_, err := Ages(0)
			Expect(err).To(HaveOccurred())
		})
	})

	Context("Days of the week", func() {
		It("Should return correct Day name given a value", func() {
			days := []struct {
				number   uint8
				expected string
			}{
				{1, "MONDAY"},
				{0, "SUNDAY"},
				{5, "FRIDAY"},
			}

			for _, day := range days {
				actual, err := WeekDays(day.number)
				Expect(err).Should(Succeed())
				Expect(actual).To(Equal(day.expected))
			}
		})

		It("Should get an error for day greater than 6", func() {
			days := []uint8{7, 8, 9}

			for _, day := range days {
				_, err := WeekDays(day)
				Expect(err).Should(MatchError(err))
			}
		})
	})
	Context("Languages", func() {
		It("Should get SPANISH", func() {
			actual, err := Lang("es")
			expected := "SPANISH"
			Expect(err).Should(Succeed())
			Expect(actual).To(Equal(expected))
		})
		It("Should get {language not supported} error", func() {
			_, err := Lang("fr")
			expectedErrorMsg := "language not supported"
			Expect(err.Error()).To(Equal(expectedErrorMsg))
			Expect(err).ShouldNot(Succeed())
		})
	})
})
