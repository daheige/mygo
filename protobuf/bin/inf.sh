#!/bin/sh
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)
packageName=helloworld
cd $root_dir/common/
# protoc --go_out=plugins=grpc:. $root_dir/common/inf.proto
# protoc --go_out=plugins=grpc:. inf.proto
protoc3 --go_out=plugins=grpc:. inf.proto

mkdir -p $root_dir/$packageName

mv $root_dir/common/inf.pb.go $root_dir/$packageName/

cd $root_dir/$packageName/
go install
