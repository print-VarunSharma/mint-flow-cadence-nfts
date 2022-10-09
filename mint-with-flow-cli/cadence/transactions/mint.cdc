import NonFungibleToken from "../contracts/standard/NonFungibleToken.cdc"
import VeNFT from "../contracts/VeNFT.cdc"

transaction(recipient: Address, name: String, description: String, thumbnail: String) {
    let minter: &VeNFT.NFTMinter

    prepare(signer: AuthAccount) {
        // borrow a reference to the NFTMinter resource in storage
        self.minter = signer.borrow<&VeNFT.NFTMinter>(from: VeNFT.MinterStoragePath)
            ?? panic("Could not borrow a reference to the NFT minter")
    }

    execute {
        // get the public account object for the recipient
        let recipient = getAccount(recipient)

        // borrow the recipient's public NFT collection reference
        let receiver = recipient
            .getCapability(VeNFT.CollectionPublicPath)!
            .borrow<&{NonFungibleToken.CollectionPublic}>()
            ?? panic("Could not get receiver reference to the NFT Collection")

        // mint the NFT and deposit it to the recipient's collection
        self.minter.mintNFT(
            recipient: receiver,
            name: name,
            description: description,
            thumbnail: thumbnail,
        )
    }
}