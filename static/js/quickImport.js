$(document).ready(function() {
    /**
    * 腾讯云kendoGrid配置对象
    * columns：表格对应的所有列
    * dataSource：数据源
    * scrollable：是否可滚动
    * selectable：是否可选。multiple表示可多选，cell表示选中单位为单元格
    * allowCopy：允许ctrl+c复制，分隔符为delimiter
    * filterable：是否可过滤
    * pageable：是否分页
    * change：选中事件处理函数
    * toolbar：工具栏
    */
    var qcloudKendoGridObj = {
        columns:[
            {field:'checkbox',title:'#',menu:false,width:30,template:'<input type="checkbox" #:data.Checked# value="#:data.HostID#" class="c-grid-checkbox"/>'},
            {field:'InnerIP',title:"内网IP",filterable:true},
            {field:"OuterIP", title: "公网IP",filterable:true},
            {field:"SN",title:'SN',filterable:true},
            {field:"AssetID",title:'固资编号',filterable:true},
            {field:"ApplicationName",title:'所属业务',filterable:true},
            {field:"SetName",title:'所属集群',filterable:true},
            {field:"ModuleName",title:'所属模块',filterable:true},
            {field:"HostName",title:'主机名称',filterable:true},
            {field:"OSName",title:'操作系统',filterable:true},
            {field:"DeadLineTime",title:'到期时间',filterable:true}
        ],
        dataSource:{
            transport: {
                read: {
                    url: "/host/getHost4QuickImport",
                    data:{ApplicationID:cookie.get('defaultAppId'),IsDistributed:false,Source:"1"},
                    method:'post',
                    dataType:"json"
                }
            },
            pageSize:20,
            serverPaging: false,
            schema: {
                data: function (response) {
                    return response.data;
                },
                total: function (response) {
                    $('.bootstrap-switch-label', '#qcloud').html('<strong>'+response.total+'台</strong>');
                    return response.total;
                }
            }
        },
        selectNum: 0,
        scrollable: true,
        selectable:"multiple cell",
        allowCopy:{delimiter : ';'},
        filterable: false,
        pageable:true,
        resizable:true,
        dataBound:function(e){
            this.selectNum = 0;
        },
        change: function(e){
//            $('.k-grid-quickDistribute', '#qcloud').attr('disabled', false);
            $('.k-grid-delete', '#qcloud').attr('disabled', false);

            var grid = $('#qcloud').data('kendoGrid');
            var selectedRows = grid.select();
            $(selectedRows).closest('tr').find('input').prop('checked', true);
            for(var i=0,len=selectedRows.length; i<len; i++){
                var data = grid.dataItem($(selectedRows[i]).closest('tr'));
                if(data.Checked!='checked'){
                    this.selectNum++;
                    data.Checked = 'checked';
                }
            }
            if(this.selectNum > 0 && !$(".host-state-switcher",'#qcloud').bootstrapSwitch('state')){
                $('.k-grid-quickDistribute', '#qcloud').attr('disabled', false);
            }else{
                $('.k-grid-quickDistribute', '#qcloud').attr('disabled',true);
            }

        },
        toolbar: [//头部工具栏kendoToolBar,可以参考ui.toolbar的api
            {text:'分配至',name:'quickDistribute',attr:{"href":"javascript:void(0);","disabled":"true"}},
            {text:'同步主机',attr:{"id":"rsyncHostFromQcloud","href":"javascript:void(0);"}},
            {text:'搜索',name:'search',template:"<div class=\"c-import-search\"><input id=\"filter-qcloud\" type=\"text\" class=\"form-control pull-left\" placeholder=\"搜索...\" /><i class=\"glyphicon glyphicon-search\"></i></div>"},
            {text:'切换',name:'switch',template:"<div class=\"switch host-state-switcher-parent\"><input type=\"checkbox\" name=\"host-state-witcher\" class=\"host-state-switcher\"></div>"}
        ]
    };

    /*初始化腾讯云表格*/
    $("#qcloud").kendoGrid(qcloudKendoGridObj);

    /**
    * 其它云kendoGrid配置对象
    * columns：表格对应的所有列
    * dataSource：数据源
    * scrollable：是否可滚动
    * selectable：是否可选。multiple表示可多选，cell表示选中单位为单元格
    * allowCopy：允许ctrl+c复制，分隔符为delimiter
    * filterable：是否可过滤
    * pageable：是否分页
    * change：选中事件处理函数
    * toolbar：工具栏
    */
    var privateKendoGridObj = {
        columns:[
            {field:'checkbox',title:'#',menu:false,width:30,template:'<input type="checkbox" #:data.Checked# value="#:data.HostID#" class="c-grid-checkbox"/>'},
            {field:'InnerIP',title:"内网IP",filterable:true},
            {field:"OuterIP", title: "公网IP",filterable:true},
            {field:"SN",title:'SN',filterable:true},
            {field:"AssetID",title:'固资编号',filterable:true},
            {field:"ApplicationName",title:'所属业务',filterable:true},
            {field:"SetName",title:'所属集群',filterable:true},
            {field:"ModuleName",title:'所属模块',filterable:true},
            {field:"HostName",title:'主机名称',filterable:true},
            {field:"OSName",title:'操作系统',filterable:true},
            {field:"DeadLineTime",title:'到期时间',filterable:true,hidden:true}
        ],
        dataSource:{
            transport: {
                read: {
                    url: "/host/getHost4QuickImport",
                    data:{ApplicationID:cookie.get('defaultAppId'),IsDistributed:false,Source:"3"},
                    method:'post',
                    dataType:"json"
                }
            },
            pageSize:20,
            serverPaging: false,
            schema: {
                data: function (response) {
                    return response.data;
                },
                total: function (response) {
                    $('.bootstrap-switch-label', '#private').html('<strong>'+response.total+'台</strong>');
                    return response.total;
                }
            }
        },
        selectNum: 0,
        scrollable: true,
        selectable:"multiple cell",
        allowCopy:{delimiter : ';'},
        filterable: false,
        pageable:true,
        resizable:true,
        dataBound:function(e){
            this.selectNum = 0;
        },
        change: function(e){
//            $('.k-grid-quickDistribute', '#private').attr('disabled', false);
            $('.k-grid-delete', '#private').attr('disabled', false);

            var grid = $('#private').data('kendoGrid');
            var selectedRows = grid.select();
            $(selectedRows).closest('tr').find('input').prop('checked', true);
            for(var i=0,len=selectedRows.length; i<len; i++){
                var data = grid.dataItem($(selectedRows[i]).parents('tr'));
                if(data.Checked!='checked'){
                    this.selectNum++;
                    data.Checked = 'checked';
                }
            }
            if(this.selectNum > 0 && !$(".host-state-switcher",'#private').bootstrapSwitch('state')){
                $('.k-grid-quickDistribute', '#private').attr('disabled', false);
            }else{
                $('.k-grid-quickDistribute', '#private').attr('disabled',true);
            }
        },
        toolbar: [//头部工具栏kendoToolBar,可以参考ui.toolbar的api
            {text:'分配至',name:'quickDistribute',attr:{"href":"javascript:void(0);","disabled":"true"}},
            {text:'导入主机',template:"<a class=\"k-button\" id=\"importPrivateHostByExcel\" href=\"javascript:void(0);\"><span class=\"\"></span>导入主机</a>"},
            // {text:'下载模板',attr:{"id":"getCCTemplate","href":"/static/excel/importToCC.xls"}},
            {text:'删除',name:'delete',attr:{"href":"javascript:void(0);","disabled":"true"}},
            {text:'搜索',template:"<div class=\"c-import-search\"><input id=\"filter-private\" type=\"text\" class=\"form-control pull-left\" placeholder=\"搜索...\" /><i class=\"glyphicon glyphicon-search\"></i></div>"},
            {text:'切换',template:"<div class=\"switch host-state-switcher-parent\"><input type=\"checkbox\" name=\"host-state-witcher\" class=\"host-state-switcher\"></div>"}
        ]
    };

    /*初始化其它云表格*/
    $("#private").kendoGrid(privateKendoGridObj);

    /**
    * 投后版kendoGrid配置对象
    * columns：表格对应的所有列
    * dataSource：数据源
    * scrollable：是否可滚动
    * selectable：是否可选。multiple表示可多选，cell表示选中单位为单元格
    * allowCopy：允许ctrl+c复制，分隔符为delimiter
    * filterable：是否可过滤
    * pageable：是否分页
    * change：选中事件处理函数
    * toolbar：工具栏
    */
    var tencentKendoGridObj = {
        columns:[
            {field:'checkbox',title:'#',menu:false,width:30,template:'<input type="checkbox" #:data.Checked# value="#:data.HostID#" class="c-grid-checkbox"/>'},
            {field:'InnerIP',title:"内网IP",filterable:true},
            {field:"OuterIP", title: "公网IP",filterable:true},
            {field:"SN",title:'SN',filterable:true},
            {field:"AssetID",title:'固资编号',filterable:true},
            {field:"ApplicationName",title:'所属业务',filterable:true},
            {field:"SetName",title:'所属集群',filterable:true},
            {field:"ModuleName",title:'所属模块',filterable:true},
            {field:"HostName",title:'主机名称',filterable:true},
            {field:"OSName",title:'操作系统',filterable:true},
            {field:"DeadLineTime",title:'到期时间',filterable:true,hidden:true}
        ],
        dataSource:{
            transport: {
                read: {
                    url: "/host/getHost4QuickImport",
                    data:{ApplicationID:cookie.get('defaultAppId'),IsDistributed:false,Source:"2"},
                    method:'post',
                    dataType:"json"
                }
            },
            pageSize:20,
            serverPaging: false,
            schema: {
                data: function (response) {
                    return response.data;
                },
                total: function (response) {
                    $('.bootstrap-switch-label', '#tencent').html('<strong>'+response.total+'台</strong>');
                    return response.total;
                }
            }
        },
        selectNum: 0,
        scrollable: true,
        selectable:"multiple cell",
        allowCopy:{delimiter : ';'},
        filterable: false,
        pageable:true,
        dataBound:function(e){
            this.selectNum = 0;
        },
        change: function(e){
//            $('.k-grid-quickDistribute', '#tencent').attr('disabled', false);
            $('.k-grid-delete', '#tencent').attr('disabled', false);

            var grid = $('#tencent').data('kendoGrid');
            var selectedRows = grid.select();
            $(selectedRows).closest('tr').find('input').prop('checked', true);
            for(var i=0,len=selectedRows.length; i<len; i++){
                var data = grid.dataItem($(selectedRows[i]).closest('tr'));
                if(data.Checked!='checked'){
                    this.selectNum++;
                    data.Checked = 'checked';
                }
            }
            if(this.selectNum > 0 && !$(".host-state-switcher",'#tencent').bootstrapSwitch('state')){
                $('.k-grid-quickDistribute', '#tencent').attr('disabled', false);
            }else{
                $('.k-grid-quickDistribute', '#tencent').attr('disabled',true);
            }
        },
        toolbar: [//头部工具栏kendoToolBar,可以参考ui.toolbar的api
            {text:'分配至',name:'quickDistribute',attr:{"href":"javascript:void(0);","disabled":"true"}},
            // {text:'更新主机',template:"<form action=\"/host/importRendHostByExcel\" enctype=\"multipart/form-data\" method=\"post\" target=\"upload_proxy\" style=\"display:inline;\"><a class=\"k-button king-file-btn filebox mt5\">导入主机<input type=\"file\" id=\"importRendHostByExcel\" name=\"importRendHostByExcel\"></a></form>"},
            // {text:'下载模板',name:'getCmdbTemplate',attr:{"id":"getCmdbTemplate","href":"/static/excel/importToCmdb.xls"}},
            {text:'更新主机',template:"<a class=\"k-button\" id=\"importRendHostByExcel\" href=\"javascript:void(0);\"><span class=\"\"></span>更新主机</a>"},
            {text:'搜索',template:"<div class=\"c-import-search\"><input id=\"filter-tencent\" type=\"text\" class=\"form-control pull-left\" placeholder=\"搜索...\" /><i class=\"glyphicon glyphicon-search\"></i></div>"},
            {text:'切换',template:"<div class=\"switch host-state-switcher-parent\"><input type=\"checkbox\" name=\"host-state-witcher\" class=\"host-state-switcher\"></div>"}
        ]
    };

    /*初始化投后版表格*/
    $("#tencent").kendoGrid(tencentKendoGridObj);

    (function(){
        /*初始化标签*/
        var current_tab = cookie.get('quick_destribute_current_tab');
        if(current_tab!=''){
            $('.nav-tabs').find('a').each(function(i, el){
                var type = $(el).attr('href');
                if(type.indexOf(current_tab)>-1){
                    $(el).click();
                    cookie.set('quick_destribute_current_tab', type.replace('#',''));
                }
            });
        }else{
            var el = $('.nav-tabs').find('a').eq(0);
            cookie.set('quick_destribute_current_tab', el.attr('href').replace('#',''));
            el.click(); 
        }

        /*标签点击事件*/
        $('.nav-tabs').on('click', function(e){
            if($(e.target).attr('href')){
                cookie.set('quick_destribute_current_tab', $(e.target).attr('href').replace('#',''));
            }
        });

        /*表格表头添加checkbox*/
        var type = ['qcloud', 'tencent', 'private'];
        for(var i=0,len=type.length; i<len; i++){
            var checkAll = $('<input type="checkbox" data-field="checkAll"/>');
            $('#'+type[i]).find('th[data-field=checkbox]').empty().append(checkAll);
        }
    })();

    /**
    * 表格工具栏的切换按钮
    * size：按钮大小
    * labelWidth：切换按钮的label宽度
    * onText：切换按钮on状态的文字
    * offText：切换按钮off状态的文字
    * onSwitchChange：状态切换时间处理函数
    */
    $(".host-state-switcher").bootstrapSwitch({
        size: 'small',
        labelWidth:'60px',
        onText: '已分配',
        offText: '未分配',
        onSwitchChange: function(e, state){
            var id = $(e.target).parents('.tab-pane').attr('id');
            var grid = $('#'+id).data('kendoGrid');
            grid.thead.find('input').attr('checked', false);
            grid.destroy();
            var gridObj = eval(id+'KendoGridObj');
            gridObj.dataSource.transport.read.data.IsDistributed = state;
            $('#'+id).kendoGrid(gridObj);
            grid = $('#'+id).data('kendoGrid');
            grid.refresh();
            $('#filter-'+id).data('data',{});
            if(state == true) {
                $('.k-grid-quickDistribute', '#'+id).attr('title', '配置平台禁止跨业务分配主机，如果实在要用，请联系原主机业务的运维同学上交后再分配');
            }
            else{
                $('.k-grid-quickDistribute', '#'+id).removeAttr('title');
            }
            /*if(id==='private'){
                if(state){
                    grid.hideColumn('AssetID');
                    grid.hideColumn('SN');
                    grid.hideColumn('OSName');
                    grid.showColumn('ApplicationName');
                    grid.showColumn('SetName');
                    grid.showColumn('ModuleName');
                }else{
                    grid.showColumn('AssetID');
                    grid.showColumn('SN');
                    grid.showColumn('OSName');
                    grid.hideColumn('ApplicationName');
                    grid.hideColumn('SetName');
                    grid.hideColumn('ModuleName');
                }
            }*/

            $('.k-grid-delete,.k-grid-quickDistribute', '#'+id).attr('disabled', true);
            if(state){
                $('.k-grid-delete', '#'+id).hide();
            }else{
                $('.k-grid-delete', '#'+id).show();
            }

            $('#filter-'+id).val('');
        }
    });
    
    /**
    * 表头搜索框输入时间处理函数
    * 支持所有字段的搜索，字段之间逻辑关系为or，搜索方式为contains，即包含
    */
    $('#filter-qcloud,#filter-tencent,#filter-private').on('keyup', function(e){
        var type = $(e.target).attr('id').split('-').pop();
        var grid = $('#'+type).data('kendoGrid');
        if(typeof JSON=='undefined'){
            $('head').append('<script type="text/javascript" src="/static/js/json2.js"></script>');
        }

        if($.isEmptyObject($(e.target).data('data'))){
            var data = grid.dataSource.data();
            $(e.target).data('data', data);
        }else{
            var data = $(e.target).data('data');
        }
        var d = JSON.parse(JSON.stringify(data));
        for(var i in d){
            d[i].Checked = '';
        }
        grid.dataSource.data(d);
        grid.refresh();
        grid.thead.find('input[type=checkbox]').prop('checked', false);

        filter = {logic: "or", filters: []};
        $searchValue = $(e.target).val();
        if ($searchValue) {
            $.each(grid.columns, function (key, column) {
                if (column.filterable) {
                    filter.filters.push({field: column.field, operator: "contains", value: $searchValue});
                }
            });
        }

        grid.dataSource.options.serverFiltering = false;
        grid.dataSource.filter(filter);
        grid.selectNum = 0;
        var query = new kendo.data.Query(grid.dataSource.data());
        $('.bootstrap-switch-label', '#'+type).html('<strong>'+query.filter(filter).data.length+'台</strong>');
        $('.k-grid-delete,.k-grid-quickDistribute', '#'+type).attr('disabled', true);
    });
    
    /**
    * 表头工具栏“分配至”点击事件处理函数
    */
    $('.k-grid-quickDistribute').on('click', function(e){
        if($(e.target).attr('disabled')=='disabled'){
            return false;
        }

        var id = $(e.target).parents('.tab-pane').attr('id');
        var grid = $('#'+id).data('kendoGrid');
        var data = grid.dataSource.data();
        if(typeof JSON=='undefined'){
            $('head').append('<script type="text/javascript" src="/static/js/json2.js"></script>');
        }
        var d = JSON.parse(JSON.stringify(data));
        var appId = [];
        var hostId = [];
        var IsDistributed = false
        for(var i=0,len=d.length; i<len; i++){
            if(d[i].Checked==='checked'){
                if($.inArray(d[i].ApplicationID, appId)==-1){
                    appId.push(d[i].ApplicationID);
                }

                hostId.push(d[i].HostID);
                
                if(d[i].ApplicationName.indexOf('资源池')>-1){
                    IsDistributed = true;
                }
            }
        }

        if(hostId.length===0){
            var noHostSelectDialog = dialog({
                content: '<div class="c-dialogdiv2"><i class="c-dialogimg-prompt"></i>请选择主机</div>'
            });
            noHostSelectDialog.show();
            setTimeout(function () {
                noHostSelectDialog.close().remove();
            }, 2000);
            return false;
        }

        var param = {};
        param['HostID'] = hostId.join(',');
        /*重构完成后，需改动*/
        param['ApplicationID'] = appId.join(',');//标签上放appId
        param['ToApplicationID'] = $('#appId').attr('ApplicationID') ? $('#appId').attr('ApplicationID') : cookie.get('defaultAppId');//标签上放appId

        var options = {
            title:'确认',
            width:300,
            content: '当前选择的主机已经被分配到其他业务使用，确认继续分配至<i class="redFont">'+ cookie.get('defaultAppName')+'</i>?',
            okValue:"继续",
            cancelValue:"我再想想",
            ok:function (){
                var d = dialog({
                    content: '<div class="c-dialogdiv2"><img class="c-dialogimg-loading" src="/static/img/loading_2_24x24.gif"></img>正在分配...</div>'
                });
                d.showModal();

                $.ajax({
                    dialog:d,
                    url:'/host/quickDistribute',
                    data:param,
                    method:'post',
                    dataType:'json',
                    success:function(response){
                        this.dialog.close().remove();
                        if(response.success && isNewUser && cookie.get('play_user_guide_quickimport')==1){
                            cookie.set('quickimport_user_guide_step', 5);
                            var distributeHostDialog = dialog({
                                title:'提示',
                                width:300,
                                height:50,
                                content: '<div class="c-dialogdiv"><i class="c-dialogimg-success"></i>恭喜，您离成功越来越近，<a href="/topology/index">点此</a>将您的业务拓扑绘制到配置平台</div>',  
                                okValue:"确定",
                                ok:function (){
                                    location.href = '/topology/index';
                                }
                            });

                            distributeHostDialog.showModal();
                            return true;
                        }else{
                            var content = response.success==true ? '<i class="c-dialogimg-success"></i>'+response.message : '<i class="c-dialogimg-prompt"></i>'+response.errInfo;
                            var d = dialog({
                                    content: '<div class="c-dialogdiv2">'+content+'</div>'
                                });
                            d.showModal();
                            setTimeout(function() {
                                d.close().remove();
                                window.location.reload();
                            }, 2500);
                        }
                        return true;
                    }
                });
            },
            cancel:function(){}
        };

        if(IsDistributed){
            options.content = '当前操作会将已勾选的<i class="redFont">'+ hostId.length +'</i>台主机分配至<i class="redFont">'+ cookie.get('defaultAppName') +'</i>的空闲机池，确认继续？';
        }

        var quickDistributeDialog = dialog(options);
        quickDistributeDialog.showModal();

        return true;
    });
    
    /**
    * 表格第一列复选框勾选事件处理函数
    */
    $('#qcloud,#tencent,#private').on('change', 'input[type=checkbox]', function(e){
        var type = $(e.target).parents('.tab-pane').attr('id');
        var grid = $('#'+type).data('kendoGrid');
        var filter = grid.dataSource.filter();

        if($(e.target).attr('data-field')==='checkAll'){
            var checkAll = $(e.target).prop('checked');
            if($.isEmptyObject($('#filter-'+type).data('data'))){
                var allData = grid.dataSource.data();
                $(e.target).data('data', allData);
            }else{
                var allData = $('#filter-'+type).data('data');
            }

            if(typeof JSON=='undefined'){
                $('head').append('<script type="text/javascript" src="/static/js/json2.js"></script>');
            }
            var query = new kendo.data.Query(allData);
            var data = query.filter(filter).data;
            var d = JSON.parse(JSON.stringify(data));
            for(var i in d){
                d[i].Checked = checkAll ? 'checked' : '';
            }

            grid.dataSource.data(d);
            grid.refresh();

            grid.selectNum = checkAll ? d.length : 0;
        }else{
            var checked = $(e.target).prop('checked');
            var data = grid.dataItem($(e.target).closest('tr'));
            data.Checked = checked ? 'checked' : '';

            if(checked){
                grid.selectNum++;
            }else{
                grid.selectNum--;
            }

            grid.thead.find('input').prop('checked', grid.selectNum === grid.dataSource._total);
        }

        if(grid.selectNum === grid.dataSource._total){
            var selectAllDialog = dialog({
                content: '<div class="c-dialogdiv2"><i class="c-dialogimg-success"></i>全选<i class="redFont">'+grid.selectNum+'</i>台主机</div>'
            });
            selectAllDialog.show();
            setTimeout(function(){selectAllDialog.close().remove();}, 2000);
        }


        var quickDistribute = $('.k-grid-quickDistribute', '#'+type);
        var gridObj = eval(type+'KendoGridObj');
        console.log(gridObj.dataSource.transport.read.data.IsDistributed);
        if(gridObj.dataSource.transport.read.data.IsDistributed == false){
            quickDistribute.attr('disabled', grid.selectNum>0 ? false : true);
        }else{
            quickDistribute.attr('disabled', true);
        }
        $('.k-grid-delete', '#'+type).attr('disabled', grid.selectNum>0 ? false : true);
    });
    
    /**
    * 表头工具栏“同步主机”点击事件处理函数
    * 从腾讯云同步主机
    */
    $('#rsyncHostFromQcloud').click(function(){
        $('.introjs-button').click();
        var d = dialog({
                content: '<div class="c-dialogdiv2"><img class="c-dialogimg-loading" src="/static/img/loading_2_24x24.gif"></img>处理中，请稍候...</div>'
            });
        d.showModal();

        $.ajax({
            url:'/host/importHostByQcloud',
            method:'post',
            dataType:'json',
            dialog:d,
            success:function(response){
                this.dialog.close().remove();

                if(response.success && isNewUser && cookie.get('play_user_guide_quickimport')==1){
                    var d = dialog({
                        title:'确认',
                        width:300,
                        content: '<div class="c-dialogdiv2"><i class="c-dialogimg-success"></i>您可以通过分配至将您勾选的云主机快速分配到当前业务.</div>',
                        okValue:"确定",
                        ok:function(){
                            if(isNewUser && cookie.get('play_user_guide_quickimport')==1){
                                cookie.set('quickimport_user_guide_step', 3);
                                window.location.reload();
                            }
                        }
                    });
                    d.showModal()
                    cookie.set('quickimport_user_guide_step', 3);
                }else{
                    if(response.success){
                        var d = dialog({
                            content: '<div class="c-dialogdiv2"><i class="c-dialogimg-success"></i>'+ response.message +'</div>'
                        });
                        d.showModal();
                        setTimeout(function(){window.location.reload();}, 2000);
                    }else{
                        if(isNewUser && cookie.get('play_user_guide_quickimport')==1){
                            var d = dialog({
                                title:'确认',
                                width:300,
                                content: '<div class="c-dialogdiv2"><i class="c-dialogimg-prompt"></i>'+ response.message +'或<a href="javascript:void" id="user_guide_import_private">导入</a>其它云主机</div>',
                                okValue:"确定",
                                ok:function(){
                                    if(isNewUser && cookie.get('play_user_guide_quickimport')==1){
                                        //cookie.set('quickimport_user_guide_step', 3);//成功才进入下一步
                                        window.location.reload();
                                    }
                                }
                            });

                            window.d = d;
                        }else{
                            var d = dialog({
                                title:'提示',
                                width:315,
                                content: '<div class="c-dialogdiv2"><i class="c-dialogimg-prompt"></i>'+ response.message +'</div>',
                                okValue:"我知道了",
                                ok:function(){
                                }
                            });
                        }
                        d.showModal();
                    }
                }

                cookie.set('quick_destribute_current_tab', 'qcloud');
                return true;
            }
        });
    });
    
    /**
    * 表头工具栏“删除”按钮点击事件处理函数
    * 将从其它云，利用excel导入的主机从配置平台彻底删除
    */
    $(".k-grid-delete").on('click', function(e){
        if($(e.target).attr('disabled')=='disabled'){
            return false;
        }

        var param = {};
        var type = $(e.target).parents('.tab-pane').attr('id');
        var grid = $('#'+type).data('kendoGrid');
        var data = grid.dataSource.data();
        if(typeof JSON=='undefined'){
            $('head').append('<script type="text/javascript" src="/static/js/json2.js"></script>');
        }
        var d = JSON.parse(JSON.stringify(data));
        var hostId = [];
        var appId = [];
        for(var i=0,len=d.length; i<len; i++){
            if(d[i].ApplicationName!=='资源池'){
                var notAllowToDeleteDialog = dialog({
                    content: '<div class="c-dialogdiv2"><i class="c-dialogimg-prompt"></i>只能删除未分配机器</div>'
                });
                notAllowToDeleteDialog.show();
                setTimeout(function () {
                    notAllowToDeleteDialog.close().remove();
                }, 2000);
                return false;
            }

            if(d[i].Checked==='checked'){
                var tmp = d[i].ApplicationID.split(',');
                if(tmp.length==1){
                    if($.inArray(d[i].ApplicationID, appId)==-1){
                        appId.push(d[i].ApplicationID);
                    }
                }else{
                    var tmp = d[i].ApplicationID.split(',');
                    for(var j=0,jlen=tmp.length; j<jlen; j++){
                       if($.inArray(tmp[j], appId)==-1){
                            appId.push(tmp[j]);
                        }
                    }
                }

                hostId.push(d[i].HostID);
            }
        }

        if(hostId.length==0){
            var noHostSelectDialog = dialog({
                content: '<div class="c-dialogdiv2"><i class="c-dialogimg-prompt"></i>请选择主机</div>'
            });
            noHostSelectDialog.show();
            setTimeout(function () {
                noHostSelectDialog.close().remove();
            }, 2000);
            return false;
        }

        param['HostID'] = hostId.join(',');
        /*重构完成后，需改动*/
        param['ApplicationID'] = appId.join(',');//标签上放appId
        var options = {
            title:'确认',
            width:300,
            content: '您勾选的<i class="redFont">' + hostId.length + '</i>台主机即将离开配置平台，确认是否继续？',
            okValue:"继续",
            cancelValue:"我再想想",
            ok:function (){
                var d = dialog({
                    content: '<div class="c-dialogdiv2"><img class="c-dialogimg-loading" src="/static/img/loading_2_24x24.gif"></img>正在删除...</div>'
                });
                d.showModal();

                $.ajax({
                    dialog:d,
                    url:'/host/delPrivateDefaultApplicationHost',
                    data:param,
                    method:'post',
                    dataType:'json',
                    success:function(response){
                        this.dialog.close().remove();
                        if(response.success){
                            var distributeHostDialog = dialog({
                                title:'提示',
                                width:300,
                                height:50,
                                content: '<div class="c-dialogdiv"><i class="c-dialogimg-success"></i>删除成功!</div>',
                                okValue:"确定",
                                ok:function (){
                                    location.href = '/host/quickImport';
                                }
                            });

                            distributeHostDialog.showModal();
                            return true;
                        }else{
                            var content = response.success==true ? '<i class="c-dialogimg-success"></i>'+response.message : '<i class="c-dialogimg-prompt"></i>'+response.message;
                            var d = dialog({
                                width: 150,
                                content: '<div class="c-dialogdiv2">'+content+'</div>'
                            });
                            d.showModal();
                            setTimeout(function() {
                                d.close().remove();
                                window.location.reload();
                            }, 2500);
                        }
                        return true;
                    }
                });
            },
            cancel:function(){}
        };

        var quickDistributeDialog = dialog(options);
        quickDistributeDialog.showModal();

        return true;
    });

    $(document.body).on('click', '#user_guide_import_private', function(e){
        d.close().remove();step1();
    });

    // $(document.body).on('change', '#importRendHostByExcel,#importPrivateHostByExcel', function(e){
    //     uploadDialog = dialog({
    //         content: '<div class="c-dialogdiv2"><img class="c-dialogimg-loading" src="/static/img/loading_2_24x24.gif"></img>正在导入...</div>'
    //     });
    //     uploadDialog.showModal();
    //     setTimeout(function(){
    //         //$(e.target).parents('form').submit();
    //         clearUpload($(e.target).attr('id'));
    //     },500);
    // });

    $('.import-page-mask').click(function(e){
        $(this).hide();
        //cookie.set('play_user_guide_quickimport', 0);//设置退出新手引导标记
    });

    $('#syncHost').click(function(){
        cookie.set('quick_destribute_current_tab', 'qcloud');
        cookie.set('quickimport_user_guide_step', 2);
        step2();
    });

    $('#importOtherHost').click(function(){
        cookie.set('quick_destribute_current_tab', 'private');
        cookie.set('quickimport_user_guide_step', 1);
        step1();
    });

    // $('#getCCTemplate').click(function(){
    //     $('.introjs-button').click();
    //     if(isNewUser && cookie.get('play_user_guide_quickimport')==1){
    //         step2();
    //     }
    // });

    // 导入私有云机器
    $('#importPrivateHostByExcel').on('click',function (e){
        var importPrivateHost = dialog({
                title:'导入主机',
                width:500,
                content: '<div class="pt10">'+
                         '<form action="/host/importPrivateHostByExcel" id="upload_form" enctype="multipart/form-data" method="post" target="upload_proxy" style="display:inline-block;">'+
                         '<lable><span class="c-gridinputmust pr10">*</span>请选择导入文件：</lable>'+
                         '<a class="k-button king-btn-mini king-file-btn filebox">选择文件'+
                         '<input type="file" id="importPrivateHost" name="importPrivateHost">'+
                         '</a>'+
                         '<span class="import-file-name ml15"></span>'+
                         '<p style="color:#666;padding:10px 0 0 5px;"></p>'+
                         '<p class="text-warning">（温馨提示：1.文件类型只支持.xls;  2.<a href="/static/excel/importToCC.xls">下载模版</a>） </p>'+
                         // '<img src="/static/img/loading_2_24x24.gif" alt="" id="loader" style="display:none;">'+
                         '</form>'+
                         '</div>',
                okValue:"导入",
                cancelValue:"关闭",
                skin:'dia-grid-batDel',
                ok:function (){
                    $("#upload_form").submit();
                    uploadDialog = dialog({
                        content: '<div class="c-dialogdiv2"><img class="c-dialogimg-loading" src="/static/img/loading_2_24x24.gif"></img>正在导入...</div>'
                    });
                    uploadDialog.showModal();
                    setTimeout(function(){
                        //clearUpload($(e.target).attr('id'));
                    },500);
                }
            });
        importPrivateHost.showModal();
        $('#importPrivateHost').on('change', function(){ 
          if (!$('.import-file-name').text($('#importPrivateHost').val().split('\\')[$('#importPrivateHost').val().split('\\').length-1])) {
               $('.import-file-name').text($('#importPrivateHost').val().split('/')[$('#importPrivateHost').val().split('/').length-1])
          };
        });
    })

    //更新投后机器
    $('#importRendHostByExcel').on('click',function (e){
        var importRendHost = dialog({
                title:'更新主机',
                width:500,
                content: '<div class="pt10">'+
                         '<form action="/host/importRendHostByExcel" id="upload_form" enctype="multipart/form-data" method="post" target="upload_proxy" style="display:inline-block;">'+
                         '<lable><span class="c-gridinputmust pr10">*</span>请选择导入文件：</lable>'+
                         '<a class="k-button king-btn-mini king-file-btn filebox">选择文件'+
                         '<input type="file" id="importRendHost" name="importRendHost">'+
                         '</a>'+
                         '<span class="import-file-name ml15"></span>'+
                         '<p style="color:#666;padding:10px 0 0 5px;"></p>'+
                         '<p class="text-warning">（温馨提示：1.文件类型只支持.xls;  2.<a href="/static/excel/importToCmdb.xls">下载模版</a>） </p>'+
                         // '<img src="/static/img/loading_2_24x24.gif" alt="" id="loader" style="display:none;">'+
                         '</form>'+
                         '</div>',
                okValue:"导入",
                cancelValue:"关闭",
                skin:'dia-grid-batDel',
                ok:function (){
                    $("#upload_form").submit();
                    uploadDialog = dialog({
                        content: '<div class="c-dialogdiv2"><img class="c-dialogimg-loading" src="/static/img/loading_2_24x24.gif"></img>正在导入...</div>'
                    });
                    uploadDialog.showModal();
                    setTimeout(function(){
                        //clearUpload($(e.target).attr('id'));
                    },500);
                }
            });
        importRendHost.showModal();
        $('#importRendHost').on('change', function(){ 
          if (!$('.import-file-name').text($('#importRendHost').val().split('\\')[$('#importRendHost').val().split('\\').length-1])) {
               $('.import-file-name').text($('#importRendHost').val().split('/')[$('#importRendHost').val().split('/').length-1])
          };
        });
    })
});


/**
* 重置上传表单，防止相同文件上传没反应
*/
function clearUpload(id){
    $("#"+id).parents('form').submit().end().remove();//移除原来的
    $("<input/>").attr("name",id).attr("id",id).attr("type","file").appendTo(".filebox");//添加新的
}

/**
* 导入主机回调函数
* 负责页面显示成功or失败的提示
*/
function uploadCallback(data){
    uploadDialog.close().remove();
    if(data.success && isNewUser && cookie.get('play_user_guide_quickimport')==1){
        var d = dialog({
            title:'确认',
            width:300,
            content: '<div class="c-dialogdiv2"><i class="c-dialogimg-success"></i>您可以通过分配至将您勾选的云主机快速分配到当前业务.</div>',
            okValue:"确定",
            ok:function(){
                if(isNewUser && cookie.get('play_user_guide_quickimport')==1){
                    cookie.set('quickimport_user_guide_step', 3);
                    window.location.reload();
                }
            }
        });
        d.showModal()
        cookie.set('quickimport_user_guide_step', 3);
    }else{
        if(data.success){
            var d = dialog({
                content: '<div class="c-dialogdiv2"><i class="c-dialogimg-success"></i>'+ data.message +'</div>'
            });
            d.showModal();
            setTimeout(function(){window.location.reload();}, 2000);
        }else{
            var d = dialog({
                title:'确认',
                width:300,
                content: '<div class="c-dialogdiv2"><i class="c-dialogimg-prompt"></i>'+ data.message +'</div>',
                okValue:"确定",
                ok:function(){}
            });

            d.showModal();
            $(".import-error-list").mCustomScrollbar({
                theme: "minimal-dark" //设置风格
            });
        }
    }
}

/**
* 新手引导第一步
*/
function step0(){
    //显示同步腾讯云主机+导入其它云主机按钮
    $('.import-page-mask').removeClass('none');
}

/**
* 新手引导第二步
*/
function step1(){
    //下载模版
    $('.c-import-block').find('[href="#private"]').trigger('click');
    window.setTimeout(function(){
        var intro = introJs();
        intro.setOptions({
            doneLabel:'我知道了',                 
            showBullets:false,
            showStepNumbers:false,
            steps: [{
            element: $('#private').get(0),
            intro: '您可以下载Excel模版，并按要求填写要导入的数据，做好导入准备。'
            }]
        });

        intro.onexit(function() {
            //cookie.set('play_user_guide_quickimport', 0);//设置退出新手引导标记
            cookie.set('quickimport_user_guide_step', 2);
        });

        intro.oncomplete(function(){
            //cookie.set('play_user_guide_quickimport', 0);//设置退出新手引导标记
            cookie.set('quickimport_user_guide_step', 2);
        });

        intro.start();
    },500);
}

/**
* 新手引导第三步
*/
function step2(){
    //导入excel
    var type = cookie.get('quick_destribute_current_tab');
    if(type=='qcloud'){
        $('.import-page-mask').addClass('none');
        $('.c-import-block').find('[href="#qcloud"]').trigger('click');
        window.setTimeout(function(){
            var intro = introJs();
            intro.setOptions({
                doneLabel:'我知道了',                 
                showBullets:false,
                showStepNumbers:false,
                steps: [{
                element: document.querySelector('#rsyncHostFromQcloud'),
                intro: '您可以通过同步功能至将您的腾讯云主机同步到分配池'
                }]
            });

            intro.onexit(function() {
                //cookie.set('play_user_guide_quickimport', 0);//设置退出新手引导标记
            });

            intro.oncomplete(function(){
                //$('#rsyncHostFromQcloud').click();
            });

            intro.start();
        },500);
    }else{
        $('.import-page-mask').addClass('none');
        $('.c-import-block').find('[href="#private"]').trigger('click');
        window.setTimeout(function(){
            var intro = introJs();
            intro.setOptions({
                doneLabel:'完成',
                showBullets:false,
                showStepNumbers:false,
                steps: [{
                element: document.querySelector('#importPrivateHostByExcel').parentNode,
                intro: '您可以通过导入功能至将您的主机导入分配池'
            }]
            });

            intro.onexit(function() {
                //cookie.set('quickimport_user_guide_step', 1);//设置退出新手引导标记
            });

            intro.oncomplete(function(){
                //$('#importToCC').click();
            });

            intro.start();
        },500);
    }
}

/**
* 新手引导第四步
*/
function step3(){
    //分配机器
    var type = cookie.get('quick_destribute_current_tab');
    $('.import-page-mask').addClass('none');

    $('.c-import-block').find('[href="#'+type+'"]').trigger('click');
    window.setTimeout(function(){
        var intro = introJs();
        intro.setOptions({
            doneLabel:'我知道了',                 
            showBullets:false,
            showStepNumbers:false,
            steps: [{
            element: $('.k-grid-quickDistribute', '#'+type).get(0),
            intro: '这里可以为已创建的业务分配主机'
            }]
        });

        intro.onexit(function() {
            //cookie.set('play_user_guide_quickimport', 0);//设置退出新手引导标记
            //cookie.set('quickimport_user_guide_step', 4);//分配主机成功后，才进入下一步
        });

        intro.oncomplete(function(){
            //cookie.set('play_user_guide_quickimport', 0);//设置退出新手引导标记
            //cookie.set('quickimport_user_guide_step', 4);//分配主机成功后，才进入下一步
        });

        intro.start();
    },500);
}

/**
* 新手引导第五步
*/
function step4(){
    //跳转
    var d = dialog({
        title:'提示',
        width:300,
        height:50,
        content: '<div class="c-dialogdiv"><i class="c-dialogimg-success"></i>恭喜，您离成功越来越近，<a href="/topology/index">点此</a>将您的业务拓扑绘制到配置平台</div>',  
        okValue:"确定",
        ok:function (){
            location.href = '/topology/index';
        }
    });

    d.showModal();
    cookie.set('quickimport_user_guide_step', 5);
}

/**
* 新手引导第六步
*/
function step5(){
    //提示已分配

    var type = cookie.get('quick_destribute_current_tab');
    $('.import-page-mask').addClass('none');

    $('.c-import-block').find('[href="#'+type+'"]').trigger('click');
    window.setTimeout(function(){
        var intro = introJs();
        intro.setOptions({
            doneLabel:'我知道了',                 
            showBullets:false,
            showStepNumbers:false,
            steps: [{
            element: $('.bootstrap-switch', '#'+type).get(0),
            intro: '点击按钮，可以查看已分配的主机'
            }]
        });

        intro.onexit(function() {
            cookie.set('play_user_guide_quickimport', 0);//设置退出新手引导标记
            cookie.set('quickimport_user_guide_step', 6);//结束引导
        });

        intro.oncomplete(function(){
            cookie.set('play_user_guide_quickimport', 0);//设置退出新手引导标记
            cookie.set('quickimport_user_guide_step', 6);//结束引导
        });

        intro.start();
    },500);
}