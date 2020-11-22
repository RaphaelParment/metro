ARG GOLANG_VERSION=1.15

  #===========================================================
  # Base
  #===========================================================
FROM golang:${GOLANG_VERSION} as base

  #===========================================================
  # Builder
  #===========================================================
FROM base as builder
WORKDIR /app
COPY . ./
ARG BUILD_VERSION
ARG APP_VERSION
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X main.build=${BUILD_VERSION} -X main.version=${APP_VERSION}" -v -o metro


  #===========================================================
  # Release
  #===========================================================
FROM alpine as release
COPY --from=builder /app/metro /metro
CMD ["./metro"]