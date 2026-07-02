package tests

import (
	"HBF-tests/tests/projects/server/namespace/suites"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

func TestGRPCSuiteUpset(t *testing.T) {
	suite.RunSuite(t, new(suites.GRPCSuiteUpsetNS))
}
