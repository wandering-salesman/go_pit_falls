package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

func GenerateGoCode(funcName string, numReturns int) string {
	var returnVars, returnVals []string
	for i := 0; i < numReturns; i++ {
		returnVars = append(returnVars, "float64")
		returnVals = append(returnVals, fmt.Sprintf("input + %d", i))
	}

	tmpl := `
func {{.FuncName}}(input float64) ({{.ReturnVars}}) {
	return {{.ReturnVals}}
}

func Benchmark{{.BenchName}}(b *testing.B) {
	for b.Loop(){
		{{.FuncName}}(0)
	}
}
`
	data := map[string]string{
		"FuncName":   funcName,
		"ReturnVars": strings.Join(returnVars, ", "),
		"ReturnVals": strings.Join(returnVals, ", "),
		"BenchName":  capitalizeFirstChar(funcName),
	}

	var buf bytes.Buffer
	template.Must(template.New("code").Parse(tmpl)).Execute(&buf, data)
	return buf.String()
}

func capitalizeFirstChar(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func main() {
	const argCount = 64
	filename := "maxreturn_test.go"
	var codeBuf bytes.Buffer

	codeBuf.WriteString("package main\n\nimport \"testing\"\n\n")

	for i := 1; i <= argCount; i++ {
		funcName := fmt.Sprintf("return%d", i)
		codeBuf.WriteString(GenerateGoCode(funcName, i))
	}

	if err := os.WriteFile(filename, codeBuf.Bytes(), 0644); err != nil {
		fmt.Println("Error generating Go file:", err)
		return
	}

	cmd := exec.Command("go", "test", "-benchmem", "-run=^$", "-bench", "^BenchmarkReturn.")
	output, err := cmd.CombinedOutput()

	var allResults bytes.Buffer
	allResults.Write(output)
	if err != nil {
		allResults.WriteString(fmt.Sprintf("Error running benchmark: %v\n", err))
	}
	if err := os.WriteFile("res.txt", allResults.Bytes(), 0644); err != nil {
		fmt.Println("Error writing results file:", err)
		return
	}
	fmt.Println("All tests completed. Results saved to res.txt.")
}
