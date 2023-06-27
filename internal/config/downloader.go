package config

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Downloader struct {
}

func NewDownloader() *Downloader {
	return &Downloader{}
}

func (d *Downloader) Download(url, outPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed download: %w", err)
	}
	defer resp.Body.Close()

	return d.saveFile(resp.Body, outPath)
}

func (d *Downloader) saveFile(r io.Reader, outPath string) error {
	err := os.MkdirAll(filepath.Dir(outPath), 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory for model: %w", err)
	}

	outFile, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("failed to create model file: %w", err)
	}

	defer outFile.Close()
	if _, err := io.Copy(outFile, r); err != nil {
		return fmt.Errorf("failed to write model file: %w", err)
	}

	return nil
}
