# AI chat llama
Llama.cpp based VoiceDock [AI chat](https://github.com/voicedock/voicedock-specs/tree/main/proto/voicedock/core/aichat/v1) implementation


> Provides gRPC API for AI chat based on [llama.cpp](https://github.com/ggerganov/llama.cpp) project.
> Provides download of new model via API.

# Usage
Run docker container on CPU:
```bash
docker run --rm \
  -v "$(pwd)/config:/data/config" \
  -v "$(pwd)/dataset:/data/dataset" \
  -p 9999:9999 \
  ghcr.io/voicedock/aichatllama:latest aichatllama
```

Run docker container on GPU (OpenCL, Nvidia CUDA):
```bash
docker run --rm \
  -v "$(pwd)/config:/data/config" \
  -v "$(pwd)/dataset:/data/dataset" \
  -e LLAMA_GPU_LAYERS=2 \
  --runtime=nvidia --gpus all \
  -p 9999:9999 \
  ghcr.io/voicedock/aichatllama:gpu aichatllama
```
Tested on NVIDIA GeForce RTX 3090.

Show more options:
```bash
docker run --rm ghcr.io/voicedock/aichatllama aichatllama -h
```
```
Usage: aichatllama [--grpcaddr GRPCADDR] [--config CONFIG] [--datadir DATADIR] [--loglevel LOGLEVEL] [--logjson] [--llamagpulayers LLAMAGPULAYERS] [--llamacontextsize LLAMACONTEXTSIZE] [--llamathreads LLAMATHREADS] [--llamatokens LLAMATOKENS] [--llamadebug]

Options:
  --grpcaddr GRPCADDR    gRPC API host:port [default: 0.0.0.0:9999, env: GRPC_ADDR]
  --config CONFIG        configuration file for models [default: /data/config/aichatllama.json, env: CONFIG]
  --datadir DATADIR      dataset directory [default: /data/dataset, env: DATA_DIR]
  --loglevel LOGLEVEL    log level: debug, info, warn, error, dpanic, panic, fatal [default: info, env: LOG_LEVEL]
  --logjson              set to true to use JSON format [env: LOG_JSON]
  --llamagpulayers LLAMAGPULAYERS
                         gpu layers for llama [env: LLAMA_GPU_LAYERS]
  --llamacontextsize LLAMACONTEXTSIZE
                         context size for llama [default: 1024, env: LLAMA_CONTEXT_SIZE]
  --llamathreads LLAMATHREADS
                         threads for llama (default MAX) [env: LLAMA_THREADS]
  --llamatokens LLAMATOKENS
                         sets number of tokens to generate for llama [default: 128, env: LLAMA_TOKENS]
  --llamadebug           debug flag for llama [env: LLAMA_DEBUG]
  --help, -h             display this help and exit
```

## API
See implementation in [proto file](https://github.com/voicedock/voicedock-specs/blob/main/proto/voicedock/core/aichat/v1/aichat_api.proto).

## FAQ
### How to add a language pack?
1. Find model (ggml format)
2. Copy link to download `.bin` model file
3. Add model to [aichatllama.json](config%2Faichatllama.json) config:
   ```json
   {
     "name": "model_name",
     "download_url": "download_url",
     "license": "license text to accept"
   }
    ```

### How to use preloaded model?
1. Add model to [aichatllama.json](config%2Faichatllama.json) config (leave "download_url" blank to disable downloads).
2. Download model
3. Save model to directory `dataset/{model_name}/model.bin` (replace `{model_name}` to name from configuration file `aichatllama.json`)


## CONTRIBUTING
Lint proto files:
```bash
docker run --rm -w "/work" -v "$(pwd):/work" bufbuild/buf:latest lint internal/api/grpc/proto
```
Generate grpc interface:
```bash
docker run --rm -w "/work" -v "$(pwd):/work" ghcr.io/voicedock/protobuilder:1.0.0 generate internal/api/grpc/proto --template internal/api/grpc/proto/buf.gen.yaml
```
Manual build CPU docker image:
```bash
docker build -t ghcr.io/voicedock/aichatllama:latest .
```
Manual build GPU docker image:
```bash
docker build -t ghcr.io/voicedock/aichatllama:gpu -f ./gpu.Dockerfile .
```

## Thanks
* [Georgi Gerganov](https://github.com/ggerganov) - AI chat llama uses [llama.cpp](https://github.com/ggerganov/llama.cpp) project
* go-skynet authors - AI chat llama uses go binding [go-llama.cpp](https://github.com/go-skynet/go-llama.cpp)