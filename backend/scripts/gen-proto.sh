protoc --proto_path=./../libs/proto --go_out=./../libs --go-grpc_out=./../libs ./../libs/proto/*/*.proto
python3 -m grpc_tools.protoc --proto_path=./../libs/proto/ \
    --python_out=./../libs/gen/python --grpc_python_out=./../libs/gen/python \
    ./../libs/proto/*/*.proto
