#!BuildTag: cis-operator:latest

# Builder image
FROM bci/golang:1.21 as builder

ARG CIS_OPERATOR=cis-operator

COPY . /"$CIS_OPERATOR"

WORKDIR /"$CIS_OPERATOR" 

RUN CGO_ENABLED=0 go build -o "$CIS_OPERATOR"

# Final image
FROM bci/bci-micro:15.5

ARG CIS_OPERATOR=cis-operator

COPY --from=builder /"$CIS_OPERATOR"/"$CIS_OPERATOR" /usr/bin
COPY --from=builder /"$CIS_OPERATOR"/pkg /pkg

USER 65535:65535

RUN ls -lha

CMD ["$CIS_OPERATOR"]

