FROM bufbuild/buf:0.43.2

RUN mkdir build

WORKDIR build

COPY . .

ENTRYPOINT buf lint