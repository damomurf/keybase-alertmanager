FROM golang:1.13.8 as builder

RUN mkdir /build
COPY . /build/

WORKDIR /build
RUN go build -mod=vendor .

FROM keybaseio/client:5.3.0-20200214172212-6d33303518-slim

COPY --from=builder /build/kbam /usr/bin/kbam
COPY default.tmpl .

RUN useradd --create-home --shell /bin/bash kbam

USER kbam

ENTRYPOINT [ "/usr/bin/kbam" ]