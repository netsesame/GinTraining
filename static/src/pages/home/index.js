loader.define(function(requires, exports, module, global) {

    var uiLoading = bui.loading({});
    uiLoading.show()
    var globaldata ;
    var uiList = bui.list({
        id: `#${module.id} .bui-scroll`,
        url: `${store.Host}getDb/`,
        handle: "li",
        method:'get',
	    children: ".bui-listview",
        field: {
            data: "data"
        },
        data: {
            "keyword": ''
        },
        page: 1,
        pageSize: 9999,
        onBeforeLoad:function(){
            uiLoading.show()
        },
        onFail:function(){
            uiLoading.hide()
        },
        template: function(data,result) {
            console.log(result)
            var html = "";
            var data = result.data.ALLKEY;
            console.log(data,result)
            globaldata = data;
            data.forEach(function (el, index) {
                html += `<li status="1"  id='` + index + `' key='${el.api_key}' exchange='${el.exchange}'>

                            <div class="bui-btn bui-box" >
                                <div class="span1">
                                        <h3 class="item-title bui-box-text-hide">username:${el.username}</h3>
                                        <h3 class="item-title bui-box-text-hide">key:${el.api_key}</h3>
                                        <h3 class="item-title bui-box-text-hide">secret:${el.secret}</h3>
                                        <p class="item-text bui-box-text-hide">${el.traders.join(',')}</p>
                                        <p style="margin-top:20px">
                                            <span class="success" style="padding:5px">${el.buy_float}</span> 
                                            <span class="danger" style="padding:5px">${el.sell_float}</span>
                                        </p>                                 
                                </div>
                                <div class="bui-thumbnail"><img src="images/${el.exchange.toLowerCase()}.png" alt=""></div>
                            </div>
                            <div class="bui-listview-menu swipeleft">
                                <div class="bui-btn primary">修改</div>
                                <div class="bui-btn danger">删除</div>
                            </div>
                        </li>`;
            });

            uiLoading.hide()
            return html;
        }
    });

    //加载初始化列表侧滑控件,模板里面已经静态渲染
    //加载初始化列表侧滑控件,模板里面已经静态渲染
    var uiListview = bui.listview({
        id: "#scrollList2",
        //data: [{ "text": "修改", "classname":"primary"}],
        callback: function(e, menu) {
            var $this = $(e.target);
            var text = $this.text();
            if( text == "修改" ){
                var url = "pages/home/edit"; 
                var itemli = $this.parents("li");
                var exchange = itemli.attr('exchange');
                var app_id   = itemli.attr('key');
                var uiEditPage;
                uiEditPage = bui.load({
                    url: url,
                    param: {
                        app_id: app_id,
                        exchange:exchange
                    }
                })
            }
            // 输出点击的按钮
            console.log(e.target)
                //关闭菜单
            menu.close();
            // 阻止冒泡
            e.stopPropagation();
        }
    });

    var uiListScroll = uiList.widget("scroll");

    $(`#${module.id} .icon-xiangshang2`).on("click", function (e) {
        uiListScroll.scrollTop();
    })
    // var n = 0;
    //搜索条的初始化
    var uiSearchbar = bui.searchbar({
        id: "#searchbar",
        onInput: function(e, keyword) {
            //实时搜索
            // console.log(++n)
        },
        onRemove: function(e, keyword) {
            //删除关键词需要做什么其它处理
            // console.log(keyword);
        },
        callback: function(e, keyword) {

            if (uiList) {

                //点击搜索清空数据
                uiList.empty();
                // 重新初始化数据
                uiList.init({
                    page: 1,
                    data: {
                        "keyword": keyword
                    }
                });

            }
        }

    });

})