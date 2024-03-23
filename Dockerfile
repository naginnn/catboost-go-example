#FROM --platform=x86_64 golang:1.22
FROM golang:1.22
RUN mkdir /app
ADD . /app/
WORKDIR /app
ENV LD_LIBRARY_PATH=/app/catboost/catboost/libs/model_interface
#ENV LD_LIBRARY_PATH=/app/catboost
#ENV LD_LIBRARY_PATH=/app/catboost
RUN go build -o main .
CMD ["/app/main"]