#FROM asia.grc.io/ccinpp-190005/rclone:1.39-amd64
FROM scratch
ADD main /
ADD cacert.pem /etc/ssl/certs/
ENTRYPOINT ["/main"]
CMD ["start"]
