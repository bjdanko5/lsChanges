sudo systemctl  stop lsChanges
cd /var/www/html/lsChanges
go build -o lsChanges
sudo cp -f /var/www/html/lsChanges/lsChanges.service /etc/systemd/system/lsChanges.service
sudo useradd -s /sbin/nologin -M lsChanges 
sudo mkdir -p /opt/lsChanges/
sudo rm  -f  /opt/lsChanges/*.html 
sudo cp  -f /var/www/html/lsChanges/*.html /opt/lsChanges/ 
sudo mv -f /var/www/html/lsChanges/lsChanges /opt/lsChanges/lsChanges
sudo chmod 755 /opt/lsChanges/lsChanges

sudo systemctl daemon-reload
sudo systemctl enable lsChanges
sudo systemctl start lsChanges