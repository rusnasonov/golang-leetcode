package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Profile struct {
	UserName string `json:"user_name"`
	NumTotal int `json:"num_total"`
	Tasks []Task `json:"stat_status_pairs"` 
}

type Task struct {
	Difficulty struct {
		Level Difficulty `json:"level"`
	} `json:"difficulty"`
	Stat struct {
		Id int `json:"frontend_question_id"`
		Title string `json:"question__title"`
		Slug string `json:"question__title_slug"`
	} `json:"stat"`

}

type LeetCode struct {
	
}

func NewLeetCode() LeetCode {
	return LeetCode{}
}

type Difficulty int

const (
	_ Difficulty = iota
	Easy
	Medium
	Hard
)

type Language int

func (l Language) String() string {
	return [...]string{"go", "python"}[l]
}

const (
	Go Language = iota
	Python
)

func (d Difficulty) String() string {
    return [...]string{"Easy", "Medium", "Hard"}[d-1]
}

func (this LeetCode) Tasks() (map[int]Task, error) {
	resp, err := http.Get("https://leetcode.com/api/problems/algorithms/")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("bad status: " + resp.Status)
	}
	defer resp.Body.Close()
	var data Profile
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	res := make(map[int]Task)
	for _, task := range data.Tasks {
		res[task.Stat.Id] = task
	}
	return res, nil
}

type File struct{
	Path string
	Number int
	Lang Language
}

func SolvedTasks(lang Language) ([]File, error) {
	var files []File
	root := fmt.Sprintf("algorithms/%s/", lang)
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, "test.go") &&
		   (strings.HasSuffix(path, ".go") || strings.HasSuffix(path, ".py")) &&
		   !info.IsDir() {
			path = strings.TrimPrefix(path, root)
			number, err := strconv.Atoi(strings.Split(path, "-")[0])
			if err != nil {
				return err
			}
			files = append(files, File{path, number, lang})
		}
        return nil
	})
	if err != nil {
		return files, err
	}
	sort.Slice(files, func(i, j int) bool{
		return files[i].Number < files[j].Number
	})
	return files, nil
}

func main() {
	leetcode := NewLeetCode()
	tasks, err := leetcode.Tasks()
	solvedTasks := map[int][]File{}

	if err != nil {
		log.Fatal(err)
	}
	solvedGo, err := SolvedTasks(Go)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range solvedGo {
		solvedTasks[t.Number] = append(solvedTasks[t.Number], t)
	}
	solvedPython, err := SolvedTasks(Python)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range solvedPython {
		solvedTasks[t.Number] = append(solvedTasks[t.Number], t)
	}
	const tpl = `
<h1>LeetCode Algorithms</h1>
<table>
<tr>
	<th>#</th>
	<th>Title</th>
	<th>Solution</th>
	<th>Difficulty</th>
</tr>
{{- range $taskNumber, $files := .SolvedTasks}}
	{{- $task := index $.Tasks $taskNumber}}
	<tr>
		<td>{{$taskNumber}}</td>
		<td>
		<a href="https://leetcode.com/problems/{{$task.Stat.Slug}}">{{$task.Stat.Title}}</a>
		</td>
		<td>
		    {{- range $files}}
				<a href="{{.Lang}}/{{.Path}}">{{ .Lang }}</a>
			{{- end}}
		</td>
		<td>{{$task.Difficulty.Level}}</td>
	</tr>
{{- end}}
</table>
	`
	f, err := os.Create("algorithms/README.md")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	funcMap := template.FuncMap{
        "ToTitle": strings.ToTitle,
    }
	t, err := template.New("table").Funcs(funcMap).Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		SolvedTasks map[int][]File
		Tasks map[int]Task
	}{solvedTasks, tasks}
	err = t.Execute(f, data)
	if err != nil {
		log.Fatal(err)
	}
}
