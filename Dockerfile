FROM scratch
COPY sendbird-go /
ENTRYPOINT ["/sendbird-go"]
