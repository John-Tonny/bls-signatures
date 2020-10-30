package blschia_test

import (
	"bytes"
	"log"
	"math/rand"
	"testing"
	"time"

	bls "github.com/dashpay/bls-signatures/go-bindings"
)

// Test around this example data https://gist.github.com/xdustinface/318c2c08c36ab12a2b1963caf1f7815c

func privateKeyFromString(hexString string) *bls.PrivateKey {
	sk, err := bls.PrivateKeyFromString(hexString)
	if err != nil {
		panic(err)
	}
	return sk
}

func publicKeyFromString(hexString string) *bls.PublicKey {
	pk, err := bls.PublicKeyFromString(hexString)
	if err != nil {
		panic(err)
	}
	return pk
}

func insecureSignatureFromString(hexString string) *bls.InsecureSignature {
	sig, err := bls.InsecureSignatureFromString(hexString)
	if err != nil {
		panic(err)
	}
	return sig
}

func hashFromString(hexString string) bls.Hash {
	hash, err := bls.HashFromString(hexString)
	if err != nil {
		panic(err)
	}
	return hash
}

type member struct {
	operatorKey *bls.PrivateKey
	skShare     *bls.PrivateKey
	blsId       bls.Hash
	sigShare    *bls.InsecureSignature
}

var quorumHash = hashFromString("75899cad8bd52eb1105efa2204ed139bbdb4bb9ddcde4c090b579b3ab4b98fb0")
var signId = hashFromString("0000000000000000000000000000000000000000000000000000000000000001")
var msgHash = hashFromString("0000000000000000000000000000000000000000000000000000000000000002")
var expectedSignHash = hashFromString("b6d8ee31bbd375dfd55d5fb4b02cfccc68709e64f4c5ffcd3895ceb46540311d")
var expectedQuorumPubKey = publicKeyFromString("97a12b918eac73718d55b7fca60435ec0b442d7e24b45cbd80f5cf2ea2e14c4164333fffdb00d27e309abbafacaa9c34")
var expectedThresholdSignature = insecureSignatureFromString("031c536960c9c44efefa4a3334d2d1b808f46abe121cd14a1d0b6ee6b6dca85fd3087d86fe5b1327e6792be53f4ed4df19c3af3aac79c0bb9dc36fc2a30ba566de6a82cd3e4af2d6654cbe02bb52769a06c1644a4c63b3c3a6fd16e5e68e5c0b")
var minimumRequiredSignatures = 6

var members = []member{
	{
		operatorKey: privateKeyFromString("5f16c2922ff97ebd252e4e5a5355dbc8f5e86d3f506be7c7dcba19c44bf7e76b"),
		skShare:     privateKeyFromString("43dd263542a8d10bc9f843b46f15c86cd87e00c8f45fe5339b30c08e3233d8e3"),
		blsId:       hashFromString("4393e2a243c3db776dcdbe2535493440d29cb65006a1e6f0f8d3f1e1488cbf1a"),
		sigShare:    insecureSignatureFromString("0888879c99852460912fd28c7a9138926c1e87fd6609fd2d3d307764e49feb85702fd8f9b3b836bc11f7ce151b769dc70b760879d26f8c33a29e24f69297f45ef028f0794e63ddb0610db7de1a608b6d6a2129ada62b845004a408f651fd44a5"),
	},
	{
		operatorKey: privateKeyFromString("0f2ac0b196da2512a23948e163810c4ac903ee4ad390dcd244fbee7d199cc73a"),
		skShare:     privateKeyFromString("5e7247ef1a0e671b8349e7be3f50ec88f1eafde90345f34520010e4aa9f18731"),
		blsId:       hashFromString("8b2d776ac75cfca1559b5616ade4eb376a6478556135276e4b3210fe170b9f59"),
		sigShare:    insecureSignatureFromString("16efc39fa5aa979245a82334856a97ebf3027bc6be7d35df117267a3c9b1618e32477fe1b8f4a23bdba346bf2b2b35ad0b395227de76827fd32eb9952e0d97b7dba275040101131a7fc38ea381a3099c2b3205177866ee4ab3119593bb58d092"),
	},
	{
		operatorKey: privateKeyFromString("2e0daa3414ebae8767acbdbf1c94c4982022eb70995f5b285fd304b9dbaa517b"),
		skShare:     privateKeyFromString("34bcc40dea17bb03ec085ac46a0ea9614b3ffc4bae8b0b292f3d7c54662b00a6"),
		blsId:       hashFromString("f2015bdbc0daa70c41a25d2450b96f433ac7d568126505e9997794bb357cf3af"),
		sigShare:    insecureSignatureFromString("8407afd2776ab9d3f9293f1519ace1a9ce8aaf07d0a6a9785ec7ae4ae47e42102c09cfb3c8655dba014d53933af6a0b41244df006575e85e333271c90fcad80410cab4962bf4bba1570770775b1282f142b526d521a38fbc14336f22dc44a683"),
	},
	{
		operatorKey: privateKeyFromString("026950917a3a33feb0f094115c17012bcece73598a175fb24cebc573f9e7d67f"),
		skShare:     privateKeyFromString("139f967b6f4af5dfe2bebf8085b6332efe31c2dc348d02e6b4745a0e7e56a469"),
		blsId:       hashFromString("5818a68f2f34e5ff7d1d43beca5ff103739dd918efda4bac7fd4ede6c53dc6af"),
		sigShare:    insecureSignatureFromString("027061a8c2d631e40882ce6919d3e5f45c4c74ac32a3bce94e5661d06cdecf2d492dfab99e9a9dd8a29a90fe8f816be30178bf9277a3751882e49bb9f08437f5f2cd9dbc12c2fdcccaf7578838e87273fd2ba87f20cf00ca5fec56822f845769"),
	},
	{
		operatorKey: privateKeyFromString("6126980e2b22cee8503e12097f82651b26ff22084dc7fed7247035f653e34556"),
		skShare:     privateKeyFromString("08442e959054d87b5de3553ef8cfd9da923241664c35c6548b5e3271a86b4a18"),
		blsId:       hashFromString("965437bdf51278f716078477a2eae595a9d2b0aa3fe69a387b30936c13c7d68e"),
		sigShare:    insecureSignatureFromString("178ece91967145b1dfc02de703dbd8049b05d626f18a71303ea0c23ee3b60a52bd61cc30fea3e4a562b20c20e0439a2f108b4e6b8a646f647afb64e3b355eba382380ef2c634f3a56de066b7a764249aba1c42c49d76d65e094e890cbbf005a3"),
	},
	{
		operatorKey: privateKeyFromString("0d8bbcc62f88e0de77fc9833985ebb41a9d74babea5daeb3e5c9ce375ac03d4a"),
		skShare:     privateKeyFromString("2698613947a156639b423ad1a9fbe45863d58540d8ebd08612504bf9cf4743ea"),
		blsId:       hashFromString("fd14695a48a35fe6a1f9894accfa83349508e350b3f743d494074fe204b17166"),
		sigShare:    insecureSignatureFromString("193da45d31d728acc92165173253fd8689301b448c81039350ae6916a72f157b00da469a7ab6d2b5db2dac216f47073d089afdcdffce25b6aac991f4745c803f164d7426b8da7d19bf699f5e5489f715ac32e539db90610d7df47121556c1a20"),
	},
	{
		operatorKey: privateKeyFromString("6ad9d9bf60484aa807ceaf9696f941dca0fcfae39e2cb62dbc898895f35c67d2"),
		skShare:     privateKeyFromString("1871c9270c8344028946eee64e79d09e4915dd3b717ffc1c9aa86faff88c0475"),
		blsId:       hashFromString("7e3c28e59ff54bf097b2f3ada4a30f5613227951116675127fc97c98405f2067"),
		sigShare:    insecureSignatureFromString("938ce6cc265fa15fe67314ac4773f18ff0b49c01eac626814abc2f836814068aeba8582d619a3e0c08dc4bafecda84a818b6a7abd350637e72a47356e5919e3a72be66316417c598338e4ab85f8d25535bd6c4a5fb16767ac470902e0cf19df0"),
	},
	{
		operatorKey: privateKeyFromString("0d0d93b08c0e7dd8f92bacffc41beb674f1a7fbb38eb26f843e254b02fe8e14e"),
		skShare:     privateKeyFromString("68409938427df3567e8948c1fff8610b5cf94872eb959c90a714b7ff0f339e88"),
		blsId:       hashFromString("5e427dba092e3d81d057c0277a9a465e036c8336c59a18f27d7c21bc51910904"),
		sigShare:    insecureSignatureFromString("985039d55a92f6fb3b324b0b9f1c7ddcd5f443d6d1ce72549f043b967ded7d56dc4320dc8569a1c41c6cfb4c150d8c61095d3b325e3308a321ceb43369fe56807fba67b6b313f072ab2cdbaa872b52a9a2e75bf396f1b2007152067f756946b7"),
	},
	{
		operatorKey: privateKeyFromString("01123e1f286104828336f8cd2fdbb5c87d3c91f73d3e0a1d7fe03dc9bffbabd0"),
		skShare:     privateKeyFromString("70d3e3ad7d22bd30e4e2ca108a3fa47f4873bda28f3b000a218339b09db6f58a"),
		blsId:       hashFromString("90976dbe492de3eda7623c7ad6523ed9f63f83c3200c383fccd9f22408e18e1b"),
		sigShare:    insecureSignatureFromString("922717d8c170862662d08a4c29943cc26bb05daff0f07b0b7c352651ec64ee5a1d032bad24dfb42243e9afe044ed25680694b183b657948a91533e9a73b6bf359ff907d5088503137edc8e271ac3d2003a4daf8d36f3f847cc87afc6f9007c72"),
	},
	{
		operatorKey: privateKeyFromString("2d65d86e8d89691fc63dacda2e55ea3f4df2754c22a80e565632552d5da81a0a"),
		skShare:     privateKeyFromString("39ea630e894b71d9f28fefb551611824f16d4b16d29fdea8bb3dbd857a6bc317"),
		blsId:       hashFromString("cd474447afd5df18a0c10c9e2d5216eace9c624119974280236a043cb4b7f8ae"),
		sigShare:    insecureSignatureFromString("02bc8e3ea8409418949fb4106e00893b49983495035b47026a1747eb6f89b05d4c1b8e357e89ddfa7c9f8145b78e0c480177842f20b98b1f7ca2ed26cfb9895380e4d86aeb60c2326bc43753a0633167a7c4ae7a526ce927ade1388a6cdc11d1"),
	},
}

func testRecovery(sigShares []*bls.InsecureSignature, ids []bls.Hash, expectedResult bool) {

	// Try recovery
	thresholdSignature, _ := bls.InsecureSignatureRecover(sigShares, ids)

	// Try verification
	verifyResult := false
	recoveredIsEqual := false
	if thresholdSignature != nil {
		verifyResult = thresholdSignature.Verify([][]byte{expectedSignHash[:]}, []*bls.PublicKey{expectedQuorumPubKey})
		recoveredIsEqual = thresholdSignature.Equal(expectedThresholdSignature)
	}

	if expectedResult != verifyResult {
		log.Panicf("verify result missmatch - expected: %t, actual: %t", expectedResult, verifyResult)
	}

	if expectedResult && !recoveredIsEqual && thresholdSignature != nil {
		log.Panicf("signature missatch - expected: %x, actual: %x", expectedThresholdSignature, thresholdSignature.Serialize())
	}

	if recoveredIsEqual && len(sigShares) < minimumRequiredSignatures {
		log.Panicf("recovery was possible with %d signatures", len(sigShares))
	}
}

func TestLLQMExample(t *testing.T) {

	signHash := bls.BuildSignHash(100, quorumHash, signId, msgHash)
	if !bytes.Equal(expectedSignHash[:], signHash[:]) {
		log.Panicf("signHash does not match: expected: %x, actual: %x", expectedSignHash, signHash)
	}

	var sigShares []*bls.InsecureSignature
	var ids []bls.Hash

	// Create and validate sigShares for each member and populate BLS-IDs from members into ids
	for i, member := range members {
		sig := member.skShare.SignInsecurePrehashed(signHash[:])
		if sig.Equal(member.sigShare) {
			sigShares = append(sigShares, sig)
		} else {
			log.Panicf("member: %d Invalid sigShare %x\n", i, sig.Serialize())
		}
		ids = append(ids, member.blsId)
	}

	testRecovery(sigShares[0:0], ids[0:0], false)
	testRecovery(sigShares[0:0], ids[0:2], false)
	testRecovery(sigShares[0:2], ids[0:0], false)
	testRecovery(sigShares[0:2], ids[0:4], false)
	testRecovery(sigShares[0:4], ids[0:2], false)

	// Validate threshold signature recovery only works with minimum expected signatures
	// and works from there up to the maximum
	rand.Seed(time.Now().UnixNano())
	for round := 0; round < 10; round++ {

		// For 10 rounds shuffle the sigShares/ids each round to try various combinations
		for i := len(sigShares) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			sigShares[i], sigShares[j] = sigShares[j], sigShares[i]
			ids[i], ids[j] = ids[j], ids[i]
		}

		for i := range sigShares {
			testRecovery(sigShares[0:i], ids[0:i], i >= minimumRequiredSignatures)
		}
	}
}
