/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Trade Finance Use Case - WORK IN  PROGRESS
 */

package main


import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SpecialtyTeaContract struct {
}


// Define the letter of credit
type TeaLot struct {
	TeaLotId		 string		`json:"teaLotId"`
	TeaVariant 		 string		`json:"teaVariant"`
	Origin           string     `json:"origin"`
	Season 		     string		`json:"season"`
	TeaEstate 		 string		`json:"teaEstate"`
	TeaMaster 		 string		`json:"teaMaster"`
	Make 			 string		`json:"make"`
	LeafType         string		`json:"leafType"`
	MadeTeaName  	 string		`json:"madeTeaName"`
	SoilType         string		`json:"soilType"`
	Characteristics  string     `json:"characteristics"`
	TeaTasteNotes    string     `json:"teaTasteNotes"`
	HealthAttributes string     `json:"healthAttributes"`
	Certified        string     `json:"certified"`
	Award            string     `json:"award"`
	LotNumber        string     `json:"lotNumber"`
	TeaTaster        string     `json:"teaTaster"`
	LotStatus        string     `json:"lotStatus"`
	LabStatus        string     `json:"labStatus"`
	TastingStatus    string     `json:"tastingStatus"`
	LabName          string     `json:"labName"`
}

type TeaPacket struct {
	TeaPacketId	         string		`json:"teaPacketId"`
	TeaLotId		     string		`json:"teaLotId"`
	Quantity             int        `json:"quantity"`
	UOM                  string     `json:"uom"`
	Status 		         string		`json:"status"`
	BuyerName 		     string		`json:"buyerName"`
	FreightNo            string	    `json:"freightNo"`
	LogisticsPartnerName string     `json:"logisticsPartnerName"`
	IndianTarrifCode     string     `json:"indianTarrifCode"`
}

func (s *SpecialtyTeaContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SpecialtyTeaContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "initLedger" {
		return s.initLedger(APIstub)
	}else if function == "createTeaLot" {
		return s.createTeaLot(APIstub, args)
	}else if function == "createTeaPacket" {
		return s.createTeaPacket(APIstub, args)
	}else if function == "updateTeaLot" {
		return s.updateTeaLot(APIstub, args)
	}else if function == "updateTeaPacket" {
		return s.updateTeaPacket(APIstub, args)
	}else if function == "initiateLabTest" {
		return s.initiateLabTest(APIstub, args)
	}else if function == "updateLabResult" {
		return s.updateLabResult(APIstub, args)
	}else if function == "initiateTeaTaste" {
		return s.initiateTeaTaste(APIstub, args)
	}else if function == "updateTeaTaste" {
		return s.updateTeaTaste(APIstub, args)
	}else if function == "getTeaLot" {
		return s.getTeaLot(APIstub, args)
	}else if function == "getTeaLotHistory" {
		return s.getTeaLotHistory(APIstub, args)
	}else if function == "queryAllTeaLots" {
		return s.queryAllTeaLots(APIstub)
	}else if function == "queryAllTeaPackets" {
		return s.queryAllTeaPackets(APIstub)
	}else if function == "getTeaPacket" {
		return s.getTeaPacket(APIstub, args)
	}else if function == "getTeaPacketHistory" {
		return s.getTeaPacketHistory(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SpecialtyTeaContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
    // Tea Lots
	teaLots := []TeaLot{
		TeaLot{TeaLotId: "1",TeaVariant: "PurpleTea",Origin: "Karbi Anglong",Season: "Spring,First Flush",TeaEstate: "Barpathar Tea Estate",TeaMaster: "P.S",Make: "Wiry",LeafType: "Two Leaf and Bud",MadeTeaName: "WildPurple",SoilType: "Clay Soil",Characteristics: "Strong",TeaTasteNotes: "Fruity, Deep Red Purple Hue",HealthAttributes: "CardioVascular; UV Rays Protection",Certified: "No",Award: "",LotNumber: "L001",TeaTaster: "HS", LotStatus: "ReadyToSale",TastingStatus: "Completed"},
		TeaLot{TeaLotId: "2",TeaVariant: "PurpleTea",Origin: "Karbi Anglong",Season: "Monsoon,Second Flush",TeaEstate: "Barpathar Tea Estate",TeaMaster: "P.S",Make: "Boutique",LeafType: "Two Leaf and Bud",MadeTeaName: "Moonsoon Rainbow",SoilType: "Clay Soil",Characteristics: "Strong",TeaTasteNotes: "Fruity",HealthAttributes: "CardioVascular; UV Rays Protection",Certified: "No",Award: "",LotNumber: "L002",TeaTaster: "HS", LotStatus: "ReadyToSale", TastingStatus: "Initiated"},
		TeaLot{TeaLotId: "3",TeaVariant: "PurpleTea",Origin: "Karbi Anglong",Season: "Autumn,Diwali",TeaEstate: "Barpathar Tea Estate",TeaMaster: "P.S",Make: "Wiry",LeafType: "Two Leaf and Bud",MadeTeaName: "Kartik Purnima Purple",SoilType: "Clay Soil",Characteristics: "Mild",TeaTasteNotes: "Sweet Purple",HealthAttributes: "CardioVascular; UV Rays Protection",Certified: "Yes",Award: "",LotNumber: "L003",TeaTaster: "HS", LotStatus: "Sold", TastingStatus: "Completed"},
		TeaLot{TeaLotId: "4",TeaVariant: "PurpleTea",Origin: "Karbi Anglong",Season: "Winter,First Flush",TeaEstate: "Barpathar Tea Estate",TeaMaster: "P.S",Make: "Shade Dry",LeafType: "Bud",MadeTeaName: "Velvet Smooth",SoilType: "Clay Soil",Characteristics: "Light",TeaTasteNotes: "Flowery",HealthAttributes: "Bone Health, Teeth Health, Collagen Building",Certified: "Yes",Award: "King of White Tea Title",LotNumber: "L004",TeaTaster: "HS", LotStatus: "ReadyToPack", TastingStatus: "Initiated"},
		}
	i := 0

	for i < len(teaLots) {
		fmt.Println("i is ", i)
		teaLotAsBytes, _ := json.Marshal(teaLots[i])
		APIstub.PutState(teaLots[i].TeaLotId, teaLotAsBytes)
		fmt.Println("Added", teaLots[i])
	}

    // Tea Packets
	teaPackets := []TeaPacket{
		TeaPacket{TeaPacketId: "1001",TeaLotId: "1",Quantity: 35,UOM: "gram",Status: "Ready",BuyerName: "Luke",FreightNo: "Requested",LogisticsPartnerName: "FedEx",IndianTarrifCode: "ITC09202120"},
		TeaPacket{TeaPacketId: "1002",TeaLotId: "1",Quantity: 35,UOM: "gram",Status: "Sold",BuyerName: "T2",FreightNo: "1263653",LogisticsPartnerName: "FedEx",IndianTarrifCode: "ITC09202120"},
 		TeaPacket{TeaPacketId: "1003",TeaLotId: "2",Quantity: 100,UOM: "gram",Status: "Sold",BuyerName: "abc",FreightNo: "1263789",LogisticsPartnerName: "FedEx",IndianTarrifCode: "ITC09202121"},
 		TeaPacket{TeaPacketId: "1004",TeaLotId: "3",Quantity: 100,UOM: "gram",Status: "Ready",BuyerName: "T2",FreightNo: "Requested",LogisticsPartnerName: "FedEx",IndianTarrifCode: "ITC09202122"},
		}
 		
	i = 0

	for i < len(teaPackets) {
		fmt.Println("i is ", i)
		tpAsBytes, _ := json.Marshal(teaPackets[i])
		APIstub.PutState(teaPackets[i].TeaPacketId, tpAsBytes)
	fmt.Println("Added", teaPackets[i])
	i = i + 1
	}

return shim.Success(nil)
}

// This function is initiate by producer 
func (s *SpecialtyTeaContract) createTeaLot(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	teaLot := TeaLot{}

	err  := json.Unmarshal([]byte(args[0]),&teaLot)
    if err != nil {
		return shim.Error("Not able to parse args into tea lot")
	}
	teaLotAsBytes, err := json.Marshal(teaLot)
    APIstub.PutState(teaLot.TeaLotId,teaLotAsBytes)
	fmt.Println("created Tea -> ", teaLot)

	return shim.Success(nil)
}

// This function is initiate by producer
func (s *SpecialtyTeaContract) createTeaPacket(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	tp := TeaPacket{}

	err  := json.Unmarshal([]byte(args[0]),&tp)
    if err != nil {
		return shim.Error("Not able to parse args into tea")
	}
	tpAsBytes, err := json.Marshal(tp)
    APIstub.PutState(tp.TeaPacketId,tpAsBytes)
	fmt.Println("created Tea Packet -> ", tp)

	return shim.Success(nil)
}

func (s *SpecialtyTeaContract) updateTeaLot(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	teaLot := TeaLot{}

	err  := json.Unmarshal([]byte(args[0]),&teaLot)
    if err != nil {
		return shim.Error("Not able to parse args into tea")
	}
	teaLotAsBytes, err := json.Marshal(teaLot)
    APIstub.PutState(teaLot.TeaLotId,teaLotAsBytes)
	fmt.Println("updated Tea Lot -> ", teaLot)

	return shim.Success(nil)
}

func (s *SpecialtyTeaContract) updateTeaPacket(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	tp := TeaPacket{}

	err  := json.Unmarshal([]byte(args[0]),&tp)
    if err != nil {
		return shim.Error("Not able to parse args into tea")
	}
	tpAsBytes, err := json.Marshal(tp)
    APIstub.PutState(tp.TeaPacketId,tpAsBytes)
	fmt.Println("created Tea Packet -> ", tp)

	return shim.Success(nil)
}

//update lab name by producer, Lab
func (s *SpecialtyTeaContract) initiateLabTest(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting only two arguments")
	}

	teaLotAsBytes, _ := APIstub.GetState(args[0])

	teaLot := TeaLot{}
	json.Unmarshal(teaLotAsBytes, &teaLot)
	teaLot.LabStatus = "Initiated"
	teaLot.LabName = args[1]
	teaLotAsBytes, _ = json.Marshal(teaLot)
    APIstub.PutState(teaLot.TeaLotId,teaLotAsBytes)
	fmt.Println("Update Tea taste-> ", teaLot)

	return shim.Success(nil)
}

func (s *SpecialtyTeaContract) updateLabResult(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting only two arguments")
	}

	teaLotAsBytes, _ := APIstub.GetState(args[0])

	teaLot := TeaLot{}
	json.Unmarshal(teaLotAsBytes, &teaLot)
	teaLot.LabStatus = "Completed"
	teaLot.Certified = args[1]
	teaLotAsBytes, _ = json.Marshal(teaLot)
    APIstub.PutState(teaLot.TeaLotId,teaLotAsBytes)
	fmt.Println("Update Tea taste-> ", teaLot)

	return shim.Success(nil)
}


func (s *SpecialtyTeaContract) initiateTeaTaste(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting only two args")
	}

	teaLotAsBytes, _ := APIstub.GetState(args[0])

	teaLot := TeaLot{}
	json.Unmarshal(teaLotAsBytes, &teaLot)
	teaLot.TastingStatus = "Initiated"
	teaLot.TeaTaster = args[1]
	teaLotAsBytes, _ = json.Marshal(teaLot)
    APIstub.PutState(teaLot.TeaLotId,teaLotAsBytes)
	fmt.Println("Update Tea taste-> ", teaLot)

	return shim.Success(nil)
}


//update tea taste notes by producer, tea taster
func (s *SpecialtyTeaContract) updateTeaTaste(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting only two arguments")
	}

	teaLotAsBytes, _ := APIstub.GetState(args[0])

	teaLot := TeaLot{}
	json.Unmarshal(teaLotAsBytes, &teaLot)
	teaLot.TastingStatus = "Completed"
	teaLot.TeaTasteNotes = args[1]
	teaLotAsBytes, _ = json.Marshal(teaLot)
    APIstub.PutState(teaLot.TeaLotId,teaLotAsBytes)
	fmt.Println("Update Tea taste-> ", teaLot)

	return shim.Success(nil)
}

func (s *SpecialtyTeaContract) getTeaLot(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	teaLotId := args[0];
	
	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	teaLotAsBytes, _ := APIstub.GetState(teaLotId)

	return shim.Success(teaLotAsBytes)
}

func (s *SpecialtyTeaContract) getTeaLotHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	teaTypeId := args[0];
	
	

	resultsIterator, err := APIstub.GetHistoryForKey(teaTypeId)
	if err != nil {
		return shim.Error("Error retrieving Tea history.")
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("Error retrieving Tea history.")
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getTeaHistory returning:\n%s\n", buffer.String())

	

	return shim.Success(buffer.Bytes())
}

func (s *SpecialtyTeaContract) getTeaPacket(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	teaPacketId := args[0];
	
	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	teaPacketAsBytes, _ := APIstub.GetState(teaPacketId)

	return shim.Success(teaPacketAsBytes)
}

func (s *SpecialtyTeaContract) getTeaPacketHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	teaTypeId := args[0];
	
	

	resultsIterator, err := APIstub.GetHistoryForKey(teaTypeId)
	if err != nil {
		return shim.Error("Error retrieving Tea packet history.")
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("Error retrieving Tea history.")
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getTeaPacketHistory returning:\n%s\n", buffer.String())

	

	return shim.Success(buffer.Bytes())
}

func (s *SpecialtyTeaContract) queryAllTeaLots(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "TL001"
	endKey := "TL999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllTeas:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SpecialtyTeaContract) queryAllTeaPackets(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "TP001"
	endKey := "TP999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllTeas:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SpecialtyTeaContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
