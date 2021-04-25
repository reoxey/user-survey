package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"client/prompt"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new survey",
	Long: `Create a new survey with max 3 definitive questions.
	A survey consist of a title and 3 questions`,
	RunE: func(cmd *cobra.Command, args []string) error {

		survey, err := prompt.Create()

		if err != nil {
			return err
		}

		payloadBuf := new(bytes.Buffer)
		err = json.NewEncoder(payloadBuf).Encode(survey)
		if err != nil {
			return err
		}

		resp, err := http.Post(BaseURL+"/surveys", "application/json", payloadBuf)
		if err != nil {
			return err
		}

		l := resp.Header.Get("Location")
		x := strings.Split(l, "/")
		fmt.Printf("Survey result:\n Location: %s\n Survey: %s\n", l, x[3])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
