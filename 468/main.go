package _468

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if strings.Contains(IP, ".") {
		splits := strings.Split(IP, ".")
		if len(splits) != 4 {
			return "Neither"
		}
		for i := 0; i < 4; i++ {
			number, err := strconv.Atoi(splits[i])
			if err != nil {
				return "Neither"
			}
			//01
			if number > 0 && splits[i][0] == '0' {
				return "Neither"
			}
			// -0 00
			if number == 0 {
				if len(splits[i]) > 1 {
					return "Neither"
				}
			}
			//-1
			if number > 255 || number < 0 {
				return "Neither"
			}
		}
		return "IPv4"
	} else if strings.Contains(IP, ":") {
		splits := strings.Split(IP, ":")
		if len(splits) != 8 {
			return "Neither"
		}
		for i := 0; i < len(splits); i++ {
			if len(splits[i]) > 4 {
				return "Neither"
			} else {
				number, err := strconv.ParseInt(splits[i], 16, 32)
				if err != nil {
					return "Neither"
				}
				//-1
				if number < 0 {
					return "Neither"
				}
				//-0
				if number == 0 {
					if splits[i][0] != '0' {
						return "Neither"
					}
				}
			}
		}
		return "IPv6"
	}
	return "Neither"
}
