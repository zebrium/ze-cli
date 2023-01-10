package common

import "testing"

func TestValidateBatchId(t *testing.T) {
	goodBatch := "BJADAa123421-_"
	badBatch := "ajhfhadsh&@3"
	err := ValidateBatchId(goodBatch)
	if err != nil {
		t.Fatalf("Batch ID: %s failed Validation, when it was valid", goodBatch)
	}
	err = ValidateBatchId(badBatch)
	if err == nil {
		t.Fatalf("Invalid Batch ID: %s was validated successfully even though it shouldnt of been", badBatch)
	}
	err = ValidateBatchId("")
	if err == nil {
		t.Fatalf("Invalid Batch ID: \"\" was validated successfully even though it shouldnt of been")
	}
}

func TestValidateAPIToken(t *testing.T) {
	goodToken := "IHBpvlLxPDMxbzuGFSnryCgJHNfynVziTtryXstRiUFTqnviEhKURQQxBrWknxEW"
	badToken := "IHBpvlLxPDMxbzuGFSnryCgJHNfynVziTtryXstRiUFTqnviEhKUEW"
	badToken2 := "IHBpvlLxaPDMxbzuGFSnryCgJHNfynVziTtryXstRiUFTnviEh!URQQxBrWknxEW"

	err := ValidateAPIToken("")
	if err == nil {
		t.Fatalf("Invalid API Token: \"\" was validated successfully even though it shouldnt of been")
	}
	err = ValidateAPIToken(goodToken)
	if err != nil {
		t.Fatalf("API Token: %s failed Validation, when it was valid", goodToken)
	}
	err = ValidateAPIToken(badToken)
	if err == nil {
		t.Fatalf("Invalid API Token: %s was validated successfully even though it shouldnt of been", badToken)
	}
	err = ValidateAPIToken(badToken2)
	if err == nil {
		t.Fatalf("Invalid API Token: %s was validated successfully even though it shouldnt of been", badToken2)
	}
}

func TestValidateAuthToken(t *testing.T) {
	goodToken := "00FAFE8422A968BFCA6C7FD08ED9DC4D5242B297"
	badToken := "00FAFE8422A968BFCA6C7FD08ED9DC4D52"

	err := ValidateAuthToken("")
	if err == nil {
		t.Fatalf("Invalid Auth Token: \"\" was validated successfully even though it shouldnt of been")
	}
	err = ValidateAuthToken(goodToken)
	if err != nil {
		t.Fatalf("Auth Token: %s failed Validation, when it was valid", goodToken)
	}
	err = ValidateAuthToken(badToken)
	if err == nil {
		t.Fatalf("Invalid Auth Token: %s was validated successfully even though it shouldnt of been", badToken)
	}
}
func TestValidateZapiUrl(t *testing.T) {
	goodUrl := "https://test.com"
	badUrl := ":/test.com"
	badurl2 := "https:/test.com"

	err := ValidateZapiUrl("")
	if err == nil {
		t.Fatalf("Invalid Auth Token: \"\" was validated successfully even though it shouldnt of been")
	}
	err = ValidateZapiUrl(goodUrl)
	if err != nil {
		t.Fatalf("URL: %s failed Validation, when it was valid", goodUrl)
	}
	err = ValidateZapiUrl(badUrl)
	if err == nil {
		t.Fatalf("Invalid Url: %s was validated successfully even though it shouldnt of been", badUrl)
	}
	err = ValidateZapiUrl(badurl2)
	if err == nil {
		t.Fatalf("Invalid Url: %s was validated successfully even though it shouldnt of been", badurl2)
	}
}

func TestValidateUpMetaData(t *testing.T) {
	//Validate Log type for streaming
	err := ValidateUpMetadata("", "", false, "", "")
	if err == nil {
		t.Fatal("Logtype and filename cannot be empty while logstash is flase")
	}
	err = ValidateUpMetadata("", "", true, "", "")
	if err != nil {
		t.Fatal("Logtype and filename can be empty while logstash is true")
	}
	err = ValidateUpMetadata("", "test123", false, "", "")
	if err != nil {
		t.Fatal("filename can be empty while logstash is false as long as logtype is set")
	}
	err = ValidateUpMetadata("", "test123", false, "123456", "")
	if err != nil {
		t.Fatal("Batch can be set if cfgs does not include ze_batch_id")
	}
	err = ValidateUpMetadata("", "test123", false, "", "ze_batch_id=1234156")
	if err != nil {
		t.Fatal("batch id can be empty if cfgs includes ze_batch_id")
	}
	err = ValidateUpMetadata("", "test123", false, "123456", "ze_batch_id=1234156")
	if err == nil {
		t.Fatal("batch id cannot be set if cfgs includes ze_batch_id")
	}
}
