FROM scratch
ADD main /
ADD cacert.pem /etc/ssl/certs/
ENTRYPOINT ["/main"]
CMD ["start"]
