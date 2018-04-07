package kata

import "regexp"

func IsValidCoordinates(coordinates string) bool {
	format, _ := regexp.Compile(`^[-+]?([0-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`)
	return format.MatchString(coordinates)
}
