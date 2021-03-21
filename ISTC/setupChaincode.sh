
echo "Installing istc chaincode to peer0.wholesaler.istc.com..."

# install chaincode
# Install code on teamaster peer
docker exec -e "CORE_PEER_LOCALMSPID=TeaMasterMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/teamaster.istc.com/users/Admin@teamaster.istc.com/msp" -e "CORE_PEER_ADDRESS=peer0.teamaster.istc.com:7051" cli peer chaincode install -n istcc -v 1.0 -p github.com/istc/go -l golang

echo "Installed istc chaincode to peer0.teamaster.istc.com"


# Install code on wholesaler peer
docker exec -e "CORE_PEER_LOCALMSPID=WholesalerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/wholesaler.istc.com/users/Admin@wholesaler.istc.com/msp" -e "CORE_PEER_ADDRESS=peer0.wholesaler.istc.com:7051" cli peer chaincode install -n istcc -v 1.0 -p github.com/istc/go -l golang

echo "Installed istc chaincode to peer0.wholesaler.istc.com"

echo "Installing istc chaincode to peer0.teamaster.istc.com...."


echo "Installing istc chaincode to peer0.istc.istc.com..."
# Install code on istc peer
docker exec -e "CORE_PEER_LOCALMSPID=ISTCMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/istc.istc.com/users/Admin@istc.istc.com/msp" -e "CORE_PEER_ADDRESS=peer0.istc.istc.com:7051" cli peer chaincode install -n istcc -v 1.0 -p github.com/istc/go -l golang

# install chaincode
# Install code on teataster peer
docker exec -e "CORE_PEER_LOCALMSPID=TeaTasterMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/teataster.istc.com/users/Admin@teataster.istc.com/msp" -e "CORE_PEER_ADDRESS=peer0.teataster.istc.com:7051" cli peer chaincode install -n istcc -v 1.0 -p github.com/istc/go -l golang

echo "Installed istc chaincode to peer0.teataster.istc.com"

# install chaincode
# Install code on lab peer
docker exec -e "CORE_PEER_LOCALMSPID=LabMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/lab.istc.com/users/Admin@lab.istc.com/msp" -e "CORE_PEER_ADDRESS=peer0.lab.istc.com:7051" cli peer chaincode install -n istcc -v 1.0 -p github.com/istc/go -l golang

echo "Installed istc chaincode to peer0.lab.istc.com"

# install chaincode
# Install code on teaboardofindia peer
docker exec -e "CORE_PEER_LOCALMSPID=TeaBoardOfIndiaMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/teaboardofindia.istc.com/users/Admin@teaboardofindia.istc.com/msp" -e "CORE_PEER_ADDRESS=peer0.teaboardofindia.istc.com:7051" cli peer chaincode install -n istcc -v 1.0 -p github.com/istc/go -l golang

echo "Installed istc chaincode to peer0.teaboardofindia.istc.com"



echo "Installed istc chaincode to peer0.teamaster.istc.com"

echo "Instantiating istc chaincode.."

docker exec -e "CORE_PEER_LOCALMSPID=TeaMasterMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/teamaster.istc.com/users/Admin@teamaster.istc.com/msp" -e "CORE_PEER_ADDRESS=peer0.teamaster.istc.com:7051" cli peer chaincode instantiate -o orderer.istc.com:7050 -C istcchannel -n istcc -l golang -v 1.0 -c '{"Args":[""]}' -P "OR ('WholesalerMSP.member','TeaMasterMSP.member','ISTCMSP.member')"

echo "Instantiated istc chaincode."

echo "Following is the docker network....."

docker ps
