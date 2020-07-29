protoc -I=./message --go_out=./message ./message/message.proto

mkdir -p ../ui/src/proto/message
#cp -R ./message/*.proto ../ui/src/proto/message/
protoc -I=./message --js_out=import_style=commonjs,binary:../ui/src/proto/message ./message/message.proto
echo -e "/*eslint-disable*/\n$(cat ../ui/src/proto/message/message_pb.js)" > ../ui/src/proto/message/message_pb.js