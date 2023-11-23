package integration

import (
	"playground/app/handlers"
	"playground/tests/common/goat"
	"playground/tests/common/test_driver"
	"testing"
)

func Test(t *testing.T) {
	suite := goat.NewSuite(t)

	suite.Test("should return 200", func(t *testing.T) {

		// ARRANGE

		driver := test_driver.New()

		driver.Start()
		defer driver.Stop()

		// ACT

		getHealthResponse, err := driver.Get("/health")

		// ASSERT

		suite.Expect(err).ToBeNil()
		suite.Expect(getHealthResponse.StatusCode).ToEqual(200)

		actualBody := DecodeJsonBody[handlers.GetHealthResponse](getHealthResponse)
		expectBody := handlers.GetHealthResponse{Status: "ok"}

		suite.Expect(actualBody).ToEqual(expectBody)
	})
}
