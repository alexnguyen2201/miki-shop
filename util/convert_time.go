package util

const (
	ALL_LIFE = 36500
)

func ConvertWarrantyDuration(days int) any {
	if days == ALL_LIFE {
		return "Trọn đời"
	}
	return days
}
