package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var rootJsonPath = "root.json"

func main() {
	rootJson, readRootJsonErr := os.Open(rootJsonPath)

	if readRootJsonErr != nil {
		fmt.Println("Error reading root.json:", readRootJsonErr)
		return
	}

	defer func(rootJson *os.File) {
		err := rootJson.Close()
		if err != nil {

		}
	}(rootJson)

	jsonAsBytes, _ := io.ReadAll(rootJson)

	var styles map[string]interface{}
	errorParsingJson := json.Unmarshal(jsonAsBytes, &styles)
	if errorParsingJson != nil {
		return
	}

	Generate(styles)
}

func Generate(styles map[string]interface{}) {
	for styleType, styleValues := range styles {
		fmt.Println(styleType)

		if styleMap, ok := styleValues.(map[string]interface{}); ok {
			for name, values := range styleMap {
				HandleStyles(styleType, name, values)
			}
		}
	}
}

func HandleStyles(styleType string, name string, style interface{}) {
	if arr, ok := style.([]interface{}); ok {
		for i, v := range arr {
			classPrefix := (i + 1) * 100
			fmt.Printf("--%s-%d: %v\n", name, classPrefix, v)
		}
	} else {
		fmt.Printf("--%s-%s: %v\n", styleType, name, style)
	}
}
