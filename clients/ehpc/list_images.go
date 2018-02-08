package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListImagesRequest struct {
	requests.RpcRequest
}

func (req *ListImagesRequest) Invoke(client *sdk.Client) (resp *ListImagesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ListImages", "ehs", "")
	resp = &ListImagesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListImagesResponse struct {
	responses.BaseResponse
	RequestId string
	OsTags    ListImagesOsInfoList
}

type ListImagesOsInfo struct {
	OsTag        string
	Platform     string
	Version      string
	Architecture string
}

type ListImagesOsInfoList []ListImagesOsInfo

func (list *ListImagesOsInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListImagesOsInfo)
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
