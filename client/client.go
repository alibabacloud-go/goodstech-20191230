// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type ClassifyCommodityRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ClassifyCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s ClassifyCommodityRequest) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityRequest) SetImageURL(v string) *ClassifyCommodityRequest {
	s.ImageURL = &v
	return s
}

type ClassifyCommodityAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ClassifyCommodityAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ClassifyCommodityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityAdvanceRequest) SetImageURLObject(v io.Reader) *ClassifyCommodityAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type ClassifyCommodityResponseBody struct {
	Data      *ClassifyCommodityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClassifyCommodityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClassifyCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityResponseBody) SetData(v *ClassifyCommodityResponseBodyData) *ClassifyCommodityResponseBody {
	s.Data = v
	return s
}

func (s *ClassifyCommodityResponseBody) SetRequestId(v string) *ClassifyCommodityResponseBody {
	s.RequestId = &v
	return s
}

type ClassifyCommodityResponseBodyData struct {
	Categories []*ClassifyCommodityResponseBodyDataCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s ClassifyCommodityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ClassifyCommodityResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityResponseBodyData) SetCategories(v []*ClassifyCommodityResponseBodyDataCategories) *ClassifyCommodityResponseBodyData {
	s.Categories = v
	return s
}

type ClassifyCommodityResponseBodyDataCategories struct {
	CategoryId   *string  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName *string  `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	Score        *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s ClassifyCommodityResponseBodyDataCategories) String() string {
	return tea.Prettify(s)
}

func (s ClassifyCommodityResponseBodyDataCategories) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityResponseBodyDataCategories) SetCategoryId(v string) *ClassifyCommodityResponseBodyDataCategories {
	s.CategoryId = &v
	return s
}

func (s *ClassifyCommodityResponseBodyDataCategories) SetCategoryName(v string) *ClassifyCommodityResponseBodyDataCategories {
	s.CategoryName = &v
	return s
}

func (s *ClassifyCommodityResponseBodyDataCategories) SetScore(v float32) *ClassifyCommodityResponseBodyDataCategories {
	s.Score = &v
	return s
}

type ClassifyCommodityResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ClassifyCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClassifyCommodityResponse) String() string {
	return tea.Prettify(s)
}

func (s ClassifyCommodityResponse) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityResponse) SetHeaders(v map[string]*string) *ClassifyCommodityResponse {
	s.Headers = v
	return s
}

func (s *ClassifyCommodityResponse) SetStatusCode(v int32) *ClassifyCommodityResponse {
	s.StatusCode = &v
	return s
}

func (s *ClassifyCommodityResponse) SetBody(v *ClassifyCommodityResponseBody) *ClassifyCommodityResponse {
	s.Body = v
	return s
}

type RecognizeFurnitureAttributeRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeFurnitureAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureAttributeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureAttributeRequest) SetImageURL(v string) *RecognizeFurnitureAttributeRequest {
	s.ImageURL = &v
	return s
}

type RecognizeFurnitureAttributeAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeFurnitureAttributeAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureAttributeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureAttributeAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeFurnitureAttributeAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeFurnitureAttributeResponseBody struct {
	Data      *RecognizeFurnitureAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeFurnitureAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureAttributeResponseBody) SetData(v *RecognizeFurnitureAttributeResponseBodyData) *RecognizeFurnitureAttributeResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeFurnitureAttributeResponseBody) SetRequestId(v string) *RecognizeFurnitureAttributeResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeFurnitureAttributeResponseBodyData struct {
	PredProbability *float32 `json:"PredProbability,omitempty" xml:"PredProbability,omitempty"`
	PredStyle       *string  `json:"PredStyle,omitempty" xml:"PredStyle,omitempty"`
	PredStyleId     *string  `json:"PredStyleId,omitempty" xml:"PredStyleId,omitempty"`
}

func (s RecognizeFurnitureAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureAttributeResponseBodyData) SetPredProbability(v float32) *RecognizeFurnitureAttributeResponseBodyData {
	s.PredProbability = &v
	return s
}

func (s *RecognizeFurnitureAttributeResponseBodyData) SetPredStyle(v string) *RecognizeFurnitureAttributeResponseBodyData {
	s.PredStyle = &v
	return s
}

func (s *RecognizeFurnitureAttributeResponseBodyData) SetPredStyleId(v string) *RecognizeFurnitureAttributeResponseBodyData {
	s.PredStyleId = &v
	return s
}

type RecognizeFurnitureAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeFurnitureAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeFurnitureAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureAttributeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureAttributeResponse) SetHeaders(v map[string]*string) *RecognizeFurnitureAttributeResponse {
	s.Headers = v
	return s
}

func (s *RecognizeFurnitureAttributeResponse) SetStatusCode(v int32) *RecognizeFurnitureAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeFurnitureAttributeResponse) SetBody(v *RecognizeFurnitureAttributeResponseBody) *RecognizeFurnitureAttributeResponse {
	s.Body = v
	return s
}

type RecognizeFurnitureSpuRequest struct {
	ImageURL *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	XLength  *float32 `json:"XLength,omitempty" xml:"XLength,omitempty"`
	YLength  *float32 `json:"YLength,omitempty" xml:"YLength,omitempty"`
	ZLength  *float32 `json:"ZLength,omitempty" xml:"ZLength,omitempty"`
}

func (s RecognizeFurnitureSpuRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureSpuRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureSpuRequest) SetImageURL(v string) *RecognizeFurnitureSpuRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeFurnitureSpuRequest) SetXLength(v float32) *RecognizeFurnitureSpuRequest {
	s.XLength = &v
	return s
}

func (s *RecognizeFurnitureSpuRequest) SetYLength(v float32) *RecognizeFurnitureSpuRequest {
	s.YLength = &v
	return s
}

func (s *RecognizeFurnitureSpuRequest) SetZLength(v float32) *RecognizeFurnitureSpuRequest {
	s.ZLength = &v
	return s
}

type RecognizeFurnitureSpuAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	XLength        *float32  `json:"XLength,omitempty" xml:"XLength,omitempty"`
	YLength        *float32  `json:"YLength,omitempty" xml:"YLength,omitempty"`
	ZLength        *float32  `json:"ZLength,omitempty" xml:"ZLength,omitempty"`
}

func (s RecognizeFurnitureSpuAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureSpuAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureSpuAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeFurnitureSpuAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeFurnitureSpuAdvanceRequest) SetXLength(v float32) *RecognizeFurnitureSpuAdvanceRequest {
	s.XLength = &v
	return s
}

func (s *RecognizeFurnitureSpuAdvanceRequest) SetYLength(v float32) *RecognizeFurnitureSpuAdvanceRequest {
	s.YLength = &v
	return s
}

func (s *RecognizeFurnitureSpuAdvanceRequest) SetZLength(v float32) *RecognizeFurnitureSpuAdvanceRequest {
	s.ZLength = &v
	return s
}

type RecognizeFurnitureSpuResponseBody struct {
	Data      *RecognizeFurnitureSpuResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeFurnitureSpuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureSpuResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureSpuResponseBody) SetData(v *RecognizeFurnitureSpuResponseBodyData) *RecognizeFurnitureSpuResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeFurnitureSpuResponseBody) SetRequestId(v string) *RecognizeFurnitureSpuResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeFurnitureSpuResponseBodyData struct {
	PredCate        *string  `json:"PredCate,omitempty" xml:"PredCate,omitempty"`
	PredCateId      *string  `json:"PredCateId,omitempty" xml:"PredCateId,omitempty"`
	PredProbability *float32 `json:"PredProbability,omitempty" xml:"PredProbability,omitempty"`
}

func (s RecognizeFurnitureSpuResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureSpuResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureSpuResponseBodyData) SetPredCate(v string) *RecognizeFurnitureSpuResponseBodyData {
	s.PredCate = &v
	return s
}

func (s *RecognizeFurnitureSpuResponseBodyData) SetPredCateId(v string) *RecognizeFurnitureSpuResponseBodyData {
	s.PredCateId = &v
	return s
}

func (s *RecognizeFurnitureSpuResponseBodyData) SetPredProbability(v float32) *RecognizeFurnitureSpuResponseBodyData {
	s.PredProbability = &v
	return s
}

type RecognizeFurnitureSpuResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeFurnitureSpuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeFurnitureSpuResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureSpuResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureSpuResponse) SetHeaders(v map[string]*string) *RecognizeFurnitureSpuResponse {
	s.Headers = v
	return s
}

func (s *RecognizeFurnitureSpuResponse) SetStatusCode(v int32) *RecognizeFurnitureSpuResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeFurnitureSpuResponse) SetBody(v *RecognizeFurnitureSpuResponseBody) *RecognizeFurnitureSpuResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("goodstech"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClassifyCommodityWithOptions(request *ClassifyCommodityRequest, runtime *util.RuntimeOptions) (_result *ClassifyCommodityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ClassifyCommodity"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClassifyCommodityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClassifyCommodity(request *ClassifyCommodityRequest) (_result *ClassifyCommodityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClassifyCommodityResponse{}
	_body, _err := client.ClassifyCommodityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClassifyCommodityAdvance(request *ClassifyCommodityAdvanceRequest, runtime *util.RuntimeOptions) (_result *ClassifyCommodityResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("goodstech"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	classifyCommodityReq := &ClassifyCommodityRequest{}
	openapiutil.Convert(request, classifyCommodityReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		classifyCommodityReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	classifyCommodityResp, _err := client.ClassifyCommodityWithOptions(classifyCommodityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = classifyCommodityResp
	return _result, _err
}

func (client *Client) RecognizeFurnitureAttributeWithOptions(request *RecognizeFurnitureAttributeRequest, runtime *util.RuntimeOptions) (_result *RecognizeFurnitureAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeFurnitureAttribute"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeFurnitureAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeFurnitureAttribute(request *RecognizeFurnitureAttributeRequest) (_result *RecognizeFurnitureAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeFurnitureAttributeResponse{}
	_body, _err := client.RecognizeFurnitureAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeFurnitureAttributeAdvance(request *RecognizeFurnitureAttributeAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeFurnitureAttributeResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("goodstech"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeFurnitureAttributeReq := &RecognizeFurnitureAttributeRequest{}
	openapiutil.Convert(request, recognizeFurnitureAttributeReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeFurnitureAttributeReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeFurnitureAttributeResp, _err := client.RecognizeFurnitureAttributeWithOptions(recognizeFurnitureAttributeReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeFurnitureAttributeResp
	return _result, _err
}

func (client *Client) RecognizeFurnitureSpuWithOptions(request *RecognizeFurnitureSpuRequest, runtime *util.RuntimeOptions) (_result *RecognizeFurnitureSpuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.XLength)) {
		body["XLength"] = request.XLength
	}

	if !tea.BoolValue(util.IsUnset(request.YLength)) {
		body["YLength"] = request.YLength
	}

	if !tea.BoolValue(util.IsUnset(request.ZLength)) {
		body["ZLength"] = request.ZLength
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeFurnitureSpu"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeFurnitureSpuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeFurnitureSpu(request *RecognizeFurnitureSpuRequest) (_result *RecognizeFurnitureSpuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeFurnitureSpuResponse{}
	_body, _err := client.RecognizeFurnitureSpuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeFurnitureSpuAdvance(request *RecognizeFurnitureSpuAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeFurnitureSpuResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("goodstech"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeFurnitureSpuReq := &RecognizeFurnitureSpuRequest{}
	openapiutil.Convert(request, recognizeFurnitureSpuReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeFurnitureSpuReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeFurnitureSpuResp, _err := client.RecognizeFurnitureSpuWithOptions(recognizeFurnitureSpuReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeFurnitureSpuResp
	return _result, _err
}
