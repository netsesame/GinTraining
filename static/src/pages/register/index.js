/**
 * 通用注册模板,包含输入交互,提交需要自己绑定验证
 * 默认模块名: pages/templates/register/index
 * @return {[object]}  [ 返回一个对象 ]
 */
loader.define(function(requires, exports, module, global) {

    var uiDropdown = null;
    var uiTimer = null;
    // 初始化数据行为存储
    var bs = bui.store({
        el: `#${module.id}`,
        scope: "register",
        data: {
           userinfo: {
               phonearea:"+86",
               username:"",
               password:"",
               code:""
           },
           passtype:"password", // 1.7.2 才支持
           btnsend: {
            text: "发送验证码",
            disabled: false
           }
        },
        methods: {
            areaInit(){
                // 选择地区
                let that =this;
                var dropdown = bui.dropdown({
                    id: "#uiDropdown",
                    data: [{
                        name:"+86", // 中国
                        value:"+86"
                    },{
                        name:"+886",// 中国台湾
                        value:"+886"
                    },{
                        name:"+852",// 中国香港
                        value:"+852"
                    },{
                        name:"+853",// 中国澳门
                        value:"+853"
                    }],
                    value:"+86",
                    //设置relative为false,二级菜单继承父层宽度
                    relative: false,
                    callback: function (e) {
                        that.userinfo.phonearea = this.value(); 
                    }
                });
                return dropdown;
            },
            submit(e){
                let userinfo = this.userinfo;

                // 加密后传输
                //userinfo.password = this.md5(userinfo.password);

                // 检测是否为空
                let canSign = this.checkEmpty(userinfo);
                console.log(userinfo)

                if( !canSign ){
                    bui.hint("请填写信息");
                    return false;
                }

                // 模拟登录登陆
                bui.ajax({
                    url: store.Host+'/?s=User_user.register',
                    data: userinfo,//接口请求的参数
                    // 可选参数
                     method: "POST"
                }).then((result)=>{
                    // 成功
                    //bui.hint("欢迎您："+result.data.name);
                    console.log(result);
                    
                    if(result.ret==200){
                        bui.hint(result.data.msg);
                       setTimeout(function(){
                        bui.back();
                       },1500)
                    }else{
                        bui.hint(result.msg);
                    }
                    // 1. 如果是在首页跳转到登录页，则使用返回后调用上一个页面的方法进行刷新
                    // bui.back({callback:function(mod){ mod.refresh() }}})
                    
                    // 2. 如果是被 bui.page局部加载进页面，则可以使用, 
                    // var dialog = bui.history.getPageDialog(module.id);
                    // dialog.close();

                },function(result,status){
                    // 失败 console.log(status)
                });
            },
            md5(password){
                // 只是示例，应该先引入md5 插件
                return password;
            },
            checkEmpty(userinfo){
                
                for(let keyname in userinfo){
                    if( userinfo[keyname] == "" ){
                        return false;
                    }
                }
                return true;
            },
            check(e){
                // 校验
                let val = e.target.value;
                let rule = e.target.getAttribute('rule');
                let tip = e.target.getAttribute('tip');
                let name = e.target.getAttribute('name');
                if( rule && new RegExp(rule).test(val)){
                    return true;
                }else {
                    // 清空输入值
                    // this.setState(`userinfo.${name}`,"");
                    tip && bui.hint(tip);
                    return false;
                }
            },
            clear(str){
                // 清空值
                this.setState(str,"");
            },
            changetype(str){
                // 改变类型
                switch(this.passtype){
                    case "text":
                        this.passtype = "password";
                        break;
                    case "password":
                        this.passtype = "text";
                        break;
                }
            },
            sendcode(){
                var userinfo = this.$data.userinfo;
                if(!userinfo.username){
                    bui.hint("请输入手机号")
                    return false;
                }
                if( this.$data.btnsend.disabled ){
                    return false;
                }
                // 计时开始
                uiTimer && uiTimer.restart();

                // 禁止点击
                this.btnsend.disabled = true;

                                // 模拟登录登陆
                                bui.ajax({
                                    url: store.Host+'/?s=User_user.sensms',
                                    data: userinfo,//接口请求的参数
                                    // 可选参数
                                     method: "POST"
                                }).then((result)=>{
                                    console.log(result)
                                    if(result.ret==200){
                                        bui.hint(result.data.msg);
                                    }
                
                                },function(result,status){
                                    // 失败 console.log(status)
                                });
               

            },
            timerInit(num){
                let that = this;
                var timer = bui.timer({
                    onEnd: function() {
                        // 重新设置
                        that.btnsend = {
                            disabled: false,
                            text: "重新获取验证码"
                        }
                    },
                    onProcess: function(e) {
                        var valWithZero = e.count < 10 ? "0" + e.count : e.count;

                        that.btnsend.text = valWithZero + "后重新获取"
                    },
                    times: num || 10
                });
                
                return timer;
            }
        },
        watch: {},
        computed: {},
        templates: {},
        mounted: function(){
            // 数据解析后执行
            // 地区
            uiDropdown = this.areaInit();
            // 初始化定时器
            uiTimer = this.timerInit();

        }
    })

    return bs;
})