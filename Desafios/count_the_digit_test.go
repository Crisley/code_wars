package desafios

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func dotest(n int, d int, exp int) {
	var ans = NbDig(n, d)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests NbDig", func() {

	It("should handle basic cases", func() {
		dotest(25, 1, 11)
		dotest(550, 5, 213)
		dotest(5750, 0, 4700)
	})
})

func TestCountTheDigit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Count the Digit")
}
