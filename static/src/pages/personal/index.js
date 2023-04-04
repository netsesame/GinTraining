
loader.define(function(require, exports, module) {

    // 初始化数据行为存储
    var bs = bui.store({
        scope: "pageMY",
        data: {
            Fullname:''
        },
        methods: {
            logout: function() {
                // 修改全局状态会有相应处理
                storage.remove("useinfo");
                store.isLogin = false;
            }
        },
        beforeMount: function() {
            // 数据解析前执行, 修改data的数据示例
            // this.$data.a = 2
        }
    })
    var userinfo = storage.get("userinfo", 0);
    bs.Fullname = userinfo.username;

    // bui.ajax({
    //     url: store.Host+"/?s=Site.tj",
    //     data: {},//接口请求的参数
    //     // 可选参数
    //     method: "GET"
    // }).then(function(result){
    //     var html = "";
    //     console.log(result)
    //     if(result.ret ==200){
    //         var list = result.data;
    //         list.forEach(element => {
    //             html +=` <li class="bui-btn">
    //             <h3 class="item-title"><span class="primary-reverse">【${element.questype}】</span>${element.c}道题</h3>
    //         </li>`;
    //         });
    //     }
    //     $(`#${module.id} #ultj`).html(html);
    // },function(result,status){
    //     // 失败 console.log(status)
    // });

    $("#clearstorage").bind('click',function(){
        storage.remove("Traders");
        storage.remove("OrderType");
        storage.remove("exchange");
        bui.hint("清理缓存成功！");
    })
})