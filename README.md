[![License](https://img.shields.io/github/license/open-telegram-server/chatengine.svg)](https://github.com/open-telegram-server/chatengine/blob/master/LICENSE)

# fedigram-server

**fedigram-server** is a fork of nebula-chat/chatengine (later renamed to teamgram)

**fedigram-server's main repo:** https://github.com/fedigram/fedigram-server

**fedigram-server's project status:** Unknown; working on this. Currently making the code build and
run from https://github.com/fedigram/fedigram-server repo.

**fedigram-server's chats:**

  * `irc.ilita.i2p`   `#fedigram-dev` (see also: https://i2pd.website/ https://geti2p.net/ )
  * `irc.libera.chat` `#fedigram-dev`
  * не хотите телегу с опенсорсным сервером поделать федеративную как федиверс? `#fedigram-dev`

## Quick start with Docker

1. Install `docker` and `docker-compose`
2. Edit `docker-compose.yml`: replace `CHATENGINE_HOST` to "127.0.0.1".
3. ? Edit `scripts/config/config.json` and specify `/data2/dc_options{ip_address,port}`. ?
4. Run `make -j$(nproc)` command in your shell at the git repo folder.
5. Now, fedigram server is listening on TCP port `12345`.

## Websites (aren't currently working)

 * http://fedigram.i2p/
 * http://www.fedigram.tranoo.com/

#### Build 

 - Get the source code　

   ```bash
   mkdir -pv $GOPATH/src/github.com/fedigram
   cd $GOPATH/src/github.com/fedigram
   #then, if you have ssh key on github
   eval `ssh-agent`
   ssh-add your-github-ssh-key
   git clone git@github.com:fedigram/fedigram-server.git
   #or, if you don't have ssh key on github
   git clone https://github.com/fedigram/fedigram-server.git
   ```

 - Build

    ```bash
    source $GOPATH/src/github.com/fedigram/fedigram-server/scripts/config/build.sh
    ```

- Init
    - OS tested: Ubuntu 20.04.6 `Linux 5.15.0-92-generic #102~20.04.1-Ubuntu SMP Mon Jan 15 13:09:14 UTC 2024 x86_64`
    - mysql tested: mysql-server-8.0 (8.0.36-0ubuntu0.20.04.1)
    - install mysql somewhere
    - then,
    
    ```bash
    cd $GOPATH/src/github.com/fedigram/fedigram-server/scripts/
    mysql -u root -p
    # at mysql: 
        CREATE DATABASE PluralityServer;
        use PluralityServer;
        source PluralityServer.sql
        source merge_20181129_201906.sql
        exit;
    ```
- Configure: edit `*.toml *.json` files at `$GOPATH/src/github.com/fedigram/fedigram-server/scripts/config/`:
    * auth_session service: `auth_session.toml`
    * document service:  `document.toml`
    * sync service: `sync.toml`
    * upload service:  `upload.toml`
    * biz_server service:  `config.json lang_pack_en.toml lang_pack_cn.toml biz_server.toml`
    * auth_key service: `server_pkcs1.key auth_key.toml`
    * session service: `session.toml`
    * frontend service: `frontend.toml`
# [OBSOLETE UNEDITED INFO BELOW]
- Run
    ```shell
    cd $GOPATH/src/github.com/fedigram/fedigram-server/service/auth_session
    ./auth_session
    
    cd $GOPATH/src/github.com/fedigram/fedigram-server/service/document
    ./document

    cd $GOPATH/src/github.com/fedigram/fedigram-server/messenger/sync
    ./sync
    
    cd $GOPATH/src/github.com/fedigram/fedigram-server/messenger/upload
    ./upload

    cd $GOPATH/src/github.com/fedigram/fedigram-server/messenger/biz_server
    ./biz_server

    cd $GOPATH/src/github.com/fedigram/fedigram-server/access/auth_key
    ./auth_key

    cd $GOPATH/src/github.com/fedigram/fedigram-server/access/session
    ./session
    
    cd $GOPATH/src/github.com/fedigram/fedigram-server/access/frontend
    ./frontend
    ```

----------------------------

# OBSOLETE/OLD INFO BELOW

### Introduction
An open source [mtproto](https://core.telegram.org/mtproto) server implemented in go language
with compatible old (layer 86) [Telegram](https://telegram.org/) clients.

## Quick start

1. Run `sudo apt install docker docker-compose` in your shell;
2. `git clone --recursive git@github.com:fedigram/fedigram-server.git && cd fedigram-server`
3. Edit `./docker-compose.yml`: replace timezone with your own. There were reports that having a wrong timezone makes chatengine fail.
4. Run `sudo make -j$(nproc)` command in your shell;
5. Now, chatengine is running on your host's TCP port `12345`;
6. Use [fedigram clients](https://github.com/fedigram/fedigram-clients) to connect to fedigram server;
7. Enjoy!

## The rest of this README is for developers

### Architecture
![Architecture](doc/image/architecture-001.jpeg)

### Documents
[Diffie–Hellman key exchange](doc/dh-key-exchange.md)

[Creating an Authorization Key](doc/Creating_an_Authorization_Key.md)

[Mobile Protocol: Detailed Description (v.1.0, DEPRECATED)](doc/Mobile_Protocol-Detailed_Description_v.1.0_DEPRECATED.md)

[Encrypted CDNs for Speed and Security](doc/cdn.md) Translate By [@steedfly](https://github.com/steedfly)

[Windows-Build](doc/windows-build.md) By [@robinfoxnan](https://github.com/robinfoxnan)

### Manual Build and Install

Note: You will probably need a VM for this as the code often uses root at MySQL and
root for filesystem write access.

```bash
git clone https://github.com/fedigram/fedigram-server
cd fedigram-server
# replace 192.168.1.100 to you own host IP.
# sed -i "" 's/CHATENGINE_HOST=127.0.0.1/CHATENGINE_HOST=192.168.1.100/g' docker-compose.yml # macOS
sed -i 's/CHATENGINE_HOST=127.0.0.1/CHATENGINE_HOST=192.168.1.100/g' docker-compose.yml # linux
make -j$(nproc)
```

#### Dependencies

 - redis
 - mysql
 - etcd

#### More

[Build document](doc/build.md)

[Build script](scripts/build.sh)

[Prerequisite script](scripts/prerequisite.sh)


#### SQL

You need all `scripts/*.sql`.

#### Compatible clients

**Important**: default signIn and signOut verify code is **12345**

[Android client for NebulaChat](https://github.com/fedigram/fedigram-clients/tree/master/Telegram-Android)

[FOSS client for NebulaChat](https://github.com/fedigram/fedigram-clients/tree/master/Telegram-FOSS)

[iOS client for NebulaChat](https://github.com/fedigram/fedigram-clients/tree/master/Telegram-iOS)

[tdesktop for NebulaChat](https://github.com/fedigram/fedigram-clients/tree/master/tdesktop)


## Original Nebula Chat author's notes

Chatengine is not a commercial project, only supports mtproto API layer 86, and
only supports private chats and small groups. 

If need enterprise edition, please PM the [author](https://t.me/benqi) or download
clients from [nebula.chat](https://nebula.chat) (default verify code is: 12345).
