package tests

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/suite"

	"HBF-tests/tests/projects/server/sb/suites"
)

func TestGRPCSuiteSBUpsert(t *testing.T) {
	suite.RunSuite(t, new(suites.GRPCSuiteUpsertSB))
}
