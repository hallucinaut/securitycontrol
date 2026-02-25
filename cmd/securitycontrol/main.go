package main

import (
	"fmt"
	"os"

	"github.com/hallucinaut/securitycontrol/pkg/control"
	"github.com/hallucinaut/securitycontrol/pkg/validate"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "validate":
		validateControls()
	case "test":
		if len(os.Args) < 3 {
			fmt.Println("Error: control ID required")
			printUsage()
			return
		}
//		testControl(os.Args[2])
	case "controls":
		listControls()
	case "report":
		generateReport()
	case "status":
		checkStatus()
	case "version":
		fmt.Printf("securitycontrol version %s\n", version)
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
	}
}

func printUsage() {
	fmt.Printf(`securitycontrol - Security Control Validation Engine

Usage:
  securitycontrol <command> [options]

Commands:
  validate     Validate all security controls
  test <id>    Test specific control
  controls     List available controls
  report       Generate validation report
  status       Check control status
  version      Show version information
  help         Show this help message

Examples:
  securitycontrol validate
  securitycontrol test ctrl-001
  securitycontrol controls
`,)
}

func validateControls() {
	fmt.Println("Security Control Validation")
	fmt.Println("==========================")
	fmt.Println()

	// Create validator
	validator := control.NewControlValidator()

	// Add common controls
	commonControls := control.CreateCommonControls()
	for _, ctrl := range commonControls {
		validator.AddControl(ctrl)
	}

	fmt.Println("Controls to Validate:")
	for i, ctrl := range commonControls {
		fmt.Printf("  [%d] %s (%s)\n", i+1, ctrl.Name, ctrl.Category)
	}
	fmt.Println()

	fmt.Println("Running Validation...")
	fmt.Println()

	// Validate controls
	for _, ctrl := range commonControls {
		result := validator.ValidateControl(ctrl.ID)
		if result != nil {
			fmt.Printf("[%s] %s\n", result.Status, ctrl.Name)
			fmt.Printf("    Effectiveness: %.1f%%\n", result.Effectiveness*100)
			fmt.Printf("    Confidence: %.1f%%\n", result.Confidence*100)
			if len(result.Issues) > 0 {
				fmt.Printf("    Issues: %d\n", len(result.Issues))
			}
			fmt.Println()
		}
	}

	fmt.Println(control.GenerateReport(validator))
}

//func testControl(controlID string) {
//	fmt.Printf("Testing Control: %s\n", controlID)
//	fmt.Println()
//
//	// Create validator
//	validator := validate.NewControlValidator()
//
//	// Add common tests
//	commonTests := validate.CreateCommonControlTests()
//	for _, test := range commonTests {
//		validator.AddControlTest(test)
//	}
//
//	// Find and run test
//	found := false
//	for _, test := range commonTests {
//		if test.ID == controlID {
//			fmt.Printf("Test: %s\n", test.Name)
//			fmt.Printf("Description: %s\n", test.Description)
//			fmt.Printf("Method: %s\n\n", test.Method)
//
//			// Run test
//			result := validator.ValidateControlTest(test)
//			fmt.Printf("Result: %s\n", result.ValidationResult)
//			fmt.Printf("Effectiveness: %.1f%%\n", result.Effectiveness*100)
//			fmt.Printf("Risk Remaining: %.1f%%\n", result.RiskRemaining*100)
//
//			if len(result.Recommendations) > 0 {
//				fmt.Println("\nRecommendations:")
//				for _, rec := range result.Recommendations {
//					fmt.Printf("  • %s\n", rec)
//				}
//			}
//
//			found = true
//			break
//		}
//	}
//
//	if !found {
//		fmt.Println("Control not found:", controlID)
//	}
//}
//
func listControls() {
	fmt.Println("Available Security Controls")
	fmt.Println("===========================")
	fmt.Println()

	controls := control.CreateCommonControls()

	fmt.Println("Controls by Category:")
	fmt.Println()

	// Group by category
	categories := map[control.ControlCategory][]control.SecurityControl{}
	for _, ctrl := range controls {
		categories[ctrl.Category] = append(categories[ctrl.Category], ctrl)
	}

	for category, ctrls := range categories {
		fmt.Printf("%s Controls:\n", category)
		for i, ctrl := range ctrls {
			fmt.Printf("  [%d] %s (%s)\n", i+1, ctrl.Name, ctrl.Status)
			fmt.Printf("      Risk Reduction: %.1f%%\n", ctrl.RiskReduction*100)
			fmt.Printf("      Owner: %s\n", ctrl.Owner)
			fmt.Println()
		}
	}

	fmt.Printf("Total Controls: %d\n", len(controls))
}

func generateReport() {
	fmt.Println("Generate Validation Report")
	fmt.Println("=========================")
	fmt.Println()

	// Create validators
	controlValidator := control.NewControlValidator()
	validateValidator := validate.NewControlValidator()

	// Add controls
	commonControls := control.CreateCommonControls()
	for _, ctrl := range commonControls {
		controlValidator.AddControl(ctrl)
	}

	// Add tests
	commonTests := validate.CreateCommonControlTests()
	for _, test := range commonTests {
		validateValidator.AddControlTest(test)
	}

	// Generate reports
	fmt.Println("=== Control Validation Report ===")
	fmt.Println(control.GenerateReport(controlValidator))

	fmt.Println("\n=== Test Validation Report ===")
	fmt.Println(validate.GenerateValidationReport(validateValidator))
}

func checkStatus() {
	fmt.Println("Security Control Status")
	fmt.Println("=======================")
	fmt.Println()

	validator := control.NewControlValidator()

	// Add common controls
	commonControls := control.CreateCommonControls()
	for _, ctrl := range commonControls {
		validator.AddControl(ctrl)
	}

	fmt.Println("Control Status Summary:")
	fmt.Println()

	// Count by status
	statusCount := make(map[control.ControlStatus]int)
	for _, ctrl := range commonControls {
		statusCount[ctrl.Status]++
	}

	for status, count := range statusCount {
		fmt.Printf("%s: %d\n", status, count)
	}

	fmt.Println()

	fmt.Println("Controls by Effectiveness:")
	fmt.Println()

	// Validate all controls
	for _, ctrl := range commonControls {
		result := validator.ValidateControl(ctrl.ID)
		if result != nil {
			fmt.Printf("[%s] %.1f%% effective - %s\n", result.Status, result.Effectiveness*100, ctrl.Name)
		}
	}
}