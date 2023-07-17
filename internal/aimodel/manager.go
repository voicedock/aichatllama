package aimodel

import (
	"fmt"
	"github.com/go-skynet/go-llama.cpp"
	"github.com/voicedock/aichatllama/internal/config"
)

type ModelManager struct {
	model       *llama.LLama
	modelPath   string
	gpuLayers   int
	contextSize int
}

func NewModelManager(gpuLayers, contextSize int) *ModelManager {
	return &ModelManager{
		gpuLayers:   gpuLayers,
		contextSize: contextSize,
	}
}

func (m *ModelManager) LazyLoadModel(cfg *config.ModelWrap) (*llama.LLama, error) {
	// unload old model from memory before loading new one
	if m.modelPath != "" && m.modelPath != cfg.ModelPath {
		m.Unload()
	}

	if m.model != nil {
		return m.model, nil
	}

	var err error
	m.model, err = llama.New(
		cfg.ModelPath,
		llama.EnableF16Memory,
		llama.SetContext(m.contextSize),
		/*llama.EnableEmbeddings,*/
		llama.SetGPULayers(m.gpuLayers),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate model: %w", err)
	}

	m.modelPath = cfg.ModelPath

	return m.model, nil
}

func (m *ModelManager) Unload() {
	if m.model != nil {
		m.modelPath = ""
		m.model.Free()
		m.model = nil
	}
}
