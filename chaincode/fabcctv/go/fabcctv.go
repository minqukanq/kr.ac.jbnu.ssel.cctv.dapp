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
 * The fabcctv smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the frame structure, with 4 properties.  Structure tags are used by encoding/json library
type Frame struct {
	Time  string `json:"time"`
	Data  string `json:"data"`
}

/*
 * The Init method is called when the Smart Contract "fabcctv" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcctv"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryFrame" {
		return s.queryFrame(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createFrame" {
		return s.createFrame(APIstub, args)
	} else if function == "queryFramesByRange" {
		return s.queryFramesByRange(APIstub, args)
	} else if function == "validateFrame" {
		return s.validateFrame(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryFrame(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	frameAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(frameAsBytes)
}

func (s *SmartContract) validateFrame(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	// TODO: change the return as simple value like true or false

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	frameAsBytes, _ := APIstub.GetState(args[0])

	if string(frameAsBytes) != args[1] {
		return shim.Success(frameAsBytes)
	}

	return shim.Success(nil)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	frames := []Frame{
		Frame{Time: "199502221420", Data: "Hash#1"},
		Frame{Time: "199502221421", Data: "Hash#2"},
		Frame{Time: "199502221422", Data: "Hash#3"},
		Frame{Time: "199502221423", Data: "Hash#4"},
		Frame{Time: "199502221424", Data: "Hash#5"},
		Frame{Time: "199502221425", Data: "Hash#6"},
	}

	i := 0
	for i < len(frames) {
		fmt.Println("i is ", i)
		frameAsBytes, _ := json.Marshal(frames[i])
		APIstub.PutState("FRAME"+strconv.Itoa(i), frameAsBytes)
		fmt.Println("Added", frames[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createFrame(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	var frame = Frame{Time: args[1], Data: args[2]}

	frameAsBytes, _ := json.Marshal(frame)
	APIstub.PutState(args[0], frameAsBytes)

	return shim.Success(frameAsBytes)
}

func (s *SmartContract) queryFramesByRange(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	startKey := "000000000000" // this value should not being changed
	endKey := args[0] // TODO: value should be defined

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

	fmt.Printf("- queryFramesByRange:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
