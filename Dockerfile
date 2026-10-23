FROM golang:1.22-bookworm AS build
RUN apt-get update && apt-get install -y --no-install-recommends libwebp-dev && rm -rf /var/lib/apt/lists/*
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 go build -o /out/media-app .

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends libwebp7 ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=build /out/media-app .
RUN mkdir -p database images
EXPOSE 8082
CMD ["./media-app"]
