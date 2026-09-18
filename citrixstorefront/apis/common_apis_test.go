// Copyright © 2026. Citrix Systems, Inc.
package apis

import (
	"strings"
	"testing"

	"github.com/citrix/citrix-daas-rest-go/citrixstorefront/models"
)

// injection is a value that would break out of a single-quoted PowerShell literal and run an arbitrary command.
const injection = "x'; Remove-Item C:\\ -Recurse -Force; '"

func TestEscapePowerShellSingleQuote(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "no quotes", input: "store", expected: "store"},
		{name: "single quote is doubled", input: "o'brien", expected: "o''brien"},
		{name: "injection attempt is neutralized", input: injection, expected: "x''; Remove-Item C:\\ -Recurse -Force; ''"},
		{name: "empty string", input: "", expected: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := escapePowerShellSingleQuote(tc.input); got != tc.expected {
				t.Errorf("escapePowerShellSingleQuote(%q) = %q, want %q", tc.input, got, tc.expected)
			}
		})
	}
}

func TestBuildAuthEscapesCredentials(t *testing.T) {
	// A quote-bearing password must not close the single-quoted literal it is embedded in.
	auth := BuildAuth("host", "admin", "p'ass", false)
	if strings.Contains(auth, "'p'ass'") {
		t.Errorf("BuildAuth left an unescaped single quote in the credential literal: %s", auth)
	}
	if !strings.Contains(auth, "'p''ass'") {
		t.Errorf("BuildAuth did not double the embedded single quote: %s", auth)
	}
}

func TestStructToStringEscapesNullableString(t *testing.T) {
	model := models.CreateSTFDeploymentRequestModel{}
	model.SetHostBaseUrl(injection)
	model.SetSiteId(1)

	got := StructToString(model)

	if strings.Contains(got, "'"+injection+"'") {
		t.Errorf("StructToString embedded the raw attribute value without escaping: %q", got)
	}
	if !strings.Contains(got, "-HostBaseUrl 'x''; Remove-Item C:\\ -Recurse -Force; '''") {
		t.Errorf("StructToString did not escape the string attribute: %q", got)
	}
	if !strings.Contains(got, "-SiteId 1") {
		t.Errorf("StructToString did not render the numeric attribute directly: %q", got)
	}
}

func TestStructToStringEscapesPlainString(t *testing.T) {
	// Internal is a plain (non-Nullable) string field and must still be escaped before interpolation.
	model := models.SetSTFRoamingInternalBeaconRequestModel{}
	model.SetInternal(injection)

	got := StructToString(model)

	if !strings.Contains(got, "-Internal 'x''; Remove-Item C:\\ -Recurse -Force; '''") {
		t.Errorf("StructToString did not escape the plain string attribute: %q", got)
	}
}
