package codesnippet

import (
	"encoding/json"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
)

const fixturesDir = "../../data/fixtures"

func TestNewCodeSnippet(t *testing.T) {
	testCases := []struct {
		name                string
		path                string
		expectedCodeSnippet CodeSnippet
		expectedErr         error
	}{
		{
			name: "success_cpp",
			path: fixturesDir + "/cpp/success.cc",
			expectedCodeSnippet: CodeSnippet{
				key: "success",
				value: codeSnippetValue{
					prefix:      "success",
					description: "success",
					body: []string{
						"#include <iostream>",
						"",
						"int main() {",
						"  std::cout << \"success\" << std::endl;",
						"  return;",
						"}",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "success_python",
			path: fixturesDir + "/python/success.py",
			expectedCodeSnippet: CodeSnippet{
				key: "success",
				value: codeSnippetValue{
					prefix:      "success",
					description: "success",
					body: []string{
						"print(\"success\")",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name:                "no such file",
			path:                "fail_file_path",
			expectedCodeSnippet: CodeSnippet{},
			expectedErr:         fs.ErrNotExist,
		},
		{
			name:                "insufficient header",
			path:                fixturesDir + "/cpp/error_insufficient_header.cc",
			expectedCodeSnippet: CodeSnippet{},
			expectedErr:         ErrInsufficientHeader,
		},
		{
			name:                "empty body",
			path:                fixturesDir + "/cpp/error_empty_body.cc",
			expectedCodeSnippet: CodeSnippet{},
			expectedErr:         ErrEmptyBody,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			codeSnippet, err := NewCodeSnippet(tc.path)
			assert.ErrorIs(t, err, tc.expectedErr)
			if assert.NotNil(t, codeSnippet) {
				assert.Equal(t, tc.expectedCodeSnippet.key, codeSnippet.key)
				assert.Equal(t, tc.expectedCodeSnippet.value.prefix, codeSnippet.value.prefix)
				assert.Equal(t, tc.expectedCodeSnippet.value.description, codeSnippet.value.description)
				assert.Equal(t, tc.expectedCodeSnippet.value.body, codeSnippet.value.body)
			}
		})
	}
}

func TestNewCodeSnippets(t *testing.T) {
	testCases := []struct {
		name        string
		filePathes  []string
		expectedLen int
	}{
		{
			name: "read cpp files",
			filePathes: []string{
				fixturesDir + "/cpp/success.cc",
				fixturesDir + "/cpp/error_insufficient_header.cc",
				fixturesDir + "/cpp/error_empty_body.cc",
				fixturesDir + "/cpp/subdir/success_subdir.cc",
				fixturesDir + "/cpp/subdir/subsubdir/success_subsubdir.cc",
				"wrong_file_path_1",
			},
			expectedLen: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			codeSnippets := NewCodeSnippets(tc.filePathes)
			assert.Equal(t, tc.expectedLen, len(codeSnippets))
		})
	}
}

func TestGetVSCodeSnippetsJSON(t *testing.T) {
	testCases := []struct {
		name              string
		inputCodeSnippets []CodeSnippet
		expectedJSONMap   map[string]interface{}
	}{
		{
			name: "read cpp files",
			inputCodeSnippets: []CodeSnippet{
				{
					key: "key1",
					value: codeSnippetValue{
						prefix:      "prefix1",
						description: "description1",
						body: []string{
							"body1_l1",
							"body1_l2",
						},
					},
				},
				{
					key: "key2",
					value: codeSnippetValue{
						prefix:      "prefix2",
						description: "description2",
						body: []string{
							"body2_l1",
							"body2_l2",
						},
					},
				},
			},
			expectedJSONMap: map[string]interface{}{
				"key1": map[string]interface{}{
					"prefix":      "prefix1",
					"description": "description1",
					"body": []string{
						"body1_l1",
						"body1_l2",
					},
				},
				"key2": map[string]interface{}{
					"prefix":      "prefix2",
					"description": "description2",
					"body": []string{
						"body2_l1",
						"body2_l2",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			acctualJSON, err := GetVSCodeSnippetsJSON(tc.inputCodeSnippets)
			assert.Nil(t, err)
			expectedJSON, err := json.Marshal(tc.expectedJSONMap)
			assert.Nil(t, err)
			assert.Equal(t, expectedJSON, acctualJSON)
		})
	}

}
