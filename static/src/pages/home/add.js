

loader.define(function(require,exports,module){

    var thispage  = this;
    var pageview = {};

    window.uiLoading = window.uiLoading ||  bui.loading({opacity:0.5});

    if(!uiLoading.isLoading()) uiLoading.show();

    // 初始化数据行为存储
    var bs = bui.store({
        scope: "page",
        data: {
           a: 1,
           m:{
            username:"",
            exchange:"",
            api_key:"",
            secret:"",
            multiplier:0,
            order_type:"",
            buy_float:0,
            sell_float:0,
            nickname:""
           },
        },
        methods: {

        },
        watch: {},
        computed: {

        },
        templates: {

        },
        beforeMount: function(){
            // 数据解析前执行, 修改data的数据示例
            // this.$data.a = 2
        },
        mounted: function(){
            // 数据解析后执行
        }
    })



    var Exchange_Select;
    var orderType_Select;
    var traders_Select;
    var uiMjInput,uiHirepriceInput;
    pageview.bindExchange = function(val){
        var tmpArray =  storage.get("exchange",0);
        var ab = function(tmpArray){
            //绑定
            Exchange_Select = bui.select({
                trigger: `#${module.id} #exchange`,
                title: "请选择",
                type: "select",
                height: 300,
                autoClose:true,
                needSearch:true,
                toggle:true,
                onChange:function(e){
                    bs.m.exchange = Exchange_Select.value();
                    console.log(bs.m.exchange )
                },        
                buttons: [{ name: "重置", className: "" }, { name: "确定", className: "primary-reverse" }],
                callback: function(e) {
        
                    var text = $(e.target).text();
                    if (text == "重置") {
                        Exchange_Select.selectNone();
                    } else {
                        Exchange_Select.clearSearch();
                        Exchange_Select.hide();
                    }
                },
                data: tmpArray
            });  
        }
        if(tmpArray){
            ab(tmpArray);
        }else{
            bui.ajax({
                url: `${store.Host}exchanges`,
                data: {
                    
                },
                // 可选参数
                method: "GET"
            }).then(function(result){
                console.log(result)
                var {exchanges} = result;
                var tmpArray = $.map(exchanges, 
                    function (value, key) { 
                        return { name: value, value: value }; 
                });
                storage.set("exchange",tmpArray);
                ab(tmpArray);
                
    
            },function(result,status){
               console.log(result,status)
            });
        }
    }

    pageview.bindOrderType = function(val){
        var tmpArray =  storage.get("OrderType",0);
        var ab = function(tmpArray){
            //绑定
            orderType_Select = bui.select({
                trigger: `#${module.id} #order_type`,
                title: "请选择",
                type: "select",
                height: 300,
                autoClose:true,
                needSearch:true,
                toggle:true,
                onChange:function(e){
                    bs.m.order_type = orderType_Select.value();
                    console.log(bs.m.exchange )
                },        
                buttons: [{ name: "重置", className: "" }, { name: "确定", className: "primary-reverse" }],
                callback: function(e) {
        
                    var text = $(e.target).text();
                    if (text == "重置") {
                        orderType_Select.selectNone();
                    } else {
                        orderType_Select.clearSearch();
                        orderType_Select.hide();
                    }
                },
                data: tmpArray
            });  
        }
        if(tmpArray){
            ab(tmpArray);
        }else{
            bui.ajax({
                url: `${store.Host}order_types`,
                data: {
                    
                },
                // 可选参数
                method: "GET"
            }).then(function(result){
                console.log(result)
                var {order_types} = result;
                var tmpArray = $.map(order_types, 
                    function (value, key) { 
                        return { name: value, value: value }; 
                });
                storage.set("OrderType",tmpArray);
                ab(tmpArray);    
            },function(result,status){
               console.log(result,status)
            });
        }
    }

    pageview.bindTraders = function(val){
        var tmpArray =  storage.get("Traders",0);
        var ab = function(tmpArray){
            //绑定
            traders_Select = bui.select({
                trigger: `#${module.id} #traders`,
                title: "请选择",
                type: "checkbox",
                height: 300,
                needSearch:true,
                toggle:true,
                onChange:function(e){
                    bs.m.traders = traders_Select.value();
                    console.log(bs.m.traders )
                },        
                buttons: [{ name: "重置", className: "" }, { name: "确定", className: "primary-reverse" }],
                callback: function(e) {        
                    var text = $(e.target).text();
                    if (text == "重置") {
                        traders_Select.selectNone();
                    } else {
                        traders_Select.clearSearch();
                        traders_Select.hide();
                    }
                },
                data: tmpArray
            });  
        }
        if(tmpArray){
            ab(tmpArray);
        }else{
            bui.ajax({
                url: `${store.Host}trades`,
                data: {
                    
                },
                // 可选参数
                method: "GET"
            }).then(function(result){
                console.log(result)
                var {trades} = result;
                var tmpArray = $.map(trades, 
                    function (value, key) { 
                        return { name: value, value: value }; 
                });
                storage.set("Traders",tmpArray);
                ab(tmpArray);
                
    
            },function(result,status){
               console.log(result,status)
            });
        }
    }

    pageview.init = function(){

        uiMjInput = bui.input({
            id:`#${module.id} #Mj`,
            callback: function (e) {
                $(e.target).hide();
            },
            showIcon:false,
            maxLength:5,
            onBlur: function(e) {
                if (e.target.value == '') { return false; }
                // 注册的时候校验只能4-18位密码
                var rule = /^[0-9\.]{1,5}$/;
                if (!rule.test(e.target.value)) {
                    bui.hint("面积格式错误");
                    return false;
                }

                return true;
            },
        })
        uiHirepriceInput = bui.input({
            id:`#${module.id} #Hireprice`,
            callback: function (e) {
                $(e.target).hide();
            },
            showIcon:false,
            maxLength:5,
            onBlur: function(e) {
                if (e.target.value == '') { return false; }
                // 注册的时候校验只能4-18位密码
                var rule = /^[0-9\.]{1,10}$/;
                if (!rule.test(e.target.value)) {
                    bui.hint("价格格式错误");
                    return false;
                }
                return true;
            },
        })

    }
        //提交
    bui.btn("#submit").submit(function(loading) {

        if(bs.$data.m.exchange==''){  bui.alert('请选择市场！');  loading.stop(); return; }
        if(bs.$data.m.api_key==''){  bui.alert('请填写apikey'); loading.stop(); return; }
        if(bs.$data.m.secret==''){  bui.alert('请填写密钥');  loading.stop(); return; }
        if(bs.$data.m.order_type==''){  bui.alert('请设置订单类型');  loading.stop(); return; }
        if(bs.$data.m.buy_float==''){  bui.alert('请设置买入浮动');   loading.stop();return; }
        if(bs.$data.m.sell_float==''){  bui.alert('请设置卖出浮动');   loading.stop();return; }
        if(bs.$data.m.nickname==''){  bui.alert('请输入备注名称');   loading.stop();return; }

        
        var data = JSON.parse(JSON.stringify(bs.$data.m));
       
        var userinfo = storage.get("userinfo", 0);
        data.username = userinfo.username;
        data.id = bui.guid();
        if(data.traders){
            data.traders = data.traders.split(',');
        }
        console.log(data);

        bui.ajax({
            url: `${store.Host}addDb/`,
            data: JSON.stringify(data),
            contentType:"application/json",
            // 可选参数
            method: "POST"
        }).then(function(result){
            console.log(result)
            loading.stop();
             if(result.code==200){
                bui.hint(result.msg,function(){
                    // 如果有登录页历史
                    if (bui.history.get("pages/main/main")) {
                        // 后退到首页
                        bui.back({
                            name: "pages/main/main"
                        })
                    } else {
                        // 插入登录页
                        bui.load({
                            url: "pages/main/main"
                        })
                    }
                });
             }else {
                bui.hint(result.msg);
             }
        },function(result,status){
            loading.stop();
            bui.hint(status);
            // 失败 console.log(status)
        });

    })


    pageview.init();
    pageview.bindExchange();
    pageview.bindOrderType();
    pageview.bindTraders();
    uiLoading.hide();

    return pageview;
})