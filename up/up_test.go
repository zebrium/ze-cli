package up

import (
	"testing"
)

var url = "https://test.local"
var auth = "xxxxxxxxxxxxxxxxxxxxxxx"

func TestCreateMap(t *testing.T) {
	testCases := []struct {
		set string
		len int
	}{
		{
			set: "",
			len: 0,
		},
		{
			set: "key1=value1",
			len: 1,
		},
		{
			set: "key1=value1,key2=value2,key3=value3",
			len: 3,
		},
	}
	for i, tc := range testCases {
		actual := createMap(tc.set)
		if len(actual) != tc.len {
			t.Fatalf("Subtest: %d failed. Incorrect number of objects created in map.  Expected: %d, Got: %d", i, tc.len, len(actual))
		}
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
	metadata, existingBatch, updatedBatch, err := generateMetadata(url, auth, "test.log", "", "",
		"", "", "", "ze_batch_id="+batch, "", "", false, "test")
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
	metadata, existingBatch, updatedBatch, err := generateMetadata(url, auth, filename, "", "", "",
		"", "", "", "", batch, false, "test")
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
	metadata, existingBatch, _, err = generateMetadata(url, auth, "test_one-1234", "", "",
		"", "", "", "", "", batch, false, "test")
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
	metadata, existingBatch, updatedBatch, err := generateMetadata(url, auth, "", logtype, "", "",
		"", "", "", "", "", false, "test")
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
	metadata, existingBatch, updatedBatch, err := generateMetadata(url, auth, "", logtype, host, svcgrp, tz,
		"", "", "", "123", false, "test")
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

func TestMetadataFilename(t *testing.T) {
	testCases := []struct {
		filename    string
		expectedLBN string
	}{
		{
			filename:    "test_one-1234.log",
			expectedLBN: "test_one-1234",
		},
		{
			filename:    "../../test_one-1234.log",
			expectedLBN: "test_one-1234",
		},
		{
			filename:    "/this/is/a/test/test_one-1234.log",
			expectedLBN: "test_one-1234",
		},
		{
			filename:    "test_one.1.log",
			expectedLBN: "test_one",
		},
		{
			filename:    "test_one",
			expectedLBN: "test_one",
		},
		{
			filename:    "../../test_one",
			expectedLBN: "test_one",
		},
		{
			filename:    "/this/is/a/test/test_one",
			expectedLBN: "test_one",
		},
	}
	for i, tc := range testCases {
		metadata, _, _, err := generateMetadata(url, auth, tc.filename, "", "", "", "", "", "", "", "test", false, "")
		if err != nil {
			t.Fatal(err)
		}
		if metadata.LogBaseName != tc.expectedLBN {
			t.Fatalf("Subtest: %d failed. Expected: %s, Actual: %s", i, tc.expectedLBN, metadata.LogBaseName)
		}
	}
}

func TestParseBatchIdFromConfig(t *testing.T) {
	testCases := []struct {
		cfgs    string
		batchId string
	}{
		{
			cfgs:    "",
			batchId: "",
		},
		{
			cfgs:    "key1=value,key2=value",
			batchId: "",
		},
		{
			cfgs:    "key1=value,key2=value,ze_batch_id=test123",
			batchId: "test123",
		},
		{
			cfgs:    "ze_batch_id=test123",
			batchId: "test123",
		},
	}
	for i, tc := range testCases {
		batchId := parseBatchIdFromConfigs(tc.cfgs)
		if batchId != tc.batchId {
			t.Fatalf("Subtest: %d failed. Expected: %s, Actual: %s", i, tc.batchId, batchId)
		}
	}
}
