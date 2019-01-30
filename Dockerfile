FROM alpine:3.9

RUN apk add -U \
	ca-certificates \
	py-pip \
	python \
 && rm -rf /var/cache/apk/* \
 && pip install --no-cache-dir --upgrade \
	pip \
	setuptools \
	wheel

ADD drone-pypi /bin/

ENTRYPOINT ["/bin/drone-pypi"]
