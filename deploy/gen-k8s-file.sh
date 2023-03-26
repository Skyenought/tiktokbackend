#!/usr/bin/env bash

service=$1
mode=$2
image=$3
port=$4
yamlName=$5
nodePort=$6

namespace=tiktokbackend
serviceAccount=find-endpoints

echo "开始创建${mode}服务：$service 的k8s 文件"
cd ../
if [ -f "cmd/$service/$mode}/$yamlName.yaml" ]; then
  rm cmd/"${service}"/"${mode}"/"${yamlName}".yaml
fi
if [ -z "$nodePort" ]; then
  cd "cmd/$service/$mode" && \
  	goctl kube deploy -replicas 3 -requestCpu 50 -requestMem 60 -limitCpu 100 -limitMem 100 \
  	-name "$service-$mode" \
  	-namespace ${namespace} -image "$image" \
  	-o "$yamlName.yaml" -port "$port" \
  	-serviceAccount "$serviceAccount"
else
  cd "cmd/$service/$mode" && \
    	goctl kube deploy -nodePort "$nodePort" -replicas 3 -requestCpu 50 -requestMem 60 -limitCpu 100 -limitMem 100 \
    	-name "$service-$mode" \
    	-namespace ${namespace} -image "${image}" \
    	-o "$yamlName.yaml" -port "${port}" \
    	-serviceAccount ${serviceAccount}
fi
