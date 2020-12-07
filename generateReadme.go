package main

import (
	"html/template"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sort"
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
	Easy Difficulty = 1
	Medium
	Hard
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

func main() {
	type File struct{
		Path string
		Number int
	}
	var files []File
	leetcode := NewLeetCode()
	tasks, err := leetcode.Tasks()
	if err != nil {
		log.Fatal(err)
	}
	root := "algorithms/go/"
    err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, "test.go") && !strings.HasSuffix(path, "go.mod") && !info.IsDir() {
			path = strings.TrimPrefix(path, root)
			number, err := strconv.Atoi(strings.Split(path, "-")[0])
			if err != nil {
				return err
			}
			files = append(files, File{path, number})
		}
        return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(files, func(i, j int) bool{
		return files[i].Number < files[j].Number
	})
	const tpl = `
<h1>LeetCode Algorithms</h1>
<table>
<tr>
	<th>#</th>
	<th>Title</th>
	<th>Solution</th>
	<th>Difficulty</th>
</tr>
{{- range .Files}}
	{{- $task := index $.Tasks .Number}}
	<tr>
		<td>{{.Number}}</td>
		<td>
		<a href="https://leetcode.com/problems/{{$task.Stat.Slug}}">{{$task.Stat.Title}}</a>
		</td>
		<td>
			<a href="go/{{.Path}}">Go</a>
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
	t, err := template.New("table").Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		Files []File
		Tasks map[int]Task
	}{files, tasks}
	err = t.Execute(f, data)
	if err != nil {
		log.Fatal(err)
	}
}
