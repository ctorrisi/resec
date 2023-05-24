# Build layer
FROM golang:1.20 AS builder
WORKDIR /go/src/resec
COPY . .
ARG RESEC_VERSION
ENV RESEC_VERSION ${RESEC_VERSION:-local-dev}
RUN echo $RESEC_VERSION
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.Version=${RESEC_VERSION}'" -a -installsuffix cgo -o /go/bin/resec  .

# Run layer
FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /go/bin/resec /resec

USER nonroot:nonroot
CMD ["/resec"]
