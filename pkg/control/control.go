// Package control provides security control definitions and management.
package control

import (
	"fmt"
	"time"
)

// ControlCategory represents a category of security control.
type ControlCategory string

const (
	CategoryPreventive ControlCategory = "preventive"
	CategoryDetective  ControlCategory = "detective"
	CategoryCorrective ControlCategory = "corrective"
	CategoryDeterrent  ControlCategory = "deterrent"
	CategoryRecovery   ControlCategory = "recovery"
)

// ControlType represents a type of security control.
type ControlType string

const (
	TypeTechnical ControlType = "technical"
	TypeAdministrative ControlType = "administrative"
	TypePhysical ControlType = "physical"
)

// ControlStatus represents a control status.
type ControlStatus string

const (
	StatusImplemented ControlStatus = "implemented"
	StatusPartiallyImplemented ControlStatus = "partially_implemented"
	StatusNotImplemented ControlStatus = "not_implemented"
	StatusDeprecated ControlStatus = "deprecated"
)

// SecurityControl represents a security control.
type SecurityControl struct {
	ID              string
	Name            string
	Description     string
	Category        ControlCategory
	Type            ControlType
	SubCategory     string
	RiskReduction   float64
	Implementation  string
	Verification    string
	Maintenance     string
	Owner           string
	Status          ControlStatus
	LastVerified    time.Time
	NextReview      time.Time
	Evidence        []string
	References      []string
}

// ControlFramework represents a security control framework.
type ControlFramework struct {
	Name          string
	Version       string
	Description   string
	Controls      []SecurityControl
	LastUpdated   time.Time
}

// ControlValidator validates security controls.
type ControlValidator struct {
	controls []SecurityControl
	results  []ControlValidationResult
}

// ControlValidationResult represents a control validation result.
type ControlValidationResult struct {
	ControlID      string
	ControlName    string
	Status         string
	Effectiveness  float64
	Confidence     float64
	Issues         []string
	Evidence       []string
	Recommendations []string
	ValidatedAt    time.Time
}

// NewControlValidator creates a new control validator.
func NewControlValidator() *ControlValidator {
	return &ControlValidator{
		controls: make([]SecurityControl, 0),
		results:  make([]ControlValidationResult, 0),
	}
}

// AddControl adds a security control.
func (v *ControlValidator) AddControl(control SecurityControl) {
	v.controls = append(v.controls, control)
}

// GetControls returns all controls.
func (v *ControlValidator) GetControls() []SecurityControl {
	return v.controls
}

// GetControlsByCategory returns controls by category.
func (v *ControlValidator) GetControlsByCategory(category ControlCategory) []SecurityControl {
	var result []SecurityControl
	for _, control := range v.controls {
		if control.Category == category {
			result = append(result, control)
		}
	}
	return result
}

// GetControlsByStatus returns controls by status.
func (v *ControlValidator) GetControlsByStatus(status ControlStatus) []SecurityControl {
	var result []SecurityControl
	for _, control := range v.controls {
		if control.Status == status {
			result = append(result, control)
		}
	}
	return result
}

// ValidateControl validates a security control.
func (v *ControlValidator) ValidateControl(controlID string) *ControlValidationResult {
	control := v.getControlByID(controlID)
	if control == nil {
		return nil
	}

	result := &ControlValidationResult{
		ControlID:   control.ID,
		ControlName: control.Name,
		Status:      "VALIDATING",
		Effectiveness: 0.0,
		Confidence:  0.0,
		Issues:      make([]string, 0),
		Evidence:    make([]string, 0),
		Recommendations: make([]string, 0),
		ValidatedAt: time.Now(),
	}

	// Validate control implementation
	effective := v.validateControlImplementation(control)
	result.Effectiveness = effective

	// Check for issues
	issues := v.identifyIssues(control)
	result.Issues = issues

	// Determine confidence
	confidence := v.calculateConfidence(control, issues)
	result.Confidence = confidence

	// Generate recommendations
	recommendations := v.generateRecommendations(control, issues)
	result.Recommendations = recommendations

	// Determine overall status
	if len(issues) == 0 && effective >= 0.9 {
		result.Status = "EFFECTIVE"
	} else if effective >= 0.7 {
		result.Status = "PARTIALLY_EFFECTIVE"
	} else {
		result.Status = "INEFFECTIVE"
	}

	v.results = append(v.results, *result)
	return result
}

// validateControlImplementation validates control implementation.
func (v *ControlValidator) validateControlImplementation(control SecurityControl) float64 {
	// In production: perform actual validation tests
	// For demo: return simulated effectiveness

	effective := 0.7 // Default effectiveness

	// Check if control is implemented
	if control.Status == StatusImplemented {
		effective = 0.9
	} else if control.Status == StatusPartiallyImplemented {
		effective = 0.6
	} else if control.Status == StatusNotImplemented {
		effective = 0.0
	}

	return effective
}

// identifyIssues identifies issues with control.
func (v *ControlValidator) identifyIssues(control SecurityControl) []string {
	var issues []string

	// Check if control has evidence
	if len(control.Evidence) == 0 {
		issues = append(issues, "No evidence provided for control implementation")
	}

	// Check if control has recent verification
	if control.LastVerified.IsZero() || control.LastVerified.Before(time.Now().AddDate(0, -6, 0)) {
		issues = append(issues, "Control not verified in last 6 months")
	}

	// Check if control has owner
	if control.Owner == "" {
		issues = append(issues, "Control owner not assigned")
	}

	return issues
}

// calculateConfidence calculates validation confidence.
func (v *ControlValidator) calculateConfidence(control SecurityControl, issues []string) float64 {
	confidence := 0.5

	// Adjust confidence based on evidence
	if len(control.Evidence) > 0 {
		confidence += 0.2
	}

	// Adjust confidence based on issues
	if len(issues) == 0 {
		confidence += 0.3
	} else {
		confidence -= float64(len(issues)) * 0.1
	}

	if confidence < 0.0 {
		confidence = 0.0
	}
	if confidence > 1.0 {
		confidence = 1.0
	}

	return confidence
}

// generateRecommendations generates recommendations for control.
func (v *ControlValidator) generateRecommendations(control SecurityControl, issues []string) []string {
	var recommendations []string

	for _, issue := range issues {
		switch {
		case issue == "No evidence provided for control implementation":
			recommendations = append(recommendations, "Provide evidence of control implementation")
		case issue == "Control not verified in last 6 months":
			recommendations = append(recommendations, "Schedule control verification")
		case issue == "Control owner not assigned":
			recommendations = append(recommendations, "Assign control owner")
		default:
			recommendations = append(recommendations, "Address: "+issue)
		}
	}

	return recommendations
}

// GetValidationResults returns all validation results.
func (v *ControlValidator) GetValidationResults() []ControlValidationResult {
	return v.results
}

// validateControl validates control.
func (v *ControlValidator) validateControl(control SecurityControl) *ControlValidationResult {
	result := &ControlValidationResult{
		ControlID:   control.ID,
		ControlName: control.Name,
		Status:      "VALIDATING",
		Effectiveness: 0.0,
		Confidence:  0.0,
		Issues:      make([]string, 0),
		Evidence:    make([]string, 0),
		Recommendations: make([]string, 0),
		ValidatedAt: time.Now(),
	}

	effective := v.validateControlImplementation(control)
	result.Effectiveness = effective

	issues := v.identifyIssues(control)
	result.Issues = issues

	confidence := v.calculateConfidence(control, issues)
	result.Confidence = confidence

	recommendations := v.generateRecommendations(control, issues)
	result.Recommendations = recommendations

	if len(issues) == 0 && effective >= 0.9 {
		result.Status = "EFFECTIVE"
	} else if effective >= 0.7 {
		result.Status = "PARTIALLY_EFFECTIVE"
	} else {
		result.Status = "INEFFECTIVE"
	}

	return result
}

// getControlByID retrieves a control by ID.
func (v *ControlValidator) getControlByID(id string) *SecurityControl {
	for i := range v.controls {
		if v.controls[i].ID == id {
			return &v.controls[i]
		}
	}
	return nil
}

// CreateCommonControls creates common security controls.
func CreateCommonControls() []SecurityControl {
	return []SecurityControl{
		{
			ID:          "ctrl-001",
			Name:        "Access Control Policy",
			Description: "Policy governing access to systems and data",
			Category:    CategoryPreventive,
			Type:        TypeAdministrative,
			SubCategory: "Access Management",
			RiskReduction: 0.3,
			Implementation: "Documented access control policy enforced through IAM",
			Verification: "Review policy documents and access logs",
			Owner:       "Security Team",
			Status:      StatusImplemented,
			LastVerified:  time.Now().AddDate(0, -3, 0),
			Evidence:    []string{"policy-access-control.pdf", "iam-configuration.json"},
			References:  []string{"NIST-800-53-AC-1"},
		},
		{
			ID:          "ctrl-002",
			Name:        "Multi-Factor Authentication",
			Description: "MFA for all user access to systems",
			Category:    CategoryPreventive,
			Type:        TypeTechnical,
			SubCategory: "Authentication",
			RiskReduction: 0.4,
			Implementation: "MFA enforced for all user accounts",
			Verification: "Test MFA enforcement",
			Owner:       "IT Operations",
			Status:      StatusImplemented,
			LastVerified:  time.Now().AddDate(0, -1, 0),
			Evidence:    []string{"mfa-configuration.json", "audit-log.json"},
			References:  []string{"NIST-800-53-IA-2"},
		},
		{
			ID:          "ctrl-003",
			Name:        "Security Monitoring",
			Description: "Continuous security monitoring of systems",
			Category:    CategoryDetective,
			Type:        TypeTechnical,
			SubCategory: "Monitoring",
			RiskReduction: 0.35,
			Implementation: "SIEM and IDS/IPS deployed",
			Verification: "Review monitoring dashboards",
			Owner:       "SOC Team",
			Status:      StatusImplemented,
			LastVerified:  time.Now().AddDate(0, -2, 0),
			Evidence:    []string{"siem-config.json", "monitoring-report.pdf"},
			References:  []string{"NIST-800-53-AU-6"},
		},
		{
			ID:          "ctrl-004",
			Name:        "Incident Response Plan",
			Description: "Documented incident response procedures",
			Category:    CategoryCorrective,
			Type:        TypeAdministrative,
			SubCategory: "Incident Response",
			RiskReduction: 0.25,
			Implementation: "IR plan documented and tested",
			Verification: "Review IR plan and test results",
			Owner:       "Security Team",
			Status:      StatusImplemented,
			LastVerified:  time.Now().AddDate(0, -4, 0),
			Evidence:    []string{"ir-plan.pdf", "test-results.pdf"},
			References:  []string{"NIST-800-53-IR-1"},
		},
	}
}

// GenerateReport generates control validation report.
func GenerateReport(validator *ControlValidator) string {
	var report string

	report += "=== Security Control Validation Report ===\n\n"

	results := validator.GetValidationResults()
	if len(results) == 0 {
		report += "No controls validated yet\n"
		return report
	}

	report += "Validation Results:\n"
	for i, result := range results {
		report += "\n[" + fmt.Sprintf("%d", i+1) + "] " + result.ControlName + "\n"
		report += "    ID: " + result.ControlID + "\n"
		report += "    Status: " + result.Status + "\n"
		report += "    Effectiveness: " + fmt.Sprintf("%.1f%%", result.Effectiveness*100) + "\n"
		report += "    Confidence: " + fmt.Sprintf("%.1f%%", result.Confidence*100) + "\n\n"

		if len(result.Issues) > 0 {
			report += "    Issues:\n"
			for j, issue := range result.Issues {
				report += "      [" + fmt.Sprintf("%d", j+1) + "] " + issue + "\n"
			}
			report += "\n"
		}

		if len(result.Recommendations) > 0 {
			report += "    Recommendations:\n"
			for j, rec := range result.Recommendations {
				report += "      [" + fmt.Sprintf("%d", j+1) + "] " + rec + "\n"
			}
			report += "\n"
		}
	}

	return report
}

// GetControl returns control.
func GetControl(validator *ControlValidator, id string) *SecurityControl {
	return validator.getControlByID(id)
}

// GetValidationResult returns validation result.
func GetValidationResult(result *ControlValidationResult) *ControlValidationResult {
	return result
}