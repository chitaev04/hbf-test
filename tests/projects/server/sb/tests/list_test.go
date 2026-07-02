package tests

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/suite"

	"HBF-tests/tests/projects/server/sb/suites"
)

func TestGRPCSuiteSBList(t *testing.T) {
	suite.RunSuite(t, new(suites.GRPCSuiteListSB))
}
