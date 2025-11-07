package cmd

import (
	"archiver/lib/vlc"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcUnpackCdm = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file using variable-lenght code",
	Run:   unpack,
}

// TODO: take extension from file
const unpackedExtension = "txt"

func unpack(_ *cobra.Command, args []string) {
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

	packed := vlc.Decode(string(data))

	fmt.Println(string(data)) // TODO: remove

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handlerErr(err)
	}
}

// TODO: refactor this
func unpackedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + unpackedExtension
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCdm)
}
