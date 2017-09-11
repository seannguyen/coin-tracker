#!/bin/bash

INSTALLATION_DIR_TEMP=/var/www/coin-tracker
INSTALLATION_DIR=/var/www/coin-tracker-temp

REMOTE_COMMANDS=`xargs << COMMANDS
    sudo apt-get install python-dev gcc

    sudo mkdir ${INSTALLATION_DIR_TEMP}
    sudo chown seannguyen:seannguyen ${INSTALLATION_DIR_TEMP}
    git clone ${INSTALLATION_DIR_TEMP}
    ssh node01 -t curl https://gist.githubusercontent.com/SeanNguyen/f3d6f427d947246df77f36c4ff4544a5/raw/449a9736406638126ffa5e72d4ed324ead509bd0/configs.py > ${INSTALLATION_DIR_TEMP}/lib/config.py

    rm -rf ${INSTALLATION_DIR}
    mv ${INSTALLATION_DIR_TEMP} ${INSTALLATION_DIR}
COMMANDS`

ssh node01 -t ${REMOTE_COMMANDS}