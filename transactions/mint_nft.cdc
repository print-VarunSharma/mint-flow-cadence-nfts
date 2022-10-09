import NonFungibleToken from 0x01
import VeNFT from 0x02

transaction(recipient: Address,name: String,ipfsLink: String) {
    let minter: &VeNFT.NFTMinter

    prepare(signer: AuthAccount) {
        self.minter = signer.borrow<&VeNFT.NFTMinter>(from: VeNFT.MinterStoragePath)
            ?? panic("Could not borrow a reference to the NFT minter")
    }

    execute {
        let recipient = getAccount(recipient)

        let receiver = recipient
            .getCapability(VeNFT.CollectionPublicPath)!
            .borrow<&{NonFungibleToken.CollectionPublic}>()
            ?? panic("Could not get receiver reference to the NFT Collection")

        self.minter.mintNFT(recipient: receiver, name: name,ipfsLink:ipfsLink)
    }
}