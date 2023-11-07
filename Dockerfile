FROM pefish/ubuntu-go:v1.20 as builder
WORKDIR /app
ENV GO111MODULE=on
COPY ./ ./
RUN make

FROM pefish/ubuntu18_04:v1.2
WORKDIR /app
COPY --from=builder /app/build/bin/linux/ /app/bin/
ENV GO_CONFIG /app/config/config.yaml
ENV GO_SECRET /app/secret/config.yaml
CMD ["/app/bin/git-util"]

# docker build -t git-util:abc .
# docker run --name git-util-dev -d -v `pwd`:/app/config -p 8000:8000 git-util:abc