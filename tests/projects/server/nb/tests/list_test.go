package tests

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/suite"

	"HBF-tests/tests/projects/server/nb/suites"
)

func TestGRPCSuiteNBList(t *testing.T) {
	suite.RunSuite(t, new(suites.GRPCSuiteListNB))
}
