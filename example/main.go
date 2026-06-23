package main

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/rootexit/rex-sdk-go-v6/rex"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexConfig"
	"github.com/rootexit/rexLib/rexCtx"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	rexConfig := rexConfig.DefaultConfig(os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	rexConfig.Protocol = "https"
	rexConfig.Endpoint = "dev-api.rootexit.com"
	rexConfig.AccessKeyID = "REx584f20043b8065b2"
	rexConfig.AccessKeySecret = "0c461fd4aed9051b4da489813e5dc3c3"
	logx.Infof("打印一下请求的accessKey :%s", rexConfig.AccessKeyID)
	logx.Infof("打印一下请求的AccessKeySecret :%s", rexConfig.AccessKeySecret)
	logx.Infof("打印一下请求的Endpoint :%s", rexConfig.Endpoint)

	sdk, err := rex.NewSdk(rexConfig)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	uuidStr := uuid.New().String()
	ctx = context.WithValue(ctx, rexCtx.CtxRequestId{}, uuidStr)
	// _, allCodes, err := sdk.BaseService.Codes(ctx, &rexTypes.CodesReq{
	// 	Lang: "zh-CN",
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// logx.Infof("%+v", allCodes)

	// _, list, err := sdk.CredentialService.QueryList(ctx, &rexTypes.CredentialCommonSearchParams{})
	// if err != nil {
	// 	panic(err)
	// }
	// logx.Infof("%+v", list)

	// testMsg := base64.StdEncoding.EncodeToString([]byte("Welcome to REx Engine!"))
	// _, encryptResult, err := sdk.KmsService.Skc.Encrypt(ctx, &rexTypes.KmsSkcEncryptReq{
	// 	Name:     "default",
	// 	BaseData: testMsg,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// logx.Infof("%+v", encryptResult)

	// _, decodeMsg, err := sdk.KmsService.Skc.Decrypt(ctx, &rexTypes.KmsSkcDecryptReq{
	// 	Name:     "default",
	// 	BaseData: encryptResult.BaseData,
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// realMsg, _ := base64.StdEncoding.DecodeString(decodeMsg.BaseData)
	// logx.Infof("%+s", realMsg)

	// testMsg := base64.StdEncoding.EncodeToString([]byte("Welcome to REx Engine!"))
	// testMsg := "Welcome to REx Engine!"
	// _, signResult, err := sdk.KmsService.Akc.Sign(ctx, &rexTypes.KmsAkcSignReq{
	// 	Name:        "default",
	// 	SignContent: testMsg,
	// 	// Kid:         "",
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// logx.Infof("%+v", signResult)

	// _, verifyResult, err := sdk.KmsService.Akc.Verify(ctx, &rexTypes.KmsAkcVerifyReq{
	// 	Name: signResult.Name,
	// 	Sign: signResult.Sign,
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// logx.Infof("%+v", verifyResult)

	// _, result, err := sdk.MasService.EmsVerificationInit(ctx, &rexTypes.EmsInitReq{
	// 	Key:     "default",
	// 	Service: "account.lilsite.com",
	// 	Type:    "login",
	// 	Mail:    "howard.huang@lilsite.com",
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// logx.Infof("%+v", result)

	// _, result, err := sdk.MasService.SmsVerificationInit(ctx, &rexTypes.SmsInitReq{
	// 	Key:     "default",
	// 	Service: "account.lilsite.com",
	// 	Type:    "login",
	// 	Zone:    "86",
	// 	Phone:   "xxxxxxxxxxx",
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// logx.Infof("%+v", result)

	// _, result, err := sdk.MasService.BehavioralVerificationInit(ctx, &rexTypes.BehavioralVerificationInitReq{
	// 	Key:     "default",
	// 	Service: "account.lilsite.com",
	// 	Type:    "login",
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// logx.Infof("%+v", result)

	//_, queryBucketResult, err := sdk.SasService.QueryBucket(context.Background(), &rexTypes.SasQueryBucketReq{
	//	BucketKey: "default",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//logx.Infof("%+v", queryBucketResult)

	//queryBucketResult, err := sdk.UpsService.UpsTagService.Query(context.Background(), &rexTypes.TagApiFormIdReq{
	//	Id: 2,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//logx.Infof("%+v", queryBucketResult)

	//queryBucketResult, err := sdk.UpsService.IndustryService.QueryList(context.Background(), &rexTypes.IndustryCommonSearchParams{
	//	Page:       1,
	//	PageSize:   5,
	//	OnlyParent: true,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//logx.Infof("%+v", queryBucketResult)

	//_, queryAccessTokenResult, err := sdk.TpasService.WechatOffiaccountService.GetAccessToken(context.Background(), &rexTypes.WechatOffiaccountGetAccessTokenReq{
	//	Key: "default",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//logx.Infof("%+v", queryAccessTokenResult)
}
