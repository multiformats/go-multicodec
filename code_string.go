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
	_ = x[Dccp-33]
	_ = x[Murmur3_128-34]
	_ = x[Murmur3_32-35]
	_ = x[Ip6-41]
	_ = x[Ip6zone-42]
	_ = x[Path-47]
	_ = x[Multicodec-48]
	_ = x[Multihash-49]
	_ = x[Multiaddr-50]
	_ = x[Multibase-51]
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
	_ = x[Bmt-214]
	_ = x[DecredBlock-224]
	_ = x[DecredTx-225]
	_ = x[IpldNs-226]
	_ = x[IpfsNs-227]
	_ = x[SwarmNs-228]
	_ = x[IpnsNs-229]
	_ = x[Zeronet-230]
	_ = x[Secp256k1Pub-231]
	_ = x[Bls12_381G1Pub-234]
	_ = x[Bls12_381G2Pub-235]
	_ = x[X25519Pub-236]
	_ = x[Ed25519Pub-237]
	_ = x[Bls12_381G1g2Pub-238]
	_ = x[DashBlock-240]
	_ = x[DashTx-241]
	_ = x[SwarmManifest-250]
	_ = x[SwarmFeed-251]
	_ = x[Udp-273]
	_ = x[P2pWebrtcStar-275]
	_ = x[P2pWebrtcDirect-276]
	_ = x[P2pStardust-277]
	_ = x[P2pCircuit-290]
	_ = x[DagJson-297]
	_ = x[Udt-301]
	_ = x[Utp-302]
	_ = x[Unix-400]
	_ = x[Thread-406]
	_ = x[P2p-421]
	_ = x[Ipfs-421]
	_ = x[Https-443]
	_ = x[Onion-444]
	_ = x[Onion3-445]
	_ = x[Garlic64-446]
	_ = x[Garlic32-447]
	_ = x[Tls-448]
	_ = x[Noise-454]
	_ = x[Quic-460]
	_ = x[Ws-477]
	_ = x[Wss-478]
	_ = x[P2pWebsocketStar-479]
	_ = x[Http-480]
	_ = x[Swhid1Snp-496]
	_ = x[Json-512]
	_ = x[Messagepack-513]
	_ = x[Libp2pPeerRecord-769]
	_ = x[Libp2pRelayRsvp-770]
	_ = x[CarIndexSorted-1024]
	_ = x[CarMultihashIndexSorted-1025]
	_ = x[Sha2_256Trunc254Padded-4114]
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
	_ = x[RsaX509Pub-4613]
	_ = x[Ed25519Priv-4864]
	_ = x[Secp256k1Priv-4865]
	_ = x[X25519Priv-4866]
	_ = x[Kangarootwelve-7425]
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
	_ = x[PoseidonBls12_381A2Fc1-46081]
	_ = x[PoseidonBls12_381A2Fc1Sc-46082]
	_ = x[ZeroxcertImprint256-52753]
	_ = x[FilCommitmentUnsealed-61697]
	_ = x[FilCommitmentSealed-61698]
	_ = x[HolochainAdrV0-8417572]
	_ = x[HolochainAdrV1-8483108]
	_ = x[HolochainKeyV0-9728292]
	_ = x[HolochainKeyV1-9793828]
	_ = x[HolochainSigV0-10645796]
	_ = x[HolochainSigV1-10711332]
	_ = x[SkynetNs-11639056]
	_ = x[ArweaveNs-11704592]
}

const _Code_name = "identitycidv1cidv2cidv3ip4tcpsha1sha2-256sha2-512sha3-512sha3-384sha3-256sha3-224shake-128shake-256keccak-224keccak-256keccak-384keccak-512blake3dccpmurmur3-128murmur3-32ip6ip6zonepathmulticodecmultihashmultiaddrmultibasednsdns4dns6dnsaddrprotobufcborrawdbl-sha2-256rlpbencodedag-pbdag-cborlibp2p-keygit-rawtorrent-infotorrent-fileleofcoin-blockleofcoin-txleofcoin-prsctpdag-josedag-coseeth-blocketh-block-listeth-tx-trieeth-txeth-tx-receipt-trieeth-tx-receipteth-state-trieeth-account-snapshoteth-storage-trieeth-receipt-log-trieeth-reciept-logbitcoin-blockbitcoin-txbitcoin-witness-commitmentzcash-blockzcash-txcaip-50streamidstellar-blockstellar-txmd4md5bmtdecred-blockdecred-txipld-nsipfs-nsswarm-nsipns-nszeronetsecp256k1-pubbls12_381-g1-pubbls12_381-g2-pubx25519-pubed25519-pubbls12_381-g1g2-pubdash-blockdash-txswarm-manifestswarm-feedudpp2p-webrtc-starp2p-webrtc-directp2p-stardustp2p-circuitdag-jsonudtutpunixthreadp2phttpsoniononion3garlic64garlic32tlsnoisequicwswssp2p-websocket-starhttpswhid-1-snpjsonmessagepacklibp2p-peer-recordlibp2p-relay-rsvpcar-index-sortedcar-multihash-index-sortedsha2-256-trunc254-paddedripemd-128ripemd-160ripemd-256ripemd-320x11p256-pubp384-pubp521-pubed448-pubx448-pubrsa-x509-pubed25519-privsecp256k1-privx25519-privkangarootwelvesm3-256blake2b-8blake2b-16blake2b-24blake2b-32blake2b-40blake2b-48blake2b-56blake2b-64blake2b-72blake2b-80blake2b-88blake2b-96blake2b-104blake2b-112blake2b-120blake2b-128blake2b-136blake2b-144blake2b-152blake2b-160blake2b-168blake2b-176blake2b-184blake2b-192blake2b-200blake2b-208blake2b-216blake2b-224blake2b-232blake2b-240blake2b-248blake2b-256blake2b-264blake2b-272blake2b-280blake2b-288blake2b-296blake2b-304blake2b-312blake2b-320blake2b-328blake2b-336blake2b-344blake2b-352blake2b-360blake2b-368blake2b-376blake2b-384blake2b-392blake2b-400blake2b-408blake2b-416blake2b-424blake2b-432blake2b-440blake2b-448blake2b-456blake2b-464blake2b-472blake2b-480blake2b-488blake2b-496blake2b-504blake2b-512blake2s-8blake2s-16blake2s-24blake2s-32blake2s-40blake2s-48blake2s-56blake2s-64blake2s-72blake2s-80blake2s-88blake2s-96blake2s-104blake2s-112blake2s-120blake2s-128blake2s-136blake2s-144blake2s-152blake2s-160blake2s-168blake2s-176blake2s-184blake2s-192blake2s-200blake2s-208blake2s-216blake2s-224blake2s-232blake2s-240blake2s-248blake2s-256skein256-8skein256-16skein256-24skein256-32skein256-40skein256-48skein256-56skein256-64skein256-72skein256-80skein256-88skein256-96skein256-104skein256-112skein256-120skein256-128skein256-136skein256-144skein256-152skein256-160skein256-168skein256-176skein256-184skein256-192skein256-200skein256-208skein256-216skein256-224skein256-232skein256-240skein256-248skein256-256skein512-8skein512-16skein512-24skein512-32skein512-40skein512-48skein512-56skein512-64skein512-72skein512-80skein512-88skein512-96skein512-104skein512-112skein512-120skein512-128skein512-136skein512-144skein512-152skein512-160skein512-168skein512-176skein512-184skein512-192skein512-200skein512-208skein512-216skein512-224skein512-232skein512-240skein512-248skein512-256skein512-264skein512-272skein512-280skein512-288skein512-296skein512-304skein512-312skein512-320skein512-328skein512-336skein512-344skein512-352skein512-360skein512-368skein512-376skein512-384skein512-392skein512-400skein512-408skein512-416skein512-424skein512-432skein512-440skein512-448skein512-456skein512-464skein512-472skein512-480skein512-488skein512-496skein512-504skein512-512skein1024-8skein1024-16skein1024-24skein1024-32skein1024-40skein1024-48skein1024-56skein1024-64skein1024-72skein1024-80skein1024-88skein1024-96skein1024-104skein1024-112skein1024-120skein1024-128skein1024-136skein1024-144skein1024-152skein1024-160skein1024-168skein1024-176skein1024-184skein1024-192skein1024-200skein1024-208skein1024-216skein1024-224skein1024-232skein1024-240skein1024-248skein1024-256skein1024-264skein1024-272skein1024-280skein1024-288skein1024-296skein1024-304skein1024-312skein1024-320skein1024-328skein1024-336skein1024-344skein1024-352skein1024-360skein1024-368skein1024-376skein1024-384skein1024-392skein1024-400skein1024-408skein1024-416skein1024-424skein1024-432skein1024-440skein1024-448skein1024-456skein1024-464skein1024-472skein1024-480skein1024-488skein1024-496skein1024-504skein1024-512skein1024-520skein1024-528skein1024-536skein1024-544skein1024-552skein1024-560skein1024-568skein1024-576skein1024-584skein1024-592skein1024-600skein1024-608skein1024-616skein1024-624skein1024-632skein1024-640skein1024-648skein1024-656skein1024-664skein1024-672skein1024-680skein1024-688skein1024-696skein1024-704skein1024-712skein1024-720skein1024-728skein1024-736skein1024-744skein1024-752skein1024-760skein1024-768skein1024-776skein1024-784skein1024-792skein1024-800skein1024-808skein1024-816skein1024-824skein1024-832skein1024-840skein1024-848skein1024-856skein1024-864skein1024-872skein1024-880skein1024-888skein1024-896skein1024-904skein1024-912skein1024-920skein1024-928skein1024-936skein1024-944skein1024-952skein1024-960skein1024-968skein1024-976skein1024-984skein1024-992skein1024-1000skein1024-1008skein1024-1016skein1024-1024poseidon-bls12_381-a2-fc1poseidon-bls12_381-a2-fc1-sczeroxcert-imprint-256fil-commitment-unsealedfil-commitment-sealedholochain-adr-v0holochain-adr-v1holochain-key-v0holochain-key-v1holochain-sig-v0holochain-sig-v1skynet-nsarweave-ns"

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
	33:       _Code_name[145:149],
	34:       _Code_name[149:160],
	35:       _Code_name[160:170],
	41:       _Code_name[170:173],
	42:       _Code_name[173:180],
	47:       _Code_name[180:184],
	48:       _Code_name[184:194],
	49:       _Code_name[194:203],
	50:       _Code_name[203:212],
	51:       _Code_name[212:221],
	53:       _Code_name[221:224],
	54:       _Code_name[224:228],
	55:       _Code_name[228:232],
	56:       _Code_name[232:239],
	80:       _Code_name[239:247],
	81:       _Code_name[247:251],
	85:       _Code_name[251:254],
	86:       _Code_name[254:266],
	96:       _Code_name[266:269],
	99:       _Code_name[269:276],
	112:      _Code_name[276:282],
	113:      _Code_name[282:290],
	114:      _Code_name[290:300],
	120:      _Code_name[300:307],
	123:      _Code_name[307:319],
	124:      _Code_name[319:331],
	129:      _Code_name[331:345],
	130:      _Code_name[345:356],
	131:      _Code_name[356:367],
	132:      _Code_name[367:371],
	133:      _Code_name[371:379],
	134:      _Code_name[379:387],
	144:      _Code_name[387:396],
	145:      _Code_name[396:410],
	146:      _Code_name[410:421],
	147:      _Code_name[421:427],
	148:      _Code_name[427:446],
	149:      _Code_name[446:460],
	150:      _Code_name[460:474],
	151:      _Code_name[474:494],
	152:      _Code_name[494:510],
	153:      _Code_name[510:530],
	154:      _Code_name[530:545],
	176:      _Code_name[545:558],
	177:      _Code_name[558:568],
	178:      _Code_name[568:594],
	192:      _Code_name[594:605],
	193:      _Code_name[605:613],
	202:      _Code_name[613:620],
	206:      _Code_name[620:628],
	208:      _Code_name[628:641],
	209:      _Code_name[641:651],
	212:      _Code_name[651:654],
	213:      _Code_name[654:657],
	214:      _Code_name[657:660],
	224:      _Code_name[660:672],
	225:      _Code_name[672:681],
	226:      _Code_name[681:688],
	227:      _Code_name[688:695],
	228:      _Code_name[695:703],
	229:      _Code_name[703:710],
	230:      _Code_name[710:717],
	231:      _Code_name[717:730],
	234:      _Code_name[730:746],
	235:      _Code_name[746:762],
	236:      _Code_name[762:772],
	237:      _Code_name[772:783],
	238:      _Code_name[783:801],
	240:      _Code_name[801:811],
	241:      _Code_name[811:818],
	250:      _Code_name[818:832],
	251:      _Code_name[832:842],
	273:      _Code_name[842:845],
	275:      _Code_name[845:860],
	276:      _Code_name[860:877],
	277:      _Code_name[877:889],
	290:      _Code_name[889:900],
	297:      _Code_name[900:908],
	301:      _Code_name[908:911],
	302:      _Code_name[911:914],
	400:      _Code_name[914:918],
	406:      _Code_name[918:924],
	421:      _Code_name[924:927],
	443:      _Code_name[927:932],
	444:      _Code_name[932:937],
	445:      _Code_name[937:943],
	446:      _Code_name[943:951],
	447:      _Code_name[951:959],
	448:      _Code_name[959:962],
	454:      _Code_name[962:967],
	460:      _Code_name[967:971],
	477:      _Code_name[971:973],
	478:      _Code_name[973:976],
	479:      _Code_name[976:994],
	480:      _Code_name[994:998],
	496:      _Code_name[998:1009],
	512:      _Code_name[1009:1013],
	513:      _Code_name[1013:1024],
	769:      _Code_name[1024:1042],
	770:      _Code_name[1042:1059],
	1024:     _Code_name[1059:1075],
	1025:     _Code_name[1075:1101],
	4114:     _Code_name[1101:1125],
	4178:     _Code_name[1125:1135],
	4179:     _Code_name[1135:1145],
	4180:     _Code_name[1145:1155],
	4181:     _Code_name[1155:1165],
	4352:     _Code_name[1165:1168],
	4608:     _Code_name[1168:1176],
	4609:     _Code_name[1176:1184],
	4610:     _Code_name[1184:1192],
	4611:     _Code_name[1192:1201],
	4612:     _Code_name[1201:1209],
	4613:     _Code_name[1209:1221],
	4864:     _Code_name[1221:1233],
	4865:     _Code_name[1233:1247],
	4866:     _Code_name[1247:1258],
	7425:     _Code_name[1258:1272],
	21325:    _Code_name[1272:1279],
	45569:    _Code_name[1279:1288],
	45570:    _Code_name[1288:1298],
	45571:    _Code_name[1298:1308],
	45572:    _Code_name[1308:1318],
	45573:    _Code_name[1318:1328],
	45574:    _Code_name[1328:1338],
	45575:    _Code_name[1338:1348],
	45576:    _Code_name[1348:1358],
	45577:    _Code_name[1358:1368],
	45578:    _Code_name[1368:1378],
	45579:    _Code_name[1378:1388],
	45580:    _Code_name[1388:1398],
	45581:    _Code_name[1398:1409],
	45582:    _Code_name[1409:1420],
	45583:    _Code_name[1420:1431],
	45584:    _Code_name[1431:1442],
	45585:    _Code_name[1442:1453],
	45586:    _Code_name[1453:1464],
	45587:    _Code_name[1464:1475],
	45588:    _Code_name[1475:1486],
	45589:    _Code_name[1486:1497],
	45590:    _Code_name[1497:1508],
	45591:    _Code_name[1508:1519],
	45592:    _Code_name[1519:1530],
	45593:    _Code_name[1530:1541],
	45594:    _Code_name[1541:1552],
	45595:    _Code_name[1552:1563],
	45596:    _Code_name[1563:1574],
	45597:    _Code_name[1574:1585],
	45598:    _Code_name[1585:1596],
	45599:    _Code_name[1596:1607],
	45600:    _Code_name[1607:1618],
	45601:    _Code_name[1618:1629],
	45602:    _Code_name[1629:1640],
	45603:    _Code_name[1640:1651],
	45604:    _Code_name[1651:1662],
	45605:    _Code_name[1662:1673],
	45606:    _Code_name[1673:1684],
	45607:    _Code_name[1684:1695],
	45608:    _Code_name[1695:1706],
	45609:    _Code_name[1706:1717],
	45610:    _Code_name[1717:1728],
	45611:    _Code_name[1728:1739],
	45612:    _Code_name[1739:1750],
	45613:    _Code_name[1750:1761],
	45614:    _Code_name[1761:1772],
	45615:    _Code_name[1772:1783],
	45616:    _Code_name[1783:1794],
	45617:    _Code_name[1794:1805],
	45618:    _Code_name[1805:1816],
	45619:    _Code_name[1816:1827],
	45620:    _Code_name[1827:1838],
	45621:    _Code_name[1838:1849],
	45622:    _Code_name[1849:1860],
	45623:    _Code_name[1860:1871],
	45624:    _Code_name[1871:1882],
	45625:    _Code_name[1882:1893],
	45626:    _Code_name[1893:1904],
	45627:    _Code_name[1904:1915],
	45628:    _Code_name[1915:1926],
	45629:    _Code_name[1926:1937],
	45630:    _Code_name[1937:1948],
	45631:    _Code_name[1948:1959],
	45632:    _Code_name[1959:1970],
	45633:    _Code_name[1970:1979],
	45634:    _Code_name[1979:1989],
	45635:    _Code_name[1989:1999],
	45636:    _Code_name[1999:2009],
	45637:    _Code_name[2009:2019],
	45638:    _Code_name[2019:2029],
	45639:    _Code_name[2029:2039],
	45640:    _Code_name[2039:2049],
	45641:    _Code_name[2049:2059],
	45642:    _Code_name[2059:2069],
	45643:    _Code_name[2069:2079],
	45644:    _Code_name[2079:2089],
	45645:    _Code_name[2089:2100],
	45646:    _Code_name[2100:2111],
	45647:    _Code_name[2111:2122],
	45648:    _Code_name[2122:2133],
	45649:    _Code_name[2133:2144],
	45650:    _Code_name[2144:2155],
	45651:    _Code_name[2155:2166],
	45652:    _Code_name[2166:2177],
	45653:    _Code_name[2177:2188],
	45654:    _Code_name[2188:2199],
	45655:    _Code_name[2199:2210],
	45656:    _Code_name[2210:2221],
	45657:    _Code_name[2221:2232],
	45658:    _Code_name[2232:2243],
	45659:    _Code_name[2243:2254],
	45660:    _Code_name[2254:2265],
	45661:    _Code_name[2265:2276],
	45662:    _Code_name[2276:2287],
	45663:    _Code_name[2287:2298],
	45664:    _Code_name[2298:2309],
	45825:    _Code_name[2309:2319],
	45826:    _Code_name[2319:2330],
	45827:    _Code_name[2330:2341],
	45828:    _Code_name[2341:2352],
	45829:    _Code_name[2352:2363],
	45830:    _Code_name[2363:2374],
	45831:    _Code_name[2374:2385],
	45832:    _Code_name[2385:2396],
	45833:    _Code_name[2396:2407],
	45834:    _Code_name[2407:2418],
	45835:    _Code_name[2418:2429],
	45836:    _Code_name[2429:2440],
	45837:    _Code_name[2440:2452],
	45838:    _Code_name[2452:2464],
	45839:    _Code_name[2464:2476],
	45840:    _Code_name[2476:2488],
	45841:    _Code_name[2488:2500],
	45842:    _Code_name[2500:2512],
	45843:    _Code_name[2512:2524],
	45844:    _Code_name[2524:2536],
	45845:    _Code_name[2536:2548],
	45846:    _Code_name[2548:2560],
	45847:    _Code_name[2560:2572],
	45848:    _Code_name[2572:2584],
	45849:    _Code_name[2584:2596],
	45850:    _Code_name[2596:2608],
	45851:    _Code_name[2608:2620],
	45852:    _Code_name[2620:2632],
	45853:    _Code_name[2632:2644],
	45854:    _Code_name[2644:2656],
	45855:    _Code_name[2656:2668],
	45856:    _Code_name[2668:2680],
	45857:    _Code_name[2680:2690],
	45858:    _Code_name[2690:2701],
	45859:    _Code_name[2701:2712],
	45860:    _Code_name[2712:2723],
	45861:    _Code_name[2723:2734],
	45862:    _Code_name[2734:2745],
	45863:    _Code_name[2745:2756],
	45864:    _Code_name[2756:2767],
	45865:    _Code_name[2767:2778],
	45866:    _Code_name[2778:2789],
	45867:    _Code_name[2789:2800],
	45868:    _Code_name[2800:2811],
	45869:    _Code_name[2811:2823],
	45870:    _Code_name[2823:2835],
	45871:    _Code_name[2835:2847],
	45872:    _Code_name[2847:2859],
	45873:    _Code_name[2859:2871],
	45874:    _Code_name[2871:2883],
	45875:    _Code_name[2883:2895],
	45876:    _Code_name[2895:2907],
	45877:    _Code_name[2907:2919],
	45878:    _Code_name[2919:2931],
	45879:    _Code_name[2931:2943],
	45880:    _Code_name[2943:2955],
	45881:    _Code_name[2955:2967],
	45882:    _Code_name[2967:2979],
	45883:    _Code_name[2979:2991],
	45884:    _Code_name[2991:3003],
	45885:    _Code_name[3003:3015],
	45886:    _Code_name[3015:3027],
	45887:    _Code_name[3027:3039],
	45888:    _Code_name[3039:3051],
	45889:    _Code_name[3051:3063],
	45890:    _Code_name[3063:3075],
	45891:    _Code_name[3075:3087],
	45892:    _Code_name[3087:3099],
	45893:    _Code_name[3099:3111],
	45894:    _Code_name[3111:3123],
	45895:    _Code_name[3123:3135],
	45896:    _Code_name[3135:3147],
	45897:    _Code_name[3147:3159],
	45898:    _Code_name[3159:3171],
	45899:    _Code_name[3171:3183],
	45900:    _Code_name[3183:3195],
	45901:    _Code_name[3195:3207],
	45902:    _Code_name[3207:3219],
	45903:    _Code_name[3219:3231],
	45904:    _Code_name[3231:3243],
	45905:    _Code_name[3243:3255],
	45906:    _Code_name[3255:3267],
	45907:    _Code_name[3267:3279],
	45908:    _Code_name[3279:3291],
	45909:    _Code_name[3291:3303],
	45910:    _Code_name[3303:3315],
	45911:    _Code_name[3315:3327],
	45912:    _Code_name[3327:3339],
	45913:    _Code_name[3339:3351],
	45914:    _Code_name[3351:3363],
	45915:    _Code_name[3363:3375],
	45916:    _Code_name[3375:3387],
	45917:    _Code_name[3387:3399],
	45918:    _Code_name[3399:3411],
	45919:    _Code_name[3411:3423],
	45920:    _Code_name[3423:3435],
	45921:    _Code_name[3435:3446],
	45922:    _Code_name[3446:3458],
	45923:    _Code_name[3458:3470],
	45924:    _Code_name[3470:3482],
	45925:    _Code_name[3482:3494],
	45926:    _Code_name[3494:3506],
	45927:    _Code_name[3506:3518],
	45928:    _Code_name[3518:3530],
	45929:    _Code_name[3530:3542],
	45930:    _Code_name[3542:3554],
	45931:    _Code_name[3554:3566],
	45932:    _Code_name[3566:3578],
	45933:    _Code_name[3578:3591],
	45934:    _Code_name[3591:3604],
	45935:    _Code_name[3604:3617],
	45936:    _Code_name[3617:3630],
	45937:    _Code_name[3630:3643],
	45938:    _Code_name[3643:3656],
	45939:    _Code_name[3656:3669],
	45940:    _Code_name[3669:3682],
	45941:    _Code_name[3682:3695],
	45942:    _Code_name[3695:3708],
	45943:    _Code_name[3708:3721],
	45944:    _Code_name[3721:3734],
	45945:    _Code_name[3734:3747],
	45946:    _Code_name[3747:3760],
	45947:    _Code_name[3760:3773],
	45948:    _Code_name[3773:3786],
	45949:    _Code_name[3786:3799],
	45950:    _Code_name[3799:3812],
	45951:    _Code_name[3812:3825],
	45952:    _Code_name[3825:3838],
	45953:    _Code_name[3838:3851],
	45954:    _Code_name[3851:3864],
	45955:    _Code_name[3864:3877],
	45956:    _Code_name[3877:3890],
	45957:    _Code_name[3890:3903],
	45958:    _Code_name[3903:3916],
	45959:    _Code_name[3916:3929],
	45960:    _Code_name[3929:3942],
	45961:    _Code_name[3942:3955],
	45962:    _Code_name[3955:3968],
	45963:    _Code_name[3968:3981],
	45964:    _Code_name[3981:3994],
	45965:    _Code_name[3994:4007],
	45966:    _Code_name[4007:4020],
	45967:    _Code_name[4020:4033],
	45968:    _Code_name[4033:4046],
	45969:    _Code_name[4046:4059],
	45970:    _Code_name[4059:4072],
	45971:    _Code_name[4072:4085],
	45972:    _Code_name[4085:4098],
	45973:    _Code_name[4098:4111],
	45974:    _Code_name[4111:4124],
	45975:    _Code_name[4124:4137],
	45976:    _Code_name[4137:4150],
	45977:    _Code_name[4150:4163],
	45978:    _Code_name[4163:4176],
	45979:    _Code_name[4176:4189],
	45980:    _Code_name[4189:4202],
	45981:    _Code_name[4202:4215],
	45982:    _Code_name[4215:4228],
	45983:    _Code_name[4228:4241],
	45984:    _Code_name[4241:4254],
	45985:    _Code_name[4254:4267],
	45986:    _Code_name[4267:4280],
	45987:    _Code_name[4280:4293],
	45988:    _Code_name[4293:4306],
	45989:    _Code_name[4306:4319],
	45990:    _Code_name[4319:4332],
	45991:    _Code_name[4332:4345],
	45992:    _Code_name[4345:4358],
	45993:    _Code_name[4358:4371],
	45994:    _Code_name[4371:4384],
	45995:    _Code_name[4384:4397],
	45996:    _Code_name[4397:4410],
	45997:    _Code_name[4410:4423],
	45998:    _Code_name[4423:4436],
	45999:    _Code_name[4436:4449],
	46000:    _Code_name[4449:4462],
	46001:    _Code_name[4462:4475],
	46002:    _Code_name[4475:4488],
	46003:    _Code_name[4488:4501],
	46004:    _Code_name[4501:4514],
	46005:    _Code_name[4514:4527],
	46006:    _Code_name[4527:4540],
	46007:    _Code_name[4540:4553],
	46008:    _Code_name[4553:4566],
	46009:    _Code_name[4566:4579],
	46010:    _Code_name[4579:4592],
	46011:    _Code_name[4592:4605],
	46012:    _Code_name[4605:4618],
	46013:    _Code_name[4618:4631],
	46014:    _Code_name[4631:4644],
	46015:    _Code_name[4644:4657],
	46016:    _Code_name[4657:4670],
	46017:    _Code_name[4670:4683],
	46018:    _Code_name[4683:4696],
	46019:    _Code_name[4696:4709],
	46020:    _Code_name[4709:4722],
	46021:    _Code_name[4722:4735],
	46022:    _Code_name[4735:4748],
	46023:    _Code_name[4748:4761],
	46024:    _Code_name[4761:4774],
	46025:    _Code_name[4774:4787],
	46026:    _Code_name[4787:4800],
	46027:    _Code_name[4800:4813],
	46028:    _Code_name[4813:4826],
	46029:    _Code_name[4826:4839],
	46030:    _Code_name[4839:4852],
	46031:    _Code_name[4852:4865],
	46032:    _Code_name[4865:4878],
	46033:    _Code_name[4878:4891],
	46034:    _Code_name[4891:4904],
	46035:    _Code_name[4904:4917],
	46036:    _Code_name[4917:4930],
	46037:    _Code_name[4930:4943],
	46038:    _Code_name[4943:4956],
	46039:    _Code_name[4956:4969],
	46040:    _Code_name[4969:4982],
	46041:    _Code_name[4982:4995],
	46042:    _Code_name[4995:5008],
	46043:    _Code_name[5008:5021],
	46044:    _Code_name[5021:5034],
	46045:    _Code_name[5034:5048],
	46046:    _Code_name[5048:5062],
	46047:    _Code_name[5062:5076],
	46048:    _Code_name[5076:5090],
	46081:    _Code_name[5090:5115],
	46082:    _Code_name[5115:5143],
	52753:    _Code_name[5143:5164],
	61697:    _Code_name[5164:5187],
	61698:    _Code_name[5187:5208],
	8417572:  _Code_name[5208:5224],
	8483108:  _Code_name[5224:5240],
	9728292:  _Code_name[5240:5256],
	9793828:  _Code_name[5256:5272],
	10645796: _Code_name[5272:5288],
	10711332: _Code_name[5288:5304],
	11639056: _Code_name[5304:5313],
	11704592: _Code_name[5313:5323],
}

func (i Code) String() string {
	if str, ok := _Code_map[i]; ok {
		return str
	}
	return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
}
