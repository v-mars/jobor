```text
swagger:
update_swag:
	swag init -pd true
	swag fmt
	#time swag init --parseDependency false --parseDepth 1

init_api:
	hz new --model_dir biz/hertz_gen -mod github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm -idl idl/api.thrift
	#hz new -I idl -idl idl/hello/hello.proto

update_api: --unset_omitempty
    hz update  --model_dir kitex_gen --unset_omitempty -idl idl/env.proto
	hz update  --model_dir kitex_gen --unset_omitempty -idl idl/sys/kv.proto
	#hz update -I idl -idl idl/*/*/*.proto
```