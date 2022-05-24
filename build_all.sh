#!/bin/bash
# author:菜鸟教程
# url:www.runoob.com
WS=$(pwd)

REPO_PATH=$(dirname $(readlink -f "$0"))

modules=("post_data_collect" "post_data_manage" "post_model_manage" "post_model_predict")

# cd $REPO_PATH/post_data_collect
# sh build.sh 
# cd ..
for module in "${modules[@]}"
do
    echo $module
    cd $REPO_PATH/$module
    source build.sh 
   # do whatever on "$i" here
done
cd $WS
# sh $REPO_PATH/post_data_manage/build.sh 
# sh $REPO_PATH/post_model_manage/build.sh 
# sh $REPO_PATH/post_model_predict/build.sh 
