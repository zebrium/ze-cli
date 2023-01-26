package common

import "testing"

func TestValidateBatchId(t *testing.T) {
	testCases := []struct {
		batch string
		valid bool
	}{
		{
			batch: "BJADAa123421-_",
			valid: true,
		},
		{
			batch: "ajhfhadsh&@3",
			valid: false,
		},
		{
			batch: "-ajhfhadsh&@3",
			valid: false,
		},
		{
			batch: "_ajhfhadsh&@3",
			valid: false,
		},
		{
			batch: "asdadfaewfdasfasgaresfasdaewfadsfaewgasfdaewr",
			valid: false,
		},
		{
			batch: "",
			valid: false,
		},
	}

	for _, tc := range testCases {
		err := ValidateBatchId(tc.batch)
		if err != nil {
			if tc.valid == true {
				t.Fatalf("Batch Id: %s was marked invalid but should of been valid", tc.batch)

			}
		} else {
			if tc.valid == false {
				t.Fatalf("Batch Id: %s was marked Valid but should of been invalid", tc.batch)
			}
		}
	}
}

func TestValidateAPIToken(t *testing.T) {
	testCases := []struct {
		token string
		valid bool
	}{
		{
			token: "IHBpvlLxPDMxbzuGFSnryCgJHNfynVziTtryXstRiUFTqnviEhKURQQxBrWknxEW",
			valid: true,
		},
		{
			token: "IHBpvlLxPDMxbzuGFSnryCgJHNfynVziTtryXstRiUFTqnviEhKUEW",
			valid: false,
		},
		{
			token: "IHBpvlLxaPDMxbzuGFSnryCgJHNfynVziTtryXstRiUFTnviEh!URQQxBrWknxEW",
			valid: false,
		},
		{
			token: "",
			valid: false,
		},
	}
	for _, tc := range testCases {
		err := ValidateAPIToken(tc.token)
		if err != nil {
			if tc.valid == true {
				t.Fatalf("Token: %s was marked invalid but should of been valid", tc.token)

			}
		} else {
			if tc.valid == false {
				t.Fatalf("Token: %s was marked Valid but should of been invalid", tc.token)
			}
		}
	}
}

func TestValidateAuthToken(t *testing.T) {
	testCases := []struct {
		token string
		valid bool
	}{
		{
			token: "00FAFE8422A968BFCA6C7FD08ED9DC4D5242B297",
			valid: true,
		},
		{
			token: "00FAFE8422A968BFCA6C7FD08ED9DC4D52",
			valid: false,
		},
		{
			token: "IHBpvlLxaPDMxbzuGFSnryCgJHNfynVziTtryXstRiUFTnviEh!URQQxBrWknxEW",
			valid: false,
		},
		{
			token: "",
			valid: false,
		},
	}
	for _, tc := range testCases {
		err := ValidateAuthToken(tc.token)
		if err != nil {
			if tc.valid == true {
				t.Fatalf("Token: %s was marked invalid but should of been valid", tc.token)

			}
		} else {
			if tc.valid == false {
				t.Fatalf("Token: %s was marked Valid but should of been invalid", tc.token)
			}
		}
	}
}
func TestValidateZapiUrl(t *testing.T) {

	testCases := []struct {
		url   string
		valid bool
	}{
		{
			url:   "https://test.com",
			valid: true,
		},
		{
			url:   ":/test.com",
			valid: false,
		},
		{
			url:   "https:/test.com",
			valid: false,
		},
		{
			url:   "",
			valid: false,
		},
	}
	for _, tc := range testCases {
		err := ValidateZapiUrl(tc.url)
		if err != nil {
			if tc.valid == true {
				t.Fatalf("URL: %s was marked invalid but should of been valid", tc.url)

			}
		} else {
			if tc.valid == false {
				t.Fatalf("URL: %s was marked Valid but should of been invalid", tc.url)
			}
		}
	}
}

func TestValidateUpMetaData(t *testing.T) {
	testCases := []struct {
		filename string
		logstash bool
		logtype  string
		batchid  string
		cfgs     string
		valid    bool
		reason   string
	}{
		{
			filename: "",
			logtype:  "",
			logstash: false,
			batchid:  "",
			cfgs:     "",
			valid:    false,
			reason:   "Logtype and filename cannot be empty while logstash is false",
		},
		{
			filename: "",
			logtype:  "",
			logstash: true,
			batchid:  "",
			cfgs:     "",
			valid:    true,
			reason:   "Logtype and filename can be empty while logstash is true",
		},
		{
			filename: "",
			logtype:  "test123",
			logstash: false,
			batchid:  "",
			cfgs:     "",
			valid:    true,
			reason:   "filename can be empty while logstash is false as long as logtype is set",
		},
		{
			filename: "",
			logtype:  "test123",
			logstash: false,
			batchid:  "123456",
			cfgs:     "",
			valid:    true,
			reason:   "Batch can be set if cfgs does not include ze_batch_id",
		},
		{
			filename: "",
			logtype:  "test123",
			logstash: false,
			batchid:  "",
			cfgs:     "ze_batch_id=1234156",
			valid:    true,
			reason:   "batch id can be empty if cfgs includes ze_batch_id",
		},
		{
			filename: "",
			logtype:  "test123",
			logstash: false,
			batchid:  "123456",
			cfgs:     "ze_batch_id=1234156",
			valid:    false,
			reason:   "batch id cannot be set if cfgs includes ze_batch_id",
		},
	}
	for _, tc := range testCases {
		err := ValidateUpMetadata(tc.filename, tc.logtype, tc.logstash, tc.batchid, tc.cfgs)
		if err != nil {
			if tc.valid == true {
				t.Fatal(tc)

			}
		} else {
			if tc.valid == false {
				t.Fatal(tc)
			}
		}
	}
}
