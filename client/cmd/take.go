package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"client/model"
	"client/prompt"
	"github.com/spf13/cobra"
)

var takeCmd = &cobra.Command{
	Use:   "take",
	Short: "Take a survey",
	Long: `Take a survey with max 3 definitive questions.
	Answer the questions with either yes or no`,
	RunE: func(cmd *cobra.Command, args []string) error {

		id, err := cmd.Flags().GetString("sid")
		if err != nil {
			return err
		}

		url := BaseURL+"/surveys"
		if id != "" {
			url += "/" + id
		}

		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var result *model.Result
		if id == "" {

			var surveys []*model.Survey
			err = json.NewDecoder(resp.Body).Decode(&surveys)
			if err != nil {
				return err
			}
			result, err = prompt.TakeFrom(surveys)

		} else {

			survey := &model.Survey{}
			err = json.NewDecoder(resp.Body).Decode(&survey)
			if err != nil {
				return err
			}
			result, err = prompt.Take(survey)
		}

		if err != nil {
			return err
		}

		payloadBuf := new(bytes.Buffer)
		err = json.NewEncoder(payloadBuf).Encode(result)
		if err != nil {
			return err
		}

		resp, err = http.Post(BaseURL+"/surveys/"+result.SurveyId+"/results", "application/json", payloadBuf)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		l := resp.Header.Get("Location")
		x := strings.Split(l, "/")
		fmt.Printf("Survey result:\n Location: %s\n Survey: %s\n Result: %s\n",
		l, x[3], x[5])
		return nil
	},
}

func init() {

	takeCmd.Flags().StringP("sid", "s", "", "Id of the Survey")

	rootCmd.AddCommand(takeCmd)
}
