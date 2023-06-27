package grpc

import "C"
import (
	"context"
	"github.com/go-skynet/go-llama.cpp"
	"github.com/voicedock/aichatllama/internal/aimodel"
	aichatv1 "github.com/voicedock/aichatllama/internal/api/grpc/gen/voicedock/extensions/aichat/v1"
	"github.com/voicedock/aichatllama/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewServerAiChat(configService *config.Service, manager *aimodel.ModelManager, logger *zap.Logger, llamaThreads int, llamaTokens int, llamaDebug bool) *ServerAiChat {
	return &ServerAiChat{
		configService: configService,
		manager:       manager,
		logger:        logger,
		llamaThreads:  llamaThreads,
		llamaTokens:   llamaTokens,
		llamaDebug:    llamaDebug,
	}
}

type ServerAiChat struct {
	configService *config.Service
	manager       *aimodel.ModelManager
	logger        *zap.Logger
	llamaThreads  int
	llamaTokens   int
	llamaDebug    bool
	aichatv1.UnimplementedAichatAPIServer
}

func (s *ServerAiChat) Generate(in *aichatv1.GenerateRequest, srv aichatv1.AichatAPI_GenerateServer) error {
	cfg := s.configService.FindDownloaded()
	if cfg == nil {
		return status.Error(codes.NotFound, "model not found")
	}

	model, err := s.manager.LazyLoadModel(cfg)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to load model: %w", err)
	}

	opts := []llama.PredictOption{
		llama.SetTokenCallback(func(token string) bool {
			s.logger.Debug("Send token", zap.String("token", token))
			err := srv.Send(&aichatv1.GenerateResponse{TokenText: token})
			if err != nil {
				s.logger.Warn("Failed to send generation result", zap.Error(err))
			}

			return true
		}),
		llama.SetTokens(s.llamaTokens),
		llama.SetThreads(s.llamaThreads),
		llama.SetTopK(90),
		llama.SetTopP(0.86),
		llama.SetStopWords("llama"),
	}

	if s.llamaDebug {
		opts = append(opts, llama.Debug)
	}

	_, err = model.Predict(in.Prompt, opts...)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to predict: %w", err)
	}

	embeds, err := model.Embeddings(in.Prompt)
	if err != nil {
		s.logger.Warn("Embeddings error", zap.Error(err))
	}

	s.logger.Debug("Embeddings", zap.Float32s("embeds", embeds))

	return nil
}

func (s *ServerAiChat) GetModels(ctx context.Context, in *aichatv1.GetModelsRequest) (*aichatv1.GetModelsResponse, error) {
	var models []*aichatv1.Model
	for _, v := range s.configService.FindAll() {
		models = append(models, &aichatv1.Model{
			Name:         v.Model.Name,
			Downloaded:   v.Downloaded,
			Downloadable: v.Downloadable(),
			License:      v.Model.License,
		})
	}

	return &aichatv1.GetModelsResponse{
		Models: models,
	}, nil
}

func (s *ServerAiChat) DownloadModel(ctx context.Context, in *aichatv1.DownloadModelRequest) (*aichatv1.DownloadModelResponse, error) {
	return &aichatv1.DownloadModelResponse{}, s.configService.Download(in.Name)
}
