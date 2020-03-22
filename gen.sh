#!/bin/sh
cd $1
DIR=.
OUT=.

if [ "" == "$DIR" ]; then
    echo "dir is empty"
    exit
fi

if [ "" == "$OUT" ]; then
    echo "out is empty"
    exit
fi

# Set parent directory to hold all the symlinks
PROTOBUF_IMPORT_DIR='protobuf-import'
mkdir -p "${PROTOBUF_IMPORT_DIR}"

# Remove any existing symlinks & empty directories 
find "${PROTOBUF_IMPORT_DIR}" -type l -delete
find "${PROTOBUF_IMPORT_DIR}" -type d -empty -delete

# Download all the required dependencies
go mod download

# Get all the modules we use and create required directory structure
go list -f "${PROTOBUF_IMPORT_DIR}/{{ .Path }}" -m all \
  | xargs -L1 dirname | sort | uniq | xargs mkdir -p

# Create symlinks
go list -f "{{ .Dir }} ${PROTOBUF_IMPORT_DIR}/{{ .Path }}" -m all \
  | xargs -L1 -- ln -s

echo "cleanup $OUT ..."
rm $OUT/*.pb*.go
rm $OUT/*.js
rm $OUT/*.dart
rm $OUT/*.swagger.json

echo "generating *.pb.go ..."
protoc -I/usr/local/include -I$DIR \
  -I./$PROTOBUF_IMPORT_DIR \
  --go_out=plugins=grpc:$OUT \
  --validate_out=lang=go:$OUT \
  --descriptor_set_out $OUT/out.pb \
  $DIR/*.proto 

protoc \
  -I./$PROTOBUF_IMPORT_DIR \
  -I /usr/local/include \
  -I $DIR \
  -I . \
  --gotag_out=xxx="bson+\"-\"":$OUT \
  $DIR/*.proto

echo "generating reverse-proxy ..."
protoc -I/usr/local/include -I $DIR \
  -I./$PROTOBUF_IMPORT_DIR \
  --grpc-gateway_out=logtostderr=true,grpc_api_configuration=$DIR/proxy.yaml:$OUT \
  $DIR/*.proto 

echo "generating swagger ..."
protoc -I/usr/local/include -I $DIR \
    -I./$PROTOBUF_IMPORT_DIR \
    --swagger_out=logtostderr=true,grpc_api_configuration=$DIR/proxy.yaml:$OUT \
    $DIR/*.proto 

echo "generating web-client ..."
protoc -I=$DIR \
  -I./$PROTOBUF_IMPORT_DIR \
  --js_out=import_style=commonjs:$OUT \
  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:$OUT \
  $DIR/*.proto 

echo "generating dart-client ..."
protoc -I=$DIR \
 -I./$PROTOBUF_IMPORT_DIR \
 --dart_out=grpc:$OUT \
 $DIR/*.proto 

rm -rf $PROTOBUF_IMPORT_DIR

echo "go get ."

mv *.go *.dart ../
