FROM ubuntu:bionic

COPY bin/tango-api .

CMD ["/tango-api"]
