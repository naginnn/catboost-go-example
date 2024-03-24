FROM --platform=aarch64 golang:1.22
#FROM golang:1.22
RUN mkdir /app
ADD . /app/
WORKDIR /app
ENV LD_LIBRARY_PATH=/app/catboost/catboost/libs/model_interface
RUN go build -o main .
CMD ["/app/main"]