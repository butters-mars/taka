#!/bin/sh
DIR=$1
OUT=$2

if [ "" == "$DIR" ]; then
    echo "dir is empty"
    exit
fi

if [ "" == "$OUT" ]; then
    echo "out is empty"
    exit
fi

echo "cleanup $OUT ..."
rm $OUT/*.go
rm $OUT/*.js
#rm $OUT/*.dart
rm $OUT/*.swagger.json

echo "generating *.pb.go ..."
protoc -I/usr/local/include -I$DIR \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:$OUT \
  --validate_out=lang=go:$OUT \
  --descriptor_set_out $OUT/out.pb \
  $DIR/*.proto 

echo "generating reverse-proxy ..."
protoc -I/usr/local/include -I $DIR \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true,grpc_api_configuration=$DIR/proxy.yaml:$OUT \
  $DIR/*.proto 

echo "generating swagger ..."
protoc -I/usr/local/include -I $DIR \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --swagger_out=logtostderr=true,grpc_api_configuration=$DIR/proxy.yaml:$OUT \
    $DIR/*.proto 

echo "generating web-client ..."
protoc -I=$DIR \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --js_out=import_style=commonjs:$OUT \
  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:$OUT \
  $DIR/*.proto 

echo "generating dart-client ..."
#protoc -I=$DIR \
#  -I$GOPATH/src \
#  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#  --dart_out=grpc:$OUT \
#  $DIR/*.proto 


echo "go get ."
D=`pwd`
cd $OUT
go get .
cd $D
