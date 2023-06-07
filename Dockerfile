FROM base:latest as build

ENV GOPROXY=https://nexus.services.mts.ru/repository/go-proxy/
ENV GOPRIVATE=gitlab.services.mts.ru
ENV GOSUMDB="sum.golang.org https://nexus.services.mts.ru/repository/go-sum"

WORKDIR /app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o baavgai ./cmd/baavgai

FROM alpine:latest

COPY --from=build /app/baavgai /baavgai
COPY internal/config/config.yaml /config.yaml
EXPOSE 8000
ENTRYPOINT ["/baavgai"]


