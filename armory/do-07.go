package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_07() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST provide design documentation demonstrating all actions and actors within the system.",
		ControlID:   "OSPS-AC-01",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_07_T01)
	if result.Tests["DO_07_T01"].Passed {
		result.ExecuteTest(DO_07_T02)
	}

	return "DO_07", result
}

func DO_07_T01() pluginkit.TestResult {
	guideLocation := Data.Rest().Insights.Repository.Security.Assessments.Self.Evidence
	found := guideLocation != ""

	testResult := pluginkit.TestResult{
		Description: "Determine whether the project's Security Insights data specifies a self assessment location.",
		Function:    utils.CallerPath(0),
		Passed:      found,
		Message:     "Self-assessment location specified in Security Insights: %t",
		Value:       guideLocation,
	}

	return testResult
}

func DO_07_T02() pluginkit.TestResult {
	_, err := makeApiCall(Data.Rest().Insights.Repository.Security.Assessments.Self.Evidence)

	testResult := pluginkit.TestResult{
		Description: "Verifying that an artifact exists at the location specified for the self assessment.",
		Function:    utils.CallerPath(0),
		Passed:      err == nil,
		Message:     "URL for Self Assessment can be reached: %t",
	}

	if err != nil {
		testResult.Value = err
	}

	return testResult
}
