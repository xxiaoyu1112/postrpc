REPO_DIR=$(dirname $(readlink -f "$0"))/..
APP_ID=post_data_collect

# RPC 服务目录
RPC_REPO_PATH=$REPO_DIR/$APP_ID/idl

if [  ! -d $RPC_REPO_PATH ]; then
  mkdir -p $RPC_REPO_PATH
fi

# RPC SDK目录
CLIENT_PATH=$REPO_DIR/rpc_sdk

# IDL 文件夹目录
IDL_ROOT=$REPO_DIR/idl
# IDL 目录
IDL_PATH=$REPO_DIR/idl/$APP_ID.proto

# 打印更新IDL
echo IDL\:$IDL_PATH
# 更新RPC服务端
protoc --proto_path=$IDL_ROOT\
    --go_out=$RPC_REPO_PATH \
    --go-grpc_out=$RPC_REPO_PATH \
      $IDL_PATH

# 更新SDK 
protoc  --proto_path=$IDL_ROOT\
  --go_out=$CLIENT_PATH \
  --go-grpc_out=$CLIENT_PATH \
      $IDL_PATH