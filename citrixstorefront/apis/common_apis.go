package apis

import (
	"fmt"
	"os/exec"
	"reflect"
	"strings"
)

func BuildAuth(remoteCompName string, username string, password string) string {
	if remoteCompName != "" {
		return fmt.Sprintf("-ComputerName  '%s' -Credential ( New-Object -TypeName System.Management.Automation.PSCredential  -ArgumentList '%s',(ConvertTo-SecureString -Force -AsPlainText '%s') )", remoteCompName, username, password)
	}
	return ""
}

func ExecuteCommand(credential string, command string, args ...string) ([]byte, error) {
	var cmdArgs []string
	if credential != "" {
		cmdArgs = append([]string{"/c", "Invoke-Command -Session (New-PSSession ", credential, " ) -ScriptBlock {", command}, args...)
		cmdArgs = append(cmdArgs, "|", "ConvertTo-Json", "}")
	} else {
		cmdArgs = append([]string{"/c", command}, args...)
		cmdArgs = append(cmdArgs, "|", "ConvertTo-Json", "|", "Write-Output")
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
						result = append(result, fmt.Sprintf("-%s %s", field.Name, stringValue.String()))
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
					strArr = append(strArr, fmt.Sprintf("'%v'", value.Index(i).String()))
				}
				result = append(result, fmt.Sprintf("-%s @(%v)", field.Name, strings.Join(strArr, ", ")))
			}
		} else {
			result = append(result, fmt.Sprintf("-%s %v", field.Name, value.Interface()))
		}
	}

	return strings.Join(result, " ")
}
