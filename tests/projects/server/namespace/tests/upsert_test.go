package tests

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/suite"

	"HBF-tests/tests/projects/server/namespace/suites"
)

func TestGRPCSuiteUpsert(t *testing.T) {
	suite.RunSuite(t, new(suites.GRPCSuiteUpsertNS))
}
