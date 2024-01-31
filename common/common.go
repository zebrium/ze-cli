// Package common Copyright © 2024 ScienceLogic Inc
package common

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"regexp"
	"strings"
)

// ValidateAuthToken validates that the auth token is correct
func ValidateAuthToken(auth string) error {
	if len(auth) == 0 {
		return errors.New("auth token must be set using the --auth flag")
	}
	if len(auth) != 40 {
		return errors.New("invalid auth token.")
	}
	return nil
}

// ValidateZapiUrl validates Zapi Url
func ValidateZapiUrl(url string) error {
	if len(url) == 0 {
		return errors.New("url must be set using the --url flag")
	}
	if !govalidator.IsURL(url) {
		return fmt.Errorf("url: %s is not a valid url", url)
	}
	return nil
}

// ValidateAPIToken validates api token
func ValidateAPIToken(api string) error {
	if len(api) == 0 {
		return errors.New("API must be set using the --api flag")
	}
	if len(api) != 64 {
		return errors.New("invalid api token")
	}
	result, err := regexp.MatchString("^[A-Za-z0-9]*$", api)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("invalid api token")
	}
	return nil
}

// ValidateUpMetadata validate the business logic around metadata
func ValidateUpMetadata(filename string, logype string, logstash bool, batchId string, cfgs string) error {
	// Make sure log type is specified
	if len(filename) == 0 {
		if len(logype) == 0 && !logstash {
			return errors.New("error: logtype must be specified for streaming with --log")
		}
	}

	if strings.Contains(cfgs, "ze_batch_id") && len(batchId) != 0 {
		return errors.New("ze_batch_id is defined in cfgs put also specified with --batch.  Please correct conflict.")
	}
	return nil
}

// ValidateBatchId Validates batchId
func ValidateBatchId(batchId string) error {
	if len(batchId) == 0 {
		return errors.New("BatchId must be set with -b")
	}
	result, err := regexp.Match("^[a-zA-Z0-9][a-zA-Z0-9_-]*$", []byte(batchId))
	if err != nil {
		return err
	}
	if !result {
		return fmt.Errorf("BatchId %s contains invalid characters.  Must contain alphanumeric characters, '_' and '-'", batchId)
	}
	if len(batchId) > 36 {
		return fmt.Errorf("BatchId: %s exceeds the max length of 36 characters", batchId)
	}
	return nil
}
