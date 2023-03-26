#!/usr/bin/env bash

service=$1
mode=$2
tag=$3
contain=$4

echo "开始生成 $service.$mode 的Dockerfile"
if [ -f ../cmd/"${service}"/"${mode}"/Dockerfile ]; then
  rm ../cmd/"${service}"/"${mode}"/Dockerfile
fi

cd ../
mkdir cmd/"${service}"/"${mode}"/etc && \
	cp etc/"${mode}"s/"${service}"-k8s.yaml cmd/"${service}"/"${mode}"/etc/"${service}".yaml && \
	cd cmd/"${service}"/"${mode}" && goctl docker -go "${service}".go

echo "开始构建镜像"
cd ../../..
pwd
if [ "$contain" = "docker" ]; then
  docker build -t "$service-$mode:$tag" -f "cmd/$service/$mode/Dockerfile" .
elif [ "$contain" = "containerd" ]; then
  nerdctl build --namespace k8s.io -t "$service-$mode:$tag" -f "cmd/$service/$mode/Dockerfile" .
else
  echo "请输入正确的container参数：docker或nerdctl"
fi

rm -rf cmd/"${service}"/"${mode}"/etc

echo "清除暂存文件完成"
