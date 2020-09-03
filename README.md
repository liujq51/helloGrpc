### 生成 pb.go

protoc --go_out=plugins=grpc:. ./proto/*.proto

### 测试运行

serve+client