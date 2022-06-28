#!/usr/bin/env bash

pushd service
  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative terraxen.proto
  result=$?
popd
if [[ "${result}" == "0" ]]; then
  echo "Updating git"
  git add service/terraxen.proto
  git add service/terraxen_grpc.pb.go
  git add service/terraxen.pb.go
  git commit -m "Updating gRPC proto"
  git push
fi