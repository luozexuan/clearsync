// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package session_key_validator

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

// SessionKeyValidatorMetaData contains all meta data concerning the SessionKeyValidator contract.
var SessionKeyValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NotImplemented\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"disable\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"enable\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"permissionKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"kernel\",\"type\":\"address\"}],\"name\":\"executionStatus\",\"outputs\":[{\"internalType\":\"ValidAfter\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"runs\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"}],\"name\":\"invalidateNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"kernel\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"lastNonce\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"invalidNonce\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"kernel\",\"type\":\"address\"}],\"name\":\"sessionData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"ValidAfter\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"ValidUntil\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"validCaller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"validateSignature\",\"outputs\":[{\"internalType\":\"ValidationData\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"ValidationData\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001657611ec0908161001b8239f35b5f80fdfe608060409080825260049081361015610016575f80fd5b5f92833560e01c9182630c9595561461049a57508163333daf92146104485781633a871cdd146103e557816346585db21461036157816352721fdd146102b75781637ecebe001461023c578382638fc925aa14610167575081639ea9bd59146100e1575063dbba225d14610088575f80fd5b346100dd5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100dd57356fffffffffffffffffffffffffffffffff811681036100dd576100da9061086c565b80f35b5080fd5b9050823461016457817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101645761011a610750565b506024359067ffffffffffffffff8211610164575061013c90369084016106fb565b5050517fd6234725000000000000000000000000000000000000000000000000000000008152fd5b80fd5b80918460207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101eb5780359067ffffffffffffffff8211610237576101b2913691016106fb565b909190601481036101ef576014116101eb576002913560601c835260016020528083203384526020528220828155826001820155015580f35b5050fd5b9192506010820361020f57506010116100dd576100da903560801c61086c565b6100da92506fffffffffffffffffffffffffffffffff9150338452836020528320541661086c565b505050fd5b8390346100dd5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100dd57809173ffffffffffffffffffffffffffffffffffffffff61028c610750565b1681528060205220548151906fffffffffffffffffffffffffffffffff8116825260801c6020820152f35b8390346100dd57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100dd578060a0926102f3610750565b6102fb61072d565b9073ffffffffffffffffffffffffffffffffffffffff809116835260016020528383209116825260205220805491600182015490600265ffffffffffff93015492815194855280831660208601528260301c169084015260601c60608301526080820152f35b919050346103e157817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e15791819261039d61072d565b90358252600260205273ffffffffffffffffffffffffffffffffffffffff83832091168252602052205465ffffffffffff825191818116835260301c166020820152f35b8280fd5b9190507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc926060843601126101645781359367ffffffffffffffff85116100dd5761016090853603011261016457506020926104419101610939565b9051908152f35b9050823461016457817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610164576024359067ffffffffffffffff8211610164575061013c90369084016106fb565b9084925060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e15767ffffffffffffffff9184358381116106f3576104e890369087016106fb565b806014116106f757806034116106f75780603a116106f7578084116106f757806054116106f7576074116106f35760548101359360a08301908111838210176106c7579060029184526014810135835260208301603482013560d01c815284840190603a83013560d01c82526105ff60608601918785013560601c835260808701948986523560601c8a526001602052878a20338b52602052878a2096518755600187019365ffffffffffff809251167fffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000865416178555511683907fffffffffffffffffffffffffffffffffffffffff000000000000ffffffffffff6bffffffffffff00000000000083549260301b169116179055565b516bffffffffffffffffffffffff7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000083549260601b1691161790555191015582602052822080546fffffffffffffffffffffffffffffffff80821681811461069b57917fffffffffffffffffffffffffffffffff00000000000000000000000000000000939160016100da9694011693849116179055146107e1565b6024866011897f4e487b7100000000000000000000000000000000000000000000000000000000835252fd5b6024866041897f4e487b7100000000000000000000000000000000000000000000000000000000835252fd5b8480fd5b8580fd5b9181601f840112156107295782359167ffffffffffffffff8311610729576020838186019501011161072957565b5f80fd5b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361072957565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361072957565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176107b457604052565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b156107e857565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f53657373696f6e4b657956616c696461746f723a20696e76616c6964206e6f6e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152fd5b6108a15f9133835282602052604083205460406fffffffffffffffffffffffffffffffff9485841694859360801c84116107e1565b3381528060205220937fffffffffffffffffffffffffffffffff0000000000000000000000000000000085549360801b1692168281178555106108e357505050565b179055565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610729570180359067ffffffffffffffff82116107295760200191813603831361072957565b610140810161094881836108e8565b601411610729573560601c905f908282526001916020948386526040808320338452875280832060028101548489528285205460801c1015610ca057858101978854928360601c8881145f14610be757506109a76101208601866108e8565b905015610b645750505b5415610b4457506109c560608201826108e8565b93909184600411806106f3577fffffffff00000000000000000000000000000000000000000000000000000000843516907f519454470000000000000000000000000000000000000000000000000000000082148015610b19575b15610a8a57505090610a31916108e8565b605593919311610164575090610a62929160758201358201926055607581860135950193818101350101918761195c565b919091610a83575065ffffffffffff610a80935460301c1691610d23565b90565b9250505090565b909195509392936103e1577f34fcd5be0000000000000000000000000000000000000000000000000000000003610b0f57610ac581856108e8565b6055116103e15790816055610add93013501946108e8565b605592919211610164575090816075610a629493013501916075605584013593019160756055830135920190876112f9565b5050505091505090565b5050847fb68df16d000000000000000000000000000000000000000000000000000000008214610a20565b949350505050610a80925065ffffffffffff808360301c16921690610d23565b6084925051907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152602660248201527f53657373696f6e4b657956616c696461746f723a207061796d6173746572206e60448201527f6f742073657400000000000000000000000000000000000000000000000000006064820152fd5b80610bf5575b5050506109b1565b610c036101208701876108e8565b601411610c9c573560601c03610c195780610bed565b6084925051907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152602760248201527f53657373696f6e4b657956616c696461746f723a207061796d6173746572206d60448201527f69736d61746368000000000000000000000000000000000000000000000000006064820152fd5b8780fd5b6084888351907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152602c60248201527f53657373696f6e4b657956616c696461746f723a2073657373696f6e206b657960448201527f206e6f7420656e61626c656400000000000000000000000000000000000000006064820152fd5b91909160043561014481013501602060243581525f917b19457468657265756d205369676e6564204d6573736167653a0a33328352603c6004209160409182516080810181811067ffffffffffffffff821117610ea8578452604180825283820192366079820111610c9c5790806038899301853781606184015260019386519783525186528580845114610e6a5750825114610e4f576001925060809150805b5afa51923d15610e42576060525273ffffffffffffffffffffffffffffffffffffffff809116911603610e3b5779ffffffffffff00000000000000000000000000000000000000007fffffffffffff00000000000000000000000000000000000000000000000000009160a01b169160d01b161790565b5050600190565b638baa579f90526004601cfd5b6001926080928660609182810151851a885201519052610dc4565b60019491507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff906080940151601b8160ff1c01875216606052610dc4565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610f025760010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b3573ffffffffffffffffffffffffffffffffffffffff811681036107295790565b15610f5757565b60846040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f53657373696f6e4b657956616c696461746f723a20746172676574206d69736d60448201527f61746368000000000000000000000000000000000000000000000000000000006064820152fd5b15610fe157565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f53657373696f6e4b657956616c696461746f723a2076616c7565206c696d697460448201527f20657863656564656400000000000000000000000000000000000000000000006064820152fd5b1561106c57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f53657373696f6e4b657956616c696461746f723a207065726d697373696f6e2060448201527f766572696669636174696f6e206661696c6564000000000000000000000000006064820152fd5b3563ffffffff811681036107295790565b359065ffffffffffff8216820361072957565b602091828252610140820190803563ffffffff811680910361072957848401528381013573ffffffffffffffffffffffffffffffffffffffff81168091036107295760409081850152808201357fffffffff000000000000000000000000000000000000000000000000000000008116809103610729576060908186015280830135608086015260808301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112156107295783019086823592019667ffffffffffffffff831161072957818302360388136107295792919081610120968760a08a015252610160870197935f905b83821061126a5750505050505065ffffffffffff908161122a60a08301611101565b1660c08501528161123d60c08301611101565b1660e085015261124f60e08201611101565b91610100921682850152013591600283101561072957015290565b90919293949889358152818a01359060068210156107295782810191909152898601358682015283019883019493929160010190611208565b90929167ffffffffffffffff84116107b4578360051b602092836040516112cc82850182610773565b809781520191810192831161072957905b8282106112ea5750505050565b813581529083019083016112dd565b9491959390929573ffffffffffffffffffffffffffffffffffffffff5f96165f52600160205260405f20335f5260205260405f209265ffffffffffff600185015416975f5b6004808801358801013581106113575750505050505050565b60048701358701600582901b8101602401359036037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7d0181121561072957828210156115de577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1853603018260051b8601351215610729576114636114546114689261141a8b60208760051b8b01358b01018373ffffffffffffffffffffffffffffffffffffffff61140883610f2f565b1615928315611613575b505050610f50565b61143f8b60448360608960051b8d01358d010135926004810135010101351115610fda565b60048b01358b010160648101906024016108e8565b8460051b880135880191611c76565b611065565b61150560028701546114c86114f46114878560051b89013589016110f0565b604051928391602083019586906024927fffffffff0000000000000000000000000000000000000000000000000000000091835260e01b1660208201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610773565b5190208260051b86013586016116be565b65ffffffffffff8b1665ffffffffffff82161161160b575b50828110156115de578060051b8501357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1863603018112156107295767ffffffffffffffff8187013511610729578086013560051b360360208288010113610729576115be9087546115b960405160208101906115a6816114c88960051b8d01358d0185611114565b5190209236908a018035906020016112a3565b611655565b156115d1576115cc90610ed5565b61133e565b5060019750505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b99505f61151d565b73ffffffffffffffffffffffffffffffffffffffff92935061164360248493836004611649950135010101610f2f565b93610f2f565b169116145f838e611412565b919091805180611666575b50501490565b91906020908180820191600595861b0101925b81518111851b9081528282519118528160405f209101938385101561169f579390611679565b50925050505f80611660565b3565ffffffffffff811681036107295790565b9065ffffffffffff60a0830160c08401826116d8826116ab565b1661183957506116e960e0916116ab565b935b0191816116f7846116ab565b16611703575b50505090565b5f9081526002602052604081203382526020526040812090600183835460301c16019083821161180c575081547fffffffffffffffffffffffffffffffffffffffff000000000000ffffffffffff1660309190911b6bffffffffffff0000000000001617815561177a9082905460301c16926116ab565b1610611788575f80806116fd565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f53657373696f6e4b657956616c696461746f723a2072756e732065786365656460448201527f65640000000000000000000000000000000000000000000000000000000000006064820152fd5b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526011600452fd5b939082611845826116ab565b16156118d8575f848152600260205260408120338252602052604081209184835497818916928315155f146118bf575061187e906116ab565b16019084821161180c5750908360e0939216955b7fffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000168685161790556116eb565b91505060e0949392506118d291506116ab565b95611892565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f53657373696f6e4b657956616c696461746f723a20696e76616c69642065786560448201527f637574696f6e2072756c650000000000000000000000000000000000000000006064820152fd5b9195939094925f955f9373ffffffffffffffffffffffffffffffffffffffff8091168552602094600186526040812033825286526040812099806004116100dd577f51945447000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008535161492878601816119f182610f2f565b1615908115611bf7575b50611a069150610f50565b8215611b4d5760441161016457611a4c92611463928592611a31606085013560248501351115610fda565b5015611b465760445b81013501602460048201359101611c76565b611aa0611adb6002890154611a60846110f0565b6040519384918883019384906024927fffffffff0000000000000000000000000000000000000000000000000000000091835260e01b1660208201520190565b0392611ad27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe094858101835282610773565b519020836116be565b9765ffffffffffff8060018301541680918b1610611b3c575b5091611b18611b24611b2f979695936115b995549660405193849182019586611114565b03908101835282610773565b5190209336916112a3565b15611b3657565b60019150565b9850611b18611af4565b6024611a3a565b5061010084013560028110156100dd57600103611b7357611a4c92611463928592611a31565b608486604051907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152602760248201527f53657373696f6e4b657956616c696461746f723a206f7065726174696f6e206d60448201527f69736d61746368000000000000000000000000000000000000000000000000006064820152fd5b905082602411611c1e5790611c0e611a0692610f2f565b16601086013560601c145f6119fb565b8380fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610729570180359067ffffffffffffffff82116107295760200191606082023603831361072957565b60049082821161072957604090848201357fffffffff00000000000000000000000000000000000000000000000000000000818116918290036107295782351603611eaa575f5b60808601611ccb8188611c22565b9050821015611e9e57611cde9087611c22565b821015611e725760608202018035808601808711611e46576024808301809211611e1c575087106107295783018501356020820135600681101561072957801580611e0f575b15611d36575050505050505050505f90565b6001811480611e02575b15611d52575050505050505050505f90565b6002811480611df5575b15611d6e575050505050505050505f90565b6003811480611de9575b15611d8a575050505050505050505f90565b86811480611ddd575b15611da5575050505050505050505f90565b60058691149283611dd0575b505050611dc657611dc190610ed5565b611cbd565b5050505050505f90565b01351490505f8481611db1565b50858301358211611d93565b50858301358210611d78565b5085830135821015611d5c565b5085830135821115611d40565b5085830135821415611d24565b6011887f4e487b71000000000000000000000000000000000000000000000000000000005f52525ffd5b6011877f4e487b71000000000000000000000000000000000000000000000000000000005f525260245ffd5b6032857f4e487b71000000000000000000000000000000000000000000000000000000005f525260245ffd5b50505050505050600190565b50505050505f9056fea164736f6c6343000815000a",
}

// SessionKeyValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use SessionKeyValidatorMetaData.ABI instead.
var SessionKeyValidatorABI = SessionKeyValidatorMetaData.ABI

// SessionKeyValidatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SessionKeyValidatorMetaData.Bin instead.
var SessionKeyValidatorBin = SessionKeyValidatorMetaData.Bin

// DeploySessionKeyValidator deploys a new Ethereum contract, binding an instance of SessionKeyValidator to it.
func DeploySessionKeyValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SessionKeyValidator, error) {
	parsed, err := SessionKeyValidatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SessionKeyValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SessionKeyValidator{SessionKeyValidatorCaller: SessionKeyValidatorCaller{contract: contract}, SessionKeyValidatorTransactor: SessionKeyValidatorTransactor{contract: contract}, SessionKeyValidatorFilterer: SessionKeyValidatorFilterer{contract: contract}}, nil
}

// SessionKeyValidator is an auto generated Go binding around an Ethereum contract.
type SessionKeyValidator struct {
	SessionKeyValidatorCaller     // Read-only binding to the contract
	SessionKeyValidatorTransactor // Write-only binding to the contract
	SessionKeyValidatorFilterer   // Log filterer for contract events
}

// SessionKeyValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type SessionKeyValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionKeyValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SessionKeyValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionKeyValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SessionKeyValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionKeyValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SessionKeyValidatorSession struct {
	Contract     *SessionKeyValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SessionKeyValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SessionKeyValidatorCallerSession struct {
	Contract *SessionKeyValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SessionKeyValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SessionKeyValidatorTransactorSession struct {
	Contract     *SessionKeyValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SessionKeyValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type SessionKeyValidatorRaw struct {
	Contract *SessionKeyValidator // Generic contract binding to access the raw methods on
}

// SessionKeyValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SessionKeyValidatorCallerRaw struct {
	Contract *SessionKeyValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// SessionKeyValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SessionKeyValidatorTransactorRaw struct {
	Contract *SessionKeyValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSessionKeyValidator creates a new instance of SessionKeyValidator, bound to a specific deployed contract.
func NewSessionKeyValidator(address common.Address, backend bind.ContractBackend) (*SessionKeyValidator, error) {
	contract, err := bindSessionKeyValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SessionKeyValidator{SessionKeyValidatorCaller: SessionKeyValidatorCaller{contract: contract}, SessionKeyValidatorTransactor: SessionKeyValidatorTransactor{contract: contract}, SessionKeyValidatorFilterer: SessionKeyValidatorFilterer{contract: contract}}, nil
}

// NewSessionKeyValidatorCaller creates a new read-only instance of SessionKeyValidator, bound to a specific deployed contract.
func NewSessionKeyValidatorCaller(address common.Address, caller bind.ContractCaller) (*SessionKeyValidatorCaller, error) {
	contract, err := bindSessionKeyValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SessionKeyValidatorCaller{contract: contract}, nil
}

// NewSessionKeyValidatorTransactor creates a new write-only instance of SessionKeyValidator, bound to a specific deployed contract.
func NewSessionKeyValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*SessionKeyValidatorTransactor, error) {
	contract, err := bindSessionKeyValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SessionKeyValidatorTransactor{contract: contract}, nil
}

// NewSessionKeyValidatorFilterer creates a new log filterer instance of SessionKeyValidator, bound to a specific deployed contract.
func NewSessionKeyValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*SessionKeyValidatorFilterer, error) {
	contract, err := bindSessionKeyValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SessionKeyValidatorFilterer{contract: contract}, nil
}

// bindSessionKeyValidator binds a generic wrapper to an already deployed contract.
func bindSessionKeyValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SessionKeyValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SessionKeyValidator *SessionKeyValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SessionKeyValidator.Contract.SessionKeyValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SessionKeyValidator *SessionKeyValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.SessionKeyValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SessionKeyValidator *SessionKeyValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.SessionKeyValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SessionKeyValidator *SessionKeyValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SessionKeyValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SessionKeyValidator *SessionKeyValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SessionKeyValidator *SessionKeyValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.contract.Transact(opts, method, params...)
}

// ExecutionStatus is a free data retrieval call binding the contract method 0x46585db2.
//
// Solidity: function executionStatus(bytes32 permissionKey, address kernel) view returns(uint48 validAfter, uint48 runs)
func (_SessionKeyValidator *SessionKeyValidatorCaller) ExecutionStatus(opts *bind.CallOpts, permissionKey [32]byte, kernel common.Address) (struct {
	ValidAfter *big.Int
	Runs       *big.Int
}, error) {
	var out []interface{}
	err := _SessionKeyValidator.contract.Call(opts, &out, "executionStatus", permissionKey, kernel)

	outstruct := new(struct {
		ValidAfter *big.Int
		Runs       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValidAfter = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Runs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ExecutionStatus is a free data retrieval call binding the contract method 0x46585db2.
//
// Solidity: function executionStatus(bytes32 permissionKey, address kernel) view returns(uint48 validAfter, uint48 runs)
func (_SessionKeyValidator *SessionKeyValidatorSession) ExecutionStatus(permissionKey [32]byte, kernel common.Address) (struct {
	ValidAfter *big.Int
	Runs       *big.Int
}, error) {
	return _SessionKeyValidator.Contract.ExecutionStatus(&_SessionKeyValidator.CallOpts, permissionKey, kernel)
}

// ExecutionStatus is a free data retrieval call binding the contract method 0x46585db2.
//
// Solidity: function executionStatus(bytes32 permissionKey, address kernel) view returns(uint48 validAfter, uint48 runs)
func (_SessionKeyValidator *SessionKeyValidatorCallerSession) ExecutionStatus(permissionKey [32]byte, kernel common.Address) (struct {
	ValidAfter *big.Int
	Runs       *big.Int
}, error) {
	return _SessionKeyValidator.Contract.ExecutionStatus(&_SessionKeyValidator.CallOpts, permissionKey, kernel)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address kernel) view returns(uint128 lastNonce, uint128 invalidNonce)
func (_SessionKeyValidator *SessionKeyValidatorCaller) Nonces(opts *bind.CallOpts, kernel common.Address) (struct {
	LastNonce    *big.Int
	InvalidNonce *big.Int
}, error) {
	var out []interface{}
	err := _SessionKeyValidator.contract.Call(opts, &out, "nonces", kernel)

	outstruct := new(struct {
		LastNonce    *big.Int
		InvalidNonce *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastNonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.InvalidNonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address kernel) view returns(uint128 lastNonce, uint128 invalidNonce)
func (_SessionKeyValidator *SessionKeyValidatorSession) Nonces(kernel common.Address) (struct {
	LastNonce    *big.Int
	InvalidNonce *big.Int
}, error) {
	return _SessionKeyValidator.Contract.Nonces(&_SessionKeyValidator.CallOpts, kernel)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address kernel) view returns(uint128 lastNonce, uint128 invalidNonce)
func (_SessionKeyValidator *SessionKeyValidatorCallerSession) Nonces(kernel common.Address) (struct {
	LastNonce    *big.Int
	InvalidNonce *big.Int
}, error) {
	return _SessionKeyValidator.Contract.Nonces(&_SessionKeyValidator.CallOpts, kernel)
}

// SessionData is a free data retrieval call binding the contract method 0x52721fdd.
//
// Solidity: function sessionData(address sessionKey, address kernel) view returns(bytes32 merkleRoot, uint48 validAfter, uint48 validUntil, address paymaster, uint256 nonce)
func (_SessionKeyValidator *SessionKeyValidatorCaller) SessionData(opts *bind.CallOpts, sessionKey common.Address, kernel common.Address) (struct {
	MerkleRoot [32]byte
	ValidAfter *big.Int
	ValidUntil *big.Int
	Paymaster  common.Address
	Nonce      *big.Int
}, error) {
	var out []interface{}
	err := _SessionKeyValidator.contract.Call(opts, &out, "sessionData", sessionKey, kernel)

	outstruct := new(struct {
		MerkleRoot [32]byte
		ValidAfter *big.Int
		ValidUntil *big.Int
		Paymaster  common.Address
		Nonce      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MerkleRoot = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ValidAfter = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ValidUntil = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Paymaster = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Nonce = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SessionData is a free data retrieval call binding the contract method 0x52721fdd.
//
// Solidity: function sessionData(address sessionKey, address kernel) view returns(bytes32 merkleRoot, uint48 validAfter, uint48 validUntil, address paymaster, uint256 nonce)
func (_SessionKeyValidator *SessionKeyValidatorSession) SessionData(sessionKey common.Address, kernel common.Address) (struct {
	MerkleRoot [32]byte
	ValidAfter *big.Int
	ValidUntil *big.Int
	Paymaster  common.Address
	Nonce      *big.Int
}, error) {
	return _SessionKeyValidator.Contract.SessionData(&_SessionKeyValidator.CallOpts, sessionKey, kernel)
}

// SessionData is a free data retrieval call binding the contract method 0x52721fdd.
//
// Solidity: function sessionData(address sessionKey, address kernel) view returns(bytes32 merkleRoot, uint48 validAfter, uint48 validUntil, address paymaster, uint256 nonce)
func (_SessionKeyValidator *SessionKeyValidatorCallerSession) SessionData(sessionKey common.Address, kernel common.Address) (struct {
	MerkleRoot [32]byte
	ValidAfter *big.Int
	ValidUntil *big.Int
	Paymaster  common.Address
	Nonce      *big.Int
}, error) {
	return _SessionKeyValidator.Contract.SessionData(&_SessionKeyValidator.CallOpts, sessionKey, kernel)
}

// ValidCaller is a free data retrieval call binding the contract method 0x9ea9bd59.
//
// Solidity: function validCaller(address , bytes ) view returns(bool)
func (_SessionKeyValidator *SessionKeyValidatorCaller) ValidCaller(opts *bind.CallOpts, arg0 common.Address, arg1 []byte) (bool, error) {
	var out []interface{}
	err := _SessionKeyValidator.contract.Call(opts, &out, "validCaller", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidCaller is a free data retrieval call binding the contract method 0x9ea9bd59.
//
// Solidity: function validCaller(address , bytes ) view returns(bool)
func (_SessionKeyValidator *SessionKeyValidatorSession) ValidCaller(arg0 common.Address, arg1 []byte) (bool, error) {
	return _SessionKeyValidator.Contract.ValidCaller(&_SessionKeyValidator.CallOpts, arg0, arg1)
}

// ValidCaller is a free data retrieval call binding the contract method 0x9ea9bd59.
//
// Solidity: function validCaller(address , bytes ) view returns(bool)
func (_SessionKeyValidator *SessionKeyValidatorCallerSession) ValidCaller(arg0 common.Address, arg1 []byte) (bool, error) {
	return _SessionKeyValidator.Contract.ValidCaller(&_SessionKeyValidator.CallOpts, arg0, arg1)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 , bytes ) pure returns(uint256)
func (_SessionKeyValidator *SessionKeyValidatorCaller) ValidateSignature(opts *bind.CallOpts, arg0 [32]byte, arg1 []byte) (*big.Int, error) {
	var out []interface{}
	err := _SessionKeyValidator.contract.Call(opts, &out, "validateSignature", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 , bytes ) pure returns(uint256)
func (_SessionKeyValidator *SessionKeyValidatorSession) ValidateSignature(arg0 [32]byte, arg1 []byte) (*big.Int, error) {
	return _SessionKeyValidator.Contract.ValidateSignature(&_SessionKeyValidator.CallOpts, arg0, arg1)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 , bytes ) pure returns(uint256)
func (_SessionKeyValidator *SessionKeyValidatorCallerSession) ValidateSignature(arg0 [32]byte, arg1 []byte) (*big.Int, error) {
	return _SessionKeyValidator.Contract.ValidateSignature(&_SessionKeyValidator.CallOpts, arg0, arg1)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes _data) payable returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactor) Disable(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _SessionKeyValidator.contract.Transact(opts, "disable", _data)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes _data) payable returns()
func (_SessionKeyValidator *SessionKeyValidatorSession) Disable(_data []byte) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.Disable(&_SessionKeyValidator.TransactOpts, _data)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes _data) payable returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactorSession) Disable(_data []byte) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.Disable(&_SessionKeyValidator.TransactOpts, _data)
}

// Enable is a paid mutator transaction binding the contract method 0x0c959556.
//
// Solidity: function enable(bytes _data) payable returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactor) Enable(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _SessionKeyValidator.contract.Transact(opts, "enable", _data)
}

// Enable is a paid mutator transaction binding the contract method 0x0c959556.
//
// Solidity: function enable(bytes _data) payable returns()
func (_SessionKeyValidator *SessionKeyValidatorSession) Enable(_data []byte) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.Enable(&_SessionKeyValidator.TransactOpts, _data)
}

// Enable is a paid mutator transaction binding the contract method 0x0c959556.
//
// Solidity: function enable(bytes _data) payable returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactorSession) Enable(_data []byte) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.Enable(&_SessionKeyValidator.TransactOpts, _data)
}

// InvalidateNonce is a paid mutator transaction binding the contract method 0xdbba225d.
//
// Solidity: function invalidateNonce(uint128 nonce) returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactor) InvalidateNonce(opts *bind.TransactOpts, nonce *big.Int) (*types.Transaction, error) {
	return _SessionKeyValidator.contract.Transact(opts, "invalidateNonce", nonce)
}

// InvalidateNonce is a paid mutator transaction binding the contract method 0xdbba225d.
//
// Solidity: function invalidateNonce(uint128 nonce) returns()
func (_SessionKeyValidator *SessionKeyValidatorSession) InvalidateNonce(nonce *big.Int) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.InvalidateNonce(&_SessionKeyValidator.TransactOpts, nonce)
}

// InvalidateNonce is a paid mutator transaction binding the contract method 0xdbba225d.
//
// Solidity: function invalidateNonce(uint128 nonce) returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactorSession) InvalidateNonce(nonce *big.Int) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.InvalidateNonce(&_SessionKeyValidator.TransactOpts, nonce)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 , uint256 ) payable returns(uint256)
func (_SessionKeyValidator *SessionKeyValidatorTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp UserOperation, arg1 [32]byte, arg2 *big.Int) (*types.Transaction, error) {
	return _SessionKeyValidator.contract.Transact(opts, "validateUserOp", userOp, arg1, arg2)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 , uint256 ) payable returns(uint256)
func (_SessionKeyValidator *SessionKeyValidatorSession) ValidateUserOp(userOp UserOperation, arg1 [32]byte, arg2 *big.Int) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.ValidateUserOp(&_SessionKeyValidator.TransactOpts, userOp, arg1, arg2)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 , uint256 ) payable returns(uint256)
func (_SessionKeyValidator *SessionKeyValidatorTransactorSession) ValidateUserOp(userOp UserOperation, arg1 [32]byte, arg2 *big.Int) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.ValidateUserOp(&_SessionKeyValidator.TransactOpts, userOp, arg1, arg2)
}
