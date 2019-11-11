package main

import (
    "os"

    "github.com/hyperledger/fabric-chaincode-go/shim"
    "github.com/hyperledger/fabric-protos-go/peer"
)

type KVStore struct {}

func (kv *KVStore) Init(stub shim.ChaincodeStubInterface) peer.Response {
    return shim.Success(nil)
}

func (kv *KVStore) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    return shim.Success(nil)
}

func main() {
    err := shim.Start(&KVStore{})
    if err != nil {
        panic(err)
    }

    os.Exit(0)
}