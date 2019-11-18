module github.com/guoger/mycontract

go 1.13

require (
	github.com/hyperledger/fabric-chaincode-go v0.0.0-00010101000000-000000000000
	github.com/hyperledger/fabric-protos-go v0.0.0-20191108205341-c09017a7950b
	google.golang.org/grpc v1.23.0
)

replace github.com/hyperledger/fabric-chaincode-go => github.com/guoger/fabric-chaincode-go v0.0.0-20191115025301-293eed878c0f
