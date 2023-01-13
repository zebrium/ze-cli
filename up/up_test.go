package up

import (
	"testing"
)

var url = "https://test.local"
var auth = "xxxxxxxxxxxxxxxxxxxxxxx"

func TestCreateMap(t *testing.T) {
	t.Log("Testing 3 Values")
	set1 := "key1=value1,key2=value2,key3=value3"
	actual := createMap(set1)
	if len(actual) != 3 {
		t.Fatalf("Incorrect number of objects created in map.  Expected: %d, Got: %d", 3, len(actual))
	}
	t.Log("Testing 0 values for array")
	set2 := ""
	actual = createMap(set2)
	if len(actual) != 0 {
		t.Fatalf("Incorrect number of objects created in map.  Expected: %d, Got: %d", 0, len(actual))
	}
}

func TestMetadataBatchPassed(t *testing.T) {
	batch := "test123456"
	version := "1.0.0"
	metadata, existingBatch, updatedBatch, err := generateMetadata(url, auth, "test.log", "", "", "", "", "", "", "", batch, false, version)
	if err != nil {
		t.Fatal(err)
	}
	if existingBatch != true {
		t.Fatal("existingBatch was set to false when it should of been true")
	}
	if metadata.Cfgs["ze_batch_id"] != batch {
		t.Fatalf("Batchids did not match what was expected. Expected: %s, Actual %s", batch, metadata.Cfgs["ze_batch_id"])
	}
	if batch != updatedBatch {
		t.Fatalf("Expected: %s Actual: %s", batch, updatedBatch)
	}

}
func TestMetadataBatchConfig(t *testing.T) {
	batch := "test123456"
	metadata, existingBatch, updatedBatch,  err := generateMetadata(url, auth, "test.log", "", "", "", "", "", "ze_batch_id="+batch, "", "", false, "test")
	if err != nil {
		t.Fatal(err)
	}
	if existingBatch != true {
		t.Fatal("existingBatch was set to false when it should of been true")
	}
	if metadata.Cfgs["ze_batch_id"] != batch {
		t.Fatalf("Batchids did not match what was expected. Expected: %s, Actual %s", batch, metadata.Cfgs["ze_batch_id"])
	}
	if batch != updatedBatch {
		t.Fatalf("Expected: %s Actual: %s", batch, updatedBatch)
	}
}

func TestMetadataFileWithNoLogType(t *testing.T) {
	batch := "test123456"
	filename := "test_one-1234.log"
	t.Log("Test with .log")
	metadata, existingBatch, updatedBatch, err := generateMetadata(url, auth, filename, "", "", "", "", "", "", "", batch, false, "test")
	if err != nil {
		t.Fatal(err)
	}
	if existingBatch != true {
		t.Fatal("existingBatch was set to False when it should of been True")
	}
	if metadata.LogBaseName != "test_one-1234" {
		t.Fatalf("LogBaseName incorrectly set.  Expected: %s, Actual: %s", "test_one", metadata.LogBaseName)
	}
	if batch != updatedBatch {
		t.Fatalf("Expected: %s Actual: %s", batch, updatedBatch)
	}
	t.Log("Test with no .")
	metadata, existingBatch,updatedBatch, err = generateMetadata(url, auth, "test_one-1234", "", "", "", "", "", "", "", batch, false, "test")
	if err != nil {
		t.Fatal(err)
	}
	if existingBatch != true {
		t.Fatal("existingBatch was set to False when it should of been True")
	}
	if metadata.LogBaseName != "test_one-1234" {
		t.Fatalf("LogBaseName incorrectly set.  Expected: %s, Actual: %s", "test_one", metadata.LogBaseName)
	}
	if metadata.Stream != "zefile" {
		t.Fatalf("Incorrect setting of Stream.  For  files, it should be zefile. Actual: %s", metadata.Stream)

	}
}

func TestMetadataStreaming(t *testing.T) {
	logtype := "peanutbutter"
	metadata, existingBatch, updatedBatch, err := generateMetadata(url, auth, "", logtype, "", "", "", "", "", "", "", false, "test")
	if err != nil {
		t.Fatal(err)
	}
	if existingBatch != true {
		t.Fatal("existingBatch was set to True when it should of been False")
	}
	if len(metadata.Cfgs["ze_batch_id"]) != 0 {
		t.Fatalf("BatchId was set even though streaming was enabled.")
	}
	if metadata.LogBaseName != logtype {
		t.Fatalf("Incorrect Log Type Set.  Expected: %s Actual: %s", logtype, metadata.LogBaseName)
	}
	if metadata.Stream != "native" {
		t.Fatalf("Incorrect setting of Stream.  For non files, it should be native. Actual: %s", metadata.Stream)

	}
	if updatedBatch != "" {
		t.Fatalf("Got %s for batchId when should of been empty", updatedBatch)
	}
}

func TestMetadataGeneral(t *testing.T) {
	logtype := "peanutbutter"
	host := "countlogula"
	svcgrp := "jelly"
	tz := "EST"
	metadata, existingBatch, updatedBatch, err := generateMetadata(url, auth, "", logtype, host, svcgrp, tz, "", "", "", "123", false, "test")
	if err != nil {
		t.Fatal(err)
	}
	if existingBatch != true {
		t.Fatal("existingBatch was set to False when it should of been True")
	}
	if metadata.Tz != tz {
		t.Fatalf("Incorrect setting of TZ.  Expected: %s Actual: %s", tz, metadata.Tz)

	}
	if metadata.TM != false {
		t.Fatalf("TM is set to true and should always be set to false")

	}
	if metadata.Ids["zid_host"] != host {
		t.Fatalf("Expected: %s Actual: %s", host, metadata.Ids["zid_host"])

	}
	if metadata.Ids["ze_deployment_name"] != svcgrp {
		t.Fatalf("Expected: %s Actual: %s", host, metadata.Ids["ze_deployment_name"])

	}
	if updatedBatch != "123" {
		t.Fatalf("Expected %s Actual %s", "123", updatedBatch)
	}
}

func TestRequestBuilder(t *testing.T) {

}
