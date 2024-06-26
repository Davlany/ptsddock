package sha_test

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestSha256(t *testing.T) {
	testData := []string{"test string", "test 2 string", "test 3 string", "Бульбозавр","random"}
	expectedData := []string{"d5579c46dfcc7f18207013e65b44e4cb4e2c2298f4ac457ba8f82743f31e930b", "608d2dd13cf69bc05ed2f8195f0a8579b7921d74fcc482a5f5e98a27481ab2fe", "1d6c7a8fa66900fc96070be9c0d8837af6d1eca470bb94012e19027800c182e1","c1be68a2e89dbf4f119f5ba03fcf5644925bed75d183d1f81bbab8be492ff5a3","a441b15fe9a3cf56661190a0b93b9dec7d04127288cc87250967cf3b52894d11"}
	for i, data := range testData {
		hashData := sha256.Sum256([]byte(data))
		if hex.EncodeToString(hashData[:]) == expectedData[i] {
			t.Logf("Test %d succeced", i)
		} else {
			t.Errorf("Test %d failed\nexpected: %s\nget: %s", i, expectedData[i], hex.EncodeToString(hashData[:]))
		}

	}

}
