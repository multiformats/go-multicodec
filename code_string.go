// Code generated by "stringer -type=Code -linecomment"; DO NOT EDIT.

package multicodec

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Identity-0]
	_ = x[Cidv1-1]
	_ = x[Cidv2-2]
	_ = x[Cidv3-3]
	_ = x[Ip4-4]
	_ = x[Tcp-6]
	_ = x[Sha1-17]
	_ = x[Sha2_256-18]
	_ = x[Sha2_512-19]
	_ = x[Sha3_512-20]
	_ = x[Sha3_384-21]
	_ = x[Sha3_256-22]
	_ = x[Sha3_224-23]
	_ = x[Shake128-24]
	_ = x[Shake256-25]
	_ = x[Keccak224-26]
	_ = x[Keccak256-27]
	_ = x[Keccak384-28]
	_ = x[Keccak512-29]
	_ = x[Blake3-30]
	_ = x[Sha2_384-32]
	_ = x[Dccp-33]
	_ = x[Murmur3X64_64-34]
	_ = x[Murmur3_32-35]
	_ = x[Ip6-41]
	_ = x[Ip6zone-42]
	_ = x[Ipcidr-43]
	_ = x[Path-47]
	_ = x[Multicodec-48]
	_ = x[Multihash-49]
	_ = x[Multiaddr-50]
	_ = x[Multibase-51]
	_ = x[Varsig-52]
	_ = x[Dns-53]
	_ = x[Dns4-54]
	_ = x[Dns6-55]
	_ = x[Dnsaddr-56]
	_ = x[Protobuf-80]
	_ = x[Cbor-81]
	_ = x[Raw-85]
	_ = x[DblSha2_256-86]
	_ = x[Rlp-96]
	_ = x[Bencode-99]
	_ = x[DagPb-112]
	_ = x[DagCbor-113]
	_ = x[Libp2pKey-114]
	_ = x[GitRaw-120]
	_ = x[TorrentInfo-123]
	_ = x[TorrentFile-124]
	_ = x[LeofcoinBlock-129]
	_ = x[LeofcoinTx-130]
	_ = x[LeofcoinPr-131]
	_ = x[Sctp-132]
	_ = x[DagJose-133]
	_ = x[DagCose-134]
	_ = x[EthBlock-144]
	_ = x[EthBlockList-145]
	_ = x[EthTxTrie-146]
	_ = x[EthTx-147]
	_ = x[EthTxReceiptTrie-148]
	_ = x[EthTxReceipt-149]
	_ = x[EthStateTrie-150]
	_ = x[EthAccountSnapshot-151]
	_ = x[EthStorageTrie-152]
	_ = x[EthReceiptLogTrie-153]
	_ = x[EthRecieptLog-154]
	_ = x[Aes128-160]
	_ = x[Aes192-161]
	_ = x[Aes256-162]
	_ = x[Chacha128-163]
	_ = x[Chacha256-164]
	_ = x[BitcoinBlock-176]
	_ = x[BitcoinTx-177]
	_ = x[BitcoinWitnessCommitment-178]
	_ = x[ZcashBlock-192]
	_ = x[ZcashTx-193]
	_ = x[Caip50-202]
	_ = x[Streamid-206]
	_ = x[StellarBlock-208]
	_ = x[StellarTx-209]
	_ = x[Md4-212]
	_ = x[Md5-213]
	_ = x[DecredBlock-224]
	_ = x[DecredTx-225]
	_ = x[Ipld-226]
	_ = x[Ipfs-227]
	_ = x[Swarm-228]
	_ = x[Ipns-229]
	_ = x[Zeronet-230]
	_ = x[Secp256k1Pub-231]
	_ = x[Dnslink-232]
	_ = x[Bls12_381G1Pub-234]
	_ = x[Bls12_381G2Pub-235]
	_ = x[X25519Pub-236]
	_ = x[Ed25519Pub-237]
	_ = x[Bls12_381G1g2Pub-238]
	_ = x[Sr25519Pub-239]
	_ = x[DashBlock-240]
	_ = x[DashTx-241]
	_ = x[SwarmManifest-250]
	_ = x[SwarmFeed-251]
	_ = x[Beeson-252]
	_ = x[Udp-273]
	_ = x[P2pWebrtcStar-275]
	_ = x[P2pWebrtcDirect-276]
	_ = x[P2pStardust-277]
	_ = x[WebrtcDirect-280]
	_ = x[Webrtc-281]
	_ = x[P2pCircuit-290]
	_ = x[DagJson-297]
	_ = x[Udt-301]
	_ = x[Utp-302]
	_ = x[Crc32-306]
	_ = x[Crc64Ecma-356]
	_ = x[Unix-400]
	_ = x[Thread-406]
	_ = x[P2p-421]
	_ = x[Https-443]
	_ = x[Onion-444]
	_ = x[Onion3-445]
	_ = x[Garlic64-446]
	_ = x[Garlic32-447]
	_ = x[Tls-448]
	_ = x[Sni-449]
	_ = x[Noise-454]
	_ = x[Quic-460]
	_ = x[QuicV1-461]
	_ = x[Webtransport-465]
	_ = x[Certhash-466]
	_ = x[Ws-477]
	_ = x[Wss-478]
	_ = x[P2pWebsocketStar-479]
	_ = x[Http-480]
	_ = x[Swhid1Snp-496]
	_ = x[Json-512]
	_ = x[Messagepack-513]
	_ = x[Car-514]
	_ = x[IpnsRecord-768]
	_ = x[Libp2pPeerRecord-769]
	_ = x[Libp2pRelayRsvp-770]
	_ = x[Memorytransport-777]
	_ = x[CarIndexSorted-1024]
	_ = x[CarMultihashIndexSorted-1025]
	_ = x[TransportBitswap-2304]
	_ = x[TransportGraphsyncFilecoinv1-2320]
	_ = x[TransportIpfsGatewayHttp-2336]
	_ = x[Multidid-3357]
	_ = x[Sha2_256Trunc254Padded-4114]
	_ = x[Sha2_224-4115]
	_ = x[Sha2_512_224-4116]
	_ = x[Sha2_512_256-4117]
	_ = x[Murmur3X64_128-4130]
	_ = x[Ripemd128-4178]
	_ = x[Ripemd160-4179]
	_ = x[Ripemd256-4180]
	_ = x[Ripemd320-4181]
	_ = x[X11-4352]
	_ = x[P256Pub-4608]
	_ = x[P384Pub-4609]
	_ = x[P521Pub-4610]
	_ = x[Ed448Pub-4611]
	_ = x[X448Pub-4612]
	_ = x[RsaPub-4613]
	_ = x[Sm2Pub-4614]
	_ = x[Ed25519Priv-4864]
	_ = x[Secp256k1Priv-4865]
	_ = x[X25519Priv-4866]
	_ = x[Sr25519Priv-4867]
	_ = x[RsaPriv-4869]
	_ = x[P256Priv-4870]
	_ = x[P384Priv-4871]
	_ = x[P521Priv-4872]
	_ = x[Kangarootwelve-7425]
	_ = x[AesGcm256-8192]
	_ = x[Silverpine-16194]
	_ = x[Sm3_256-21325]
	_ = x[Blake2b8-45569]
	_ = x[Blake2b16-45570]
	_ = x[Blake2b24-45571]
	_ = x[Blake2b32-45572]
	_ = x[Blake2b40-45573]
	_ = x[Blake2b48-45574]
	_ = x[Blake2b56-45575]
	_ = x[Blake2b64-45576]
	_ = x[Blake2b72-45577]
	_ = x[Blake2b80-45578]
	_ = x[Blake2b88-45579]
	_ = x[Blake2b96-45580]
	_ = x[Blake2b104-45581]
	_ = x[Blake2b112-45582]
	_ = x[Blake2b120-45583]
	_ = x[Blake2b128-45584]
	_ = x[Blake2b136-45585]
	_ = x[Blake2b144-45586]
	_ = x[Blake2b152-45587]
	_ = x[Blake2b160-45588]
	_ = x[Blake2b168-45589]
	_ = x[Blake2b176-45590]
	_ = x[Blake2b184-45591]
	_ = x[Blake2b192-45592]
	_ = x[Blake2b200-45593]
	_ = x[Blake2b208-45594]
	_ = x[Blake2b216-45595]
	_ = x[Blake2b224-45596]
	_ = x[Blake2b232-45597]
	_ = x[Blake2b240-45598]
	_ = x[Blake2b248-45599]
	_ = x[Blake2b256-45600]
	_ = x[Blake2b264-45601]
	_ = x[Blake2b272-45602]
	_ = x[Blake2b280-45603]
	_ = x[Blake2b288-45604]
	_ = x[Blake2b296-45605]
	_ = x[Blake2b304-45606]
	_ = x[Blake2b312-45607]
	_ = x[Blake2b320-45608]
	_ = x[Blake2b328-45609]
	_ = x[Blake2b336-45610]
	_ = x[Blake2b344-45611]
	_ = x[Blake2b352-45612]
	_ = x[Blake2b360-45613]
	_ = x[Blake2b368-45614]
	_ = x[Blake2b376-45615]
	_ = x[Blake2b384-45616]
	_ = x[Blake2b392-45617]
	_ = x[Blake2b400-45618]
	_ = x[Blake2b408-45619]
	_ = x[Blake2b416-45620]
	_ = x[Blake2b424-45621]
	_ = x[Blake2b432-45622]
	_ = x[Blake2b440-45623]
	_ = x[Blake2b448-45624]
	_ = x[Blake2b456-45625]
	_ = x[Blake2b464-45626]
	_ = x[Blake2b472-45627]
	_ = x[Blake2b480-45628]
	_ = x[Blake2b488-45629]
	_ = x[Blake2b496-45630]
	_ = x[Blake2b504-45631]
	_ = x[Blake2b512-45632]
	_ = x[Blake2s8-45633]
	_ = x[Blake2s16-45634]
	_ = x[Blake2s24-45635]
	_ = x[Blake2s32-45636]
	_ = x[Blake2s40-45637]
	_ = x[Blake2s48-45638]
	_ = x[Blake2s56-45639]
	_ = x[Blake2s64-45640]
	_ = x[Blake2s72-45641]
	_ = x[Blake2s80-45642]
	_ = x[Blake2s88-45643]
	_ = x[Blake2s96-45644]
	_ = x[Blake2s104-45645]
	_ = x[Blake2s112-45646]
	_ = x[Blake2s120-45647]
	_ = x[Blake2s128-45648]
	_ = x[Blake2s136-45649]
	_ = x[Blake2s144-45650]
	_ = x[Blake2s152-45651]
	_ = x[Blake2s160-45652]
	_ = x[Blake2s168-45653]
	_ = x[Blake2s176-45654]
	_ = x[Blake2s184-45655]
	_ = x[Blake2s192-45656]
	_ = x[Blake2s200-45657]
	_ = x[Blake2s208-45658]
	_ = x[Blake2s216-45659]
	_ = x[Blake2s224-45660]
	_ = x[Blake2s232-45661]
	_ = x[Blake2s240-45662]
	_ = x[Blake2s248-45663]
	_ = x[Blake2s256-45664]
	_ = x[Skein256_8-45825]
	_ = x[Skein256_16-45826]
	_ = x[Skein256_24-45827]
	_ = x[Skein256_32-45828]
	_ = x[Skein256_40-45829]
	_ = x[Skein256_48-45830]
	_ = x[Skein256_56-45831]
	_ = x[Skein256_64-45832]
	_ = x[Skein256_72-45833]
	_ = x[Skein256_80-45834]
	_ = x[Skein256_88-45835]
	_ = x[Skein256_96-45836]
	_ = x[Skein256_104-45837]
	_ = x[Skein256_112-45838]
	_ = x[Skein256_120-45839]
	_ = x[Skein256_128-45840]
	_ = x[Skein256_136-45841]
	_ = x[Skein256_144-45842]
	_ = x[Skein256_152-45843]
	_ = x[Skein256_160-45844]
	_ = x[Skein256_168-45845]
	_ = x[Skein256_176-45846]
	_ = x[Skein256_184-45847]
	_ = x[Skein256_192-45848]
	_ = x[Skein256_200-45849]
	_ = x[Skein256_208-45850]
	_ = x[Skein256_216-45851]
	_ = x[Skein256_224-45852]
	_ = x[Skein256_232-45853]
	_ = x[Skein256_240-45854]
	_ = x[Skein256_248-45855]
	_ = x[Skein256_256-45856]
	_ = x[Skein512_8-45857]
	_ = x[Skein512_16-45858]
	_ = x[Skein512_24-45859]
	_ = x[Skein512_32-45860]
	_ = x[Skein512_40-45861]
	_ = x[Skein512_48-45862]
	_ = x[Skein512_56-45863]
	_ = x[Skein512_64-45864]
	_ = x[Skein512_72-45865]
	_ = x[Skein512_80-45866]
	_ = x[Skein512_88-45867]
	_ = x[Skein512_96-45868]
	_ = x[Skein512_104-45869]
	_ = x[Skein512_112-45870]
	_ = x[Skein512_120-45871]
	_ = x[Skein512_128-45872]
	_ = x[Skein512_136-45873]
	_ = x[Skein512_144-45874]
	_ = x[Skein512_152-45875]
	_ = x[Skein512_160-45876]
	_ = x[Skein512_168-45877]
	_ = x[Skein512_176-45878]
	_ = x[Skein512_184-45879]
	_ = x[Skein512_192-45880]
	_ = x[Skein512_200-45881]
	_ = x[Skein512_208-45882]
	_ = x[Skein512_216-45883]
	_ = x[Skein512_224-45884]
	_ = x[Skein512_232-45885]
	_ = x[Skein512_240-45886]
	_ = x[Skein512_248-45887]
	_ = x[Skein512_256-45888]
	_ = x[Skein512_264-45889]
	_ = x[Skein512_272-45890]
	_ = x[Skein512_280-45891]
	_ = x[Skein512_288-45892]
	_ = x[Skein512_296-45893]
	_ = x[Skein512_304-45894]
	_ = x[Skein512_312-45895]
	_ = x[Skein512_320-45896]
	_ = x[Skein512_328-45897]
	_ = x[Skein512_336-45898]
	_ = x[Skein512_344-45899]
	_ = x[Skein512_352-45900]
	_ = x[Skein512_360-45901]
	_ = x[Skein512_368-45902]
	_ = x[Skein512_376-45903]
	_ = x[Skein512_384-45904]
	_ = x[Skein512_392-45905]
	_ = x[Skein512_400-45906]
	_ = x[Skein512_408-45907]
	_ = x[Skein512_416-45908]
	_ = x[Skein512_424-45909]
	_ = x[Skein512_432-45910]
	_ = x[Skein512_440-45911]
	_ = x[Skein512_448-45912]
	_ = x[Skein512_456-45913]
	_ = x[Skein512_464-45914]
	_ = x[Skein512_472-45915]
	_ = x[Skein512_480-45916]
	_ = x[Skein512_488-45917]
	_ = x[Skein512_496-45918]
	_ = x[Skein512_504-45919]
	_ = x[Skein512_512-45920]
	_ = x[Skein1024_8-45921]
	_ = x[Skein1024_16-45922]
	_ = x[Skein1024_24-45923]
	_ = x[Skein1024_32-45924]
	_ = x[Skein1024_40-45925]
	_ = x[Skein1024_48-45926]
	_ = x[Skein1024_56-45927]
	_ = x[Skein1024_64-45928]
	_ = x[Skein1024_72-45929]
	_ = x[Skein1024_80-45930]
	_ = x[Skein1024_88-45931]
	_ = x[Skein1024_96-45932]
	_ = x[Skein1024_104-45933]
	_ = x[Skein1024_112-45934]
	_ = x[Skein1024_120-45935]
	_ = x[Skein1024_128-45936]
	_ = x[Skein1024_136-45937]
	_ = x[Skein1024_144-45938]
	_ = x[Skein1024_152-45939]
	_ = x[Skein1024_160-45940]
	_ = x[Skein1024_168-45941]
	_ = x[Skein1024_176-45942]
	_ = x[Skein1024_184-45943]
	_ = x[Skein1024_192-45944]
	_ = x[Skein1024_200-45945]
	_ = x[Skein1024_208-45946]
	_ = x[Skein1024_216-45947]
	_ = x[Skein1024_224-45948]
	_ = x[Skein1024_232-45949]
	_ = x[Skein1024_240-45950]
	_ = x[Skein1024_248-45951]
	_ = x[Skein1024_256-45952]
	_ = x[Skein1024_264-45953]
	_ = x[Skein1024_272-45954]
	_ = x[Skein1024_280-45955]
	_ = x[Skein1024_288-45956]
	_ = x[Skein1024_296-45957]
	_ = x[Skein1024_304-45958]
	_ = x[Skein1024_312-45959]
	_ = x[Skein1024_320-45960]
	_ = x[Skein1024_328-45961]
	_ = x[Skein1024_336-45962]
	_ = x[Skein1024_344-45963]
	_ = x[Skein1024_352-45964]
	_ = x[Skein1024_360-45965]
	_ = x[Skein1024_368-45966]
	_ = x[Skein1024_376-45967]
	_ = x[Skein1024_384-45968]
	_ = x[Skein1024_392-45969]
	_ = x[Skein1024_400-45970]
	_ = x[Skein1024_408-45971]
	_ = x[Skein1024_416-45972]
	_ = x[Skein1024_424-45973]
	_ = x[Skein1024_432-45974]
	_ = x[Skein1024_440-45975]
	_ = x[Skein1024_448-45976]
	_ = x[Skein1024_456-45977]
	_ = x[Skein1024_464-45978]
	_ = x[Skein1024_472-45979]
	_ = x[Skein1024_480-45980]
	_ = x[Skein1024_488-45981]
	_ = x[Skein1024_496-45982]
	_ = x[Skein1024_504-45983]
	_ = x[Skein1024_512-45984]
	_ = x[Skein1024_520-45985]
	_ = x[Skein1024_528-45986]
	_ = x[Skein1024_536-45987]
	_ = x[Skein1024_544-45988]
	_ = x[Skein1024_552-45989]
	_ = x[Skein1024_560-45990]
	_ = x[Skein1024_568-45991]
	_ = x[Skein1024_576-45992]
	_ = x[Skein1024_584-45993]
	_ = x[Skein1024_592-45994]
	_ = x[Skein1024_600-45995]
	_ = x[Skein1024_608-45996]
	_ = x[Skein1024_616-45997]
	_ = x[Skein1024_624-45998]
	_ = x[Skein1024_632-45999]
	_ = x[Skein1024_640-46000]
	_ = x[Skein1024_648-46001]
	_ = x[Skein1024_656-46002]
	_ = x[Skein1024_664-46003]
	_ = x[Skein1024_672-46004]
	_ = x[Skein1024_680-46005]
	_ = x[Skein1024_688-46006]
	_ = x[Skein1024_696-46007]
	_ = x[Skein1024_704-46008]
	_ = x[Skein1024_712-46009]
	_ = x[Skein1024_720-46010]
	_ = x[Skein1024_728-46011]
	_ = x[Skein1024_736-46012]
	_ = x[Skein1024_744-46013]
	_ = x[Skein1024_752-46014]
	_ = x[Skein1024_760-46015]
	_ = x[Skein1024_768-46016]
	_ = x[Skein1024_776-46017]
	_ = x[Skein1024_784-46018]
	_ = x[Skein1024_792-46019]
	_ = x[Skein1024_800-46020]
	_ = x[Skein1024_808-46021]
	_ = x[Skein1024_816-46022]
	_ = x[Skein1024_824-46023]
	_ = x[Skein1024_832-46024]
	_ = x[Skein1024_840-46025]
	_ = x[Skein1024_848-46026]
	_ = x[Skein1024_856-46027]
	_ = x[Skein1024_864-46028]
	_ = x[Skein1024_872-46029]
	_ = x[Skein1024_880-46030]
	_ = x[Skein1024_888-46031]
	_ = x[Skein1024_896-46032]
	_ = x[Skein1024_904-46033]
	_ = x[Skein1024_912-46034]
	_ = x[Skein1024_920-46035]
	_ = x[Skein1024_928-46036]
	_ = x[Skein1024_936-46037]
	_ = x[Skein1024_944-46038]
	_ = x[Skein1024_952-46039]
	_ = x[Skein1024_960-46040]
	_ = x[Skein1024_968-46041]
	_ = x[Skein1024_976-46042]
	_ = x[Skein1024_984-46043]
	_ = x[Skein1024_992-46044]
	_ = x[Skein1024_1000-46045]
	_ = x[Skein1024_1008-46046]
	_ = x[Skein1024_1016-46047]
	_ = x[Skein1024_1024-46048]
	_ = x[Xxh32-46049]
	_ = x[Xxh64-46050]
	_ = x[Xxh3_64-46051]
	_ = x[Xxh3_128-46052]
	_ = x[PoseidonBls12_381A2Fc1-46081]
	_ = x[PoseidonBls12_381A2Fc1Sc-46082]
	_ = x[Urdca2015Canon-46083]
	_ = x[Ssz-46337]
	_ = x[SszSha2_256Bmt-46338]
	_ = x[JsonJcs-46593]
	_ = x[Iscc-52225]
	_ = x[ZeroxcertImprint256-52753]
	_ = x[NonstandardSig-53248]
	_ = x[Es256k-53479]
	_ = x[Bls12381G1Sig-53482]
	_ = x[Bls12381G2Sig-53483]
	_ = x[Eddsa-53485]
	_ = x[Eip191-53649]
	_ = x[Jwk_jcsPub-60241]
	_ = x[FilCommitmentUnsealed-61697]
	_ = x[FilCommitmentSealed-61698]
	_ = x[Plaintextv2-7367777]
	_ = x[HolochainAdrV0-8417572]
	_ = x[HolochainAdrV1-8483108]
	_ = x[HolochainKeyV0-9728292]
	_ = x[HolochainKeyV1-9793828]
	_ = x[HolochainSigV0-10645796]
	_ = x[HolochainSigV1-10711332]
	_ = x[SkynetNs-11639056]
	_ = x[ArweaveNs-11704592]
	_ = x[SubspaceNs-11770128]
	_ = x[KumandraNs-11835664]
	_ = x[Es256-13636096]
	_ = x[Es284-13636097]
	_ = x[Es512-13636098]
	_ = x[Rs256-13636101]
}

const _Code_name = "identitycidv1cidv2cidv3ip4tcpsha1sha2-256sha2-512sha3-512sha3-384sha3-256sha3-224shake-128shake-256keccak-224keccak-256keccak-384keccak-512blake3sha2-384dccpmurmur3-x64-64murmur3-32ip6ip6zoneipcidrpathmulticodecmultihashmultiaddrmultibasevarsigdnsdns4dns6dnsaddrprotobufcborrawdbl-sha2-256rlpbencodedag-pbdag-cborlibp2p-keygit-rawtorrent-infotorrent-fileleofcoin-blockleofcoin-txleofcoin-prsctpdag-josedag-coseeth-blocketh-block-listeth-tx-trieeth-txeth-tx-receipt-trieeth-tx-receipteth-state-trieeth-account-snapshoteth-storage-trieeth-receipt-log-trieeth-reciept-logaes-128aes-192aes-256chacha-128chacha-256bitcoin-blockbitcoin-txbitcoin-witness-commitmentzcash-blockzcash-txcaip-50streamidstellar-blockstellar-txmd4md5decred-blockdecred-txipldipfsswarmipnszeronetsecp256k1-pubdnslinkbls12_381-g1-pubbls12_381-g2-pubx25519-pubed25519-pubbls12_381-g1g2-pubsr25519-pubdash-blockdash-txswarm-manifestswarm-feedbeesonudpp2p-webrtc-starp2p-webrtc-directp2p-stardustwebrtc-directwebrtcp2p-circuitdag-jsonudtutpcrc32crc64-ecmaunixthreadp2phttpsoniononion3garlic64garlic32tlssninoisequicquic-v1webtransportcerthashwswssp2p-websocket-starhttpswhid-1-snpjsonmessagepackcaripns-recordlibp2p-peer-recordlibp2p-relay-rsvpmemorytransportcar-index-sortedcar-multihash-index-sortedtransport-bitswaptransport-graphsync-filecoinv1transport-ipfs-gateway-httpmultididsha2-256-trunc254-paddedsha2-224sha2-512-224sha2-512-256murmur3-x64-128ripemd-128ripemd-160ripemd-256ripemd-320x11p256-pubp384-pubp521-pubed448-pubx448-pubrsa-pubsm2-pubed25519-privsecp256k1-privx25519-privsr25519-privrsa-privp256-privp384-privp521-privkangarootwelveaes-gcm-256silverpinesm3-256blake2b-8blake2b-16blake2b-24blake2b-32blake2b-40blake2b-48blake2b-56blake2b-64blake2b-72blake2b-80blake2b-88blake2b-96blake2b-104blake2b-112blake2b-120blake2b-128blake2b-136blake2b-144blake2b-152blake2b-160blake2b-168blake2b-176blake2b-184blake2b-192blake2b-200blake2b-208blake2b-216blake2b-224blake2b-232blake2b-240blake2b-248blake2b-256blake2b-264blake2b-272blake2b-280blake2b-288blake2b-296blake2b-304blake2b-312blake2b-320blake2b-328blake2b-336blake2b-344blake2b-352blake2b-360blake2b-368blake2b-376blake2b-384blake2b-392blake2b-400blake2b-408blake2b-416blake2b-424blake2b-432blake2b-440blake2b-448blake2b-456blake2b-464blake2b-472blake2b-480blake2b-488blake2b-496blake2b-504blake2b-512blake2s-8blake2s-16blake2s-24blake2s-32blake2s-40blake2s-48blake2s-56blake2s-64blake2s-72blake2s-80blake2s-88blake2s-96blake2s-104blake2s-112blake2s-120blake2s-128blake2s-136blake2s-144blake2s-152blake2s-160blake2s-168blake2s-176blake2s-184blake2s-192blake2s-200blake2s-208blake2s-216blake2s-224blake2s-232blake2s-240blake2s-248blake2s-256skein256-8skein256-16skein256-24skein256-32skein256-40skein256-48skein256-56skein256-64skein256-72skein256-80skein256-88skein256-96skein256-104skein256-112skein256-120skein256-128skein256-136skein256-144skein256-152skein256-160skein256-168skein256-176skein256-184skein256-192skein256-200skein256-208skein256-216skein256-224skein256-232skein256-240skein256-248skein256-256skein512-8skein512-16skein512-24skein512-32skein512-40skein512-48skein512-56skein512-64skein512-72skein512-80skein512-88skein512-96skein512-104skein512-112skein512-120skein512-128skein512-136skein512-144skein512-152skein512-160skein512-168skein512-176skein512-184skein512-192skein512-200skein512-208skein512-216skein512-224skein512-232skein512-240skein512-248skein512-256skein512-264skein512-272skein512-280skein512-288skein512-296skein512-304skein512-312skein512-320skein512-328skein512-336skein512-344skein512-352skein512-360skein512-368skein512-376skein512-384skein512-392skein512-400skein512-408skein512-416skein512-424skein512-432skein512-440skein512-448skein512-456skein512-464skein512-472skein512-480skein512-488skein512-496skein512-504skein512-512skein1024-8skein1024-16skein1024-24skein1024-32skein1024-40skein1024-48skein1024-56skein1024-64skein1024-72skein1024-80skein1024-88skein1024-96skein1024-104skein1024-112skein1024-120skein1024-128skein1024-136skein1024-144skein1024-152skein1024-160skein1024-168skein1024-176skein1024-184skein1024-192skein1024-200skein1024-208skein1024-216skein1024-224skein1024-232skein1024-240skein1024-248skein1024-256skein1024-264skein1024-272skein1024-280skein1024-288skein1024-296skein1024-304skein1024-312skein1024-320skein1024-328skein1024-336skein1024-344skein1024-352skein1024-360skein1024-368skein1024-376skein1024-384skein1024-392skein1024-400skein1024-408skein1024-416skein1024-424skein1024-432skein1024-440skein1024-448skein1024-456skein1024-464skein1024-472skein1024-480skein1024-488skein1024-496skein1024-504skein1024-512skein1024-520skein1024-528skein1024-536skein1024-544skein1024-552skein1024-560skein1024-568skein1024-576skein1024-584skein1024-592skein1024-600skein1024-608skein1024-616skein1024-624skein1024-632skein1024-640skein1024-648skein1024-656skein1024-664skein1024-672skein1024-680skein1024-688skein1024-696skein1024-704skein1024-712skein1024-720skein1024-728skein1024-736skein1024-744skein1024-752skein1024-760skein1024-768skein1024-776skein1024-784skein1024-792skein1024-800skein1024-808skein1024-816skein1024-824skein1024-832skein1024-840skein1024-848skein1024-856skein1024-864skein1024-872skein1024-880skein1024-888skein1024-896skein1024-904skein1024-912skein1024-920skein1024-928skein1024-936skein1024-944skein1024-952skein1024-960skein1024-968skein1024-976skein1024-984skein1024-992skein1024-1000skein1024-1008skein1024-1016skein1024-1024xxh-32xxh-64xxh3-64xxh3-128poseidon-bls12_381-a2-fc1poseidon-bls12_381-a2-fc1-scurdca-2015-canonsszssz-sha2-256-bmtjson-jcsiscczeroxcert-imprint-256nonstandard-siges256kbls-12381-g1-sigbls-12381-g2-sigeddsaeip-191jwk_jcs-pubfil-commitment-unsealedfil-commitment-sealedplaintextv2holochain-adr-v0holochain-adr-v1holochain-key-v0holochain-key-v1holochain-sig-v0holochain-sig-v1skynet-nsarweave-nssubspace-nskumandra-nses256es284es512rs256"

var _Code_map = map[Code]string{
	0:        _Code_name[0:8],
	1:        _Code_name[8:13],
	2:        _Code_name[13:18],
	3:        _Code_name[18:23],
	4:        _Code_name[23:26],
	6:        _Code_name[26:29],
	17:       _Code_name[29:33],
	18:       _Code_name[33:41],
	19:       _Code_name[41:49],
	20:       _Code_name[49:57],
	21:       _Code_name[57:65],
	22:       _Code_name[65:73],
	23:       _Code_name[73:81],
	24:       _Code_name[81:90],
	25:       _Code_name[90:99],
	26:       _Code_name[99:109],
	27:       _Code_name[109:119],
	28:       _Code_name[119:129],
	29:       _Code_name[129:139],
	30:       _Code_name[139:145],
	32:       _Code_name[145:153],
	33:       _Code_name[153:157],
	34:       _Code_name[157:171],
	35:       _Code_name[171:181],
	41:       _Code_name[181:184],
	42:       _Code_name[184:191],
	43:       _Code_name[191:197],
	47:       _Code_name[197:201],
	48:       _Code_name[201:211],
	49:       _Code_name[211:220],
	50:       _Code_name[220:229],
	51:       _Code_name[229:238],
	52:       _Code_name[238:244],
	53:       _Code_name[244:247],
	54:       _Code_name[247:251],
	55:       _Code_name[251:255],
	56:       _Code_name[255:262],
	80:       _Code_name[262:270],
	81:       _Code_name[270:274],
	85:       _Code_name[274:277],
	86:       _Code_name[277:289],
	96:       _Code_name[289:292],
	99:       _Code_name[292:299],
	112:      _Code_name[299:305],
	113:      _Code_name[305:313],
	114:      _Code_name[313:323],
	120:      _Code_name[323:330],
	123:      _Code_name[330:342],
	124:      _Code_name[342:354],
	129:      _Code_name[354:368],
	130:      _Code_name[368:379],
	131:      _Code_name[379:390],
	132:      _Code_name[390:394],
	133:      _Code_name[394:402],
	134:      _Code_name[402:410],
	144:      _Code_name[410:419],
	145:      _Code_name[419:433],
	146:      _Code_name[433:444],
	147:      _Code_name[444:450],
	148:      _Code_name[450:469],
	149:      _Code_name[469:483],
	150:      _Code_name[483:497],
	151:      _Code_name[497:517],
	152:      _Code_name[517:533],
	153:      _Code_name[533:553],
	154:      _Code_name[553:568],
	160:      _Code_name[568:575],
	161:      _Code_name[575:582],
	162:      _Code_name[582:589],
	163:      _Code_name[589:599],
	164:      _Code_name[599:609],
	176:      _Code_name[609:622],
	177:      _Code_name[622:632],
	178:      _Code_name[632:658],
	192:      _Code_name[658:669],
	193:      _Code_name[669:677],
	202:      _Code_name[677:684],
	206:      _Code_name[684:692],
	208:      _Code_name[692:705],
	209:      _Code_name[705:715],
	212:      _Code_name[715:718],
	213:      _Code_name[718:721],
	224:      _Code_name[721:733],
	225:      _Code_name[733:742],
	226:      _Code_name[742:746],
	227:      _Code_name[746:750],
	228:      _Code_name[750:755],
	229:      _Code_name[755:759],
	230:      _Code_name[759:766],
	231:      _Code_name[766:779],
	232:      _Code_name[779:786],
	234:      _Code_name[786:802],
	235:      _Code_name[802:818],
	236:      _Code_name[818:828],
	237:      _Code_name[828:839],
	238:      _Code_name[839:857],
	239:      _Code_name[857:868],
	240:      _Code_name[868:878],
	241:      _Code_name[878:885],
	250:      _Code_name[885:899],
	251:      _Code_name[899:909],
	252:      _Code_name[909:915],
	273:      _Code_name[915:918],
	275:      _Code_name[918:933],
	276:      _Code_name[933:950],
	277:      _Code_name[950:962],
	280:      _Code_name[962:975],
	281:      _Code_name[975:981],
	290:      _Code_name[981:992],
	297:      _Code_name[992:1000],
	301:      _Code_name[1000:1003],
	302:      _Code_name[1003:1006],
	306:      _Code_name[1006:1011],
	356:      _Code_name[1011:1021],
	400:      _Code_name[1021:1025],
	406:      _Code_name[1025:1031],
	421:      _Code_name[1031:1034],
	443:      _Code_name[1034:1039],
	444:      _Code_name[1039:1044],
	445:      _Code_name[1044:1050],
	446:      _Code_name[1050:1058],
	447:      _Code_name[1058:1066],
	448:      _Code_name[1066:1069],
	449:      _Code_name[1069:1072],
	454:      _Code_name[1072:1077],
	460:      _Code_name[1077:1081],
	461:      _Code_name[1081:1088],
	465:      _Code_name[1088:1100],
	466:      _Code_name[1100:1108],
	477:      _Code_name[1108:1110],
	478:      _Code_name[1110:1113],
	479:      _Code_name[1113:1131],
	480:      _Code_name[1131:1135],
	496:      _Code_name[1135:1146],
	512:      _Code_name[1146:1150],
	513:      _Code_name[1150:1161],
	514:      _Code_name[1161:1164],
	768:      _Code_name[1164:1175],
	769:      _Code_name[1175:1193],
	770:      _Code_name[1193:1210],
	777:      _Code_name[1210:1225],
	1024:     _Code_name[1225:1241],
	1025:     _Code_name[1241:1267],
	2304:     _Code_name[1267:1284],
	2320:     _Code_name[1284:1314],
	2336:     _Code_name[1314:1341],
	3357:     _Code_name[1341:1349],
	4114:     _Code_name[1349:1373],
	4115:     _Code_name[1373:1381],
	4116:     _Code_name[1381:1393],
	4117:     _Code_name[1393:1405],
	4130:     _Code_name[1405:1420],
	4178:     _Code_name[1420:1430],
	4179:     _Code_name[1430:1440],
	4180:     _Code_name[1440:1450],
	4181:     _Code_name[1450:1460],
	4352:     _Code_name[1460:1463],
	4608:     _Code_name[1463:1471],
	4609:     _Code_name[1471:1479],
	4610:     _Code_name[1479:1487],
	4611:     _Code_name[1487:1496],
	4612:     _Code_name[1496:1504],
	4613:     _Code_name[1504:1511],
	4614:     _Code_name[1511:1518],
	4864:     _Code_name[1518:1530],
	4865:     _Code_name[1530:1544],
	4866:     _Code_name[1544:1555],
	4867:     _Code_name[1555:1567],
	4869:     _Code_name[1567:1575],
	4870:     _Code_name[1575:1584],
	4871:     _Code_name[1584:1593],
	4872:     _Code_name[1593:1602],
	7425:     _Code_name[1602:1616],
	8192:     _Code_name[1616:1627],
	16194:    _Code_name[1627:1637],
	21325:    _Code_name[1637:1644],
	45569:    _Code_name[1644:1653],
	45570:    _Code_name[1653:1663],
	45571:    _Code_name[1663:1673],
	45572:    _Code_name[1673:1683],
	45573:    _Code_name[1683:1693],
	45574:    _Code_name[1693:1703],
	45575:    _Code_name[1703:1713],
	45576:    _Code_name[1713:1723],
	45577:    _Code_name[1723:1733],
	45578:    _Code_name[1733:1743],
	45579:    _Code_name[1743:1753],
	45580:    _Code_name[1753:1763],
	45581:    _Code_name[1763:1774],
	45582:    _Code_name[1774:1785],
	45583:    _Code_name[1785:1796],
	45584:    _Code_name[1796:1807],
	45585:    _Code_name[1807:1818],
	45586:    _Code_name[1818:1829],
	45587:    _Code_name[1829:1840],
	45588:    _Code_name[1840:1851],
	45589:    _Code_name[1851:1862],
	45590:    _Code_name[1862:1873],
	45591:    _Code_name[1873:1884],
	45592:    _Code_name[1884:1895],
	45593:    _Code_name[1895:1906],
	45594:    _Code_name[1906:1917],
	45595:    _Code_name[1917:1928],
	45596:    _Code_name[1928:1939],
	45597:    _Code_name[1939:1950],
	45598:    _Code_name[1950:1961],
	45599:    _Code_name[1961:1972],
	45600:    _Code_name[1972:1983],
	45601:    _Code_name[1983:1994],
	45602:    _Code_name[1994:2005],
	45603:    _Code_name[2005:2016],
	45604:    _Code_name[2016:2027],
	45605:    _Code_name[2027:2038],
	45606:    _Code_name[2038:2049],
	45607:    _Code_name[2049:2060],
	45608:    _Code_name[2060:2071],
	45609:    _Code_name[2071:2082],
	45610:    _Code_name[2082:2093],
	45611:    _Code_name[2093:2104],
	45612:    _Code_name[2104:2115],
	45613:    _Code_name[2115:2126],
	45614:    _Code_name[2126:2137],
	45615:    _Code_name[2137:2148],
	45616:    _Code_name[2148:2159],
	45617:    _Code_name[2159:2170],
	45618:    _Code_name[2170:2181],
	45619:    _Code_name[2181:2192],
	45620:    _Code_name[2192:2203],
	45621:    _Code_name[2203:2214],
	45622:    _Code_name[2214:2225],
	45623:    _Code_name[2225:2236],
	45624:    _Code_name[2236:2247],
	45625:    _Code_name[2247:2258],
	45626:    _Code_name[2258:2269],
	45627:    _Code_name[2269:2280],
	45628:    _Code_name[2280:2291],
	45629:    _Code_name[2291:2302],
	45630:    _Code_name[2302:2313],
	45631:    _Code_name[2313:2324],
	45632:    _Code_name[2324:2335],
	45633:    _Code_name[2335:2344],
	45634:    _Code_name[2344:2354],
	45635:    _Code_name[2354:2364],
	45636:    _Code_name[2364:2374],
	45637:    _Code_name[2374:2384],
	45638:    _Code_name[2384:2394],
	45639:    _Code_name[2394:2404],
	45640:    _Code_name[2404:2414],
	45641:    _Code_name[2414:2424],
	45642:    _Code_name[2424:2434],
	45643:    _Code_name[2434:2444],
	45644:    _Code_name[2444:2454],
	45645:    _Code_name[2454:2465],
	45646:    _Code_name[2465:2476],
	45647:    _Code_name[2476:2487],
	45648:    _Code_name[2487:2498],
	45649:    _Code_name[2498:2509],
	45650:    _Code_name[2509:2520],
	45651:    _Code_name[2520:2531],
	45652:    _Code_name[2531:2542],
	45653:    _Code_name[2542:2553],
	45654:    _Code_name[2553:2564],
	45655:    _Code_name[2564:2575],
	45656:    _Code_name[2575:2586],
	45657:    _Code_name[2586:2597],
	45658:    _Code_name[2597:2608],
	45659:    _Code_name[2608:2619],
	45660:    _Code_name[2619:2630],
	45661:    _Code_name[2630:2641],
	45662:    _Code_name[2641:2652],
	45663:    _Code_name[2652:2663],
	45664:    _Code_name[2663:2674],
	45825:    _Code_name[2674:2684],
	45826:    _Code_name[2684:2695],
	45827:    _Code_name[2695:2706],
	45828:    _Code_name[2706:2717],
	45829:    _Code_name[2717:2728],
	45830:    _Code_name[2728:2739],
	45831:    _Code_name[2739:2750],
	45832:    _Code_name[2750:2761],
	45833:    _Code_name[2761:2772],
	45834:    _Code_name[2772:2783],
	45835:    _Code_name[2783:2794],
	45836:    _Code_name[2794:2805],
	45837:    _Code_name[2805:2817],
	45838:    _Code_name[2817:2829],
	45839:    _Code_name[2829:2841],
	45840:    _Code_name[2841:2853],
	45841:    _Code_name[2853:2865],
	45842:    _Code_name[2865:2877],
	45843:    _Code_name[2877:2889],
	45844:    _Code_name[2889:2901],
	45845:    _Code_name[2901:2913],
	45846:    _Code_name[2913:2925],
	45847:    _Code_name[2925:2937],
	45848:    _Code_name[2937:2949],
	45849:    _Code_name[2949:2961],
	45850:    _Code_name[2961:2973],
	45851:    _Code_name[2973:2985],
	45852:    _Code_name[2985:2997],
	45853:    _Code_name[2997:3009],
	45854:    _Code_name[3009:3021],
	45855:    _Code_name[3021:3033],
	45856:    _Code_name[3033:3045],
	45857:    _Code_name[3045:3055],
	45858:    _Code_name[3055:3066],
	45859:    _Code_name[3066:3077],
	45860:    _Code_name[3077:3088],
	45861:    _Code_name[3088:3099],
	45862:    _Code_name[3099:3110],
	45863:    _Code_name[3110:3121],
	45864:    _Code_name[3121:3132],
	45865:    _Code_name[3132:3143],
	45866:    _Code_name[3143:3154],
	45867:    _Code_name[3154:3165],
	45868:    _Code_name[3165:3176],
	45869:    _Code_name[3176:3188],
	45870:    _Code_name[3188:3200],
	45871:    _Code_name[3200:3212],
	45872:    _Code_name[3212:3224],
	45873:    _Code_name[3224:3236],
	45874:    _Code_name[3236:3248],
	45875:    _Code_name[3248:3260],
	45876:    _Code_name[3260:3272],
	45877:    _Code_name[3272:3284],
	45878:    _Code_name[3284:3296],
	45879:    _Code_name[3296:3308],
	45880:    _Code_name[3308:3320],
	45881:    _Code_name[3320:3332],
	45882:    _Code_name[3332:3344],
	45883:    _Code_name[3344:3356],
	45884:    _Code_name[3356:3368],
	45885:    _Code_name[3368:3380],
	45886:    _Code_name[3380:3392],
	45887:    _Code_name[3392:3404],
	45888:    _Code_name[3404:3416],
	45889:    _Code_name[3416:3428],
	45890:    _Code_name[3428:3440],
	45891:    _Code_name[3440:3452],
	45892:    _Code_name[3452:3464],
	45893:    _Code_name[3464:3476],
	45894:    _Code_name[3476:3488],
	45895:    _Code_name[3488:3500],
	45896:    _Code_name[3500:3512],
	45897:    _Code_name[3512:3524],
	45898:    _Code_name[3524:3536],
	45899:    _Code_name[3536:3548],
	45900:    _Code_name[3548:3560],
	45901:    _Code_name[3560:3572],
	45902:    _Code_name[3572:3584],
	45903:    _Code_name[3584:3596],
	45904:    _Code_name[3596:3608],
	45905:    _Code_name[3608:3620],
	45906:    _Code_name[3620:3632],
	45907:    _Code_name[3632:3644],
	45908:    _Code_name[3644:3656],
	45909:    _Code_name[3656:3668],
	45910:    _Code_name[3668:3680],
	45911:    _Code_name[3680:3692],
	45912:    _Code_name[3692:3704],
	45913:    _Code_name[3704:3716],
	45914:    _Code_name[3716:3728],
	45915:    _Code_name[3728:3740],
	45916:    _Code_name[3740:3752],
	45917:    _Code_name[3752:3764],
	45918:    _Code_name[3764:3776],
	45919:    _Code_name[3776:3788],
	45920:    _Code_name[3788:3800],
	45921:    _Code_name[3800:3811],
	45922:    _Code_name[3811:3823],
	45923:    _Code_name[3823:3835],
	45924:    _Code_name[3835:3847],
	45925:    _Code_name[3847:3859],
	45926:    _Code_name[3859:3871],
	45927:    _Code_name[3871:3883],
	45928:    _Code_name[3883:3895],
	45929:    _Code_name[3895:3907],
	45930:    _Code_name[3907:3919],
	45931:    _Code_name[3919:3931],
	45932:    _Code_name[3931:3943],
	45933:    _Code_name[3943:3956],
	45934:    _Code_name[3956:3969],
	45935:    _Code_name[3969:3982],
	45936:    _Code_name[3982:3995],
	45937:    _Code_name[3995:4008],
	45938:    _Code_name[4008:4021],
	45939:    _Code_name[4021:4034],
	45940:    _Code_name[4034:4047],
	45941:    _Code_name[4047:4060],
	45942:    _Code_name[4060:4073],
	45943:    _Code_name[4073:4086],
	45944:    _Code_name[4086:4099],
	45945:    _Code_name[4099:4112],
	45946:    _Code_name[4112:4125],
	45947:    _Code_name[4125:4138],
	45948:    _Code_name[4138:4151],
	45949:    _Code_name[4151:4164],
	45950:    _Code_name[4164:4177],
	45951:    _Code_name[4177:4190],
	45952:    _Code_name[4190:4203],
	45953:    _Code_name[4203:4216],
	45954:    _Code_name[4216:4229],
	45955:    _Code_name[4229:4242],
	45956:    _Code_name[4242:4255],
	45957:    _Code_name[4255:4268],
	45958:    _Code_name[4268:4281],
	45959:    _Code_name[4281:4294],
	45960:    _Code_name[4294:4307],
	45961:    _Code_name[4307:4320],
	45962:    _Code_name[4320:4333],
	45963:    _Code_name[4333:4346],
	45964:    _Code_name[4346:4359],
	45965:    _Code_name[4359:4372],
	45966:    _Code_name[4372:4385],
	45967:    _Code_name[4385:4398],
	45968:    _Code_name[4398:4411],
	45969:    _Code_name[4411:4424],
	45970:    _Code_name[4424:4437],
	45971:    _Code_name[4437:4450],
	45972:    _Code_name[4450:4463],
	45973:    _Code_name[4463:4476],
	45974:    _Code_name[4476:4489],
	45975:    _Code_name[4489:4502],
	45976:    _Code_name[4502:4515],
	45977:    _Code_name[4515:4528],
	45978:    _Code_name[4528:4541],
	45979:    _Code_name[4541:4554],
	45980:    _Code_name[4554:4567],
	45981:    _Code_name[4567:4580],
	45982:    _Code_name[4580:4593],
	45983:    _Code_name[4593:4606],
	45984:    _Code_name[4606:4619],
	45985:    _Code_name[4619:4632],
	45986:    _Code_name[4632:4645],
	45987:    _Code_name[4645:4658],
	45988:    _Code_name[4658:4671],
	45989:    _Code_name[4671:4684],
	45990:    _Code_name[4684:4697],
	45991:    _Code_name[4697:4710],
	45992:    _Code_name[4710:4723],
	45993:    _Code_name[4723:4736],
	45994:    _Code_name[4736:4749],
	45995:    _Code_name[4749:4762],
	45996:    _Code_name[4762:4775],
	45997:    _Code_name[4775:4788],
	45998:    _Code_name[4788:4801],
	45999:    _Code_name[4801:4814],
	46000:    _Code_name[4814:4827],
	46001:    _Code_name[4827:4840],
	46002:    _Code_name[4840:4853],
	46003:    _Code_name[4853:4866],
	46004:    _Code_name[4866:4879],
	46005:    _Code_name[4879:4892],
	46006:    _Code_name[4892:4905],
	46007:    _Code_name[4905:4918],
	46008:    _Code_name[4918:4931],
	46009:    _Code_name[4931:4944],
	46010:    _Code_name[4944:4957],
	46011:    _Code_name[4957:4970],
	46012:    _Code_name[4970:4983],
	46013:    _Code_name[4983:4996],
	46014:    _Code_name[4996:5009],
	46015:    _Code_name[5009:5022],
	46016:    _Code_name[5022:5035],
	46017:    _Code_name[5035:5048],
	46018:    _Code_name[5048:5061],
	46019:    _Code_name[5061:5074],
	46020:    _Code_name[5074:5087],
	46021:    _Code_name[5087:5100],
	46022:    _Code_name[5100:5113],
	46023:    _Code_name[5113:5126],
	46024:    _Code_name[5126:5139],
	46025:    _Code_name[5139:5152],
	46026:    _Code_name[5152:5165],
	46027:    _Code_name[5165:5178],
	46028:    _Code_name[5178:5191],
	46029:    _Code_name[5191:5204],
	46030:    _Code_name[5204:5217],
	46031:    _Code_name[5217:5230],
	46032:    _Code_name[5230:5243],
	46033:    _Code_name[5243:5256],
	46034:    _Code_name[5256:5269],
	46035:    _Code_name[5269:5282],
	46036:    _Code_name[5282:5295],
	46037:    _Code_name[5295:5308],
	46038:    _Code_name[5308:5321],
	46039:    _Code_name[5321:5334],
	46040:    _Code_name[5334:5347],
	46041:    _Code_name[5347:5360],
	46042:    _Code_name[5360:5373],
	46043:    _Code_name[5373:5386],
	46044:    _Code_name[5386:5399],
	46045:    _Code_name[5399:5413],
	46046:    _Code_name[5413:5427],
	46047:    _Code_name[5427:5441],
	46048:    _Code_name[5441:5455],
	46049:    _Code_name[5455:5461],
	46050:    _Code_name[5461:5467],
	46051:    _Code_name[5467:5474],
	46052:    _Code_name[5474:5482],
	46081:    _Code_name[5482:5507],
	46082:    _Code_name[5507:5535],
	46083:    _Code_name[5535:5551],
	46337:    _Code_name[5551:5554],
	46338:    _Code_name[5554:5570],
	46593:    _Code_name[5570:5578],
	52225:    _Code_name[5578:5582],
	52753:    _Code_name[5582:5603],
	53248:    _Code_name[5603:5618],
	53479:    _Code_name[5618:5624],
	53482:    _Code_name[5624:5640],
	53483:    _Code_name[5640:5656],
	53485:    _Code_name[5656:5661],
	53649:    _Code_name[5661:5668],
	60241:    _Code_name[5668:5679],
	61697:    _Code_name[5679:5702],
	61698:    _Code_name[5702:5723],
	7367777:  _Code_name[5723:5734],
	8417572:  _Code_name[5734:5750],
	8483108:  _Code_name[5750:5766],
	9728292:  _Code_name[5766:5782],
	9793828:  _Code_name[5782:5798],
	10645796: _Code_name[5798:5814],
	10711332: _Code_name[5814:5830],
	11639056: _Code_name[5830:5839],
	11704592: _Code_name[5839:5849],
	11770128: _Code_name[5849:5860],
	11835664: _Code_name[5860:5871],
	13636096: _Code_name[5871:5876],
	13636097: _Code_name[5876:5881],
	13636098: _Code_name[5881:5886],
	13636101: _Code_name[5886:5891],
}

func (i Code) String() string {
	if str, ok := _Code_map[i]; ok {
		return str
	}
	return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
}
