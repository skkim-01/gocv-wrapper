#!/bin/bash

mkdir -p target/android 

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
rm -f target/android/*


CGO_LDFLAGS=-L/Users/bobkim/opencv-libs/OpenCV-android-sdk/sdk/native/libs -lpthread -ldl -ldliaPlugin -lHeteroPlugin -lMKLDNNPlugin -lmyriadPlugin -linference_engine -lclDNNPlugin -lopencv_core -lopencv_pvl -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video -lopencv_dnn -lopencv_calib3d

CGO_LDFLAGS=

gomobile bind --target android -v -o target/android/gocv-wrapper.aar ./src/ifaces
