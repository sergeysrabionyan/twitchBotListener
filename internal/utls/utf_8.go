package utls

// RemoveInvalidUTF8Bytes Костыль из-за того, что твич иногда присылает не utf-8, а какой-то мусор в конце строки
func RemoveInvalidUTF8Bytes(input string) string {
	var result []byte
	for _, v := range input {
		if v < 32 || v > 127 {
			continue
		}
		result = append(result, byte(v))
	}
	return string(result)
}
