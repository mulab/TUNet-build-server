# Deprecated!

# TUNet Build Server

## 准备工作

0. 根据config.json.example设置config.json,相对项目根目录 建议使用绝对路径
0. 按照TUNet自动化发布流程的说明设定build.properties并准备release.keystore

## 使用说明

0. 编译: `go build`
0. 运行: `./tunet_build_server`
0. 登录: 使用config.json中设定的用户名密码登录
0. 构建: 访问/build点击build按钮进行构建
0. 下载: 访问/download进行下载
