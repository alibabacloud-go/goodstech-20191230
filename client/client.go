// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/tea"
	oss "github.com/aliyun/alibabacloud-oss-sdk/golang/client"
	ossutil "github.com/aliyun/alibabacloud-oss-sdk/util/golang/service"
	rpcutil "github.com/aliyun/alibabacloud-rpc-util-sdk/golang/service"
	openplatform "github.com/aliyun/alibabacloud-sdk/openplatform-20191219/golang/client"
	credential "github.com/aliyun/credentials-go/credentials"
	fileform "github.com/aliyun/tea-fileform/golang/service"
	util "github.com/aliyun/tea-util/golang/service"
	"io"
)

type Config struct {
	AccessKeyId          *string `json:"accessKeyId" xml:"accessKeyId"`
	AccessKeySecret      *string `json:"accessKeySecret" xml:"accessKeySecret"`
	Type                 *string `json:"type" xml:"type"`
	SecurityToken        *string `json:"securityToken" xml:"securityToken"`
	Endpoint             *string `json:"endpoint" xml:"endpoint" require:"true"`
	Protocol             *string `json:"protocol" xml:"protocol"`
	RegionId             *string `json:"regionId" xml:"regionId" require:"true"`
	UserAgent            *string `json:"userAgent" xml:"userAgent"`
	ReadTimeout          *int    `json:"readTimeout" xml:"readTimeout"`
	ConnectTimeout       *int    `json:"connectTimeout" xml:"connectTimeout"`
	HttpProxy            *string `json:"httpProxy" xml:"httpProxy"`
	HttpsProxy           *string `json:"httpsProxy" xml:"httpsProxy"`
	NoProxy              *string `json:"noProxy" xml:"noProxy"`
	Socks5Proxy          *string `json:"socks5Proxy" xml:"socks5Proxy"`
	Socks5NetWork        *string `json:"socks5NetWork" xml:"socks5NetWork"`
	MaxIdleConns         *int    `json:"maxIdleConns" xml:"maxIdleConns"`
	EndpointType         *string `json:"endpointType" xml:"endpointType"`
	OpenPlatformEndpoint *string `json:"openPlatformEndpoint" xml:"openPlatformEndpoint"`
}

func (s Config) String() string {
	return tea.Prettify(s)
}

func (s Config) GoString() string {
	return s.String()
}

func (s *Config) SetAccessKeyId(v string) *Config {
	s.AccessKeyId = &v
	return s
}

func (s *Config) SetAccessKeySecret(v string) *Config {
	s.AccessKeySecret = &v
	return s
}

func (s *Config) SetType(v string) *Config {
	s.Type = &v
	return s
}

func (s *Config) SetSecurityToken(v string) *Config {
	s.SecurityToken = &v
	return s
}

func (s *Config) SetEndpoint(v string) *Config {
	s.Endpoint = &v
	return s
}

func (s *Config) SetProtocol(v string) *Config {
	s.Protocol = &v
	return s
}

func (s *Config) SetRegionId(v string) *Config {
	s.RegionId = &v
	return s
}

func (s *Config) SetUserAgent(v string) *Config {
	s.UserAgent = &v
	return s
}

func (s *Config) SetReadTimeout(v int) *Config {
	s.ReadTimeout = &v
	return s
}

func (s *Config) SetConnectTimeout(v int) *Config {
	s.ConnectTimeout = &v
	return s
}

func (s *Config) SetHttpProxy(v string) *Config {
	s.HttpProxy = &v
	return s
}

func (s *Config) SetHttpsProxy(v string) *Config {
	s.HttpsProxy = &v
	return s
}

func (s *Config) SetNoProxy(v string) *Config {
	s.NoProxy = &v
	return s
}

func (s *Config) SetSocks5Proxy(v string) *Config {
	s.Socks5Proxy = &v
	return s
}

func (s *Config) SetSocks5NetWork(v string) *Config {
	s.Socks5NetWork = &v
	return s
}

func (s *Config) SetMaxIdleConns(v int) *Config {
	s.MaxIdleConns = &v
	return s
}

func (s *Config) SetEndpointType(v string) *Config {
	s.EndpointType = &v
	return s
}

func (s *Config) SetOpenPlatformEndpoint(v string) *Config {
	s.OpenPlatformEndpoint = &v
	return s
}

type RecognizeFurnitureAttributeRequest struct {
	ImageURL *string `json:"ImageURL" xml:"ImageURL" require:"true"`
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

type RecognizeFurnitureAttributeResponse struct {
	RequestId *string                                  `json:"RequestId" xml:"RequestId" require:"true"`
	Data      *RecognizeFurnitureAttributeResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s RecognizeFurnitureAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureAttributeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureAttributeResponse) SetRequestId(v string) *RecognizeFurnitureAttributeResponse {
	s.RequestId = &v
	return s
}

func (s *RecognizeFurnitureAttributeResponse) SetData(v *RecognizeFurnitureAttributeResponseData) *RecognizeFurnitureAttributeResponse {
	s.Data = v
	return s
}

type RecognizeFurnitureAttributeResponseData struct {
	PredStyleId     *string  `json:"PredStyleId" xml:"PredStyleId" require:"true"`
	PredStyle       *string  `json:"PredStyle" xml:"PredStyle" require:"true"`
	PredProbability *float32 `json:"PredProbability" xml:"PredProbability" require:"true"`
}

func (s RecognizeFurnitureAttributeResponseData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureAttributeResponseData) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureAttributeResponseData) SetPredStyleId(v string) *RecognizeFurnitureAttributeResponseData {
	s.PredStyleId = &v
	return s
}

func (s *RecognizeFurnitureAttributeResponseData) SetPredStyle(v string) *RecognizeFurnitureAttributeResponseData {
	s.PredStyle = &v
	return s
}

func (s *RecognizeFurnitureAttributeResponseData) SetPredProbability(v float32) *RecognizeFurnitureAttributeResponseData {
	s.PredProbability = &v
	return s
}

type RecognizeFurnitureAttributeAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject" xml:"ImageURLObject" require:"true"`
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

type RecognizeFurnitureSpuRequest struct {
	ImageURL *string  `json:"ImageURL" xml:"ImageURL" require:"true"`
	XLength  *float32 `json:"XLength" xml:"XLength" require:"true"`
	YLength  *float32 `json:"YLength" xml:"YLength" require:"true"`
	ZLength  *float32 `json:"ZLength" xml:"ZLength" require:"true"`
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

type RecognizeFurnitureSpuResponse struct {
	RequestId *string                            `json:"RequestId" xml:"RequestId" require:"true"`
	Data      *RecognizeFurnitureSpuResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s RecognizeFurnitureSpuResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureSpuResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureSpuResponse) SetRequestId(v string) *RecognizeFurnitureSpuResponse {
	s.RequestId = &v
	return s
}

func (s *RecognizeFurnitureSpuResponse) SetData(v *RecognizeFurnitureSpuResponseData) *RecognizeFurnitureSpuResponse {
	s.Data = v
	return s
}

type RecognizeFurnitureSpuResponseData struct {
	PredCateId      *string  `json:"PredCateId" xml:"PredCateId" require:"true"`
	PredCate        *string  `json:"PredCate" xml:"PredCate" require:"true"`
	PredProbability *float32 `json:"PredProbability" xml:"PredProbability" require:"true"`
}

func (s RecognizeFurnitureSpuResponseData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureSpuResponseData) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureSpuResponseData) SetPredCateId(v string) *RecognizeFurnitureSpuResponseData {
	s.PredCateId = &v
	return s
}

func (s *RecognizeFurnitureSpuResponseData) SetPredCate(v string) *RecognizeFurnitureSpuResponseData {
	s.PredCate = &v
	return s
}

func (s *RecognizeFurnitureSpuResponseData) SetPredProbability(v float32) *RecognizeFurnitureSpuResponseData {
	s.PredProbability = &v
	return s
}

type RecognizeFurnitureSpuAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject" xml:"ImageURLObject" require:"true"`
	XLength        *float32  `json:"XLength" xml:"XLength" require:"true"`
	YLength        *float32  `json:"YLength" xml:"YLength" require:"true"`
	ZLength        *float32  `json:"ZLength" xml:"ZLength" require:"true"`
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

type ClassifyCommodityRequest struct {
	ImageURL *string `json:"ImageURL" xml:"ImageURL" require:"true"`
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

type ClassifyCommodityResponse struct {
	RequestId *string                        `json:"RequestId" xml:"RequestId" require:"true"`
	Data      *ClassifyCommodityResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s ClassifyCommodityResponse) String() string {
	return tea.Prettify(s)
}

func (s ClassifyCommodityResponse) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityResponse) SetRequestId(v string) *ClassifyCommodityResponse {
	s.RequestId = &v
	return s
}

func (s *ClassifyCommodityResponse) SetData(v *ClassifyCommodityResponseData) *ClassifyCommodityResponse {
	s.Data = v
	return s
}

type ClassifyCommodityResponseData struct {
	Categories []*ClassifyCommodityResponseDataCategories `json:"Categories" xml:"Categories" require:"true" type:"Repeated"`
}

func (s ClassifyCommodityResponseData) String() string {
	return tea.Prettify(s)
}

func (s ClassifyCommodityResponseData) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityResponseData) SetCategories(v []*ClassifyCommodityResponseDataCategories) *ClassifyCommodityResponseData {
	s.Categories = v
	return s
}

type ClassifyCommodityResponseDataCategories struct {
	Score        *float32 `json:"Score" xml:"Score" require:"true"`
	CategoryName *string  `json:"CategoryName" xml:"CategoryName" require:"true"`
	CategoryId   *string  `json:"CategoryId" xml:"CategoryId" require:"true"`
}

func (s ClassifyCommodityResponseDataCategories) String() string {
	return tea.Prettify(s)
}

func (s ClassifyCommodityResponseDataCategories) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityResponseDataCategories) SetScore(v float32) *ClassifyCommodityResponseDataCategories {
	s.Score = &v
	return s
}

func (s *ClassifyCommodityResponseDataCategories) SetCategoryName(v string) *ClassifyCommodityResponseDataCategories {
	s.CategoryName = &v
	return s
}

func (s *ClassifyCommodityResponseDataCategories) SetCategoryId(v string) *ClassifyCommodityResponseDataCategories {
	s.CategoryId = &v
	return s
}

type ClassifyCommodityAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject" xml:"ImageURLObject" require:"true"`
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

type Client struct {
	Endpoint             string
	RegionId             string
	Protocol             string
	UserAgent            string
	EndpointType         string
	ReadTimeout          int
	ConnectTimeout       int
	HttpProxy            string
	HttpsProxy           string
	Socks5Proxy          string
	Socks5NetWork        string
	NoProxy              string
	MaxIdleConns         int
	OpenPlatformEndpoint string
	Credential           credential.Credential
}

func NewClient(config *Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *Config) (_err error) {
	if util.IsUnset(tea.ToMap(config)) {
		_err = tea.NewSDKError(map[string]interface{}{
			"name":    "ParameterMissing",
			"message": "'config' can not be unset",
		})
		return _err
	}

	if util.Empty(tea.StringValue(config.RegionId)) {
		_err = tea.NewSDKError(map[string]interface{}{
			"name":    "ParameterMissing",
			"message": "'config.regionId' can not be empty",
		})
		return _err
	}

	if util.Empty(tea.StringValue(config.Endpoint)) {
		_err = tea.NewSDKError(map[string]interface{}{
			"name":    "ParameterMissing",
			"message": "'config.endpoint' can not be empty",
		})
		return _err
	}

	if util.Empty(tea.StringValue(config.Type)) {
		config.Type = tea.String("access_key")
	}

	credentialConfig := &credential.Config{
		AccessKeyId:     config.AccessKeyId,
		Type:            config.Type,
		AccessKeySecret: config.AccessKeySecret,
		SecurityToken:   config.SecurityToken,
	}
	client.Credential, _err = credential.NewCredential(credentialConfig)
	if _err != nil {
		return _err
	}

	client.Endpoint = tea.StringValue(config.Endpoint)
	client.Protocol = tea.StringValue(config.Protocol)
	client.RegionId = tea.StringValue(config.RegionId)
	client.UserAgent = tea.StringValue(config.UserAgent)
	client.ReadTimeout = tea.IntValue(config.ReadTimeout)
	client.ConnectTimeout = tea.IntValue(config.ConnectTimeout)
	client.HttpProxy = tea.StringValue(config.HttpProxy)
	client.HttpsProxy = tea.StringValue(config.HttpsProxy)
	client.NoProxy = tea.StringValue(config.NoProxy)
	client.Socks5Proxy = tea.StringValue(config.Socks5Proxy)
	client.Socks5NetWork = tea.StringValue(config.Socks5NetWork)
	client.MaxIdleConns = tea.IntValue(config.MaxIdleConns)
	client.EndpointType = tea.StringValue(config.EndpointType)
	client.OpenPlatformEndpoint = tea.StringValue(config.OpenPlatformEndpoint)
	return nil
}

func (client *Client) _request(action string, protocol string, method string, authType string, query map[string]interface{}, body map[string]interface{}, runtime *util.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_err = tea.Validate(runtime)
	if _err != nil {
		return nil, _err
	}
	_runtime := map[string]interface{}{
		"timeouted":      "retry",
		"readTimeout":    util.DefaultNumber(tea.IntValue(runtime.ReadTimeout), client.ReadTimeout),
		"connectTimeout": util.DefaultNumber(tea.IntValue(runtime.ConnectTimeout), client.ConnectTimeout),
		"httpProxy":      util.DefaultString(tea.StringValue(runtime.HttpProxy), client.HttpProxy),
		"httpsProxy":     util.DefaultString(tea.StringValue(runtime.HttpsProxy), client.HttpsProxy),
		"noProxy":        util.DefaultString(tea.StringValue(runtime.NoProxy), client.NoProxy),
		"maxIdleConns":   util.DefaultNumber(tea.IntValue(runtime.MaxIdleConns), client.MaxIdleConns),
		"retry": map[string]interface{}{
			"retryable":   tea.BoolValue(runtime.Autoretry),
			"maxAttempts": util.DefaultNumber(tea.IntValue(runtime.MaxAttempts), 3),
		},
		"backoff": map[string]interface{}{
			"policy": util.DefaultString(tea.StringValue(runtime.BackoffPolicy), "no"),
			"period": util.DefaultNumber(tea.IntValue(runtime.BackoffPeriod), 1),
		},
		"ignoreSSL": tea.BoolValue(runtime.IgnoreSSL),
	}

	_resp := make(map[string]interface{})
	for _retryTimes := 0; tea.AllowRetry(_runtime["retry"], _retryTimes); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], _retryTimes)
			if _backoffTime > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (map[string]interface{}, error) {
			request_ := tea.NewRequest()
			request_.Protocol = util.DefaultString(client.Protocol, protocol)
			request_.Method = method
			request_.Pathname = "/"
			request_.Query = rpcutil.Query(tea.ToMap(map[string]interface{}{
				"Action":         action,
				"Format":         "json",
				"RegionId":       client.RegionId,
				"Timestamp":      rpcutil.GetTimestamp(),
				"Version":        "2019-12-30",
				"SignatureNonce": util.GetNonce(),
			}, query))
			if !util.IsUnset(body) {
				tmp := util.AnyifyMapValue(rpcutil.Query(body))
				request_.Body = tea.ToReader(util.ToFormString(tmp))
			}

			request_.Headers = map[string]string{
				"host":       rpcutil.GetHost("goodstech", client.RegionId, client.Endpoint),
				"user-agent": client.GetUserAgent(),
			}
			if !util.EqualString(authType, "Anonymous") {
				accessKeyId, _err := client.GetAccessKeyId()
				if _err != nil {
					return nil, _err
				}

				accessKeySecret, _err := client.GetAccessKeySecret()
				if _err != nil {
					return nil, _err
				}

				request_.Query["SignatureMethod"] = "HMAC-SHA1"
				request_.Query["SignatureVersion"] = "1.0"
				request_.Query["AccessKeyId"] = accessKeyId
				request_.Query["Signature"] = rpcutil.GetSignature(request_, accessKeySecret)
			}

			response_, _err := tea.DoRequest(request_, _runtime)
			if _err != nil {
				return nil, _err
			}
			obj, _err := util.ReadAsJSON(response_.Body)
			if _err != nil {
				return nil, _err
			}

			res := util.AssertAsMap(obj)
			if util.Is4xx(response_.StatusCode) || util.Is5xx(response_.StatusCode) {
				_err = tea.NewSDKError(map[string]interface{}{
					"message": res["Message"],
					"data":    res,
					"code":    res["Code"],
				})
				return nil, _err
			}

			_result = res
			return _result, _err
		}()
		if !tea.Retryable(_err) {
			break
		}
	}

	return _resp, _err
}

func (client *Client) RecognizeFurnitureAttribute(request *RecognizeFurnitureAttributeRequest, runtime *util.RuntimeOptions) (_result *RecognizeFurnitureAttributeResponse, _err error) {
	_result = &RecognizeFurnitureAttributeResponse{}
	_body, _err := client._request("RecognizeFurnitureAttribute", "HTTPS", "POST", "AK", nil, tea.ToMap(request), runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeFurnitureAttributeAdvance(request *RecognizeFurnitureAttributeAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeFurnitureAttributeResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return nil, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return nil, _err
	}

	authConfig := &openplatform.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        tea.String(client.Protocol),
		RegionId:        tea.String(client.RegionId),
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return nil, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("goodstech"),
		RegionId: tea.String(client.RegionId),
	}
	authResponse, _err := authClient.AuthorizeFileUpload(authRequest, runtime)
	if _err != nil {
		return nil, _err
	}

	// Step 1: request OSS api to upload file
	ossConfig := &oss.Config{
		AccessKeyId:     authResponse.AccessKeyId,
		AccessKeySecret: tea.String(accessKeySecret),
		Type:            tea.String("access_key"),
		Endpoint:        tea.String(rpcutil.GetEndpoint(tea.StringValue(authResponse.Endpoint), tea.BoolValue(authResponse.UseAccelerate), client.EndpointType)),
		Protocol:        tea.String(client.Protocol),
		RegionId:        tea.String(client.RegionId),
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return nil, _err
	}

	fileObj := &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader := &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest := &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return
	}
	// Step 2: request final api
	recognizeFurnitureAttributereq := &RecognizeFurnitureAttributeRequest{}
	rpcutil.Convert(request, recognizeFurnitureAttributereq)
	recognizeFurnitureAttributereq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeFurnitureAttributeResp, _err := client.RecognizeFurnitureAttribute(recognizeFurnitureAttributereq, runtime)
	if _err != nil {
		return nil, _err
	}

	_result = recognizeFurnitureAttributeResp
	return _result, _err
}

func (client *Client) RecognizeFurnitureSpu(request *RecognizeFurnitureSpuRequest, runtime *util.RuntimeOptions) (_result *RecognizeFurnitureSpuResponse, _err error) {
	_result = &RecognizeFurnitureSpuResponse{}
	_body, _err := client._request("RecognizeFurnitureSpu", "HTTPS", "POST", "AK", nil, tea.ToMap(request), runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeFurnitureSpuAdvance(request *RecognizeFurnitureSpuAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeFurnitureSpuResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return nil, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return nil, _err
	}

	authConfig := &openplatform.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        tea.String(client.Protocol),
		RegionId:        tea.String(client.RegionId),
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return nil, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("goodstech"),
		RegionId: tea.String(client.RegionId),
	}
	authResponse, _err := authClient.AuthorizeFileUpload(authRequest, runtime)
	if _err != nil {
		return nil, _err
	}

	// Step 1: request OSS api to upload file
	ossConfig := &oss.Config{
		AccessKeyId:     authResponse.AccessKeyId,
		AccessKeySecret: tea.String(accessKeySecret),
		Type:            tea.String("access_key"),
		Endpoint:        tea.String(rpcutil.GetEndpoint(tea.StringValue(authResponse.Endpoint), tea.BoolValue(authResponse.UseAccelerate), client.EndpointType)),
		Protocol:        tea.String(client.Protocol),
		RegionId:        tea.String(client.RegionId),
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return nil, _err
	}

	fileObj := &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader := &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest := &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return
	}
	// Step 2: request final api
	recognizeFurnitureSpureq := &RecognizeFurnitureSpuRequest{}
	rpcutil.Convert(request, recognizeFurnitureSpureq)
	recognizeFurnitureSpureq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeFurnitureSpuResp, _err := client.RecognizeFurnitureSpu(recognizeFurnitureSpureq, runtime)
	if _err != nil {
		return nil, _err
	}

	_result = recognizeFurnitureSpuResp
	return _result, _err
}

func (client *Client) ClassifyCommodity(request *ClassifyCommodityRequest, runtime *util.RuntimeOptions) (_result *ClassifyCommodityResponse, _err error) {
	_result = &ClassifyCommodityResponse{}
	_body, _err := client._request("ClassifyCommodity", "HTTPS", "GET", "AK", tea.ToMap(request), nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClassifyCommodityAdvance(request *ClassifyCommodityAdvanceRequest, runtime *util.RuntimeOptions) (_result *ClassifyCommodityResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return nil, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return nil, _err
	}

	authConfig := &openplatform.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        tea.String(client.Protocol),
		RegionId:        tea.String(client.RegionId),
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return nil, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("goodstech"),
		RegionId: tea.String(client.RegionId),
	}
	authResponse, _err := authClient.AuthorizeFileUpload(authRequest, runtime)
	if _err != nil {
		return nil, _err
	}

	// Step 1: request OSS api to upload file
	ossConfig := &oss.Config{
		AccessKeyId:     authResponse.AccessKeyId,
		AccessKeySecret: tea.String(accessKeySecret),
		Type:            tea.String("access_key"),
		Endpoint:        tea.String(rpcutil.GetEndpoint(tea.StringValue(authResponse.Endpoint), tea.BoolValue(authResponse.UseAccelerate), client.EndpointType)),
		Protocol:        tea.String(client.Protocol),
		RegionId:        tea.String(client.RegionId),
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return nil, _err
	}

	fileObj := &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader := &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest := &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return
	}
	// Step 2: request final api
	classifyCommodityreq := &ClassifyCommodityRequest{}
	rpcutil.Convert(request, classifyCommodityreq)
	classifyCommodityreq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	classifyCommodityResp, _err := client.ClassifyCommodity(classifyCommodityreq, runtime)
	if _err != nil {
		return nil, _err
	}

	_result = classifyCommodityResp
	return _result, _err
}

func (client *Client) GetUserAgent() (_result string) {
	userAgent := util.GetUserAgent(client.UserAgent)
	_result = userAgent
	return _result
}

func (client *Client) GetAccessKeyId() (_result string, _err error) {
	if util.IsUnset(client.Credential) {
		_result = ""
		return _result, _err
	}

	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return "", _err
	}

	_result = accessKeyId
	return _result, _err
}

func (client *Client) GetAccessKeySecret() (_result string, _err error) {
	if util.IsUnset(client.Credential) {
		_result = ""
		return _result, _err
	}

	secret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return "", _err
	}

	_result = secret
	return _result, _err
}
