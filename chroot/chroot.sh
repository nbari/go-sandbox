#!/bin/sh

# Specify HERE the apps you want into the enviroment
APPS="/bin/sh /bin/bash /bin/ls /bin/mkdir /bin/mv /bin/pwd /bin/rm"

# Sanity check
if [ "$1" = "" ] ; then
    echo "     Usage: $0 destinationPath"
    exit
fi

# Go to the destination directory
if [ ! -d $1 ]; then mkdir $1; fi
cd $1

# Create Directories
mkdir etc
mkdir bin
mkdir usr
mkdir usr/bin
mkdir usr/lib
mkdir usr/lib/system

# Add some users to ./etc/paswd
grep /etc/passwd -e "^root" > etc/passwd
grep /etc/group -e "^root" > etc/group

# Copy the apps and the related libs
for prog in $APPS;  do
    cp $prog ./$prog

    # obtain a list of related libraryes
    LIBS=`otool -L $prog | grep version | awk '{ print $1 }'`
    for l in $LIBS; do
        cp $l ./$l
    done
done
