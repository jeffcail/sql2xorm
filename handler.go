package main

import (
	"encoding/json"
	"go/format"
	"io"
	"net/http"
	"strings"
	"text/template"
)

type RequestBody struct {
	SQL string `json:"sql"`
}

type Response struct {
	Struct string `json:"struct"`
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有域
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	body := r.Body
	bytes, _ := io.ReadAll(body)

	rb := new(RequestBody)
	err := json.Unmarshal(bytes, rb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	table, err := parseSQLTable(rb.SQL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	structTmpl, err := template.New("struct").Funcs(template.FuncMap{
		"Mapper": func(s string) string { return strings.Title(s) },
		"Type":   sqlTypeToGoType,
		"Tag":    func(table Table, col Column) string { return col.Tag },
		"lower":  toUpperCamelCase,
	}).Parse(structTemplateContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tableNameTmpl, err := template.New("tableName").Funcs(template.FuncMap{
		"Mapper": func(s string) string { return strings.Title(s) },
	}).Parse(tableNameTemplateContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var structStr strings.Builder
	err = structTmpl.Execute(&structStr, struct {
		Tables []Table
	}{
		Tables: []Table{table},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var tableNameStr strings.Builder
	err = tableNameTmpl.Execute(&tableNameStr, struct {
		Tables []Table
	}{
		Tables: []Table{table},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Combine both parts
	s := structStr.String()
	contains := strings.Contains(s, "ENGINE")
	if contains {
		index := strings.LastIndex(s, "ENGINE")
		s = s[:index]
	}

	combinedStr := s + tableNameStr.String()

	// Format the combined code
	formattedStruct, err := format.Source([]byte(combinedStr))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	res := new(Response)
	res.Struct = string(formattedStruct)
	marshal, _ := json.Marshal(res)
	_, _ = w.Write(marshal)
}
