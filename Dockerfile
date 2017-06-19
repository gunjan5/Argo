FROM golang:1.8 
RUN mkdir /code 
ADD ./code /code 
WORKDIR /code 
RUN go build -o main . 
CMD ["/code/main"]

