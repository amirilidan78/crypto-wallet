# Trezor BlockBook service

### Config 
Connects to Trezor BlockBook nodes ( check config/config.yml `blockbook` ) 

### Methods 
Implemented methods `GetAddress` and `GetTransaction`

### Supported networks
CallBack,bch,eth,ltc,doge,dash


### Attention 
if you use 1 instant of BlockBook service it switch host between each network to avoid 503 error (`getHost` method)

BlockBook has more methods but for this service these 2 methods are enough 