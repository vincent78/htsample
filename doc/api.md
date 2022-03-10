[toc]

# 前言

host = **http://localhost:8080**

# API

## 所有用户
```
➜  bin  curl http://localhost:8080/accountList
{"code":200,"error":null,"Context":[{"id":1,"code":"a001","name":"bob123","remark":""},{"id":2,"code":"a002","name":"alice456","remark":""}]}
```

## 按code查询用户
```
➜  bin  curl http://localhost:8080/accountByCode\?code\=a001
{"code":200,"error":null,"Context":{"id":1,"code":"a001","name":"bob123","remark":""}}
```

## 按code查询用户余额
```
➜  bin  curl http://localhost:8080/accountBalanceByCode\?code\=a001
{"code":200,"error":null,"Context":{"id":1,"code":"a001","curr":"usd","balance":9980}}
```

## 查询交易记录
```
➜  bin  curl http://localhost:8080/paymentList                     
{
	"code": 200,
	"error": null,
	"Context": [{
		"id": 3,
		"account": "a001",
		"ptype": "outgoing",
		"curr": "usb",
		"balance": 10,
		"frozen": 0,
		"token": "58464319d000105",
		"remark": "",
		"createAt": 1646489553463,
		"createBy": "admin"
	}, {
		"id": 4,
		"account": "a002",
		"ptype": "incoming",
		"curr": "usb",
		"balance": 10,
		"frozen": 0,
		"token": "58464319d000105",
		"remark": "",
		"createAt": 1646489553223,
		"createBy": "admin"
	}, {
		"id": 1,
		"account": "a001",
		"ptype": "outgoing",
		"curr": "usb",
		"balance": 10,
		"frozen": 0,
		"token": "584642f4c000105",
		"remark": "",
		"createAt": 1646489547457,
		"createBy": "admin"
	}, {
		"id": 2,
		"account": "a002",
		"ptype": "incoming",
		"curr": "usb",
		"balance": 10,
		"frozen": 0,
		"token": "584642f4c000105",
		"remark": "",
		"createAt": 1646489547267,
		"createBy": "admin"
	}]
}
```

## 按token查询交易记录
```
➜  bin  curl http://localhost:8080/paymentListByToken\?token\=584642f4c000105
{"code":200,"error":null,"Context":[{"id":1,"account":"a001","ptype":"outgoing","curr":"usb","balance":10,"frozen":0,"token":"584642f4c000105","remark":"","createAt":1646489547457,"createBy":"admin"},{"id":2,"account":"a002","ptype":"incoming","curr":"usb","balance":10,"frozen":0,"token":"584642f4c000105","remark":"","createAt":1646489547267,"createBy":"admin"}]}
```

## 按用户code查询交易记录
```
➜  bin  curl http://localhost:8080/paymentListByAccount\?account\=a001
{"code":200,"error":null,"Context":[{"id":3,"account":"a001","ptype":"outgoing","curr":"usb","balance":10,"frozen":0,"token":"58464319d000105","remark":"","createAt":1646489553463,"createBy":"admin"},{"id":1,"account":"a001","ptype":"outgoing","curr":"usb","balance":10,"frozen":0,"token":"584642f4c000105","remark":"","createAt":1646489547457,"createBy":"admin"}]}
```

## 交易
```
➜  bin  curl -X POST 'http://127.0.0.1:8080/transfer' -H 'Content-Type: application/json' --data '{"account1":"a001","curr1":"usb","num1":"10","account2":"a002","curr2":"usb","num2":"10"}'
{"code":200,"error":null,"Context":"584683994000105"}
```



