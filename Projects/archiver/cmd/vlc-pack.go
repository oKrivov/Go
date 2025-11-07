package cmd

import (
	"archiver/lib/vlc"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcPackCdm = &cobra.Command{
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
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handlerErr(err)
	}

	packed := vlc.Encode(string(data))

	fmt.Println(string(data)) // TODO: remove

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handlerErr(err)
	}
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcPackCdm)
}
