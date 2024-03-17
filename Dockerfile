ARG base="ubuntu:jammy"
FROM $base
USER root
LABEL MAINTAINER Author <vsoch>

# install go 20.10
RUN apt-get update && apt-get install -y wget python3-pip
RUN wget https://go.dev/dl/go1.20.10.linux-amd64.tar.gz  && tar -xvf go1.20.10.linux-amd64.tar.gz && \
         mv go /usr/local && rm go1.20.10.linux-amd64.tar.gz

ENV PATH=/usr/local/go/bin:$PATH
WORKDIR /code
COPY . /code
RUN make build-all && \
    cp ./bin/fracta* /usr/bin/

ENTRYPOINT ["fractal"]

# Anticipate different running contexts
EXPOSE 80
EXPOSE 8080
EXPOSE 9092
EXPOSE 443
EXPOSE 50051
CMD ["leader", "--force-exit"]
