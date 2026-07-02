package tests

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/suite"

	"HBF-tests/tests/projects/server/ag/suites"
)

func TestGRPCSuiteAGList(t *testing.T) {
	suite.RunSuite(t, new(suites.GRPCSuiteListAG))
}
