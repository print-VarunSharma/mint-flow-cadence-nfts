package test

import (
	"testing"

	"github.com/onflow/cadence"
	"github.com/stretchr/testify/assert"
)

func TestDeployingContracts(t *testing.T) {
	blockchain := newBlockchain(t)
	standardContractsAddress := deployStandardContracts(t, blockchain)
	deployVeNftContract(t, blockchain, standardContractsAddress)
}

func TestGettingTotalSupply(t *testing.T) {
	blockchain := newBlockchain(t)
	standardContractAddress := deployStandardContracts(t, blockchain)
	veNftAddress, _ := deployVeNftContract(t, blockchain, standardContractAddress)
	totalSupply := getTotalSupply(t, blockchain, veNftNftAddress)
	assert.Equal(t, totalSupply, cadence.NewUInt64(0))
}

func TestMintingAnVeNFT(t *testing.T) {
	blockchain := newBlockchain(t)
	standardContractAddress := deployStandardContracts(t, blockchain)
	veNftAddress, signer := deployVeNftContract(t, blockchain, standardContractAddress)

	setUpTestAccountToReceiveVeNfts(t, blockchain, standardContractAddress, veNftAddress, signer)

	totalSupply := getTotalSupply(t, blockchain, veNftAddress)
	assert.EqualValues(t, cadence.NewUInt64(0), totalSupply)

	mint(t, blockchain, standardContractAddress, veNftAddress, signer, veNftAddress, "test name", "test description", "test thumbnail")

	totalSupply = getTotalSupply(t, blockchain, veNftAddress)
	assert.EqualValues(t, cadence.NewUInt64(1), totalSupply)
}

func TestGetveNft(t *testing.T) {
	type veNftMetadata struct {
		veNftID uint64
		ResourceID    uint64
		Name          string
		Description   string
		Thumbnail     string
		Owner         string
		NftType       string
	}

	blockchain := newBlockchain(t)
	standardContractAddress := deployStandardContracts(t, blockchain)
	veNftAddress, signer := deployVeNftContract(t, blockchain, standardContractAddress)

	setUpTestAccountToReceiveveNfts(t, blockchain, standardContractAddress, veNftAddress, signer)

	testName := "test name"
	testDescription := "test description"
	testThumbnail := "test thumbnail"
	mint(t, blockchain, standardContractAddress, veNftAddress, signer, veNftAddress, testName, testDescription, testThumbnail)

	veNftID := uint64(0)
	veNft := getveNft(t, blockchain, standardContractAddress, veNftAddress, veNftAddress, veNftID)
	veNftAsCadenceStruct := veNft.(cadence.Struct)

	var actualMetadata veNftMetadata
	for fieldIndex, field := range veNftAsCadenceStruct.StructType.Fields {
		fieldValue := veNftAsCadenceStruct.Fields[fieldIndex].ToGoValue()
		switch field.Identifier {
		case "veNftID":
			actualMetadata.veNftID = fieldValue.(uint64)
		case "resourceID":
			actualMetadata.ResourceID = fieldValue.(uint64)
		case "name":
			actualMetadata.Name = fieldValue.(string)
		case "description":
			actualMetadata.Description = fieldValue.(string)
		case "thumbnail":
			actualMetadata.Thumbnail = fieldValue.(string)
		case "owner":
			addressBytes := fieldValue.([8]uint8)
			addressString := cadence.NewAddress(addressBytes).Hex()
			actualMetadata.Owner = addressString
		case "type":
			actualMetadata.NftType = fieldValue.(string)
		}
	}

	expectedMetadata := veNftMetadata{
		veNftID: uint64(0),
		ResourceID:  uint64(31),
		Name:        testName,
		Description: testDescription,
		Thumbnail:   "ipfs://" + testThumbnail,
		Owner:       veNftAddress.Hex(),
		NftType:     "A." + veNftAddress.Hex() + ".VeNFT.NFT",
	}

	assert.Equal(t, actualMetadata, expectedMetadata)