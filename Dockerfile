FROM golang:1.19-alpine as BUILD
WORKDIR /src/app
COPY . .
RUN CGO_ENABLED=0 go build -o gonyast

FROM gcr.io/distroless/static-debian11
WORKDIR /src/app 
COPY --from=BUILD /src/app/gonyast .
ENTRYPOINT [ "/src/app/gonyast" ]