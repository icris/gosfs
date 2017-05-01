FROM scratch
EXPOSE 8000

COPY sfs /sfs
COPY data /data
WORKDIR /data
CMD ["/sfs"]
