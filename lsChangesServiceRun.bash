systemctl  stop lsChanges
cd /var/www/html/lsChanges
go build -o lsChanges
cp -f /var/www/html/lsChanges/lsChanges.service /etc/systemd/system/lsChanges.service
nanuseradd -s /sbin/nologin -M lsChanges 
mkdir -p /opt/lsChanges/
cp  -f /var/www/html/lsChanges/*.html /opt/lsChanges/ 
mv -f /var/www/html/lsChanges/lsChanges /opt/lsChanges/lsChanges
chmod 755 /opt/lsChanges/lsChanges

systemctl daemon-reload
systemctl enable lsChanges
systemctl start lsChanges