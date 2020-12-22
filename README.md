[![License](https://img.shields.io/github/license/nebula-chat-fork/chatengine.svg)](https://github.com/nebula-chat-fork/chatengine/blob/master/LICENSE)

# PluralityServer

**PluralityServer** is a fork of https://github.com/nebula-chat/chatengine

**PluralityServer's main repo:** https://github.com/PluralityNET/PluralityServer

**PluralityServer's status:** Limited functionality (as in public NebulaChat's chatengine). Currently making the code build and run from Plurality repo.

----------------------------

# OLDER INFO BELOW

# NebulaChat - Open source [mtproto](https://core.telegram.org/mtproto) server written in golang
> open source mtproto server implemented in golang with compatible telegram client.

### Introduction
Open source [mtproto](https://core.telegram.org/mtproto) server written in golang

### Architecture
![Architecture](doc/image/architecture-001.jpeg)

### Documents
[Diffie–Hellman key exchange](doc/dh-key-exchange.md)

[Creating an Authorization Key](doc/Creating_an_Authorization_Key.md)

[Mobile Protocol: Detailed Description (v.1.0, DEPRECATED)](doc/Mobile_Protocol-Detailed_Description_v.1.0_DEPRECATED.md)

[Encrypted CDNs for Speed and Security](doc/cdn.md) [@steedfly](https://github.com/steedfly)翻译

### Quick start with Docker

1. Install `docker` and `docker-compose`
2. Edit `docker-compose.yml`: replace `CHATENGINE_HOST` to your own host IP
3. Run `make` command in your shell.
4. Now, `Chatengine` is running on your host port `12345`.

#### Docker run demo

```shell
git clone https://github.com/PluralityNET/PluralityServer
cd PluralityServer
# replace 192.168.1.100 to you own host IP.
sed -i "" 's/CHATENGINE_HOST=127.0.0.1/CHATENGINE_HOST=192.168.1.100/g' docker-compose.yml # macOS
# sed -i 's/CHATENGINE_HOST=127.0.0.1/CHATENGINE_HOST=192.168.1.100/g' docker-compose.yml # linux
make
```

### Manual Build and Install

Note: You will probably need a VM for this as the code often uses root at MySQL and root for filesystem write access.

#### Depends

- redis
- mysql
- etcd

#### Build

- Get source code　

```
mkdir -p $GOPATH/src/github.com/PluralityNET/
cd $GOPATH/src/github.com/PluralityNET/
git clone https://github.com/PluralityNET/PluralityServer.git
```

- Build
    ```
    build frontend
        cd $GOPATH/src/github.com/PluralityNET/PluralityServer/access/frontend
        go build
    
    build auth_key
        cd $GOPATH/src/github.com/PluralityNET/PluralityServer/access/auth_key
        go build

    build auth_session
        cd $GOPATH/src/github.com/PluralityNET/PluralityServer/service/auth_session
        go build
        
    build sync
        cd $GOPATH/src/github.com/PluralityNET/PluralityServer/messenger/sync
        go build
    
    build upload
        cd $GOPATH/src/github.com/PluralityNET/PluralityServer/messenger/upload
        go build
    
    build document
        cd $GOPATH/src/github.com/PluralityNET/PluralityServer/service/document
        go build

    build biz_server
        cd $GOPATH/src/github.com/PluralityNET/PluralityServer/messenger/biz_server
        go build
        
    build session
        cd $GOPATH/src/github.com/PluralityNET/PluralityServer/access/session
        go build
    ```

- Init
    - configure mysql passwordless login for OS user `root` for mysql user `root@localhost`;
    
    - then,
    
    ```shell
    cd $GOPATH/src/github.com/PluralityNET/PluralityServer/scripts/
    mysql -u root
        CREATE DATABASE PluralityServer;
        use PluralityServer;
        source PluralityServer.sql
        source merge_20181129_201906.sql
        exit;
    ```
- Run
    ```shell
    cd $GOPATH/src/github.com/PluralityNET/PluralityServer/service/auth_session
    ./auth_session
    
    cd $GOPATH/src/github.com/PluralityNET/PluralityServer/service/document
    ./document

    cd $GOPATH/src/github.com/PluralityNET/PluralityServer/messenger/sync
    ./sync
    
    cd $GOPATH/src/github.com/PluralityNET/PluralityServer/messenger/upload
    ./upload

    cd $GOPATH/src/github.com/PluralityNET/PluralityServer/messenger/biz_server
    ./biz_server

    cd $GOPATH/src/github.com/PluralityNET/PluralityServer/access/auth_key
    ./auth_key

    cd $GOPATH/src/github.com/PluralityNET/PluralityServer/access/session
    ./session
    
    cd $GOPATH/src/github.com/PluralityNET/PluralityServer/access/frontend
    ./frontend
    ```

#### More

[Build document](doc/build.md)

[Build script](scripts/build.sh)

[Prerequisite script](scripts/prerequisite.sh)


### Compatible clients

**Important**: default signIn and signOut verify code is **12345**

[Android client for NebulaChat](https://github.com/nebula-chat/clients/tree/master/Telegram-Android)

[FOSS client for NebulaChat](https://github.com/nebula-chat/clients/tree/master/Telegram-FOSS)

[iOS client for NebulaChat](https://github.com/nebula-chat/clients/tree/master/Telegram-iOS)

[tdesktop for NebulaChat](https://github.com/nebula-chat/clients/tree/master/tdesktop)


## Feedback

PluralityNET's chat: `irc.ilita.i2p` `#plurality`

Nebula Chat's chats:

 * English: https://t.me/entelegramd
 * Chinese: https://t.me/cntelegramd
