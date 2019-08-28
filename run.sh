#!/bin/bash

confName=$(grep wgConfig: bravo.yaml|awk '{print $2}')
if [[ ! -f /etc/wireguard/${confName}.conf ]]; then
	echo "bravo: unable to find /etc/wireguard/${confName}.conf, please make sure to name your Wireguard config correctly\n"
	exit 1
fi
echo 'bravo: enter your password to run this script as a superuser (sudo)'
sudo -l
sudo ./bravo &
