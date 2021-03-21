rm -R crypto-config/*

./bin/cryptogen generate --config=crypto-config.yaml

rm config/*

./bin/configtxgen -profile ISTCOrgOrdererGenesis -outputBlock ./config/genesis.block

./bin/configtxgen -profile ISTCOrgChannel -outputCreateChannelTx ./config/istcchannel.tx -channelID istcchannel
