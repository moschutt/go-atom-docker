FROM go_atom

ENV PATH=$PATH:/usr/local/go/bin

USER root
RUN curl -L https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz > /tmp/go.tar.gz && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz

RUN apt-get update && \
    apt-get install -y golint \
       gdb
USER atom
RUN mkdir -p /home/atom/go

WORKDIR /home/atom/go
COPY atom-requirements.txt /tmp
#RUN rm ~/package-lock.json
RUN apm install --packages-file /tmp/atom-requirements.txt
RUN echo "PATH=$PATH:$(go env GOPATH)/bin" >> ~/.bashrc
RUN echo "GOPATH=$(go env GOPATH)" >> ~/.bashrc

CMD ["/usr/bin/atom", "-f"]
