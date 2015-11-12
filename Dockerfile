FROM alpine:3.2
RUN apk add -U ca-certificates python py-pip && rm -rf /var/cache/apk/*
ADD drone-pypi /bin/
ENTRYPOINT ["/bin/drone-pypi"]
