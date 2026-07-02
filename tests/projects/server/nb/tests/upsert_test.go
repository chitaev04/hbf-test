package tests

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/suite"

	"HBF-tests/tests/projects/server/nb/suites"
)

func TestGRPCSuiteNBUpsert(t *testing.T) {
	suite.RunSuite(t, new(suites.GRPCSuiteUpsertNB))
}
