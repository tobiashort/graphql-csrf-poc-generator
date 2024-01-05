package main

import (
	"bufio"
	"encoding/json"
	"html/template"
	"io"
	"os"
)

func must(err error) {
  if err != nil {
    panic(err)
  }
}

func must2[T any](val T, err error) T {
  must(err)
  return val
}

type G struct {
  Query string `json:"query"`
  Variables map[string]interface{} `json:"variables"`
  OperationName string `json:"operationName"`
}

type D struct {
  Query string
  Variables string
  OperationName string
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  bs := must2(io.ReadAll(reader))
  g := G{}
  must(json.Unmarshal(bs, &g))
  d := D{}
  d.Query = g.Query
  d.OperationName = g.OperationName
  d.Variables = string(must2(json.Marshal(g.Variables)))
  tmpl := must2(template.New("form").Parse(`
<form method="POST" action="TODO">
  <input type="hidden" name="query" value="{{ .Query }}">
  <input type="hidden" name="variables" value="{{ .Variables }}">
  <input type="hidden" name="operationName" value="{{ .OperationName }}">
</form>
<script>
  document.querySelector("form").submit();
</script>
`))
  must(tmpl.Execute(os.Stdout, d))
}
