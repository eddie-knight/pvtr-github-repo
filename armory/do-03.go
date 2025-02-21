package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_03() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST provide user guides for all basic functionality.",
		ControlID:   "OSPS-DO-03",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_03_T01)
	if result.Tests["DO_03_T01"].Passed {
		result.ExecuteTest(DO_03_T02)
	}
	return "DO_03", result
}

func DO_03_T01() pluginkit.TestResult {
	guideLocation := Data.Rest().Insights.Project.Documentation.DetailedGuide
	found := guideLocation != ""

	testResult := pluginkit.TestResult{
		Description: "Determine whether the project's Security Insights data specifies a detailed-guide location.",
		Function:    utils.CallerPath(0),
		Passed:      found,
		Message:     fmt.Sprintf("Detailed Guide docs location specified in Security Insights: %t", found),
		Value:       guideLocation,
	}

	return testResult
}

func DO_03_T02() pluginkit.TestResult {
	_, err := makeApiCall(Data.Rest().Insights.Project.Documentation.DetailedGuide)

	testResult := pluginkit.TestResult{
		Description: "Verifying that an artifact exists at the location specified for the detailed-guide.",
		Function:    utils.CallerPath(0),
		Passed:      err == nil,
		Message:     fmt.Sprintf("URL for Detailed Guide can be reached: %t", err == nil),
	}

	if err != nil {
		testResult.Value = err
	}

	return testResult
}
