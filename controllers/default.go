/*
 * @Description:
 * @Author: Moqi
 * @Email: str@li.cm
 * @Github: https://github.com/strugglerx
 * @LastEditors: Moqi
 * @Date: 2019-03-09 14:33:56
 * @LastEditTime: 2019-03-09 15:46:43
 */

package controllers

import (
	"encoding/json"
	"server/utils"
)

type Base_ struct {
	Code   string `json:"code"`
	Status int64  `json:"status"`
}

//自定义响应生成结构体
type CustomResponse struct {
	Base_
	Data interface{} `json:"data,omitempty"`
}

//网页通知响应
type MainResponse struct {
	Base_
	Message interface{} `json:"message,omitempty"`
}

//eip响应生成结构体
type CustomEipResponse struct {
	Code    string      `json:"code"`
	Status  int64       `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Encrypt string      `json:"encrypt,omitempty"`
}

func (res *CustomResponse) JsonFormat() string {
	datas := utils.WxJsonMarshal(res)
	return string(datas)
}

func (res *MainResponse) JsonParse() string {
	datas, _ := json.Marshal(res)
	return string(datas)
}

func (res *CustomEipResponse) JsonFormat() string {
	//datas, _ := json.Marshal(res)
	datas := utils.WxJsonMarshal(res)
	return string(datas)
}

/**
 * @description:Api接口响应内容
 * @return:
 */

//api有数据响应
func ApiResponse(data interface{}) string {
	var r *CustomResponse = &CustomResponse{Base_{"1000", 0}, data}
	return r.JsonFormat()
}

//数据库查询不正确响应
func ApiDbFail() string {
	var r *CustomResponse = &CustomResponse{Base_{"1001", -1}, nil}
	return r.JsonFormat()
}

//参数不正确响应
func ApiFail() string {
	var r *CustomResponse = &CustomResponse{Base_{"1002", -2}, nil}
	return r.JsonFormat()
}

/**
 * @description:带message的响应
 * @param {type}
 * @return:
 */
func MsgResponse(message interface{}) string {
	var r *MainResponse = &MainResponse{Base_{"1000", 0}, message}
	return r.JsonParse()
}

func MsgDbFail(message interface{}) string {
	var r *MainResponse = &MainResponse{Base_{"1001", -1}, message}
	return r.JsonParse()
}

func MsgFail(message interface{}) string {
	var r *MainResponse = &MainResponse{Base_{"1002", -2}, message}
	return r.JsonParse()
}
