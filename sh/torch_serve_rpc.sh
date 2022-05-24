REPO_DIR=$(dirname "$0")/..
# RPC GroupID
APP_GROUP=torch_serve
# RPC APP ID
APP_ID_1=inference
APP_ID_2=management
# RPC 客户端路径
CLIENT_PATH=$REPO_DIR/rpc_sdk/torch_serve
# 判断SDK目录是否存在
if [  -d $CLIENT_PATH ]; then
  rm -r $CLIENT_PATH
fi
mkdir $CLIENT_PATH
# IDL根路径
IDL_ROOT=$REPO_DIR/idl/external
# IDL路径
IDL_PATH_1=$REPO_DIR/idl/external/$APP_ID_1.proto
IDL_PATH_2=$REPO_DIR/idl/external/$APP_ID_2.proto
# 输出更新RPC的IDL
echo IDL\:$IDL_PATH_1,$IDL_PATH_2

# 更新APP1
protoc  --proto_path=$IDL_ROOT\
  --go_out=$CLIENT_PATH \
  --go_opt=paths=source_relative\
  --go-grpc_out=$CLIENT_PATH\
  --go-grpc_opt=paths=source_relative\
      $IDL_PATH_1 

# 更新APP2 
protoc  --proto_path=$IDL_ROOT\
  --go_out=$CLIENT_PATH \
  --go_opt=paths=source_relative\
  --go-grpc_out=$CLIENT_PATH\
  --go-grpc_opt=paths=source_relative\
      $IDL_PATH_2 