package tests

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/suite"

	"HBF-tests/tests/projects/server/hb/suites"
)

func TestGRPCSuiteHBList(t *testing.T) {
	suite.RunSuite(t, new(suites.GRPCSuiteListHB))
}
