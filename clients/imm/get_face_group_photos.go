package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetFaceGroupPhotosRequest struct {
	requests.RpcRequest
	MaxKeys int    `position:"Query" name:"MaxKeys"`
	Marker  string `position:"Query" name:"Marker"`
	GroupId string `position:"Query" name:"GroupId"`
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
}

func (req *GetFaceGroupPhotosRequest) Invoke(client *sdk.Client) (resp *GetFaceGroupPhotosResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "GetFaceGroupPhotos", "imm", "")
	resp = &GetFaceGroupPhotosResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetFaceGroupPhotosResponse struct {
	responses.BaseResponse
	RequestId  common.String
	Photos     common.String
	NextMarker common.String
}
