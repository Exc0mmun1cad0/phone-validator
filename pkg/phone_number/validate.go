package phonenumber

import (
	"fmt"
	"regexp"
)

func IsValidPhoneNum(phoneNum string) (bool, error) {
	for _, opCode := range opCodes {
		for _, format := range formats {
			isValid, err := regexp.MatchString(
				fmt.Sprintf(format, opCode),
				phoneNum,
			)
			if err != nil {
				return false, fmt.Errorf("cannot validate phone number: %w", err)
			}
			if isValid {
				return true, nil
			}
		}
	}

	return false, nil
}
