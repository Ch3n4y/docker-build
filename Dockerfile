FROM alpine

LABEL maintainer="Chaney <csl@live.com>"

# Copy python requirements file
COPY conf/requirements.txt /tmp/requirements.txt
RUN apk add --no-cache \
    python3 \
    bash \
    supervisor && \
    python3 -m ensurepip && \
    rm -r /usr/lib/python*/ensurepip && \
    pip3 install --upgrade pip setuptools && \
    pip3 install -r /tmp/requirements.txt && \
    rm -r /root/.cache

# Custom Supervisord config
COPY conf/supervisord.conf /etc/supervisord.conf


COPY ./app /app
WORKDIR /app

CMD ["/usr/bin/supervisord"]