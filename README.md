➜  bui tree
.
├── src                                                 ## 应用目录
|   ├── index.html                                      ## 首页路由结构
|   ├── index.js                                        ## 路由初始化
|   ├── pages                                           ## 页面存放地址
|   │   ├── main
|   │   |   └── main.html                               ## 默认打开的main模板，代表首页内容（可配）
|   │   |   └── main.js                                 ## 默认打开的main模块，代表首页初始化（可配）
|   │   └── components (可选)                            ## 组件目录
|   │       └── slide                                   ## slide 组件
|   │           └── index.html                          ## slide 模板
|   │           └── index.js                            ## slide 模块
|   ├── js                                              ## 框架及插件目录
|   |   ├── zepto.js                                    ## zepto库
|   |   ├── bui.js                                      ## bui库
|   │   ├── config  (可选)                               ## 公共配置的文件目录
|   │   |   └── global.js
|   │   └── plugins (可选)                               ## 插件目录，插件必须存放在这里面才不会被二次编译
|   │   |   └── map                                     ## 地图插件
|   │   |       └── map.css
|   │   |       └── map.js
|   │   └── platform (可选)                              ## 原生的平台框架
|   │       └── cordova.js
|   ├── css                                             ## 应用的样式目录
|   |   ├── bui.css
|   |   └── style.css
|   ├── less（可选）                                     ## 使用less写样式，会覆盖掉style.css
|   |   ├── _common.less
|   |   └── style.less
|   ├── scss（可选）                                     ## 使用scss写样式，会覆盖掉style.css
|   |   ├── _common.scss
|   |   └── style.scss
│   ├── font                                            ## bui.css用到的字体图标
│   └── images                                          ## 应用的图片目录
│
├── app.json                                            ## 编译的配置
├── gulpfile.js
├── package.json
