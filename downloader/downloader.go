package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func Download(url string) error {
	filename := path.Base(url)
	fmt.Println("Writing : ", url, " to ", filename)
	out, err := os.Create("./" + filename)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
