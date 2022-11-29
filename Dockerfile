FROM scratch
MAINTAINER Entire studio <dev@entire.studio>
ADD hello /
ENTRYPOINT ["/hello"]
