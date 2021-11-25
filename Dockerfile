FROM ubuntu:18.04

RUN apt-get update
RUN apt-get install ca-certificates -y

COPY bin /usr/local/bin
CMD ["./usr/local/bin/impact-analysis-server"]
