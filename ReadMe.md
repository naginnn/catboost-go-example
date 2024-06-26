## Usage/Examples

This example uses a compiled library
[CatBoost C library](https://catboost.ai/en/docs/concepts/c-plus-plus-api_dynamic-c-pluplus-wrapper)

Clone this repository
```
git clone https://github.com/naginnn/catboost-go-example.git
```

Specify the model path (.cbm format only) in main.go
```
model, err := catboost.Load("testmodel")
```
As an example to run in Docker (uses the catboost aarch64-linux folder):
```
docker build -t catimage . && docker run -it catimage
```
For local use the absolute path (uses the catlibosx aarch64 OSX folder):
```
export DYLD_LIBRARY_PATH={PathToProj}/catlibosx/catboost/libs/model_interface
go run main.go
```

If you want to compile for your architecture, you can change the --target-platform option a short example:
```
git clone --depth 1 --branch v1.1.1 https://github.com/catboost/catboost.git
cd catboost/catboost/libs/model_interface && ../../../ya make -r . --target-platform CLANG14-DARWIN-ARM64
```
P.s Versions older than v1.1.1 are no longer supported by the yamake utility
see [CatBoost C library](https://catboost.ai/en/docs/concepts/c-plus-plus-api_dynamic-c-pluplus-wrapper)
