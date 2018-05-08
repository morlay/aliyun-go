package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type FetchLibrariesRequest struct {
	requests.RpcRequest
	Size      int    `position:"Query" name:"Size"`
	StoreName string `position:"Query" name:"StoreName"`
	Page      int    `position:"Query" name:"Page"`
}

func (req *FetchLibrariesRequest) Invoke(client *sdk.Client) (resp *FetchLibrariesResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "FetchLibraries", "cloudphoto", "")
	resp = &FetchLibrariesResponse{}
	err = client.DoAction(req, resp)
	return
}

type FetchLibrariesResponse struct {
	responses.BaseResponse
	Code       common.String
	Message    common.String
	TotalCount common.Integer
	RequestId  common.String
	Action     common.String
	Libraries  FetchLibrariesLibraryList
}

type FetchLibrariesLibrary struct {
	LibraryId common.String
	Ctime     common.Long
}

type FetchLibrariesLibraryList []FetchLibrariesLibrary

func (list *FetchLibrariesLibraryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FetchLibrariesLibrary)
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
