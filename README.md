mem
=
mysql error map(mem)
-

# 简述
> 解决使用 github.com/go-sql-driver/mysql 包时的错误处理问题。  
> 要获取 MySQLError 结构类型，通常则要引用 github.com/go-sql-driver/mysql。  
> 本工具可以让你用下划线 import 包时 —— 即只执行该报中的init函数，不把其类型和变量导入。也可以拿到 MySQLError 结构类型。  

* 使用时只需要如下即可
```
	import _ "github.com/go-sql-driver/mysql"
```
* 不需要
```
	import "github.com/go-sql-driver/mysql"
```
* 这样只会执行此包的init()函数，不必浪费多余的内存空间。


# 安装
```
	go get github.com/mosalut/mem
```

# 依赖
## 隐式依赖
```
	import _ "github.com/go-sql-driver/mysql"
```
* 不引入 mysql 包，也不会用到 mem 包
* 不依赖其他第三方包

# 使用
```
	[...], err := <某个数据库操作造成的MySQL错误>
	if err != nil {
		dbError := &mysql.dbErr{} // 注意该结构体实现了 error 接口
		err := dbError.mapping(err) // 这里的实参err 必须是个和MySQLError结构相同的类型，之前所谓的隐式依赖就在于此。
		if err != nil {
			// 处理执行该函数本身的错误，通常是一个JSON转换错误。
		}
		fmt.Println("MySQL错误号：", dbError.Number)
		fmt.Println("MySQL字符串：", dbError.Message)
		fmt.Println("MySQL字符串：", dbError.Error())
		return
	}
```
* 如果发生错误，但并没有错误号，例如：
```
	db, err := sql.Open("mysql", "乱写一通，让他出错")
	if err != nil {
		return err
	}
	defer db.Close()
```
则Number为0，Message为原来的 error 实现类型的 Error() 函数返回结果。  
* 可参考 mem_test.go 文件
