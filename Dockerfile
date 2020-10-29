FROM golang:1.14

WORKDIR /golife

COPY . .

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y libgl1-mesa-dev xorg-dev
RUN make build

CMD ["go", "test", "-timeout=6h", "-benchtime=2m", "-bench=.", "./pkg/life/life_test.go"]
