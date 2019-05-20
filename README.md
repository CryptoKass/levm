# LEVM - Little Ethereum Virtual Machine

Run smart contracts on the Ethereum Virtual Machine (EVM) without the Ethereum blockchain. 

I created this to make it a little easier to use the EVM in other projects, and for testing smart contracts.

`levm.New(...)` creates a small wrapper for the ethereum virtual machine, and sets up the EVM with default parameters and the required stateDB.

# ↓ Installation

This project only has two dependencies. (go-ethereum contains the EVM, and splashkeys is an ecdsa key library)

## Dependencies
- `go get github.com/ethereum/go-ethereum`
- `go get github.com/CryptoKass/splashkeys`

## Install
- `go get github.com/CryptoKass/levm`


# ⏰ Quick start

The Example is really short and it will be your best bet for getting started.

- Install `go get github.com/CryptoKass/levm && go get github.com/CryptoKass/splashkeys && go get github.com/ethereum/go-ethereum` 
- Then See Example at bottom of README.md

# 🔨 Usage 

**Basic Usage:**
- Create a new LEVM instance: 
    - `levm.New(dbPath string, blockNumber *big.Int, origin common.Address)`
    - e.g. `lvm := levm.New("./db", big.NewInt(0), fromAddr)`
- Load Contract Bin and ABi:
    -  `tools.LoadContract(abiPath, binPath string)`
    - e.g. `abiObject, binData, err := tools.LoadContract("contract/example_sol_Example.abi", "contract/example_sol_Example.bin")`
- Deploy a contract:
    - `lvm.DeployContract(fromAddr common.Address, contractData []byte)`
    - e.g. `code, addr, gas, err := lvm.DeployContract(fromAddr, binData)`
- Call a contract:
    - `levm.CallContract(callerAddr, contractAddr common.Address, value *big.Int, inputs []byte)`
    - e.g. `output, err := lvm.CallContract(fromAddr, addr, big.NewInt(0), inputs)`
    - Inputs can be created using the ABI.Pack method using a loaded ABI (see contract loading above).


**I reccommend you use the CallContractABI method to call contract methods by name.**

- Call a contract using the ABI:
    `CallContractABI(callerAddr, contractAddr common.Address, value *big.Int, abiObject abi.ABI, funcName string, args ...interface{})`



# 📑 Example 
This is Example is located in the example sub-directory:

```go
func main() {

	//make a new address evm
	fromAddr = tools.NewRandomAddress()

	//Load a contract from file
	abiObject, binData, err := tools.LoadContract("contract/example_sol_Example.abi", "contract/example_sol_Example.bin")
	fmt.Println("Abi\n", abiObject.Methods)

	// create new LEVM instance
	lvm := levm.New("./db", big.NewInt(0), fromAddr)

	// create a new account and set the balance
	// (needs enough balance to cover gas cost)
	lvm.NewAccount(fromAddr, big.NewInt(1e18))

	// deploy a contract
	code, addr, gas, err := lvm.DeployContract(fromAddr, binData)
	fmt.Println("contract code length:", len(code))
	fmt.Printf("contract address: %x\n", addr)
	fmt.Println("unused gas:", gas)
	fmt.Println("errors:", err)

}
```

contact: kasscrypto@gmail.com 👍

**HELP WANTED** would be great if you would create an issue if you find a bug.