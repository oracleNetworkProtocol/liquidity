# !/bin/bash
# make install
rm -rf home1
NODENAME1="home1"

CHAIN_ID=chain-a
PASSWORD="123123123"
DENOM=plug
liquidityd init $NODENAME1 --chain-id $CHAIN_ID --home $NODENAME1

liquidityd keys add mywallet  --home $NODENAME1 
liquidityd keys add mywallet1  --home $NODENAME1 

sed -i ''  's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' "$(pwd)/"$NODENAME1/config/config.toml
sed -i ''  "s/\"stake\"/\"$DENOM\"/g" "$(pwd)/"$NODENAME1/config/genesis.json
sed -i ''  's/pruning = "syncable"/pruning = "nothing"/g' "$(pwd)/"$NODENAME1/config/app.toml

liquidityd add-genesis-account $(liquidityd keys show mywallet -a --home $NODENAME1)  "10100000000plug,10000000000onp" --home $NODENAME1 &&\
liquidityd add-genesis-account $(liquidityd keys show mywallet1 -a --home $NODENAME1) "10000000000plug,10000000000onp" --home $NODENAME1

liquidityd gentx mywallet 1000000plug --keyring-dir $NODENAME1 --home $NODENAME1 --chain-id $CHAIN_ID

liquidityd collect-gentxs --home $NODENAME1