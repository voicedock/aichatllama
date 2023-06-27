package main

import (
	"github.com/alexflint/go-arg"
	"github.com/voicedock/aichatllama/internal/aimodel"
	grpcapi "github.com/voicedock/aichatllama/internal/api/grpc"
	aichatv1 "github.com/voicedock/aichatllama/internal/api/grpc/gen/voicedock/extensions/aichat/v1"
	"github.com/voicedock/aichatllama/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"runtime"
)

var cfg AppConfig
var logger *zap.Logger

func init() {
	arg.MustParse(&cfg)
	logger = initLogger(cfg.LogLevel, cfg.LogJson)
}

func main() {
	defer logger.Sync()

	if cfg.LlamaThreads == 0 {
		cfg.LlamaThreads = runtime.NumCPU()
	}

	logger.Info(
		"Starting AI Chat llama",
		zap.Int("llama_threads", cfg.LlamaThreads),
		zap.Int("llama_context_size", cfg.LlamaContextSize),
		zap.Int("llama_gpu_layers", cfg.LlamaGpuLayers),
		zap.Int("llama_tokens", cfg.LlamaTokens),
		zap.Bool("llama_debug", cfg.LlamaDebug),
		zap.String("data_dir", cfg.DataDir),
		zap.String("config", cfg.Config),
	)

	lis, err := net.Listen("tcp", cfg.GrpcAddr)
	if err != nil {
		logger.Fatal("Failed to listen gRPC server", zap.Error(err))
	}

	dl := config.NewDownloader()
	cr := config.NewReader(cfg.Config)
	cs := config.NewService(cr, dl, logger, cfg.DataDir)
	cs.LoadConfig()
	mm := aimodel.NewModelManager(cfg.LlamaGpuLayers, cfg.LlamaContextSize)
	defer mm.Unload()

	srv := grpcapi.NewServerAiChat(cs, mm, logger, cfg.LlamaThreads, cfg.LlamaTokens, cfg.LlamaDebug)

	s := grpc.NewServer()
	aichatv1.RegisterAichatAPIServer(s, srv)
	reflection.Register(s)

	logger.Info("gRPC server listen", zap.String("addr", cfg.GrpcAddr))
	err = s.Serve(lis)
	if err != nil {
		logger.Fatal("gRPC server error", zap.Error(err))
	}
}
