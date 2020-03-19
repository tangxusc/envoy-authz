FROM scratch
COPY cmd/main /
EXPOSE 9999
CMD ["/main"]