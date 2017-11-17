FROM alpine 

COPY hash-svc /

ENTRYPOINT ["/bin/sh","-c"]
CMD ["/hash-svc"]
