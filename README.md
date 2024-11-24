## 使用 make 命令创建新模型

* 模型文件
```bash
go run main.go make model model_name
```
* 迁移文件
```bash
go run main.go make migration add_model_names_table
```
* 请求验证文件
```bash
go run main.go make request model_name
```
* 控制器文件
```bash
go run main.go make apicontroller v1/model_name
```