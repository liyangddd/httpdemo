FROM cicd-test-cn-beijing.cr.volces.com/test1/docker-gen:0.14.1-debian AS docker-gen
# Install Forego + docker-gen
COPY gohttp_demo /usr/local/bin/gohttp_demo

WORKDIR /usr/local/bin/

ENTRYPOINT ["./gohttp_demo"]