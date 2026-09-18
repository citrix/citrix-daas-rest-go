// Copyright © 2026. Citrix Systems, Inc.
package apis

import (
	"fmt"
	"os/exec"
	"reflect"
	"strings"
)

// escapePowerShellSingleQuote makes a value safe to embed inside a single-quoted PowerShell literal.
// A literal single quote is represented by doubling it, so a quote in an input value can no longer close the literal and inject commands.
func escapePowerShellSingleQuote(value string) string {
	return strings.ReplaceAll(value, "'", "''")
}

func BuildAuth(remoteCompName string, username string, password string, disableSSL bool) string {
	if remoteCompName == "" {
		return ""
	}

	escapedRemoteCompName := escapePowerShellSingleQuote(remoteCompName)
	escapedUsername := escapePowerShellSingleQuote(username)
	escapedPassword := escapePowerShellSingleQuote(password)

	if strings.Contains(remoteCompName, "https") {
		if disableSSL {
			return fmt.Sprintf("-ConnectionUri '%s' -Credential ( New-Object -TypeName System.Management.Automation.PSCredential  -ArgumentList '%s',(ConvertTo-SecureString -Force -AsPlainText '%s') ) -SessionOption (New-PSSessionOption -SkipCACheck -SkipCNCheck -SkipRevocationCheck) -Authentication Negotiate", escapedRemoteCompName, escapedUsername, escapedPassword)
		} else {
			return fmt.Sprintf("-ConnectionUri '%s' -Credential ( New-Object -TypeName System.Management.Automation.PSCredential  -ArgumentList '%s',(ConvertTo-SecureString -Force -AsPlainText '%s') )  -Authentication Negotiate", escapedRemoteCompName, escapedUsername, escapedPassword)
		}
	} else {
		return fmt.Sprintf("-ComputerName  '%s' -Credential ( New-Object -TypeName System.Management.Automation.PSCredential  -ArgumentList '%s',(ConvertTo-SecureString -Force -AsPlainText '%s') )", escapedRemoteCompName, escapedUsername, escapedPassword)
	}
}

func ExecuteCommand(credential string, command string, args ...string) ([]byte, error) {
	return ExecuteCommandBase(credential, 2, command, args...)
}

func ExecuteCommandWithDepth(credential string, jsonDepth int, command string, args ...string) ([]byte, error) {
	return ExecuteCommandBase(credential, jsonDepth, command, args...)
}

func ExecuteCommandBase(credential string, jsonDepth int, command string, args ...string) ([]byte, error) {
	var cmdArgs []string
	if credential != "" {
		cmdArgs = append([]string{"/c", "Invoke-Command -Session (New-PSSession ", credential, " ) -ScriptBlock { try { $ErrorActionPreference='Stop'; return $command = ", command}, args...)
		cmdArgs = append(cmdArgs, "|", fmt.Sprintf("ConvertTo-Json -Depth %d", jsonDepth))
		cmdArgs = append(cmdArgs, "} catch {Write-Error $_.Exception.Message} }")
	} else {
		cmdArgs = append([]string{"/c", command}, args...)
		cmdArgs = append(cmdArgs, "|", fmt.Sprintf("ConvertTo-Json -Depth %d", jsonDepth), "|", "Write-Output")
	}
	cmd := exec.Command("powershell", cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Errorf("error executing command: %v", command)
		return nil, fmt.Errorf("error executing command: %s \n Error Message:\n %s", command, string(output))
	}
	fmt.Println(strings.TrimSuffix(string(output), "\n"))
	return output, nil
}

// StructToString renders a request model as PowerShell parameter arguments. Every string value is
// wrapped in a single-quoted literal with embedded quotes doubled, so an attribute value can never
// close the literal and inject commands into the remote StoreFront session.
func StructToString(s interface{}) string {
	v := reflect.ValueOf(s)
	t := v.Type()

	var result []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if value.Kind() == reflect.Ptr && !value.IsNil() {
			value = value.Elem()
		}

		if value.Kind() == reflect.Struct {
			if value.FieldByName("isSet").Bool() {
				switch field.Type.Name() {
				case "NullableBool":
					boolValue := reflect.Indirect(value.FieldByName("value"))
					if boolValue.IsValid() {
						result = append(result, fmt.Sprintf("-%s $%t", field.Name, boolValue.Bool()))
					}

				case "NullableInt", "NullableInt32", "NullableInt64":
					intValue := reflect.Indirect(value.FieldByName("value"))
					if intValue.IsValid() {
						result = append(result, fmt.Sprintf("-%s %d", field.Name, intValue.Int()))
					}
				case "NullableFloat32", "NullableFloat64":
					floatValue := reflect.Indirect(value.FieldByName("value"))
					if floatValue.IsValid() {
						result = append(result, fmt.Sprintf("-%s %f", field.Name, floatValue.Float()))
					}

				case "NullableString":
					stringValue := reflect.Indirect(value.FieldByName("value"))
					if stringValue.IsValid() && stringValue.String() != "" {
						result = append(result, fmt.Sprintf("-%s '%s'", field.Name, escapePowerShellSingleQuote(stringValue.String())))
					}
				default:
					result = append(result, fmt.Sprintf("-%s %v", field.Name, value.Interface()))
				}
			}
		} else if value.Kind() == reflect.Slice && value.Type().Elem().Kind() == reflect.String {
			if value.IsNil() {
				continue
			} else {
				var strArr []string
				for i := 0; i < value.Len(); i++ {
					strArr = append(strArr, fmt.Sprintf("'%v'", escapePowerShellSingleQuote(value.Index(i).String())))
				}
				result = append(result, fmt.Sprintf("-%s @(%v)", field.Name, strings.Join(strArr, ", ")))
			}
		} else if value.Kind() == reflect.String {
			result = append(result, fmt.Sprintf("-%s '%s'", field.Name, escapePowerShellSingleQuote(value.String())))
		} else {
			result = append(result, fmt.Sprintf("-%s %v", field.Name, value.Interface()))
		}
	}

	return strings.Join(result, " ")
}
