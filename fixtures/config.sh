export FABRIC_CFG_PATH=$PWD

rm -rf artifacts/*
rm -rf crypto-config/

./bin/cryptogen generate --config=./crypto-config.yaml

./bin/configtxgen -profile FourOrgsOrdererGenesis -outputBlock ./artifacts/orderer.genesis.block

./bin/configtxgen -profile FourOrgsChannel -outputCreateChannelTx ./artifacts/privateledger.channel.tx -channelID privateledger

./bin/configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./artifacts/Org1MSPanchors.tx -channelID privateledger -asOrg Org1MSP

./bin/configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./artifacts/Org2MSPanchors.tx -channelID privateledger -asOrg Org2MSP

./bin/configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./artifacts/Org3MSPanchors.tx -channelID privateledger -asOrg Org3MSP

./bin/configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./artifacts/Org4MSPanchors.tx -channelID privateledger -asOrg Org4MSP
