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