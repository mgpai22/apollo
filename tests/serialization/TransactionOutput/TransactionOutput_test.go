package transactionoutput_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/Salvionied/cbor/v2"
	"github.com/SundaeSwap-finance/apollo/serialization/Address"
	"github.com/SundaeSwap-finance/apollo/serialization/PlutusData"
	"github.com/SundaeSwap-finance/apollo/serialization/Transaction"
	"github.com/SundaeSwap-finance/apollo/serialization/TransactionOutput"
	"github.com/SundaeSwap-finance/apollo/serialization/Value"
)

func TestTransactionOutputWithDatumHash(t *testing.T) {
	cborHex := "83583911a65ca58a4e9c755fa830173d2a5caed458ac0c73f97db7faae2e7e3b52563c5410bff6a0d43ccebb7c37e1f69f5eb260552521adff33b9c21a00895440582070c5d760293d3d92bfa7369e472891ab36041cbf81fd5ed103462fb7c03f2a6e"
	cborBytes, _ := hex.DecodeString(cborHex)
	txOut := TransactionOutput.TransactionOutput{}
	err := txOut.UnmarshalCBOR(cborBytes)
	if err != nil {
		t.Errorf("Error while unmarshaling")
	}
	outBytes, err := cbor.Marshal(txOut)
	if err != nil {
		t.Errorf("Error while marshaling")
	}

	if hex.EncodeToString(outBytes) != cborHex {
		t.Errorf("Invalid Reserialization")
	}
}

func TestPostAlonzo(t *testing.T) {
	txO := TransactionOutput.TransactionOutput{}
	cborHex := "d8799fd8799fd8799f581c37dce7298152979f0d0ff71fb2d0c759b298ac6fa7bc56b928ffc1bcffd8799fd8799fd8799f581cf68864a338ae8ed81f61114d857cb6a215c8e685aa5c43bc1f879cceffffffffd8799fd8799f581c37dce7298152979f0d0ff71fb2d0c759b298ac6fa7bc56b928ffc1bcffd8799fd8799fd8799f581cf68864a338ae8ed81f61114d857cb6a215c8e685aa5c43bc1f879cceffffffffd87a80d8799fd8799f581c25f0fc240e91bd95dcdaebd2ba7713fc5168ac77234a3d79449fc20c47534f4349455459ff1b00002cc16be02b37ff1a001e84801a001e8480ff"
	decoded_cbor, _ := hex.DecodeString(cborHex)
	var pd PlutusData.PlutusData
	cbor.Unmarshal(decoded_cbor, &pd)
	txO.IsPostAlonzo = true
	decoded_address, _ := Address.DecodeAddress("addr1wynp362vmvr8jtc946d3a3utqgclfdl5y9d3kn849e359hsskr20n")
	txO.PostAlonzo = TransactionOutput.TransactionOutputAlonzo{}
	txO.PostAlonzo.Address = decoded_address
	txO.PostAlonzo.Amount = Value.PureLovelaceValue(1000000).ToAlonzoValue()
	d := PlutusData.DatumOptionLiteral(&pd)
	txO.PostAlonzo.Datum = &d
	resultHex := "a300581d712618e94cdb06792f05ae9b1ec78b0231f4b7f4215b1b4cf52e6342de01821a000f4240a0028201d81858e8d8799fd8799fd8799f581c37dce7298152979f0d0ff71fb2d0c759b298ac6fa7bc56b928ffc1bcffd8799fd8799fd8799f581cf68864a338ae8ed81f61114d857cb6a215c8e685aa5c43bc1f879cceffffffffd8799fd8799f581c37dce7298152979f0d0ff71fb2d0c759b298ac6fa7bc56b928ffc1bcffd8799fd8799fd8799f581cf68864a338ae8ed81f61114d857cb6a215c8e685aa5c43bc1f879cceffffffffd87a80d8799fd8799f581c25f0fc240e91bd95dcdaebd2ba7713fc5168ac77234a3d79449fc20c47534f4349455459ff1b00002cc16be02b37ff1a001e84801a001e8480ff"
	cborred, _ := cbor.Marshal(txO)
	if hex.EncodeToString(cborred) != resultHex {
		fmt.Println(hex.EncodeToString(cborred))
		t.Errorf("Invalid marshaling")
	}

}

func TestDeSerializeTxWithPostAlonzoOut(t *testing.T) {
	cborHex := "84a500838258205628043acaccaf3e07ce6d93bec8da6ae013d2546aa1f491c68dfa2942e6aab401825820250cb6fab4bab5fe0746748cdb8dd42b545328ecc8109e16cd56c0ca9382c7bb028258205e9344d4529b623cb1e17b5a041f58f8275e0fdea54c52a7dc73e0d47ff2fe1a010183a300581d712618e94cdb06792f05ae9b1ec78b0231f4b7f4215b1b4cf52e6342de01821a00e4e1c0a0028201d81858bfd8799fd8799f4040ffd8799f581cf43a62fdc3965df486de8a0d32fe800963589c41b38946602a0dc5354441474958ffd8799f581cfd011feb9dc34f85e58e56838989816343f5c62619a82f6a089f05484c414749585f4144415f4e4654ff1903e51b002904d642c7b27c1b7fffffffffffffff581c37dce7298152979f0d0ff71fb2d0c759b298ac6fa7bc56b928ffc1bcd8799f581cf68864a338ae8ed81f61114d857cb6a215c8e685aa5c43bc1f879cceff1a009896801a4d6fd4bcff82583901bb2ff620c0dd8b0adc19e6ffadea1a150c85d1b22d05e2db10c55c613b8c8a100c16cf62b9c2bacc40453aaa67ced633993f2b4eec5b88e41a000fea4c8258390137dce7298152979f0d0ff71fb2d0c759b298ac6fa7bc56b928ffc1bcf68864a338ae8ed81f61114d857cb6a215c8e685aa5c43bc1f879cce821a0633d59aab581c10a49b996e2402269af553a8a96fb8eb90d79e9eca79e2b4223057b6a1444745524f1a001e8480581c25f0fc240e91bd95dcdaebd2ba7713fc5168ac77234a3d79449fc20ca147534f43494554591b00000019e1ae3741581c279c909f348e533da5808898f87f9a14bb2c3dfbbacccd631d927a3fa144534e454b1928b0581c29d222ce763455e3d7a09a665ce554f00ac89d2e99a1a83d267170c6a1434d494e1a0cb30355581c533bb94a8850ee3ccbe483106489399112b74c905342cb1792a797a0a144494e44591a156f14e4581c5d16cc1a177b5d9ba9cfa9793b07e60f1fb70fea1f8aef064415d114a1434941471b0000002e921a6381581c8a1cfae21368b8bebbbed9800fec304e95cce39a2a57dc35e2e3ebaaa1444d494c4b05581c8b4e239aef4d1d1bc5dd628ff3ce34d392d632e5cda83e42d6fcb1cca14b586572636865723234393301581cd480f68af028d6324ad77df489176e7f5e5d793e09a6b133392ff2f6aa524e7563617374496e63657074696f6e31343101524e7563617374496e63657074696f6e32303601524e7563617374496e63657074696f6e33323101524e7563617374496e63657074696f6e33383501524e7563617374496e63657074696f6e34303001524e7563617374496e63657074696f6e36333701524e7563617374496e63657074696f6e36373001524e7563617374496e63657074696f6e37383701524e7563617374496e63657074696f6e38333301524e7563617374496e63657074696f6e38373001581ce3ff4ab89245ede61b3e2beab0443dbcc7ea8ca2c017478e4e8990e2a549746170707930333831014974617070793034313901497461707079313430390149746170707931343437014974617070793135353001581cf0ff48bbb7bbe9d59a40f1ce90e9e9d0ff5002ec48f232b49ca0fb9aa24a626c7565646573657274014a6d6f6e74626c616e636f01021a000342dd031a05fd33e3081a05fd32b7a1049ffff5f6"

	decoded_cbor, _ := hex.DecodeString(cborHex)
	var tx Transaction.Transaction

	err := cbor.Unmarshal(decoded_cbor, &tx)
	if err != nil {
		t.Error("Error while unmarshaling", err)
	}
	fmt.Println(tx.TransactionBody.Outputs[0].PostAlonzo.Datum)
	remarshaled, err := cbor.Marshal(tx)
	if err != nil {
		t.Error("Error While remarshaling", err)
	}
	if hex.EncodeToString(remarshaled) != cborHex {
		t.Error("Error while reserializing", hex.EncodeToString(remarshaled))
	}

}
