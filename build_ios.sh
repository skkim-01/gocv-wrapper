#!/bin/bash

mkdir -p target/ios

# currently ignore backup script
#TM=$(date +%s)

#FIND_FILE=$(ls target/android | grep wallet)

#if [ -z $FIND_FILE ]; then
#   echo BACKUP: FILE NOT FOUND
#else
#   echo BACKUP: target/android/$TM
#   mkdir -p target/android/$TM
#   mv target/android/wallet* target/android/$TM
#fi

# remove previous library files
rm -rf target/ios/*



#CGO_LDFLAGS=-L/opencv/lib -L/deployment_tools/inference_engine/lib/ubuntu_16.04/intel64 -lpthread -ldl -ldliaPlugin -lHeteroPlugin -lMKLDNNPlugin -lmyriadPlugin -linference_engine -lclDNNPlugin -lopencv_core -lopencv_pvl -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video -lopencv_dnn -lopencv_calib3d

#CGO_LDFLAGS=

CGO_ENABLED=1 gomobile bind --target ios -v -o target/ios/gocv-wrapper.framework ./src/ifaces