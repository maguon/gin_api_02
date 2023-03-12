package utils

func GetGenderStr(gender int8) string {
	if gender == 1 {
		return "男"
	} else {
		return "女"
	}
}

func GetStatusStr(status int8) string {
	if status == 1 {
		return "可用"
	} else {
		return "停用"
	}
}
