loader.define(function (require, exports, module) {

    var pageview = {};
    // 从参数来判断用户是弹窗加载,还是跳转加载.
    var param = bui.history.getParams(module.id);
  
    pageview.init = function () {
        //按钮在tab外层,需要传id
        var tab = bui.tab({
            id: "#tabFoot",
            menu: "#tabFootNav",
            animate: true,
            swipe:false //不允许侧滑
            ,onBeforeTo: function(e) {
            }
        })

        var uiSlideTab = bui.tab({
            id: "#uiSlideTabChild",
        })

        tab.on('to',function(){
            var index = this.index();
            // 加载delay属性的组件
            loader.delay({
                id:`#maintab${index}`
            })            
        })
    }
    
    pageview.init();
   
    return pageview;
})