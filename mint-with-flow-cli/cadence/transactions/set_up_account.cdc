import NonFungibleToken from "../contracts/standard/NonFungibleToken.cdc"
import VeNFT from "../contracts/VeNFT.cdc"

transaction {
    prepare(signer: AuthAccount) {
        if signer.borrow<&VeNFT.Collection>(from: VeNFT.CollectionStoragePath) == nil {
            // create a new empty collection
            let collection <- VeNFT.createEmptyCollection()
            
            // save it to the account
            signer.save(<- collection, to: VeNFT.CollectionStoragePath)

            // Creates a public capability for the collection so that other users can publicly access electable attributes.
            // The pieces inside of the brackets specify the type of the linked object, and only expose the fields and
            // functions on those types.
            signer.link<&VeNFT.Collection{NonFungibleToken.CollectionPublic, VeNFT.VeNFTCollectionPublic}>(
                VeNFT.CollectionPublicPath, target: VeNFT.CollectionStoragePath
            )
        }
    }
}