/**
 * 通用登录模板,包含输入交互,提交需要自己绑定验证
 * 默认模块名: pages/login/login
 * @return {[object]}  [ 返回一个对象 ]
 */
loader.define(function(require, exports, module) {


    var pageview = {};
    // 从参数来判断用户是弹窗加载,还是跳转加载.
    var param = bui.history.getParams(module.id);

    pageview.bind = function() {

        // 手机号,帐号是同个样式名, 获取值的时候,取的是最后一个focus的值
        var userInput = bui.input({
            id: `#${module.id} .user-input`,
            callback: function(e) {
                // 清空数据
                this.empty();
            }
        })

        // 密码显示或者隐藏
        var password = bui.input({
            id: `#${module.id} .password-input`,
            iconClass: ".icon-eye",
            onBlur: function(e) {
                if (e.target.value == '') { return false; }
                // 注册的时候校验只能4-18位密码
                var rule = /^[a-zA-Z0-9_-]{2,18}$/;
                if (!rule.test(e.target.value)) {
                    bui.hint("密码只能由2-18位字母或者数字上下横杠组成");
                    return false;
                }
                return true;
            },
            callback: function(e) {
                //切换类型
                this.toggleType();
                //
                $(e.target).toggleClass("active")
            }
        })
        $(`#${module.id} #register`).on('click', function() {
            bui.load({ url: 'pages/register/index.html' });
        })
        $(`#${module.id} #login`).on('click', function(argument) {
            //href="pages/main/main.html"
            var u = userInput.value();
            var p = password.value();
            if (u == "" || p == "") {
                bui.hint("账号密码不能为空！");
                return false;
            }
            var uiLoading = bui.loading({
                appendTo: "#loading",
                width: 40,
                height: 40,
                text: 'loading'
            });

            bui.ajax({
                url: `${store.Host}login`,

                data: JSON.stringify({
                    username: u,
                    password: p,
                }),
                contentType: "application/json",
                // 可选参数
                method: "POST"
            }).then(function(result) {
                console.log(result)
                    //storage.set("userinfo", result);
                if (result && result.code != 200) {
                    bui.alert(result.msg);
                    return;
                }

                if (result && result.code == 200) {
                    store.isLogin = true;
                    // 保存本地信息, 应该保存 result 信息, 一般里面会包含token
                    storage.set("userinfo", result.data);
                    console.log(param);
                    if (param.type === "page") {
                        bui.trigger("loginsuccess", result.data)
                    } else {
                        // 登录后跳转首页, main 已经给了 pages/login/login , 这里的首页模块名是路径名了.
                        bui.hint(result.msg, function() {
                            bui.load({ url: "pages/main/main.html", param: {} });
                        });
                    }
                }
                // 成功
            }, function(result, status) {
                console.log(result, status);
            });
        })
    }
    console.log()

    pageview.init = function() {

        // 绑定事件
        this.bind();
    }


    // 初始化
    pageview.init();

    // 输出模块
    return pageview;
})