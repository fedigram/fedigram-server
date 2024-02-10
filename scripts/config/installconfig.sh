skip() { true; }
skip auth_key.toml
cp -v auth_session.toml ../../service/auth_session
cp -v biz_server.toml ../../messenger/biz_server
skip config.json
cp -v document.toml ../../service/document
skip frontend.toml
skip lang_pack_en.toml
skip lang_pack_ru.toml
skip session.toml
skip sync.toml
skip upload.toml
echo "done"
