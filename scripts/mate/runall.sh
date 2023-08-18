function run() {
  cd $GOPATH/src/github.com/fedigram/fedigram-server/$1
  mate-terminal -e ./$2 --title=$2 &
}

echo "spawning..."
run service/auth_session	auth_session
run service/document		document
run messenger/sync		sync
run messenger/upload		upload
run messenger/biz_server	biz_server
run access/auth_key		auth_key
run access/session		session
run access/frontend		frontend
echo "spawn done."
