#!/bin/sh
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)
proto_bin=$(which "protoc3")
if [ -z $proto_bin ]; then
    echo 'Please install protoc3!'
    exit
fi

server_dir=$root_dir/server/pb
protos_dir=$root_dir/protos
if [ ! -d $server_dir ];then
    mkdir -p $server_dir
fi

echo "build go grpc pb file"

cd $root_dir
$proto_bin -I $protos_dir --go_out=plugins=grpc:$server_dir $protos_dir/*.proto
cd $server_dir
go install #编译

echo "make proto success!"
exit 0
