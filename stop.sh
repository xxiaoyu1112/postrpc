source build_all.sh
ports=(50051 50052 50053 50054)


for port in "${ports[@]}"
do
pIDa=`lsof -i :$port | grep -v "PID" `
if [ "$pIDa" != "" ];
then
   fuser -k -n tcp $port   
   echo "close port: $port"
fi
 
done