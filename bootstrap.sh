source build_all.sh
ports=(50051 50052 50053 50054)

modules=("post_data_collect" "post_data_manage" "post_model_manage" "post_model_predict")

echo "stop old service"

for port in "${ports[@]}"
do
pIDa=`lsof -i :$port | grep -v "PID" `
if [ "$pIDa" != "" ];
then
   fuser -k -n tcp $port   
   echo "close port: $port"
fi
 
done


for module in "${modules[@]}"
do
    $REPO_PATH/$module/out/$module &
   # do whatever on "$i" here
done