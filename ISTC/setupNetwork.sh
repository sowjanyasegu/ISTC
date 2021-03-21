echo "Setting up the network.."

echo "Creating channel genesis block.."

# Create the channel
docker exec -e "CORE_PEER_LOCALMSPID=ISTCMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/istc.istc.com/users/Admin@istc.istc.com/msp" -e "CORE_PEER_ADDRESS=peer0.istc.istc.com:7051" cli peer channel create -o orderer.istc.com:7050 -c istcchannel -f /etc/hyperledger/configtx/istcchannel.tx


sleep 5

echo "Channel genesis block created."

echo "peer0.istc.istc.com joining the channel..."
# Join peer0.istc.istc.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=ISTCMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/istc.istc.com/users/Admin@istc.istc.com/msp" -e "CORE_PEER_ADDRESS=peer0.istc.istc.com:7051" cli peer channel join -b istcchannel.block

echo "peer0.istc.istc.com joined the channel"

echo "peer0.teamaster.istc.com joining the channel..."

# Join peer0.teamaster.istc.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=TeaMasterMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/teamaster.istc.com/users/Admin@teamaster.istc.com/msp" -e "CORE_PEER_ADDRESS=peer0.teamaster.istc.com:7051" cli peer channel join -b istcchannel.block

echo "peer0.teamaster.istc.com joined the channel"

echo "peer0.wholesaler.istc.com joining the channel..."
# Join peer0.wholesaler.istc.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=WholesalerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/wholesaler.istc.com/users/Admin@wholesaler.istc.com/msp" -e "CORE_PEER_ADDRESS=peer0.wholesaler.istc.com:7051" cli peer channel join -b istcchannel.block


echo "peer0.wholesaler.istc.com joined the channel"

echo "peer0.teataster.istc.com joining the channel..."
# Join peer0.teataster.istc.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=TeaTasterMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/teataster.istc.com/users/Admin@teataster.istc.com/msp" -e "CORE_PEER_ADDRESS=peer0.teataster.istc.com:7051" cli peer channel join -b istcchannel.block


echo "peer0.teataster.istc.com joined the channel"

echo "peer0.lab.istc.com joining the channel..."
# Join peer0.lab.istc.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=LabMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/lab.istc.com/users/Admin@lab.istc.com/msp" -e "CORE_PEER_ADDRESS=peer0.lab.istc.com:7051" cli peer channel join -b istcchannel.block


echo "peer0.lab.istc.com joined the channel"

echo "peer0.teaboardofindia.istc.com joining the channel..."
# Join peer0.teaboardofindia.istc.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=TeaBoardOfIndiaMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/teaboardofindia.istc.com/users/Admin@teaboardofindia.istc.com/msp" -e "CORE_PEER_ADDRESS=peer0.teaboardofindia.istc.com:7051" cli peer channel join -b istcchannel.block


echo "peer0.teaboardofindia.istc.com joined the channel"



echo "Following is the docker network....."

docker ps
