// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/flare-foundation/flare/utils/units"
	"github.com/flare-foundation/flare/vms/platformvm/reward"
)

var (
	localGenesisConfigJSON = `{
		"networkID": 12345,
		"allocations": [],
		"startTime": 1630987200,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [],
		"initialStakers": [],
		"cChainGenesis": "",
		"message": "flare"
	}`

	// localCChainGenesis is the C-Chain genesis block used for the local
	// network.
	localCChainGenesis = `{
		"config": {
			"chainId": 4294967295,
			"homesteadBlock": 0,
			"daoForkBlock": 0,
			"daoForkSupport": true,
			"eip150Block": 0,
			"eip150Hash": "0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0",
			"eip155Block": 0,
			"eip158Block": 0,
			"byzantiumBlock": 0,
			"constantinopleBlock": 0,
			"petersburgBlock": 0,
			"istanbulBlock": 0,
			"muirGlacierBlock": 0,
			"apricotPhase1BlockTimestamp": 0,
			"apricotPhase2BlockTimestamp": 0,
			"apricotPhase3BlockTimestamp": 0,
			"apricotPhase4BlockTimestamp": 0,
			"apricotPhase5BlockTimestamp": 0
		},
		"nonce": "0x0",
		"timestamp": "0x0",
		"extraData": "0x00",
		"gasLimit": "0x5f5e100",
		"difficulty": "0x0",
		"mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
		"coinbase": "0x0100000000000000000000000000000000000000",
		"alloc": {
			"1000000000000000000000000000000000000001": {
				"balance": "0x0",
				"code": "0x608060405234801561001057600080fd5b50600436106100b45760003560e01c8063cfd1fdad11610071578063cfd1fdad1461015f578063eaebf6d3146101a2578063ec7424a0146101c7578063f417c9d8146101cf578063f5f59a4a146101d7578063f64b6fda146101df576100b4565b806329be4db2146100b95780634b8a125f146100e85780635f8c940d146100f057806371c5ecb1146100f857806371e24574146101155780637ff6faa61461013b575b600080fd5b6100d6600480360360208110156100cf57600080fd5b503561024f565b60408051918252519081900360200190f35b6100d661033f565b6100d6610347565b6100d66004803603602081101561010e57600080fd5b503561034c565b6100d66004803603602081101561012b57600080fd5b50356001600160a01b0316610364565b610143610379565b604080516001600160a01b039092168252519081900360200190f35b61018e6004803603608081101561017557600080fd5b508035906020810135906040810135906060013561037f565b604080519115158252519081900360200190f35b6101c5600480360360408110156101b857600080fd5b508035906020013561041d565b005b6100d66104c1565b6100d66104c7565b6100d66104cd565b6101c5600480360360208110156101f557600080fd5b81019060208101813564010000000081111561021057600080fd5b82018360208201111561022257600080fd5b8035906020019184600183028401116401000000008311171561024457600080fd5b5090925090506104d2565b60006001821161025e57600080fd5b3360009081526020819052604090206009015460001983019081111561028357600080fd5b33600090815260208190526040812060038306600381106102a057fe5b6003908102919091016002015433600090815260208190526040812091935091600019850106600381106102d057fe5b60030201600101549050816040516020018082815260200191505060405160208183030381529060405280519060200120811461030c57600080fd5b33600090815260208190526040812060036000198601066003811061032d57fe5b60030201549290921895945050505050565b636184740081565b600381565b600281611a40811061035d57600080fd5b0154905081565b60006020819052908152604090206009015481565b61dead81565b6000605a63618473ff19420104851461039757600080fd5b33600081815260208181526040808320600981018a905581516060810183528981528084018990529182018790529383529190529060038706600381106103da57fe5b6003020160008201518160000155602082015181600101556040820151816002015590505060015485111561041157506001610415565b5060005b949350505050565b6001821161042a57600080fd5b605a63618473ff19420104821461044057600080fd5b600154821161044e57600080fd5b334114801561045e57504161dead145b156104bd576001829055806002611a40600019850106611a40811061047f57fe5b0155604080518381526020810183905281517f8ffd19aa79a62d0764e560d21b1245698310783be781d7d80b38233d4d7d288c929181900390910190a15b5050565b60015481565b611a4081565b605a81565b7f5a4fad455fbfa0bb0f22d912bbfa4ef3d0887bb6933bffaf0f1f3e9fe1a12ca142838360405180848152602001806020018281038252848482818152602001925080828437600083820152604051601f909101601f1916909201829003965090945050505050a1505056fea264697066735822122093876c1f2ec0c4be61ccfd1ee6b192af3d6daff617d10ecf09f29bea82c4570564736f6c63430007060033"
			},
			"1000000000000000000000000000000000000002": {
				"balance": "0x0",
				"code": "0x608060405234801561001057600080fd5b50600436106101cf5760003560e01c80638be2fb8611610104578063d38bfff4116100a2578063e9de7d6011610071578063e9de7d6014610322578063ecdda0dd14610335578063ed21b6e414610348578063ee323c921461035f576101cf565b8063d38bfff4146102e9578063d48a38df146102fc578063dded1b4714610304578063e371aef01461030c576101cf565b8063a6817ace116100de578063a6817ace146102be578063be0522e0146102c6578063c373a08e146102ce578063c9f960eb146102e1576101cf565b80638be2fb861461029b5780638ccf77a0146102a35780639d6a890f146102ab576101cf565b806362da19a511610171578063689c49991161014b578063689c49991461026557806372993615146102785780637fec8d3814610280578063870196b814610288576101cf565b806362da19a51461023c578063639031431461024457806363d4a53a1461024c576101cf565b80635042916c116101ad5780635042916c146102025780635aa6e675146102175780635d36b1901461022c57806360f7ac9714610234576101cf565b806310663750146101d45780631d76dea1146101f25780634f6a77b5146101fa575b600080fd5b6101dc610372565b6040516101e991906127a1565b60405180910390f35b6101dc610378565b6101dc61037e565b61021561021036600461249c565b610384565b005b61021f6103da565b6040516101e99190612592565b6102156103e9565b61021f6104ab565b6101dc6104ba565b6101dc6104c0565b6102546104d0565b6040516101e9959493929190612698565b61021561027336600461240d565b61050b565b6101dc61083a565b6101dc610840565b61021561029636600461249c565b6108eb565b6101dc610a3b565b610215610a41565b6102156102b93660046123f1565b610a9c565b6101dc610aa1565b61021f610aa7565b6102156102dc3660046123f1565b610ab6565b61021f610b5b565b6102156102f73660046123f1565b610b9a565b6101dc610c5c565b6101dc610c62565b610314610c68565b6040516101e992919061277f565b61021561033036600461249c565b610c89565b6102546103433660046124b4565b610dc2565b61035061111f565b6040516101e993929190612623565b61021561036d3660046123f1565b6112e4565b60055481565b60095481565b60085481565b6000546001600160a01b031633146103d5576040805162461bcd60e51b815260206004820152600f60248201526e6f6e6c7920676f7665726e616e636560881b604482015290519081900360640190fd5b600c55565b6000546001600160a01b031681565b6001546001600160a01b03163314610438576040805162461bcd60e51b815260206004820152600d60248201526c1b9bdd0818db185a5b585a5b9d609a1b604482015290519081900360640190fd5b600054600154604080516001600160a01b03938416815292909116602083015280517f434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d9281900390910190a160018054600080546001600160a01b03199081166001600160a01b03841617909155169055565b6001546001600160a01b031681565b60045481565b60006104ca6113fb565b90505b90565b600b546060908190819081906000906104fa90600160c01b90046001600160401b03166001610dc2565b945094509450945094509091929394565b6000546001600160a01b0316331461055c576040805162461bcd60e51b815260206004820152600f60248201526e6f6e6c7920676f7665726e616e636560881b604482015290519081900360640190fd5b604080518082019091526008815267746f6f206d616e7960c01b60208201528190600a8211156105a85760405162461bcd60e51b815260040161059f919061276c565b60405180910390fd5b506105b1611406565b60005b818110156108345760008484838181106105ca57fe5b6105e092602060409092020190810191506123f1565b6001600160a01b031614156040518060400160405280600c81526020016b61646472657373207a65726f60a01b8152509061062e5760405162461bcd60e51b815260040161059f919061276c565b5060105460005b818110156106d7576010818154811061064a57fe5b6000918252602090912001546001600160a01b031686868581811061066b57fe5b61068192602060409092020190810191506123f1565b6001600160a01b031614156040518060400160405280600b81526020016a647570206164647265737360a81b815250906106ce5760405162461bcd60e51b815260040161059f919061276c565b50600101610635565b5060108585848181106106e657fe5b6106fc92602060409092020190810191506123f1565b81546001810183556000928352602090922090910180546001600160a01b0319166001600160a01b0390921691909117905584848381811061073a57fe5b905060400201602001356011600087878681811061075457fe5b61076a92602060409092020190810191506123f1565b6001600160a01b03166001600160a01b03168152602001908152602001600020819055506000601260008787868181106107a057fe5b6107b692602060409092020190810191506123f1565b6001600160a01b031681526020810191909152604001600020557f86d03f430c7616021073d7a71766f632f1ce19f289aa989534d9f4732253eb598585848181106107fd57fe5b61081392602060409092020190810191506123f1565b6001604051610823929190612737565b60405180910390a1506001016105b4565b50505050565b60075481565b6002546000906001600160a01b031661088557610885306040518060400160405280600e81526020016d696e666c6174696f6e207a65726f60901b81525060006114b6565b60026001609c1b01331461089857600080fd5b6108a061167a565b905060006108ac611ee5565b90504781146108e7576108e7306040518060400160405280600e81526020016d6f7574206f662062616c616e636560901b81525060006114b6565b5090565b6000546001600160a01b0316331461093c576040805162461bcd60e51b815260206004820152600f60248201526e6f6e6c7920676f7665726e616e636560881b604482015290519081900360640190fd5b60085461094c90606e6064611f0a565b811115604051806040016040528060118152602001700dac2f040dad2dce840e8dede40d0d2ced607b1b815250906109975760405162461bcd60e51b815260040161059f919061276c565b5060408051808201909152601081526f6d6178206d696e74206973207a65726f60801b6020820152816109dd5760405162461bcd60e51b815260040161059f919061276c565b5062015180600a54014211604051806040016040528060128152602001711d1a5b594819d85c081d1bdbc81cda1bdc9d60721b81525090610a315760405162461bcd60e51b815260040161059f919061276c565b5060085542600a55565b60035481565b6000546001600160a01b03163314610a92576040805162461bcd60e51b815260206004820152600f60248201526e6f6e6c7920676f7665726e616e636560881b604482015290519081900360640190fd5b610a9a611406565b565bfe5b50565b600a5481565b6002546001600160a01b031681565b6000546001600160a01b03163314610b07576040805162461bcd60e51b815260206004820152600f60248201526e6f6e6c7920676f7665726e616e636560881b604482015290519081900360640190fd5b600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f1f95fb40be3a947982072902a887b521248d1d8931a39eb38f84f4d6fd758b699181900360200190a150565b600f5460009060ff16610b8857600f805460ff191660011790556000610b7f612012565b91506104cd9050565b506000546001600160a01b03166104cd565b6000546001600160a01b03163314610beb576040805162461bcd60e51b815260206004820152600f60248201526e6f6e6c7920676f7665726e616e636560881b604482015290519081900360640190fd5b600054604080516001600160a01b039283168152918316602083015280517f434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d9281900390910190a1600080546001600160a01b039092166001600160a01b0319928316179055600180549091169055565b600c5481565b60065481565b600b546001600160c01b03811690600160c01b90046001600160401b031682565b60025460408051808201909152600d81526c3737ba1034b7333630ba34b7b760991b602082015233916001600160a01b03168214610cda5760405162461bcd60e51b815260040161059f919061276c565b5060085482111560405180604001604052806007815260200166746f6f2062696760c81b81525090610d1f5760405162461bcd60e51b815260040161059f919061276c565b5042610d296113fb565b10604051806040016040528060098152602001683a37b79037b33a32b760b91b81525090610d6a5760405162461bcd60e51b815260040161059f919061276c565b508115610dbe5742600955600454610d829083612032565b6004556040517f4c4f1efc376f31abeb51b72c5f9ed81cf4016591312bb02337e58149dcfaaab490610db59084906127a1565b60405180910390a15b5050565b606080606080600060148054905087106040518060400160405280601081526020016f0e6e8c2e4e840d2dcc8caf040d0d2ced60831b81525090610e195760405162461bcd60e51b815260040161059f919061276c565b506014546000908888011115610e3457601454889003610e36565b865b9050806001600160401b0381118015610e4e57600080fd5b50604051908082528060200260200182016040528015610e78578160200160208202803683370190505b509550806001600160401b0381118015610e9157600080fd5b50604051908082528060200260200182016040528015610ebb578160200160208202803683370190505b509450806001600160401b0381118015610ed457600080fd5b50604051908082528060200260200182016040528015610f0857816020015b6060815260200190600190039081610ef35790505b509350806001600160401b0381118015610f2157600080fd5b50604051908082528060200260200182016040528015610f4b578160200160208202803683370190505b50925060005b818110156111035760006014828b0181548110610f6a57fe5b6000918252602080832090910154808352601390915260409091205489519192506001600160c01b031690899084908110610fa157fe5b6020026020010181815250506013600082815260200190815260200160002060000160189054906101000a90046001600160401b03166001600160401b0316878381518110610fec57fe5b6020908102919091018101919091526000828152601382526040908190206002908101805483516001821615610100026000190190911692909204601f810185900485028301850190935282825290929091908301828280156110905780601f1061106557610100808354040283529160200191611090565b820191906000526020600020905b81548152906001019060200180831161107357829003601f168201915b50505050508683815181106110a157fe5b60200260200101819052506013600082815260200190815260200160002060010160009054906101000a90046001600160a01b03168583815181106110e257fe5b6001600160a01b039092166020928302919091019091015250600101610f51565b5050600b549497939650919450926001600160c01b0316919050565b60105460609081908190806001600160401b038111801561113f57600080fd5b50604051908082528060200260200182016040528015611169578160200160208202803683370190505b509350806001600160401b038111801561118257600080fd5b506040519080825280602002602001820160405280156111ac578160200160208202803683370190505b509250806001600160401b03811180156111c557600080fd5b506040519080825280602002602001820160405280156111ef578160200160208202803683370190505b50915060005b818110156112dd5760006010828154811061120c57fe5b9060005260206000200160009054906101000a90046001600160a01b031690508086838151811061123957fe5b60200260200101906001600160a01b031690816001600160a01b03168152505060116000826001600160a01b03166001600160a01b031681526020019081526020016000205485838151811061128b57fe5b60200260200101818152505060126000826001600160a01b03166001600160a01b03168152602001908152602001600020548483815181106112c957fe5b6020908102919091010152506001016111f5565b5050909192565b6000546001600160a01b03163314611335576040805162461bcd60e51b815260206004820152600f60248201526e6f6e6c7920676f7665726e616e636560881b604482015290519081900360640190fd5b60408051808201909152600e81526d696e666c6174696f6e207a65726f60901b60208201526001600160a01b0382166113815760405162461bcd60e51b815260040161059f919061276c565b506002546040517f4bdd1012a7d55ed9afad8675a125e1b68c7c15f712c0f3d5cddac69c3b979805916113bf9184916001600160a01b031690612752565b60405180910390a1600280546001600160a01b0319166001600160a01b038316179055600854610a9e576a295be96e6406697200000060085550565b600954620143700190565b60105460005b81811015610dbe576010805460009190600019810190811061142a57fe5b600091825260209091200154601080546001600160a01b039092169250908061144f57fe5b600082815260208120820160001990810180546001600160a01b03191690559091019091556040517f86d03f430c7616021073d7a71766f632f1ce19f289aa989534d9f4732253eb59916114a591849190612737565b60405180910390a15060010161140c565b600083836040516020016114cb9291906125a6565b60408051808303601f190181529181528151602092830120600081815260139093529120805491925090600160c01b90046001600160401b03166115b25760148054600180820183556000929092527fce6d7b5282bd9a3661ae061feed1dbda4e52ab073b1f9285be6e155d9c38d4ec01839055810180546001600160a01b0319166001600160a01b038716179055611565846040612095565b805161157b916002840191602090910190612359565b5060145460018201805467ffffffffffffffff60a01b1916600160a01b6000199093016001600160401b0316929092029190911790555b8054436001600160c01b038181166001600160401b03600160c01b80860482166001019091160291909316176001600160c01b0319169190911782556040517f7a459ed083a9b267865360013a5ad6dbc07e5befe6e4f71671c940fdd4206bee9161162391889190889088906125eb565b60405180910390a1600b80546001600160c01b0319811660016001600160c01b0392831681018316919091178084559301549216600160a01b9092046001600160401b0316600160c01b0291909117905550505050565b600060035443141561168e575060006104cd565b43600355600d544790811115611a645760006116b7600e54600d5461203290919063ffffffff16565b90508082141561182c57600e546005546116d19082612032565b6005556040517fa42d823c276ad1990284418c303209194a75fa95a901f19752a9f65a407ffa8c906117049083906127a1565b60405180910390a1600260009054906101000a90046001600160a01b03166001600160a01b031663c611c2c5826040518263ffffffff1660e01b81526004016000604051808303818588803b15801561175c57600080fd5b505af19350505050801561176e575060015b6117dd5761177a6127b0565b806117855750611797565b611791308260006114b6565b506117d8565b6117d8306040518060400160405280601d81526020017f756e6b6e6f776e206572726f722e20726563656976654d696e74696e6700000081525060006114b6565b611826565b6006546117ea9082612032565b6006556040517f12773bf711e11ec0b058c3856d441d726d2dc89113706c4f4175571f1e830c5a9061181d9083906127a1565b60405180910390a15b50611a62565b8082101561188e57600061184b600d548461214990919063ffffffff16565b60075490915061185b9082612032565b6007556040517f3fe36bcb00188390b2b40f1ab66c58f660aea67fe98b9f80667f692e1a9ab3689061181d9083906127a1565b600e5460055461189d91612032565b600555600e54600d546000916118be916118b8908690612149565b90612149565b6007549091506118ce9082612032565b600755600e546040517fa42d823c276ad1990284418c303209194a75fa95a901f19752a9f65a407ffa8c91611902916127a1565b60405180910390a17f3fe36bcb00188390b2b40f1ab66c58f660aea67fe98b9f80667f692e1a9ab3688160405161193991906127a1565b60405180910390a1600260009054906101000a90046001600160a01b03166001600160a01b031663c611c2c5600e546040518263ffffffff1660e01b81526004016000604051808303818588803b15801561199357600080fd5b505af1935050505080156119a5575060015b611a14576119b16127b0565b806119bc57506119ce565b6119c8308260006114b6565b50611a0f565b611a0f306040518060400160405280601d81526020017f756e6b6e6f776e206572726f722e20726563656976654d696e74696e6700000081525060006114b6565b611a60565b600e54600654611a2391612032565b600655600e546040517f12773bf711e11ec0b058c3856d441d726d2dc89113706c4f4175571f1e830c5a91611a57916127a1565b60405180910390a15b505b505b60105460005b81811015611e8357600060108281548110611a8157fe5b60009182526020808320909101546001600160a01b031680835260129091526040909120549091508015611b0d576001600160a01b0382166000908152601260205260409081902060001983019055517f9895eddb1e8569b1dae526135aa5cab97f982fdc3b0ff7e17920c95e3b9bda6290611b0090849084906125d2565b60405180910390a1611e79565b6001600160a01b038216600090815260116020526040812054905a90506204a768811015611b77577f9b5c4be38598cb8d8b6e07727d2303d1d9fc2dfc31ad323170f5ea4dcc1f914a858703604051611b6691906127a1565b60405180910390a150505050611e83565b620493df1981018215801590611b8c57508083105b15611b945750815b846001600160a01b0316636d0e8c34826040518263ffffffff1660e01b8152600401602060405180830381600088803b158015611bd057600080fd5b5087f193505050508015611c01575060408051601f3d908101601f19168201909252611bfe9181019061247c565b60015b611e3857611c0d6127b0565b80611c185750611c9f565b611c2586825a86036114b6565b856001600160a01b031663e22fdece6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015611c6057600080fd5b505af1158015611c74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c98919061247c565b5050611e33565b60005a9050600084118015611cbd575083611cba8483612149565b10155b15611d9057611cf1866040518060400160405280600a8152602001696f7574206f662067617360b01b8152508386036114b6565b6000866001600160a01b031663e22fdece6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015611d2e57600080fd5b505af1158015611d42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d66919061247c565b905080611d8a57600c546001600160a01b0388166000908152601260205260409020555b50611e31565b611dbc86604051806040016040528060078152602001663ab735b737bbb760c91b8152508386036114b6565b856001600160a01b031663e22fdece6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015611df757600080fd5b505af1158015611e0b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e2f919061247c565b505b505b611e75565b507fe7aa66356adbd5e839ef210626f6d8f6f72109c17fadf4c4f9ca82b315ae79b4855a8403604051611e6c9291906125d2565b60405180910390a15b5050505b5050600101611a6a565b50611e8c6121a6565b92508215611ed657600e8390556040517f34f843cef0df42035141347873da1758a6643358831b5ba5b1580be947644f9290611ec99085906127a1565b60405180910390a1611edc565b6000600e555b505047600d5590565b60006104ca600754611f0460065460055461214990919063ffffffff16565b90612032565b6000808211611f53576040805162461bcd60e51b815260206004820152601060248201526f4469766973696f6e206279207a65726f60801b604482015290519081900360640190fd5b83611f605750600061200b565b83830283858281611f6d57fe5b041415611f8657828181611f7d57fe5b0491505061200b565b6000838681611f9157fe5b0490506000848781611f9f57fe5b0690506000858781611fad57fe5b0490506000868881611fbb57fe5b069050612003611fd588611fcf86856121bf565b90612218565b611f04611fe286866121bf565b611f04611fef89876121bf565b611f048d611ffd8c8b6121bf565b906121bf565b955050505050505b9392505050565b600073fffec6c83c8bf5c3f4ae0ccf8c45ce20e4560bd76104ca8161227f565b60008282018381101561208c576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b90505b92915050565b60606000839050828151116120ad578391505061208f565b6000836001600160401b03811180156120c557600080fd5b506040519080825280601f01601f1916602001820160405280156120f0576020820181803683370190505b50905060005b848110156121405782818151811061210a57fe5b602001015160f81c60f81b82828151811061212157fe5b60200101906001600160f81b031916908160001a9053506001016120f6565b50949350505050565b6000828211156121a0576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b60006104ca60055460045461214990919063ffffffff16565b6000826121ce5750600061208f565b828202828482816121db57fe5b041461208c5760405162461bcd60e51b815260040180806020018281038252602181526020018061286a6021913960400191505060405180910390fd5b600080821161226e576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b81838161227757fe5b049392505050565b600154600160a01b900460ff16156122d5576040805162461bcd60e51b8152602060048201526014602482015273696e697469616c6973656420213d2066616c736560601b604482015290519081900360640190fd5b6001805460ff60a01b1916600160a01b179055600054604080516001600160a01b039283168152918316602083015280517f434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d9281900390910190a1600080546001600160a01b039092166001600160a01b0319928316179055600180549091169055565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928261238f57600085556123d5565b82601f106123a857805160ff19168380011785556123d5565b828001600101855582156123d5579182015b828111156123d55782518255916020019190600101906123ba565b506108e79291505b808211156108e757600081556001016123dd565b600060208284031215612402578081fd5b813561208c81612854565b6000806020838503121561241f578081fd5b82356001600160401b0380821115612435578283fd5b818501915085601f830112612448578283fd5b813581811115612456578384fd5b86602060408302850101111561246a578384fd5b60209290920196919550909350505050565b60006020828403121561248d578081fd5b8151801515811461208c578182fd5b6000602082840312156124ad578081fd5b5035919050565b600080604083850312156124c6578182fd5b50508035926020909101359150565b6000815180845260208085019450808401835b8381101561250d5781516001600160a01b0316875295820195908201906001016124e8565b509495945050505050565b6000815180845260208085019450808401835b8381101561250d5781518752958201959082019060010161252b565b60008151808452815b8181101561256c57602081850181015186830182015201612550565b8181111561257d5782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b6001600160a01b03831681526040602082018190526000906125ca90830184612547565b949350505050565b6001600160a01b03929092168252602082015260400190565b600060018060a01b0386168252846020830152608060408301526126126080830185612547565b905082606083015295945050505050565b606080825284519082018190526000906020906080840190828801845b828110156126655781516001600160a01b031684529284019290840190600101612640565b505050838103828501526126798187612518565b915050828103604084015261268e8185612518565b9695505050505050565b600060a082526126ab60a0830188612518565b6020838203818501526126be8289612518565b848103604086015287518082529092508183019082810284018301838a01865b8381101561270c57601f198784030185526126fa838351612547565b948601949250908501906001016126de565b50508681036060880152612720818a6124d5565b955050505050508260808301529695505050505050565b6001600160a01b039290921682521515602082015260400190565b6001600160a01b0392831681529116602082015260400190565b60006020825261200b6020830184612547565b6001600160c01b039290921682526001600160401b0316602082015260400190565b90815260200190565b60e01c90565b600060443d10156127c0576104cd565b600481823e6308c379a06127d482516127aa565b146127de576104cd565b6040513d600319016004823e80513d6001600160401b03816024840111818411171561280d57505050506104cd565b8284019250825191508082111561282757505050506104cd565b503d8301602082840101111561283f575050506104cd565b601f01601f1916810160200160405291505090565b6001600160a01b0381168114610a9e57600080fdfe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a26469706673582212204565821f65608e6e4fdb7076f78e5f829e603a82de85005cd73fca61d73e634f64736f6c63430007060033"
			},
			"1000000000000000000000000000000000000003": {
				"balance": "0x0",
				"code": "0x608060405234801561001057600080fd5b50600436106101165760003560e01c80639d6a890f116100a2578063c373a08e11610071578063c373a08e1461051b578063c5adc53914610541578063c9f960eb1461066b578063d38bfff414610673578063ffacb84e1461069957610116565b80639d6a890f146104205780639d986f91146104465780639ec2b58114610472578063b39c68581461051357610116565b806371e1fad9116100e957806371e1fad9146102fd57806376794efb146103055780637ac420ad146103a85780638ab63380146103e05780638c9d28b61461041857610116565b80635aa6e6751461011b5780635d36b1901461013f57806360848b441461014957806360f7ac97146102f5575b600080fd5b6101236106f1565b604080516001600160a01b039092168252519081900360200190f35b610147610700565b005b6101476004803603608081101561015f57600080fd5b81359190810190604081016020820135600160201b81111561018057600080fd5b82018360208201111561019257600080fd5b803590602001918460208302840111600160201b831117156101b357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561020257600080fd5b82018360208201111561021457600080fd5b803590602001918460208302840111600160201b8311171561023557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561028457600080fd5b82018360208201111561029657600080fd5b803590602001918460208302840111600160201b831117156102b757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506107c2945050505050565b610123610dbe565b610123610dcd565b6101476004803603604081101561031b57600080fd5b810190602081018135600160201b81111561033557600080fd5b82018360208201111561034757600080fd5b803590602001918460208302840111600160201b8311171561036857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505091359250610ddc915050565b6103ce600480360360208110156103be57600080fd5b50356001600160a01b0316610ebd565b60408051918252519081900360200190f35b610147600480360360608110156103f657600080fd5b506001600160a01b038135811691602081013582169160409091013516610ed8565b610123610f68565b6101476004803603602081101561043657600080fd5b50356001600160a01b0316610f77565b6101476004803603604081101561045c57600080fd5b506001600160a01b038135169060200135610f79565b6101476004803603602081101561048857600080fd5b810190602081018135600160201b8111156104a257600080fd5b8201836020820111156104b457600080fd5b803590602001918460208302840111600160201b831117156104d557600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611029945050505050565b610123611177565b6101476004803603602081101561053157600080fd5b50356001600160a01b0316611186565b6101476004803603606081101561055757600080fd5b81359190810190604081016020820135600160201b81111561057857600080fd5b82018360208201111561058a57600080fd5b803590602001918460208302840111600160201b831117156105ab57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156105fa57600080fd5b82018360208201111561060c57600080fd5b803590602001918460208302840111600160201b8311171561062d57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061122b945050505050565b6101236115e6565b6101476004803603602081101561068957600080fd5b50356001600160a01b031661160b565b6106a16116cd565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156106dd5781810151838201526020016106c5565b505050509050019250505060405180910390f35b6000546001600160a01b031681565b6001546001600160a01b0316331461074f576040805162461bcd60e51b815260206004820152600d60248201526c1b9bdd0818db185a5b585a5b9d609a1b604482015290519081900360640190fd5b600054600154604080516001600160a01b03938416815292909116602083015280517f434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d9281900390910190a160018054600080546001600160a01b03199081166001600160a01b03841617909155169055565b8251825160408051808201909152601a815279082e4e4c2f240d8cadccee8d0e640c8de40dcdee840dac2e8c6d60331b60208201529082146108825760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561084757818101518382015260200161082f565b50505050905090810190601f1680156108745780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50815181146040518060400160405280601a815260200179082e4e4c2f240d8cadccee8d0e640c8de40dcdee840dac2e8c6d60331b815250906109065760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561084757818101518382015260200161082f565b506002546040516313968ea760e31b81526020600482018181528751602484015287516000946001600160a01b031693639cb47538938a939283926044019180860191028083838b5b8381101561096757818101518382015260200161094f565b505050509050019250505060006040518083038186803b15801561098a57600080fd5b505afa15801561099e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156109c757600080fd5b8101908080516040519392919084600160201b8211156109e657600080fd5b9083019060208201858111156109fb57600080fd5b82518660208202830111600160201b82111715610a1757600080fd5b82525081516020918201928201910280838360005b83811015610a44578181015183820152602001610a2c565b5050505091909101604090815233600090815260056020529081205495965093506000199250839150505b85811015610c9e576000898281518110610a8557fe5b60200260200101519050806001901b851660001415610b2e5783610b2e573360009081526007602052604090205460ff1615610ac45760019350610b2e565b604080518082018252600f81526e139bdd081dda1a5d195b1a5cdd1959608a1b6020808301918252925162461bcd60e51b8152600481019384528251602482015282519293928392604490920191908083836000831561084757818101518382015260200161082f565b600019831415610bd157858281518110610b4457fe5b60200260200101516001600160a01b031663f72cab28338d6040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015610ba257600080fd5b505af1158015610bb6573d6000803e3d6000fd5b505050506040513d6020811015610bcc57600080fd5b505192505b858281518110610bdd57fe5b60200260200101516001600160a01b03166355f7b69b338d8c8681518110610c0157fe5b60200260200101518c8781518110610c1557fe5b6020026020010151886040518663ffffffff1660e01b815260040180866001600160a01b0316815260200185815260200184815260200183815260200182815260200195505050505050600060405180830381600087803b158015610c7957600080fd5b505af1158015610c8d573d6000803e3d6000fd5b505060019093019250610a6f915050565b5088336001600160a01b03167fa32444a31df2f9a116229eec3193d223a6bad89f7670ff17b8e5c7014a377da1868a8a4260405180806020018060200180602001858152602001848103845288818151815260200191508051906020019060200280838360005b83811015610d1d578181015183820152602001610d05565b50505050905001848103835287818151815260200191508051906020019060200280838360005b83811015610d5c578181015183820152602001610d44565b50505050905001848103825286818151815260200191508051906020019060200280838360005b83811015610d9b578181015183820152602001610d83565b5050505090500197505050505050505060405180910390a3505050505050505050565b6001546001600160a01b031681565b6004546001600160a01b031690565b600454604080518082019091526016815275566f7465722077686974656c6973746572206f6e6c7960501b6020820152906001600160a01b03163314610e635760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561084757818101518382015260200161082f565b5060005b8251811015610eb857816001901b1960056000858481518110610e8657fe5b6020908102919091018101516001600160a01b0316825281019190915260400160002080549091169055600101610e67565b505050565b6001600160a01b031660009081526005602052604090205490565b6000546001600160a01b03163314610f29576040805162461bcd60e51b815260206004820152600f60248201526e6f6e6c7920676f7665726e616e636560881b604482015290519081900360640190fd5b600280546001600160a01b039485166001600160a01b031991821617909155600480549385169382169390931790925560038054919093169116179055565b6002546001600160a01b031690565bfe5b600454604080518082019091526016815275566f7465722077686974656c6973746572206f6e6c7960501b6020820152906001600160a01b031633146110005760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561084757818101518382015260200161082f565b506001600160a01b0390911660009081526005602052604090208054600190921b919091179055565b60035460408051808201909152601081526f4654534f4d616e61676572206f6e6c7960801b6020820152906001600160a01b031633146110aa5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561084757818101518382015260200161082f565b5060065460005b8181101561110857600060076000600684815481106110cc57fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff19169115159190911790556001016110b1565b5050805160005b818110156111635760016007600085848151811061112957fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191691151591909117905560010161110f565b508151610eb8906006906020850190611809565b6003546001600160a01b031690565b6000546001600160a01b031633146111d7576040805162461bcd60e51b815260206004820152600f60248201526e6f6e6c7920676f7665726e616e636560881b604482015290519081900360640190fd5b600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f1f95fb40be3a947982072902a887b521248d1d8931a39eb38f84f4d6fd758b699181900360200190a150565b8151815160408051808201909152601a815279082e4e4c2f240d8cadccee8d0e640c8de40dcdee840dac2e8c6d60331b60208201529082146112ae5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561084757818101518382015260200161082f565b506002546040516313968ea760e31b81526020600482018181528651602484015286516000946001600160a01b031693639cb475389389939283926044019180860191028083838b5b8381101561130f5781810151838201526020016112f7565b505050509050019250505060006040518083038186803b15801561133257600080fd5b505afa158015611346573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561136f57600080fd5b8101908080516040519392919084600160201b82111561138e57600080fd5b9083019060208201858111156113a357600080fd5b82518660208202830111600160201b821117156113bf57600080fd5b82525081516020918201928201910280838360005b838110156113ec5781810151838201526020016113d4565b50505050919091016040908152336000908152600560205290812054959650935083925050505b8481101561150e57600087828151811061142957fe5b60200260200101519050806001901b8416600014156114645782611464573360009081526007602052604090205460ff1615610ac457600192505b84828151811061147057fe5b60200260200101516001600160a01b03166327bd2ad5338b8a868151811061149457fe5b60200260200101516040518463ffffffff1660e01b815260040180846001600160a01b031681526020018381526020018281526020019350505050600060405180830381600087803b1580156114e957600080fd5b505af11580156114fd573d6000803e3d6000fd5b505060019093019250611413915050565b5086336001600160a01b03167f90c022ade239639b1f8c4ebb8a76df5e03a7129df46cf9bcdae3c1450ea35434858842604051808060200180602001848152602001838103835286818151815260200191508051906020019060200280838360005b83811015611588578181015183820152602001611570565b50505050905001838103825285818151815260200191508051906020019060200280838360005b838110156115c75781810151838201526020016115af565b505050509050019550505050505060405180910390a350505050505050565b600073fffec6c83c8bf5c3f4ae0ccf8c45ce20e4560bd76116068161172f565b905090565b6000546001600160a01b0316331461165c576040805162461bcd60e51b815260206004820152600f60248201526e6f6e6c7920676f7665726e616e636560881b604482015290519081900360640190fd5b600054604080516001600160a01b039283168152918316602083015280517f434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d9281900390910190a1600080546001600160a01b039092166001600160a01b0319928316179055600180549091169055565b6060600680548060200260200160405190810160405280929190818152602001828054801561172557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611707575b5050505050905090565b600154600160a01b900460ff1615611785576040805162461bcd60e51b8152602060048201526014602482015273696e697469616c6973656420213d2066616c736560601b604482015290519081900360640190fd5b6001805460ff60a01b1916600160a01b179055600054604080516001600160a01b039283168152918316602083015280517f434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d9281900390910190a1600080546001600160a01b039092166001600160a01b0319928316179055600180549091169055565b82805482825590600052602060002090810192821561185e579160200282015b8281111561185e57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611829565b5061186a92915061186e565b5090565b5b8082111561186a576000815560010161186f56fea2646970667358221220c1d73dd5d713527b4b18f6ca6d3d915fe356fa5425b5776172fa3bf91a5476e564736f6c63430007060033"
			},
			"1000000000000000000000000000000000000004": {
				"balance": "0x0",
				"code": "0x6080604052348015600f57600080fd5b5060043610603c5760003560e01c806336e880e2146041578063b66cf7c0146059578063b7ab4db5146075575b600080fd5b604760b1565b60408051918252519081900360200190f35b607360048036036020811015606d57600080fd5b503560b7565b005b607b60bc565b6040518082606080838360005b83811015609e5781810151838201526020016088565b5050505090500191505060405180910390f35b60005481565b600055565b60c260d0565b600181526002602082015290565b6040518060600160405280600390602082028036833750919291505056fea2646970667358221220a9b055891df80e44514ef4d7c5cc65c7cd9e9fe886282327f4bebff94e23c1c164736f6c63430007060033"
			},
			"0xc783df8a850f42e7f7e57013759c285caa701eb6": {
				"balance": "0x314dc6448d9338c15B0a00000000"
			},
			"6b0dd034A2FD67b932F10E3dBA1d2bbD39348695": { "balance": "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff" }
		},
		"number": "0x0",
		"gasUsed": "0x0",
		"parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000"
	}`

	// LocalParams are the params used for local networks
	LocalParams = Params{
		TxFeeConfig: TxFeeConfig{
			TxFee:                 units.MilliAvax,
			CreateAssetTxFee:      units.MilliAvax,
			CreateSubnetTxFee:     100 * units.MilliAvax,
			CreateBlockchainTxFee: 100 * units.MilliAvax,
		},
		StakingConfig: StakingConfig{
			UptimeRequirement: .8, // 80%
			MinValidatorStake: 2 * units.KiloAvax,
			MaxValidatorStake: 3 * units.MegaAvax,
			MinDelegatorStake: 25 * units.Avax,
			MinDelegationFee:  20000, // 2%
			MinStakeDuration:  24 * time.Hour,
			MaxStakeDuration:  365 * 24 * time.Hour,
			RewardConfig: reward.Config{
				MaxConsumptionRate: .12 * reward.PercentDenominator,
				MinConsumptionRate: .10 * reward.PercentDenominator,
				MintingPeriod:      365 * 24 * time.Hour,
				SupplyCap:          720 * units.MegaAvax,
			},
		},
	}
)
