FROM golang:1.22-alpine AS stage1
WORKDIR /project/vegn/

COPY go.* .
RUN go mod download

COPY . . 
RUN go build -o ./cmd/vegnExecutableFile ./cmd/main.go

#stage2 (only copying the executable and enviornment files to reduce the size,the stage1 image will be replaced by stage2)
FROM scratch
COPY --from=stage1 /project/vegn/cmd/vegnExecutableFile /project/vegn/
COPY --from=stage1 /project/vegn/.env /project/vegn/
COPY --from=stage1 /project/vegn/templates /project/vegn/templates/

WORKDIR /project/vegn/

EXPOSE 8080
ENTRYPOINT [ "/project/vegn/vegnExecutableFile" ]