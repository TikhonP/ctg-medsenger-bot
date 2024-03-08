# Install GO
curl -OL https://go.dev/dl/go1.22.1.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz

./build.sh

sudo cp ctg_medsenger_bot.conf /etc/supervisor/conf.d/
sudo cp ctg_nginx.conf /etc/nginx/sites-enabled/
sudo supervisorctl update
sudo systemctl restart nginx
sudo certbot --nginx -d ctg.ai.medsenger.ru