FROM golang:1.13.8 as builder

RUN mkdir /build
COPY . /build/

WORKDIR /build
RUN go build -mod=vendor .

FROM keybaseio/client:5.6.0-20200702122450-d407f6ad44-slim

COPY --from=builder /build/kbam /usr/bin/kbam
COPY default.tmpl .

RUN useradd --create-home --shell /bin/bash kbam

USER kbam

ENTRYPOINT [ "/usr/bin/kbam" ]