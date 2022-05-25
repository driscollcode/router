package mockgen

import (
	"github.com/driscollcode/router/testing-examples/mockgen/mock"
	myHandler "github.com/driscollcode/router/testing-examples/mockgen/my-handler"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unit Tests")
}

var _ = Describe("Router unit tests", func() {
	var (
		mockController *gomock.Controller
		mockRequest    *mock.MockRequest
	)

	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())
		mockRequest = mock.NewMockRequest(mockController)
	})

	AfterEach(func() {
		mockController.Finish()
	})

	Context("Testing a handler", func() {
		When("I call my application code handler", func() {
			When("a required argument - userId - is not supplied", func() {
				It("should return an error", func() {
					mockRequest.EXPECT().ArgExists("userId").Return(false)
					mockRequest.EXPECT().Error("No user Id supplied")

					myHandler.Handler(mockRequest)
				})
			})
			When("the user ID argument is supplied", func() {
				It("should return that argument back with the word 'processed' after it", func() {
					mockRequest.EXPECT().ArgExists("userId").Return(true)
					mockRequest.EXPECT().GetArg("userId").Return("Bob")
					mockRequest.EXPECT().Success("Bob processed")

					myHandler.Handler(mockRequest)
				})
			})
		})
	})
})
