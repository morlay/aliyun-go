package kms

import (
	"encoding/json"
)

func (c *KmsClient) CreateKey(req *CreateKeyArgs) (resp *CreateKeyResponse, err error) {
	resp = &CreateKeyResponse{}
	err = c.Request("CreateKey", req, resp)
	return
}

type CreateKeyKeyMetadata struct {
	CreationDate string
	Description  string
	KeyId        string
	KeyState     string
	KeyUsage     string
	DeleteDate   string
	Creator      string
	Arn          string
}
type CreateKeyArgs struct {
	Description string
	KeyUsage    string
	STSToken    string
}
type CreateKeyResponse struct {
	RequestId   string
	KeyMetadata CreateKeyKeyMetadata
}

func (c *KmsClient) GenerateDataKey(req *GenerateDataKeyArgs) (resp *GenerateDataKeyResponse, err error) {
	resp = &GenerateDataKeyResponse{}
	err = c.Request("GenerateDataKey", req, resp)
	return
}

type GenerateDataKeyArgs struct {
	KeyId             string
	KeySpec           string
	NumberOfBytes     int
	STSToken          string
	EncryptionContext string
}
type GenerateDataKeyResponse struct {
	CiphertextBlob string
	KeyId          string
	Plaintext      string
	RequestId      string
}

func (c *KmsClient) EnableKey(req *EnableKeyArgs) (resp *EnableKeyResponse, err error) {
	resp = &EnableKeyResponse{}
	err = c.Request("EnableKey", req, resp)
	return
}

type EnableKeyArgs struct {
	KeyId    string
	STSToken string
}
type EnableKeyResponse struct {
	RequestId string
}

func (c *KmsClient) DescribeKey(req *DescribeKeyArgs) (resp *DescribeKeyResponse, err error) {
	resp = &DescribeKeyResponse{}
	err = c.Request("DescribeKey", req, resp)
	return
}

type DescribeKeyKeyMetadata struct {
	CreationDate string
	Description  string
	KeyId        string
	KeyState     string
	KeyUsage     string
	DeleteDate   string
	Creator      string
	Arn          string
}
type DescribeKeyArgs struct {
	KeyId    string
	STSToken string
}
type DescribeKeyResponse struct {
	RequestId   string
	KeyMetadata DescribeKeyKeyMetadata
}

func (c *KmsClient) Decrypt(req *DecryptArgs) (resp *DecryptResponse, err error) {
	resp = &DecryptResponse{}
	err = c.Request("Decrypt", req, resp)
	return
}

type DecryptArgs struct {
	CiphertextBlob    string
	STSToken          string
	EncryptionContext string
}
type DecryptResponse struct {
	Plaintext string
	KeyId     string
	RequestId string
}

func (c *KmsClient) Encrypt(req *EncryptArgs) (resp *EncryptResponse, err error) {
	resp = &EncryptResponse{}
	err = c.Request("Encrypt", req, resp)
	return
}

type EncryptArgs struct {
	KeyId             string
	Plaintext         string
	STSToken          string
	EncryptionContext string
}
type EncryptResponse struct {
	CiphertextBlob string
	KeyId          string
	RequestId      string
}

func (c *KmsClient) ScheduleKeyDeletion(req *ScheduleKeyDeletionArgs) (resp *ScheduleKeyDeletionResponse, err error) {
	resp = &ScheduleKeyDeletionResponse{}
	err = c.Request("ScheduleKeyDeletion", req, resp)
	return
}

type ScheduleKeyDeletionArgs struct {
	KeyId               string
	PendingWindowInDays int
	STSToken            string
}
type ScheduleKeyDeletionResponse struct {
	RequestId string
}

func (c *KmsClient) DisableKey(req *DisableKeyArgs) (resp *DisableKeyResponse, err error) {
	resp = &DisableKeyResponse{}
	err = c.Request("DisableKey", req, resp)
	return
}

type DisableKeyArgs struct {
	KeyId    string
	STSToken string
}
type DisableKeyResponse struct {
	RequestId string
}

func (c *KmsClient) CancelKeyDeletion(req *CancelKeyDeletionArgs) (resp *CancelKeyDeletionResponse, err error) {
	resp = &CancelKeyDeletionResponse{}
	err = c.Request("CancelKeyDeletion", req, resp)
	return
}

type CancelKeyDeletionArgs struct {
	KeyId    string
	STSToken string
}
type CancelKeyDeletionResponse struct {
	RequestId string
}

func (c *KmsClient) DescribeRegions(req *DescribeRegionsArgs) (resp *DescribeRegionsResponse, err error) {
	resp = &DescribeRegionsResponse{}
	err = c.Request("DescribeRegions", req, resp)
	return
}

type DescribeRegionsRegion struct {
	RegionId string
}
type DescribeRegionsArgs struct {
	STSToken string
}
type DescribeRegionsResponse struct {
	RequestId string
	Regions   DescribeRegionsRegionList
}

type DescribeRegionsRegionList []DescribeRegionsRegion

func (list *DescribeRegionsRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsRegion)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

func (c *KmsClient) ListKeys(req *ListKeysArgs) (resp *ListKeysResponse, err error) {
	resp = &ListKeysResponse{}
	err = c.Request("ListKeys", req, resp)
	return
}

type ListKeysKey struct {
	KeyId  string
	KeyArn string
}
type ListKeysArgs struct {
	PageNumber int
	PageSize   int
	STSToken   string
}
type ListKeysResponse struct {
	TotalCount int
	PageNumber int
	PageSize   int
	RequestId  string
	Keys       ListKeysKeyList
}

type ListKeysKeyList []ListKeysKey

func (list *ListKeysKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListKeysKey)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
