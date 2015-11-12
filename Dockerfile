FROM alpine:3.2

RUN apk add -U \
	ca-certificates \
	py-pip \
	python \
 && rm -rf /var/cache/apk/* \
 && pip install --no-cache-dir --upgrade pip

ADD drone-pypi /bin/

ENTRYPOINT ["/bin/drone-pypi"]
