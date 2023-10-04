#!BuildTag: cis-operator:latest

# Builder image
FROM opensuse/bci/golang:stable as builder

ARG CIS_OPERATOR=cis-operator

COPY . /"$CIS_OPERATOR"
RUN ls -lha /"$CIS_OPERATOR"

WORKDIR /"$CIS_OPERATOR" 

RUN CGO_ENABLED=0 go build -o "$CIS_OPERATOR"


# Final image
FROM opensuse/bci/bci-micro:latest

ARG CIS_OPERATOR=cis-operator

COPY --from=builder /"$CIS_OPERATOR"/"$CIS_OPERATOR" /usr/bin
COPY --from=builder /"$CIS_OPERATOR"/pkg /pkg

USER 65535:65535

CMD ["$CIS_OPERATOR"]

