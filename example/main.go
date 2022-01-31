package main

import (
	"fmt"
	"math/big"

	"github.com/sledro/levm"

	"github.com/sledro/levm/tools"
)

var (
	fromAddr = tools.NewRandomAddress()
)

func main() {

	//make a new address evm
	fromAddr = tools.NewRandomAddress()

	//Load a contract from file
	abiObject, binData, err := tools.LoadContract("./contract/example_sol_Example.abi", "./contract/example_sol_Example.bin")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Abi Methods: ", abiObject.Methods)

	// create new LEVM instance
	lvm := levm.New("./db", big.NewInt(8000000), fromAddr)

	// create a new account and set the balance
	// (needs enough balance to cover gas cost)
	lvm.NewAccount(fromAddr, big.NewInt(5000000000000))

	// deploy a contract
	code, addr, gas, err := lvm.DeployContract(fromAddr, binData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("contract code length:", len(code))
	fmt.Printf("contract address: %x\n", addr)
	fmt.Println("unused gas:", gas)
	fmt.Println("address:", lvm.GetAccount(fromAddr))

	// call a contract: set
	setOutput, err := lvm.CallContractABI(fromAddr, addr, big.NewInt(0), abiObject, "SetValA", big.NewInt(1))
	if err != nil {
		fmt.Println("set error : ", err)
	}
	fmt.Println("set output:", setOutput)

	// call a contract: get
	getOutput, err := lvm.CallContractABI(fromAddr, addr, big.NewInt(0), abiObject, "vala")
	if err != nil {
		fmt.Println("get error : ", err)
	}
	fmt.Println("get output:", getOutput)
}
