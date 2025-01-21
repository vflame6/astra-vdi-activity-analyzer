#!/bin/bash
#
# This script installs the agent as a service on the system.

if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit 1
fi

echo "Checking the configuration file (config.json)"
if [ ! -f ./config.json ]; then
    echo "Configuration file not found!"
    exit 1
fi
cat ./config.json
read -p "Is the configuration correct? (n/Y)" config_correct
case $config_correct in
  [Yy]* ) echo "[*] Proceeding...";;
  * ) echo "Please review the configuration file."; exit;;
esac

mkdir /etc/astra-dlp
mkdir /etc/astra-dlp/data
cp agent-astra /etc/astra-dlp/
cp config.json /etc/astra-dlp/
chown -R root: /etc/astra-dlp/
chmod -R 700 /etc/astra-dlp/

echo "[*] Running the agent in register mode"
cd /etc/astra-dlp/
/etc/astra-dlp/astra-agent --register

echo "[*] Installing the service"

cat << EOF > /etc/systemd/system/astra-agent.service
[Unit]
Description=Astra DLP agent
After=multi-user.target

[Service]
Type=simple
Restart=always
RestartSec=60
User=root
ExecStart=/etc/astra-dlp/astra-agent

[Install]
WantedBy=graphical.target
EOF

systemctl start astra-agent.service
systemctl enable astra-agent.service
