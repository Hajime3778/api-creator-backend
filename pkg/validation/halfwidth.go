package validation

import "regexp"

// IsHalfWidthOnly 半角英数字 汎用記号のみ家判定します
func IsHalfWidthOnly(str string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9!-/:-?@¥[-{-~]+$`)
	return re.MatchString(str)
}
