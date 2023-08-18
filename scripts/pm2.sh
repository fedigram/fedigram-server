pm2 stop all
sleep 10
cd /home/anhttn/go/src/github.com/fedigram/fedigram-server/service/auth_session
pm2 start ./auth_session
sleep 1
cd /home/anhttn/go/src/github.com/fedigram/fedigram-server/service/document
pm2 start ./document
sleep 1
cd /home/anhttn/go/src/github.com/fedigram/fedigram-server/messenger/sync
pm2 start ./sync
sleep 1
cd /home/anhttn/go/src/github.com/fedigram/fedigram-server/messenger/upload
pm2 start ./upload
sleep 1
cd /home/anhttn/go/src/github.com/fedigram/fedigram-server/messenger/biz_server
pm2 start ./biz_server
sleep 1
cd /home/anhttn/go/src/github.com/fedigram/fedigram-server/access/auth_key
pm2 start ./auth_key
sleep 1
cd /home/anhttn/go/src/github.com/fedigram/fedigram-server/access/session
pm2 start ./session
sleep 1
cd /home/anhttn/go/src/github.com/fedigram/fedigram-server/access/frontend
pm2 start ./frontend
sleep 1
