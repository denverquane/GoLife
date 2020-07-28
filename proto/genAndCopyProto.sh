protoc -I=./message --go_out=./message ./message/message.proto

mkdir -p ../ui/src/proto/message
cp -R ./message/*.proto ../ui/src/proto/message/