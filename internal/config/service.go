package config

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

type Service struct {
	reader     *Reader
	downloader *Downloader
	logger     *zap.Logger
	config     []*ModelWrap
	idxConfig  map[string]*ModelWrap
	dataDir    string
}

func NewService(
	confReader *Reader,
	downloader *Downloader,
	logger *zap.Logger,
	dataDir string,
) *Service {
	return &Service{
		reader:     confReader,
		downloader: downloader,
		logger:     logger,
		config:     []*ModelWrap{},
		idxConfig:  make(map[string]*ModelWrap),
		dataDir:    dataDir,
	}
}

func (s *Service) LoadConfig() error {
	items, err := s.reader.ReadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	models := make([]*ModelWrap, 0, len(items))
	idx := map[string]*ModelWrap{}
	for _, v := range items {
		if _, ok := idx[v.Name]; ok {
			s.logger.Warn("Model with same name skipped", zap.String("name", v.Name))
			continue
		}

		model := s.WrapModel(v)
		idx[v.Name] = model
		models = append(
			models,
			model,
		)
	}

	s.config = models
	s.idxConfig = idx

	return nil
}

func (s *Service) WrapModel(model *Model) *ModelWrap {
	ret := &ModelWrap{
		Model:     model,
		ModelPath: filepath.Join(s.dataDir, model.Name, "model.bin"),
	}

	_, err := os.Stat(ret.ModelPath)
	ret.Downloaded = !os.IsNotExist(err)

	return ret
}

func (s *Service) FindAll() []*ModelWrap {
	return s.config
}

func (s *Service) Download(name string) error {
	model, ok := s.idxConfig[name]
	if !ok {
		return errors.New("model configuration for download not found")
	}

	if !model.Downloadable() {
		return errors.New("model is not downloadable")
	}

	downloadUrl := filepath.Join(s.dataDir, model.Model.Name, "model.bin")
	err := s.downloader.Download(model.Model.DownloadUrl, downloadUrl)
	if err != nil {
		return fmt.Errorf("can't download model: %w", err)
	}

	return s.LoadConfig()
}

func (s *Service) FindDownloaded() *ModelWrap {
	for _, v := range s.config {
		if v.Downloaded {
			return v
		}
	}

	return nil
}
