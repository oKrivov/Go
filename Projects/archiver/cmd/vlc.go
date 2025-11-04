package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCdm = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-lenght code",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptypath = errors.New("path to file is not specified")

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handlerErr(ErrEmptypath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handlerErr(err)
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		handlerErr(err)
	}

	// packed := Encode(data)
	packed := ""
	fmt.Println(string(data)) // TODO: remove

	err = ioutil.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handlerErr(err)
	}
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcCdm)
}
