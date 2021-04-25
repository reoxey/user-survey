package prompt

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"

	"client/model"
)

func Create() (survey *model.Survey, err error) {

	validateTitle := func(input string) error {
		if len(input) < 3 {
			return fmt.Errorf("title is too short")
		}
		return nil
	}

	validateFAQ := func(input string) error {
		if len(input) < 10 {
			return fmt.Errorf("question is too short")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     "Title:",
		Templates: templates,
		Validate:  validateTitle,
	}

	title, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	var questions []string
	for i:=0; i<3; i++ {
		prompt = promptui.Prompt{
			Label:     fmt.Sprintf("Question %d: ", i),
			Templates: templates,
			Validate:  validateFAQ,
		}

		q, err := prompt.Run()
		if err != nil {
			return nil, err
		}

		questions = append(questions, q)
	}

	return &model.Survey{
		Title: title,
		Questions: questions,
	}, err
}

func Take(survey *model.Survey) (*model.Result, error) {

	if survey == nil {
		return nil, fmt.Errorf("no results")
	}

	fmt.Println("Please fill the Survey.")
	fmt.Printf("Title: %s", survey.Title)

	var qas []*model.QA
	for i, q := range survey.Questions {
		prompt := promptui.Select{
			Label:     fmt.Sprintf("Question %d: %s", i, q),
			Items: []string{"Yes", "No"},
		}

		_, f, err := prompt.Run()
		if err != nil {
			return nil, err
		}

		ans := false
		if f == "Yes" {
			ans = true
		}

		qas = append(qas, &model.QA{
			Question: q,
			Answer:   ans,
		})
	}

	return &model.Result{
		Title:    survey.Title,
		SurveyId: survey.Id,
		Qas:      qas,
	}, nil
}

func TakeFrom(surveys []*model.Survey) (*model.Result, error) {

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Title | cyan }} ({{ .Id | red }})",
		Inactive: "  {{ .Title | cyan }} ({{ .Id | red }})",
		Selected: "\U0001F336 {{ .Title | red | cyan }}",
		Details: `
--------- Survey ----------
{{ "Title:" | faint }}	{{ .Title }}
{{range $i, $f := .Questions}}
	{{ $i | faint }}	{{ $f }}
{{end}}`,
	}

	searcher := func(input string, index int) bool {
		survey := surveys[index]
		name := strings.Replace(strings.ToLower(survey.Title), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Survey Titles",
		Items:     surveys,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return Take(surveys[i])
}

func ShowFrom(surveys []*model.Result) (*model.Result, error) {

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Title | cyan }} ({{ .Id | red }})",
		Inactive: "  {{ .Title | cyan }} ({{ .Id | red }})",
		Selected: "\U0001F336 {{ .Title | red | cyan }}",
		Details: `
--------- Survey ----------
{{ "Title:" | faint }}	{{ .Title }}
{{range $i, $f := .Qas}}
	{{ $i | faint }}	{{ $f.Question }}
	{{ $f.Answer | bold }}
{{end}}`,
	}

	searcher := func(input string, index int) bool {
		survey := surveys[index]
		name := strings.Replace(strings.ToLower(survey.Title), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Survey Titles",
		Items:     surveys,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return surveys[i], nil
}
