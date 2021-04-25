package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"client/model"
	"client/prompt"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show results of a survey",
	Long: `Show results of a survey`,
	RunE: func(cmd *cobra.Command, args []string) error {

		sid, err := cmd.Flags().GetString("sid")
		if err != nil {
			return err
		}
		rid, err := cmd.Flags().GetString("rid")
		if err != nil {
			return err
		}

		url := BaseURL+"/surveys/"+sid+"/results"
		if rid != "" {
			url += "/" + rid
		}

		resp, err := http.Get(url)
		if err != nil {
			return err
		}

		var result *model.Result
		if rid == "" {
			var surveys []*model.Result
			err = json.NewDecoder(resp.Body).Decode(&surveys)
			if err != nil {
				return err
			}
			result, err = prompt.ShowFrom(surveys)
		} else {
			err = json.NewDecoder(resp.Body).Decode(&result)
			if err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}

		fmt.Printf("Survey result:\n %s \n\n", strModel{result})
		return nil
	},
}

// strModel embeds model.Survey struct to extend its functionality
type strModel struct {
	*model.Result
}

// String formats the model.Survey data embedded in strModel
func (x strModel) String() string {
	if x.Result == nil {
		return "No result"
	}
	s := fmt.Sprintf("#%s => Title: %s.", x.Id, x.Title)

	sb := strings.Builder{}
	for i, qa := range x.Qas {
		sb.WriteString(fmt.Sprintf("\n Q %d: %s \n A: %v", i+1, qa.Question, qa.Answer))
	}

	return s + sb.String()
}

func init() {

	showCmd.Flags().StringP("sid", "s", "", "Id of the Survey")
	showCmd.MarkFlagRequired("sid")

	showCmd.Flags().StringP("rid", "r", "", "Id of the Result")

	rootCmd.AddCommand(showCmd)
}
