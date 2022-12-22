package rollinghash

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func Test_computeChunkHashes(t *testing.T) {
	//tests := []struct {
	//	name  string
	//	testData  *os.File
	//	error error
	//}{
	//	{},
	//}
	//
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		decodeText, err := ComputeDelta(tt)
	//		if !reflect.DeepEqual(err, tt.wantErr) {
	//			t.Errorf("Test Interpret error got = %v, want %v", err, tt.wantErr)
	//		}
	//		if strings.Compare(decodeText, tt.decodedText) == -1 {
	//			t.Errorf("Test Interpret got  = %v, want %v", decodeText, tt.decodedText)
	//		}
	//	})
	//}
}

func Test_ComputeDelta(t *testing.T) {

	tests := []struct {
		name             string
		originalFilename string
		updatedFilename  string
		delta            []byte
		error            error
	}{
		{
			name:             "Empty file",
			originalFilename: "./testData/empty.txt",
			updatedFilename:  "./testData/empty.txt",
			delta:            []byte{},
			error:            nil,
		},
		{
			name:             "Happy path with text",
			originalFilename: "./testData/bigfile.txt",
			updatedFilename:  "./testData/updated_bigfile.txt",
			delta:            []byte("2Oxpd4Rp7pR35SqX9i1ETGS5VuIKScfXUa9hlrYs0 Kel4PGMexzRtHy1bJ1gX9fCRxzgS5wHgdnIpOsRaDjnV\nfK9HdmioMlLp86thJiwBZoCeXXxZCuSvaHTT73JuhG62STZaZVPscX57A2F0edcwMzDOfhHkUZaL0cymURxlSGz9Q33OY VIVOOK\nHNy6v6dDNm4wxoN2n Mxc 3V9puGtXwv5WO7XYEktPTV66GNxiZfO9bw63WE6KnEETBt494iATKb1NUFgM3Jq3LGQ FN6tNyulAM\ncHSYZSRcoOrPyre17STKzGtpKWaq7vDmWryWkefL9qYscUSMhxQHi4lYVAqLm43HONlbW4Vp8ZL8f1cuZ48AkZSDOJTTRvXhPrHZ\nII1KXJxO7CImWKPI8w0ivU Hfgasfnm5hhf3ynUlAMRlCcpnTp6sHYBt4tp3mEtJTPsAnvHs4PTaD2jhpvgoyQUPzg ZZpFhkJmC\nqnXq0lSz5ghiGAlA0GqHuHqWfN Z4fGQYdw8NgygLECrw0V5uwix zol7WMR9rFIgVrttHeXROSbBZT1vzw79ShdzLIDWe asP6IFd4bLfUP6Wwang\nyr3IoaPAL6lUwPvbLQkeZFfoYmAryuBTfpicteccOzxT3dUyWLUvELqvfcoqOyONbqU3YX4 dtmJ4Y6U1UPg06lx37veeu51Cc6X\nvNttnAe1UtOHvLYzSVyLJUZJsp2I5OKP  a6RgCnZqRNHSJnk4YZ1ase4PNB6DiTDuBmvUyhjzkTFahCOjYYbhhiI5xoMBPuq0Sl\ndQe8Jpo1jKYtqXXbrslYRnSUYr KEtwuOA5er2vsiSgqn0bf33JLIBLksyXHN6JRVCCYHywClB1HO1ZbXFhBPHScfxMupzQLZlww\n1ZK1ltgudv9wjqxj8CkpRioLxR8nVfDgpFds52 HLCio4yS pwseVpTqg0EsCVXfVsIv8 TlwhI8njLiZbO2jzGiVTrgPBktfSia\n HORYc5gKgXw 82Oxpd4Rp7pR35SqX9i1ETGS5VuIKScfXUa9hlrYs0 Kel4PGMexzRtHy1bJ1gX9fCRxzgS5wHgdnIpOsRaDjnV\nfK9HdmioMlLp86thJiwBZoCeXXxZCuSvaHTT73JuhG62STZaZVPscX57A2F0edcwMzDOfhHkUZaL0cymURxlSGz9Q33OY VIVOOK\nHNy6v6dDNm4wxoN2n Mxc 3V9puGtXwv5WO7XYEktPTV66GNxiZfO9bw63WE6KnEETBt494iATKb1NUFgM3Jq3LGQ FN6tNyulAM\ncHSYZSRcoOrPyre17STKzGtpKWaq7vDmWryWkefL9qYscUSMhxQHi4lYVAqLm43HONlbW4Vp8ZL8f1cuZ48AkZSDOJTTRvXhPrHZ\nII1KXJxO7CImWKPI8w0ivU Hfgasfnm5hhf3ynUlAMRlCcpnTp6sHYBt4tp3mEtJTPsAnvHs4PTaD2jhpvgoyQUPzg ZZpFhkJmC\nqnXq0lSz5ghiGAlA0GqHuHqWfN Z4fGQYdw8NgygLECrw0V5uwix zol7WMR9rFIgVrttHeXROSbBZT1vzw79ShdzLIDWe asP6I"),
			error:            nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original, err := os.Open(tt.originalFilename)
			if err != nil {
				log.Fatal(err)
			}
			defer original.Close()

			updated, err := os.Open(tt.updatedFilename)
			if err != nil {
				log.Fatal(err)
			}
			defer updated.Close()
			delta, err := ComputeDelta(original, updated)
			if err != tt.error {
				t.Errorf("error got = %v, want %v", err, tt.error)
			}

			if bytes.Compare(delta, tt.delta) != 0 {
				t.Errorf("delta got = %v, want %v", string(delta), string(tt.delta))
			}
		})
	}
}
