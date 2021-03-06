
# 字节跳动青训营项目文档说明                              

## 队名：提瓦特大陆旅游队

# 1. 项目介绍
实现一个极简版的抖音，通过利用Go语言编程，常用框架，数据库，对象储存等内容，实现极简版抖音的视频Feed流，视频投稿，点赞列表，用户评论，关注和粉丝列表等基本功能。使用户能够按投稿时间倒序显示视频，并且可以自己拍视频进行投稿，简易化的用户注册流程。在刷视频的同时用户可以对视频进行点赞评论，并在个人主页显示点赞记录，同时实现可视化的关注和粉丝列表功能。
# 2. 成员分工
## 2.1  项目总体分配策划：王健霖（队长）
## 2.2 主要功能分配:
### 罗小其、王健霖

视频Feed流、视频投稿、个人信息
支持所有用户刷抖音，按投稿时间倒序推出，登录用户可以自己拍视频投稿，查看自己的基本信息和投稿列表，注册用户流程简化。

### 罗小其、刘芮彤

点赞列表、用户评论
看点赞视频列表、以及用户评论功能


### 罗小其

关注列表、粉丝列表

登录用户可以关注其他用户，能够在个人信息页查看本人的关注数和粉丝数，点击打开关注列表和粉丝列表。


### 罗小其、王健霖

bug和流程优化
修复若干bug和优化用户使用体验

## 2.3  项目展示文档：王健霖（介绍文档） 刘芮彤（汇报文档）
1. 源码设计说明
模块文件夹
模块说明

模块包含函数或文件
实现功能
 
 
auth
 
 
用于初始化用户登录信息和权限

func CreateToken(username string) (string,error)
创建用户权限 


func CheckToken(tokenString string) (string,error)
检查注册的用户权限
 
 
common
创建用户，视频等状态信息结构体，应用对象储存空间

entity_common.go
用户信息，视频信息和状态反馈信息结构体


upload.go
创建并上传对象储存空间


utils.go
拷贝用户信息结构体
 
conf
 
基本配置信息
config.ini
mysql,redis,token的配置
 
 
 
 
 
controller
 
 
实现对用户注册，用户点赞关注，视频点赞评论等功能的创建和控制
comment.go
创建视频评论列表和不同视频的用户评论状态（实现评论的增删等）


favourite.go
实现用户对视频的点赞功能，创建视频的点赞列表


feed.go
检验用户登录状态来控制视频流


publish.go
实现登录用户的视频上传功能，并创建登录用户发布视频列表，显示投稿视频


relation.go
实现关注用户操作，创建关注列表和粉丝列表


user.go
上传合法的用户注册信息，并控制用户登录界面
 
dao
 
初始化数据库
db.go
初始化数据库文件


redis.go
初始化key-value数据库储存系统
 
 
 
logic
 
对视频，用户及对应条件下信息的获取
comment.go
通过视频获取相关评论列表


favorite.go
获取用户点赞过的视频列表


feed.go
获取对应条件下的用户视频列表


relation.go
获取关注列表和粉丝列表


user.go
获取作者的用户信息
 
 
 
models
实现具体功能
comment.go 
增删获取评论列表


user.go
获取用户注册信息


user_favorite.go 

判断用户点赞状态，视频点赞情况，创建用户点赞列表


user_follower.go

获取用户粉丝，获取用户关注列表，粉丝数，关注数，判断用户之间是否关注，以及关注取关操作


vedio.go
获取保存视频信息
setting
基础配置信息
setting.go
总配置结构体，对mysql,redis,token进行相关配置
 
# 3. 项目地源码地址
https://github.com/wjl110/Teyvat
# 4. 演示视频
[飞书文档](https://bytedancecampus1.feishu.cn/docx/doxcnIU04DwZIXunXgvm20R2eJg)
# 5. 演示说明：
  
  1、测试注册登录接口   ——>  测试视频流接口  ——> 测试关注和点赞功能 ——>  用户中心查看点赞视频列表

  2、测试关注他人功能   ——>  展示关注和粉丝列表   ——> 测试发布视频（上传至阿里云oss） ——> 在作品栏中查看用户已发布视频  

  3、自己关注自己方便测试粉丝接口  ——> 刷新后粉丝数量才会变化——>测试评论区功能（刷新后呈现）
  

## 环境配置

所需依赖 
github.com/aliyun/aliyun-oss-go-sdk v2.2.4+incompatible   
github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f  // indirect    
github.com/gin-gonic/gin v1.7.7    
github.com/go-playground/universal-translator v0.18.0 // indirect    
github.com/go-playground/validator/v10 v10.8.0 // indirect    
github.com/go-redis/redis/v8 v8.11.5   
github.com/golang-jwt/jwt v3.2.1+incompatible   
github.com/jinzhu/now v1.1.5 // indirect    
github.com/kr/pretty v0.3.0 // indirect   
github.com/rogpeppe/go-internal v1.8.0 // indirect    
github.com/satori/go.uuid v1.2.0 // indirect   
golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect   
golang.org/x/text v0.3.7 // indirect   
golang.org/x/time v0.0.0-20220411224347-583f2d630306 // indirect  
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect   
gopkg.in/ini.v1 v1.66.4   
gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect  
gorm.io/driver/mysql v1.3.3   
gorm.io/gorm v1.23.5   
   
  
  
* Go 1.14.14
* MySql 8.0 
* redis任意
* 阿里云oss
* 端口8090
