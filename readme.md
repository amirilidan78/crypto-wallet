# Crypto wallet 
simple crypto wallet repository for generating wallets and signing transactions with private key 

### Packages and services 
- Mysql  https://www.mysql.com/
- Golang https://go.dev/
- Trust wallet core https://developer.trustwallet.com/wallet-core

### Port bindings  
- Database 50001:3306
- Http api 50002:50001

### Quick start 
- `docker-compose up -d `
- `docker exec -it api-crypto-wallet-container go run ./app/migrate.go`
- send post request `localhost:50002/api/v1/address/new` for create or getting wallet

### Sample 

![post](./samples/post.png)

![db](./samples/db.png)

### Supported coins 
- AE
- AION
- ALGO
- BNB
- BTC
- BCH
- BTG
- CLO
- ADA
- ATOM
- DASH
- DCR
- DGB
- DOGE
- EOS
- ETH
- ETC
- FIO
- GO
- GRS
- ICX
- IOTX
- Kava
- KIN
- LTC
- MONA
- NAS
- NULS
- NANO
- NEAR
- NIM
- ONT
- POA
- QTUM
- XRP
- SOL
- XLM
- TON
- XTZ
- THETA
- TT
- NEO
- TOMO
- TRX
- VET
- VIA
- WAN
- ZEC
- XZC
- ZIL
- ZEL
- RVN
- WAVES
- LUNA
- ONE
- KSM
- IOU
- FIL
- ERD

*Add coins in `pkg/twallet/hdWallet.go` from twallet package*


### TODOs 
- <del>Generate wallet</del>
- Sign transaction