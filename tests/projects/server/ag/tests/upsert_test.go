package tests

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/suite"

	"HBF-tests/tests/projects/server/ag/suites"
)

func TestGRPCSuiteAGUpsert(t *testing.T) {
	suite.RunSuite(t, new(suites.GRPCSuiteUpsertAG))
}
