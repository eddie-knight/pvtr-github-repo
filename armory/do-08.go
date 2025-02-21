package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_08() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST include a policy that defines a threshold for remediation of SCA findings related to vulnerabilities and licenses.",
		ControlID:   "OSPS-DO-08",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_08_T01)
	if result.Tests["DO_08_T01"].Passed {
		result.ExecuteTest(DO_08_T02)
	}

	return "DO_08", result
}

func DO_08_T01() pluginkit.TestResult {
	policyLocation := Data.Rest().Insights.Repository.Documentation.DependencyManagement
	found := policyLocation != ""

	testResult := pluginkit.TestResult{
		Description: "Determine whether the project's Security Insights data specifies a dependency-management policy location.",
		Function:    utils.CallerPath(0),
		Passed:      found,
		Message:     fmt.Sprintf("Dependency Management Policy docs location specified in Security Insights: %t", found),
		Value:       policyLocation,
	}

	return testResult
}

func DO_08_T02() pluginkit.TestResult {
	_, err := makeApiCall(Data.Rest().Insights.Repository.Documentation.DependencyManagement)

	testResult := pluginkit.TestResult{
		Description: "Verifying that an artifact exists at the location specified for the dependency-management policy.",
		Function:    utils.CallerPath(0),
		Passed:      err == nil,
		Message:     "URL for Dependency Management Policy can be reached",
	}

	if err != nil {
		testResult.Value = err
	}

	return testResult
}
