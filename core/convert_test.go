package core
import "testing"
// go test -v -timeout 10m -run TestM2TS2MKV
func TestM2TS2MKV(t *testing.T) {
	if err := m2ts2mkv("C:\\Users\\zhang\\Github\\makemkv\\00003.m2ts"); err != nil {
		t.Error(err)
	}
}
