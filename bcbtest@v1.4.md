## 接口

```http
GET http://192.168.1.99:8089/api/v1/cases
```

描述：查看所有用例



```http
GET http://192.168.1.99:8089/api/v1/reports?mode=1
```

描述：查看所有用例测试报告，mode=1（完整模式）



```http
GET http://192.168.1.99:8089/api/v1/report/30001
```

描述：查看单个用例详细测试报告



```http
GET http://192.168.1.99:8089/api/v1/start
{
	code: 200,
	data: "success."
}
```

描述：启动测试engine，开始测试



## 数据库

### 地址

```yaml
  dbIP: "192.168.1.99"
  dbPort: 3306
  dbUser: "root"
  dbPwd: "123456"
  dbName: "bcbtest"
```



### 表描述

**report**

comment: 测试报告记录表



**report_case**

comment: 记录具体测试用例报告结果，每次跑用例前会清空此表



**report_case_his**

comment: 跑测试用例之前，将report_case中的数据迁移到此表中做备份



**test_case**

comment: 测试用例描述表，启动server时会清空此表



**test_case_suite**

comment: 测试用例层级关系表，启动server时会清空此表



**test_case_kind**

comment: 测试用例集分类表，启动server时会清空此表



## excel格式



根据工作表中第一行的非空列数来确定层级关系



### 三层结构

![image-20200715102117135](https://sayming.oss-cn-beijing.aliyuncs.com/GIChain/2020-07-15/image-20200715102117135.png)



**总列数**: 5 	

第一列：第一级测试组

第二列：第二级测试组

第三列：第三级测试组

第四列：测试用例描述

第五列：预期结果



### 两层结构

![image-20200715102640334](https://sayming.oss-cn-beijing.aliyuncs.com/GIChain/2020-07-15/image-20200715102640334.png)



**总列数**：4

第一列：第一级测试组

第二列：第二级测试组（第三级测试组）

第三列：测试用例描述

第四列：预期结果



## 问题

1. repeatCount从哪获取