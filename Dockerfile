FROM alpine
ADD go-srv /go-srv
ENTRYPOINT [ "/go-srv" ]
