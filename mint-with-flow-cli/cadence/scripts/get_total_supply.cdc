import VeNFT from "../contracts/VeNFT.cdc"

pub fun main(): UInt64 {
  return VeNFT.totalSupply
}