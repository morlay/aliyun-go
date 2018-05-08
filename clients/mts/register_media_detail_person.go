package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RegisterMediaDetailPersonRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Images               string `position:"Query" name:"Images"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Category             string `position:"Query" name:"Category"`
	PersonName           string `position:"Query" name:"PersonName"`
}

func (req *RegisterMediaDetailPersonRequest) Invoke(client *sdk.Client) (resp *RegisterMediaDetailPersonResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "RegisterMediaDetailPerson", "mts", "")
	resp = &RegisterMediaDetailPersonResponse{}
	err = client.DoAction(req, resp)
	return
}

type RegisterMediaDetailPersonResponse struct {
	responses.BaseResponse
	RequestId            common.String
	RegisteredPersonages RegisterMediaDetailPersonRegisteredPersonageList
	FailedImages         RegisterMediaDetailPersonFailedImageList
}

type RegisterMediaDetailPersonRegisteredPersonage struct {
	PersonName common.String
	FaceId     common.String
	Target     common.String
	Quality    common.String
	Gender     common.String
	ImageId    common.String
	ImageFile  RegisterMediaDetailPersonImageFile
}

type RegisterMediaDetailPersonImageFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type RegisterMediaDetailPersonFailedImage struct {
	Code       common.String
	Success    common.String
	ImageFile1 RegisterMediaDetailPersonImageFile1
}

type RegisterMediaDetailPersonImageFile1 struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type RegisterMediaDetailPersonRegisteredPersonageList []RegisterMediaDetailPersonRegisteredPersonage

func (list *RegisterMediaDetailPersonRegisteredPersonageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RegisterMediaDetailPersonRegisteredPersonage)
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

type RegisterMediaDetailPersonFailedImageList []RegisterMediaDetailPersonFailedImage

func (list *RegisterMediaDetailPersonFailedImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RegisterMediaDetailPersonFailedImage)
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
