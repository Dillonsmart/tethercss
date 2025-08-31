package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var rootJsonPath = "root.json"
var baseIndentation = "  "

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

	errorClearing := clearCSSFile()
	if errorClearing != nil {
		return
	}

	Generate(styles)
	fmt.Println("CSS file generated successfully.")
}

func Generate(styles map[string]interface{}) {

	errorStarting := startRootCSS()
	if errorStarting != nil {
		return
	}

	for styleType, styleValues := range styles {
		if styleMap, ok := styleValues.(map[string]interface{}); ok {
			for name, values := range styleMap {
				generateRootVariables(styleType, name, values)
			}
		}
	}

	errorEnding := endRootCSS()
	if errorEnding != nil {
		return
	}

	for styleType, styleValues := range styles {
		if styleMap, ok := styleValues.(map[string]interface{}); ok {
			for name, values := range styleMap {
				generateClasses(styleType, name, values)
			}
		}
	}
}

func generateRootVariables(styleType string, name string, style interface{}) {
	var lines []string

	if arr, ok := style.([]interface{}); ok {
		for i, v := range arr {
			classPrefix := (i + 1) * 100
			lines = append(lines, baseIndentation+fmt.Sprintf("--%s-%d: %v;\n", name, classPrefix, v))
		}
	} else {
		lines = append(lines, baseIndentation+fmt.Sprintf("--%s-%s: %v;\n", styleType, name, style))
	}

	err := writeRulesToFile(lines)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func generateClasses(styleType string, name string, style interface{}) {
	switch styleType {
	case "colors":
		generateColorClasses("text", "color", name, style)
	default:
		// Handle unknown style types if necessary
	}
}

func generateColorClasses(classPrefix string, cssProperty string, name string, style interface{}) {
	var lines []string

	var i = 100

	if arr, ok := style.([]interface{}); ok {
		for _ = range arr {
			className := fmt.Sprintf(".%s-%s-%d {\n", classPrefix, name, i)
			classProperty := fmt.Sprintf("  %s: var(--%s-%d);\n}\n", cssProperty, name, i)
			lines = append(lines, className)
			lines = append(lines, classProperty)
			i += 100
		}
	}

	err := writeRulesToFile(lines)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func clearCSSFile() error {
	err := os.WriteFile("../tether.css", []byte(""), 0644)
	if err != nil {
		return err
	}
	return nil
}

func startRootCSS() error {
	lines := []string{":root {\n"}
	err := writeToFile(lines)
	if err != nil {
		return err
	}
	return nil
}

func endRootCSS() error {
	lines := []string{"}\n"}
	err := writeToFile(lines)
	if err != nil {
		return err
	}
	return nil
}

func writeToFile(lines []string) error {
	file, err := os.OpenFile("../tether.css", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file while writing:", err)
		}
	}(file)

	for _, line := range lines {
		if _, err := file.WriteString(line); err != nil {
			return err
		}
	}

	return nil
}

func writeRulesToFile(lines []string) error {
	file, err := os.OpenFile("../tether.css", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file while writing rules:", err)
		}
	}(file)

	for _, line := range lines {
		if _, err := file.WriteString(line); err != nil {
			return err
		}
	}

	return nil
}
