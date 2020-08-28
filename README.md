# ChindeoTest

- 使用 THINKPHP 6.0 开发的一个后台系统，由于  THINKPHP 单元测试长时间没更新，所以选择使用 go 写一个通用的测试。

- 使用 [httpexpect](https://github.com/gavv/httpexpect) 写的 RESTFUL-API 端元测试
- 使用 [faker](https://github.com/azumads/faker) 做为数据填充
- 使用 [air](https://github.com/cosmtrek/air) 作为单元测试自动运行工具
- 使用 [gormt](https://github.com/xxjwxc/gormt) mysql 生成 go struct 工具

```shell script
gormt
```

修改 `time.Time` 为 	`sql.NullTime`
修改 `TimeType      string` 为 	`TimeType      model.MenuTimeType`

```shell script
#全部执行
go test -v 
#部分执行
go test -v --run Test***
```

### 注意事项
**需要设置环境变量 TRAVIS_BUILD_DIR 为项目绝对路径**

