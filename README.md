# securitycontrol - Security Control Validation Engine

[![Go](https://img.shields.io/badge/Go-1.21-blue)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

**Validate and test security controls to ensure they are working effectively.**

Go beyond configuration checks to validate that security controls actually work.

## 🚀 Features

- **Control Validation**: Test security controls to verify effectiveness
- **Control Management**: Manage security controls and frameworks
- **Validation Testing**: Execute validation tests on controls
- **Effectiveness Scoring**: Score control effectiveness
- **Issue Detection**: Identify control gaps and issues
- **Report Generation**: Generate validation reports

## 📦 Installation

### Build from Source

```bash
git clone https://github.com/hallucinaut/securitycontrol.git
cd securitycontrol
go build -o securitycontrol ./cmd/securitycontrol
sudo mv securitycontrol /usr/local/bin/
```

### Install via Go

```bash
go install github.com/hallucinaut/securitycontrol/cmd/securitycontrol@latest
```

## 🎯 Usage

### Validate Controls

```bash
# Validate all security controls
securitycontrol validate
```

### Test Specific Control

```bash
# Test a specific control
securitycontrol test ctrl-001
```

### List Controls

```bash
# List available controls
securitycontrol controls
```

### Generate Report

```bash
# Generate validation report
securitycontrol report
```

### Check Status

```bash
# Check control status
securitycontrol status
```

### Programmatic Usage

```go
package main

import (
    "fmt"
    "github.com/hallucinaut/securitycontrol/pkg/control"
    "github.com/hallucinaut/securitycontrol/pkg/validate"
)

func main() {
    // Create control validator
    validator := control.NewControlValidator()
    
    // Add controls
    commonControls := control.CreateCommonControls()
    for _, ctrl := range commonControls {
        validator.AddControl(ctrl)
    }
    
    // Validate controls
    for _, ctrl := range commonControls {
        result := validator.ValidateControl(ctrl.ID)
        if result != nil {
            fmt.Printf("Control: %s\n", result.ControlName)
            fmt.Printf("Status: %s\n", result.Status)
            fmt.Printf("Effectiveness: %.1f%%\n", result.Effectiveness*100)
        }
    }
    
    // Create test validator
    testValidator := validate.NewControlValidator()
    
    // Add tests
    commonTests := validate.CreateCommonControlTests()
    for _, test := range commonTests {
        testValidator.AddControlTest(test)
    }
    
    // Run validation
    results := testValidator.Validate()
    fmt.Printf("Validation Results: %d\n", len(results))
}
```

## 🔍 Control Categories

### Preventive Controls
Controls that prevent security incidents.

- Access Control Policy
- Multi-Factor Authentication
- Network Segmentation
- Encryption

### Detective Controls
Controls that detect security incidents.

- Security Monitoring
- Intrusion Detection
- Log Analysis
- Anomaly Detection

### Corrective Controls
Controls that correct security incidents.

- Incident Response
- Backup and Recovery
- System Restoration
- Data Recovery

### Deterrent Controls
Controls that discourage security violations.

- Security Awareness
- Warning Messages
- Security Policies
- Legal Agreements

### Recovery Controls
Controls that restore systems after incidents.

- Disaster Recovery
- Business Continuity
- System Redundancy
- Failover Systems

## 🧪 Validation Methods

| Method | Description | Example |
|--------|-------------|---------|
| Documentation | Review policy documents | Check access control policy |
| Interview | Interview control owners | Discuss incident response procedures |
| Observation | Observe control operation | Watch monitoring systems |
| Testing | Execute control tests | Test authentication |
| Automation | Automated validation scripts | Run security scanners |

## 📊 Control Status

| Status | Description | Action |
|--------|-------------|--------|
| Implemented | Control is active | Verify effectiveness |
| Partially Implemented | Control is partially active | Complete implementation |
| Not Implemented | Control is not active | Plan implementation |
| Deprecated | Control is outdated | Replace with modern control |

## 🏥 Effectiveness Scoring

| Score | Status | Action |
|-------|--------|--------|
| ≥90% | EFFECTIVE | Maintain current state |
| 70-89% | PARTIALLY_EFFECTIVE | Improve implementation |
| <70% | INEFFECTIVE | Significant improvement needed |

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test -v ./pkg/control -run TestValidateControl
```

## 📋 Example Output

```
$ securitycontrol validate

Security Control Validation
==========================

Controls to Validate:
  [1] Access Control Policy (preventive)
  [2] Multi-Factor Authentication (preventive)
  [3] Security Monitoring (detective)
  [4] Incident Response Plan (corrective)

Running Validation...

[EFFECTIVE] Access Control Policy
    Effectiveness: 90.0%
    Confidence: 85.0%

[PARTIALLY_EFFECTIVE] Multi-Factor Authentication
    Effectiveness: 75.0%
    Confidence: 80.0%

=== Security Control Validation Report ===

Validation Results:

[1] Access Control Policy
    ID: ctrl-001
    Status: EFFECTIVE
    Effectiveness: 90.0%
    Confidence: 85.0%
```

## 🏗️ Architecture

```
securitycontrol/
├── cmd/
│   └── securitycontrol/
│       └── main.go          # CLI entry point
├── pkg/
│   ├── control/
│   │   ├── control.go      # Control definitions
│   │   └── control_test.go # Unit tests
│   └── validate/
│       ├── validate.go     # Control validation
│       └── validate_test.go # Unit tests
└── README.md
```

## 🔒 Security Use Cases

- **Control Testing**: Validate security controls are working
- **Compliance Audits**: Test controls for compliance requirements
- **Security Assessments**: Evaluate control effectiveness
- **Risk Management**: Assess residual risk after controls
- **Continuous Monitoring**: Ongoing control validation
- **Third-Party Validation**: Independent control verification

## 🛡️ Best Practices

1. **Validate controls regularly** - Don't just configure them
2. **Test in isolation** - Validate individual controls
3. **Document test results** - Keep validation records
4. **Use multiple methods** - Combine testing approaches
5. **Measure effectiveness** - Quantify control performance
6. **Continuous validation** - Ongoing testing and monitoring

## 📄 License

MIT License

## 🙏 Acknowledgments

- Security control frameworks (NIST, CIS, ISO 27001)
- Security practitioners
- Control validation researchers

## 🔗 Resources

- [NIST SP 800-53](https://csrc.nist.gov/publications/detail/sp/800-53/rev-5/final)
- [CIS Controls](https://www.cisecurity.org/controls)
- [ISO 27001 Controls](https://www.iso.org/standard/54534.html)
- [OWASP Security Controls](https://owasp.org/www-project-security-controls/)

---

**Built with ❤️ by [hallucinaut](https://github.com/hallucinaut)**