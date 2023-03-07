# 项目名称
PROJECT=cfx_serv
# 当前路径
CURRENT=$(shell pwd)
# 编译目录
BUILD_DIR=build
# 入口文件
MAIN_PATH=$(CURRENT)/cfx/main.go


common:
    # 准备好编译目录结构
	rm -rf $(BUILD_DIR)
	mkdir -p $(BUILD_DIR)/cfx/config
	mkdir -p $(BUILD_DIR)/logs
	cp -rf cfx/config $(BUILD_DIR)/cfx/
    export GO111MODULE=on
    export GOPROXY=https://goproxy.cn
#	export GOPATH=$(CURRENT) && go build -o $(CURRENT)/$(BUILD_DIR)/api $(CURRENT)/code/cmd/api/main.go
# 	export GOPATH=$(CURRENT) && go build -o $(CURRENT)/$(BUILD_DIR)/cmd $(CURRENT)/code/cmd/cmd/main.go


dev: common
	echo "dev" > $(BUILD_DIR)/env.conf

clean:
	rm -rf $(BUILD_DIR)

http:
	cd cfx && go run main.go
localBin: cfx/common
	cd cfx && go build -o $(CURRENT)/$(BUILD_DIR)/$(PROJECT) $(MAIN_PATH)

# cp -rf cfx/config/* $(BUILD_DIR)/cfx/ 两者效果是不一样的，需要的话，可以自己查看下
cfxServer:
	rm -rf $(BUILD_DIR)
	mkdir -p $(BUILD_DIR)/cfx/config
	mkdir -p $(BUILD_DIR)/logs
	cp -rf cfx/config $(BUILD_DIR)/cfx/
	cd cfx && go build -o $(CURRENT)/$(BUILD_DIR)/$(PROJECT) $(MAIN_PATH)






