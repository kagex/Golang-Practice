package parsinglogfiles

import "regexp"

// Компилируем регулярные выражения один раз при загрузке пакета.
var (
	// 1. Начало строки ^, затем один из тегов в скобках. Символы [ и ] нужно экранировать.
	validLineRegex = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	// 2. Разделитель начинается с <, заканчивается на >, а внутри от 0 до бесконечности символов ~, *, =, -.
	// Дефис стоит в начале [], поэтому его не нужно экранировать.
	splitLogLineRegex = regexp.MustCompile(`<[-~=*]*>`)

	// 3. Кавычки, затем любые символы, затем "password" (флаг (?i) игнорирует регистр), любые символы и кавычка.
	quotedPasswordRegex = regexp.MustCompile(`"(?i).*password.*"`)

	// 4. Текст "end-of-line" и следующие за ним цифры (\d+).
	removeEndOfLineRegex = regexp.MustCompile(`end-of-line\d+`)
    
	// 5. Строка "User", затем один или более пробельных символов \s+,
	// затем захватывающая группа () с именем пользователя (любой набор непробельных символов \S+).
	tagWithUserNameRegex = regexp.MustCompile(`User\s+(\S+)`)
)

// IsValidLine проверяет, начинается ли строка с валидного тега.
func IsValidLine(text string) bool {
	return validLineRegex.MatchString(text)
}

// SplitLogLine разбивает строку по кастомному разделителю.
// Второй аргумент -1 означает, что нужно вернуть все возможные части.
func SplitLogLine(text string) []string {
	return splitLogLineRegex.Split(text, -1)
}

// CountQuotedPasswords считает строки, в которых слово "password" находится внутри кавычек.
func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if quotedPasswordRegex.MatchString(line) {
			count++
		}
	}
	return count
}

// RemoveEndOfLineText удаляет мусорный текст "end-of-line" вместе с номером.
func RemoveEndOfLineText(text string) string {
	return removeEndOfLineRegex.ReplaceAllString(text, "")
}

// TagWithUserName ищет упоминание пользователя и добавляет тег в начало строки.
func TagWithUserName(lines []string) []string {
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		// FindStringSubmatch возвращает слайс, где индекс 0 — полное совпадение,
		// а индекс 1 — содержимое первой захватывающей скобки (имя пользователя).
		match := tagWithUserNameRegex.FindStringSubmatch(line)
		if match != nil {
			username := match[1]
			result = append(result, "[USR] "+username+" "+line)
		} else {
			result = append(result, line)
		}
	}
	return result
}