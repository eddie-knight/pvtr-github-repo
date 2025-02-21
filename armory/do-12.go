package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_12() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST contain instructions to verify the integrity and authenticity of the release assets, including the expected identity of the person or process authoring the software release.",
		ControlID:   "OSPS-DO-12",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_12_T01)

	if result.Tests["DO_12_T01"].Passed {
		result.ExecuteTest(DO_12_T02)
	}

	return "DO_12", result
}

func DO_12_T01() pluginkit.TestResult {

	guideLocation := Data.Rest().Insights.Project.Documentation.SignatureVerification

	testResult := pluginkit.TestResult{
		Description: "Ensure the project's Security Insights data specifies a signature verification guide location.",
		Function:    utils.CallerPath(0),
		Passed:      guideLocation != "",
		Message:     fmt.Sprintf("Signature verification guide location specified in Security Insights: %v", guideLocation != ""),
		Value:       guideLocation,
	}

	return testResult
}

func DO_12_T02() pluginkit.TestResult {

	_, err := makeApiCall(Data.Rest().Insights.Project.Documentation.SignatureVerification)

	testResult := pluginkit.TestResult{
		Description: "Check if the signature verification guide is accessible via HTTP.",
		Function:    utils.CallerPath(0),
		Passed:      err == nil,
		Message:     fmt.Sprintf("Signature verification guide is accessible via HTTP: %v", err == nil),
		Value:       Data.Rest().Insights.Project.Documentation.SignatureVerification,
	}
	return testResult
}
