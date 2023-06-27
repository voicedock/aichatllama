# AI chat llama
Llama.cpp based VoiceDock [AI chat](https://github.com/voicedock/voicedock-specs/tree/main/proto/voicedock/extensions/aichat/v1) implementation


> Provides gRPC API for AI chat based on [llama.cpp](https://github.com/ggerganov/llama.cpp) project.
> Provides download of new model via API.

# Usage
Build docker image:
```bash
docker build -t ghcr.io/voicedock/aichatllama:latest .
```
Run docker container:
```bash
docker run --rm \
  -v "$(pwd)/config:/data/config" \
  -v "$(pwd)/dataset:/data/dataset" \
  -p 9999:9999 \
  ghcr.io/voicedock/aichatllama:latest aichatllama
```
## API
See implementation in [proto file](https://github.com/voicedock/voicedock-specs/blob/main/proto/voicedock/extensions/aichat/v1/aichat_api.proto).

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
Create protobuilder docker image:
```bash
cd ci/protobuilder && \
docker build -t protobuilder .
```
Lint proto files:
```bash
docker run --rm -w "/work" -v "$(pwd):/work" bufbuild/buf:latest lint internal/api/grpc/proto
```
Generate grpc interface:
```bash
docker run --rm -w "/work" -v "$(pwd):/work" protobuilder generate internal/api/grpc/proto --template internal/api/grpc/proto/buf.gen.yaml
```

## Thanks
* [Georgi Gerganov](https://github.com/ggerganov) - AI chat llama uses [llama.cpp](https://github.com/ggerganov/llama.cpp) project
* go-skynet authors - AI chat llama uses go binding [go-llama.cpp](https://github.com/go-skynet/go-llama.cpp)