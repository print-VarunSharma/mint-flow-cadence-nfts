import NonFungibleToken from  0x01
import VeNFT from 0x02


pub fun main(address:Address) : [VeNFT.NftData] {
    let account = getAccount(address)
    let nft = VeNFT.getNft(address: address)
    return nft
}