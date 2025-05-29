package phonenumber

import "strings"

func NormalizePhoneNum(phoneNum string) string {
	for _, sym := range []string{" ", "-", "(", ")"} {
		phoneNum = strings.ReplaceAll(phoneNum, sym, "")
	}
	if strings.HasPrefix(phoneNum, "8") {
		phoneNum = strings.Replace(phoneNum, "8", "+7", 1)
	}
	phoneNum = phoneNum[:2] + "-" + phoneNum[2:5] + "-" + phoneNum[5:8] + "-" + phoneNum[8:]

	return phoneNum
}
