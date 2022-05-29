package codesnippet

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"strings"
)

// ファイルの情報を格納する構造体
type CodeSnippet struct {
	key, prefix, description string
	body                     []string
}

// ファイルの情報を取得
func NewCodeSnippet(path string) (CodeSnippet, error) {
	snippet := CodeSnippet{}

	data, err := os.Open(path)
	defer data.Close()
	if err != nil {
		return CodeSnippet{}, err
	}

	lang, err := getProgrammingLang(path)
	if err != nil {
		return CodeSnippet{}, err
	}

	header_start_flag := false
	header_end_flag := false

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		text := scanner.Text()
		if !header_start_flag {
			if strings.HasPrefix(text, lang.commentStart) {
				header_start_flag = true
			}
		} else if !header_end_flag {
			// headerの終端であるか確認
			if strings.HasPrefix(text, lang.commentEnd) {
				if snippet.key == "" || snippet.prefix == "" || snippet.description == "" {
					// headerの終端まできてるのに、key, prefix, descriptionのいずれかが空の場合は空のCodeSnippetをリターンする
					return CodeSnippet{}, ErrInsufficientHeader
				}
				header_end_flag = true
			}
			//
			if strings.HasPrefix(text, "key:") {
				snippet.key = strings.TrimPrefix(text, "key: ")
			} else if strings.HasPrefix(text, "prefix:") {
				snippet.prefix = strings.TrimPrefix(text, "prefix: ")
			} else if strings.HasPrefix(text, "description:") {
				snippet.description = strings.TrimPrefix(text, "description: ")
			}
		} else {
			snippet.body = append(snippet.body, text)
		}
	}

	// bodyが空の場合は空CodeSnippetをリターン
	if len(snippet.body) == 0 {
		return CodeSnippet{}, ErrEmptyBody
	}

	return snippet, err
}

func NewCodeSnippets(filePathes []string) []CodeSnippet {
	snippets := []CodeSnippet{}

	for _, filePath := range filePathes {
		snippet, err := NewCodeSnippet(filePath)
		if errors.Is(err, fs.ErrNotExist) ||
			errors.Is(err, ErrEmptyBody) ||
			errors.Is(err, ErrInsufficientHeader) {
			continue
		}
		snippets = append(snippets, snippet)
	}

	return snippets
}

func GetVSCodeSnippetsJSON(snippets []CodeSnippet) ([]byte, error) {
	jsonMap := map[string]interface{}{}

	for _, snippet := range snippets {
		jsonMap[snippet.key] = map[string]interface{}{
			"prefix":      snippet.prefix,
			"body":        snippet.body,
			"description": snippet.description,
		}
	}

	return json.Marshal(jsonMap)
}
