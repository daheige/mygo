#!/bin/bash
#运维上线需执行该脚本
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

mkdir -p $root_dir/runtime/logs
chmod 777 -R $root_dir/runtime
