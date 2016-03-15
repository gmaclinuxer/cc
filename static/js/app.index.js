$(document).ready(function() {

    function format(state) {
        return state.id+"<span class='select-info'>("+state.text+")</span>";
    }

    $("#grid").kendoGrid({
        scrollable: false,
//            toolbar: [{text:'新增业务',name:'newapp',url: "/App/newapp",id:'newapp'},],
        dataBound: function(e) {
            $("#grid").css("opacity","1")
        }
    });
    var grid=$('#grid').data('kendoGrid'),
        wrapper=grid.wrapper;//获取grid;
    wrapper.find(".k-grid-toolbar").on('click.kendoGrid','.k-grid-test',function(e){
        window.location.href="/app/newapp";
    });
    //投后公司业务，不能创建业务
    if(source=='tencent')
    {
        $("#newapp").hide();
    }

    $('.c-searchyw-delete').click(function(e){
        var gridBatDel = dialog({
            title:'确认',
            width:250,
            content: '是否删除选中业务',
            okValue:"确定",
            cancelValue:"取消",
            ok:function (){
                var tr = $(e.target).closest('tr');
                var ApplicationID = tr.find('.ApplicationID').val();
                $.post("/app/delete",
                    {ApplicationID:ApplicationID}
                    ,function(result) {
                        // rere = $.parseJSON(result);
                        rere = result;
                        if(rere.success == false)
                        {
                            showWindows(rere.errInfo,'notice');
                            return;
                        }
                        else
                        {
                            showWindows("删除成功！",'success');
                            setTimeout(function(){window.location.reload();},1000);
                            return;
                        }
                    });

            },
            cancel: function () {
            }
        });
        gridBatDel.showModal();
    })

    $('.c-searchyw-save').click(function(){
        var tr = $(this).closest('tr');
        var nameDom = tr.find('.business_name');
        var businessName = $.trim(nameDom.find('input').val());
        var operaerVal = tr.find('.operaer').select2('val');
        operaerVal.sort();
        $.unique(operaerVal);
        var ApplicationID = tr.find('.ApplicationID').val();
        var operaDom = tr.find('.operaer');
        var ApplicationName = $("#AppName").val();
        if((businessName.length> 32) || (businessName.length== 0)) {
            var diaCopyMsg = dialog({
                quickClose: true,
                align: 'left',
                padding:'5px 5px 5px 10px',
                skin: 'c-Popuplayer-remind-left',
                content: '<span style="color:#fff">业务名称不合法</span>'
            });
            diaCopyMsg.show(tr.find('.business_name').find('input').get(0));
            return ;
        }
        if(operaerVal.length==0) {
            var diaCopyMsg = dialog({
                quickClose: true,
                align: 'left',
                padding:'5px 5px 5px 10px',
                skin: 'c-Popuplayer-remind-left',
                content: '<span style="color:#fff">请配置运维人员</span>'
            });
            diaCopyMsg.show(tr.find('.select2-container-multi input').get(0));
            return ;
        }
        if(operaerVal.length >24) {
            var diaCopyMsg = dialog({
                quickClose: true,
                align: 'left',
                padding:'5px 5px 5px 10px',
                skin: 'c-Popuplayer-remind-left',
                content: '<span style="color:#fff">请配置24个以下运维人员</span>'
            });
            diaCopyMsg.show(tr.find('.select2-container-multi input').get(0));
            return ;
        }
        $.post("/app/edit",
            {ApplicationName:businessName,Maintainers:operaerVal.join(';'),ApplicationID:ApplicationID}
            ,function(result) {
                rere = $.parseJSON(result);
                if(rere.success == false)
                {
                    showWindows(rere.errInfo,'notice');
                    return;
                }
                else
                {
                    showWindows('修改成功','success');
                    setTimeout(function(){window.location.reload();},1000);
                    return;
                }
            });
    })

    $('.c-searchyw-edit').click(function(){
        var tr = $(this).closest('tr');
        $(this).hide();
        $(this).siblings().filter('a[name="deletes"]').hide();
        $(this).siblings().filter('a[name="saves"]').show();
        $(this).siblings().filter('a[name="cancels"]').show();
        var business_value=tr.find('.business_name').text(); //业务名称值
        var operaer_value=tr.find('.Maintainers').val();  //运维人员值
        var arr = (operaer_value).split(";");
        var selectHtml ='';
        //for(var i=0;i<arr.length;i++){
        //    selectHtml += '<option selected="selected" value="'+arr[i]+'">'+arr[i]+'</option>';
        //}
        //获取运维列表
        $.post("/app/getMainterners",function(result)
        {
            var reval=$.parseJSON(result);
            var selectHtml2 ='';
            var data = new Array();
            if(reval.success != false)
            {
                var uinlist =$.parseJSON(result).uinList;
                var UserNameList =$.parseJSON(result).UserNameList;
                for(var i=0;i<uinlist.length;i++){
                //    selectHtml2 += '<option  value="'+uinlist[i]+'">'+uinlist[i]+'</option>';
                    var ud = { id: uinlist[i], text: UserNameList[i] };
                    data.push(ud);
                }
            }
//            tr.find('.operaer').html(''+ '<select multiple="true" selected="selected" style="width:100%;">'+selectHtml+selectHtml2+'</select>');
            tr.find('.operaer ').select2({
                placeholder: "选择运维人员",
                allowClear: true,
                data: data,
                multiple: true,
                allowClear: true,
                formatResult: format,
                formatSelection: format
            }).select2('val', arr);
        });
        tr.find('.business_name').html('<input type="text" value="'+business_value+'" style="width:100%;height:36px;"  class="k-input k-textbox business_name_input">');
        tr.find(".business_name .k-input").attr('placeholder','请输入业务名');
        tr.find(".business_name .k-input").attr('maxlength','32');
    })

    $('.c-searchyw-cancel').click(function(){
        var tr = $(this).closest('tr');
        ApplicationName = tr.find('.ApplicationName').val();
        Maintainers = tr.find('.Maintainers').val();
        tr.find('.business_name').html(ApplicationName);
        tr.find('.operaer').html(Maintainers);
        $(this).hide();
        $(this).siblings().filter('a[name="edits"]').show();
        $(this).siblings().filter('a[name="deletes"]').show();
        $(this).siblings().filter('a[name="saves"]').hide();
    })

    function startIntro1(){
        var intro = introJs();
        intro.setOptions({
            exitOnOverlayClick:false,
            skipLabel:'我知道了',
            doneLabel:'我知道了',
            prevLabel:'上一步',
            nextLabel:'下一步',
            showBullets:false,
            showButtons:true,
            showStepNumbers:false,
            tooltipClass:"first_creat_msg module-msg",
            steps: [
                {
                    element: document.querySelector('#appname'),
                    intro: "我们将腾讯云中的项目叫做业务，一个开发商可以有多个业务"
                },
                {
                    element: document.querySelector('#maintainers'),
                    intro: "我们将腾讯云中的项目协作者叫做业务运维，在配置平台里，业务运维是业务的上帝"
                },
                {
                    element: document.querySelector('#newapp'),
                    intro: "如果您需要新建业务，点击这里，也许您会发现惊喜",
                    position: 'left'
                },
                {
                    element: document.querySelector('#hostdis'),
                    intro: "这里可以为已创建的业务分配主机",
                    position: 'right'
                },
                {
                    element: document.querySelector('#apptopo'),
                    intro: "快速把业务拓扑树绘制到配置平台，就在这里",
                    position: 'right'
                },
                {
                    element: document.querySelector('#hostmng'),
                    intro: "主机查询以及让主机在拓扑树上跳来跳去，就靠这里了",
                    position: 'right'
                }
            ]
        });
        intro.onexit(function() {
            cookie.set('play_user_guide_imqcloud_index', 0);     //退出新手指引
        });

        intro.oncomplete(function(){
            cookie.set('play_user_guide_imqcloud_index', 0);   //退出新手指引
        });
        intro.start();
    }
    if(newUser)
    {
        startIntro1();
    }


});

function showWindows(Msg,level)
{
    if(level=='success')
    {
        var d = dialog({
            width: 150,
            content: '<div class="c-dialogdiv2"><i class="c-dialogimg-success"></i>'+Msg+'</div>'
        });
        d.show();
    }
    else if(level =='error')
    {
        var d = dialog({
            title:'错误',
            width:300,
            okValue:"确定",
            ok:function(){},
            content: '<div class="c-dialogdiv2"><i class="c-dialogimg-failure"></i>'+Msg+'</div>'
        });
        d.showModal();
    }
    else
    {
        var d = dialog({
            title:'警告',
            width:300,
            okValue:"确定",
            ok:function(){},
            content: '<div class="c-dialogdiv2"><i class="c-dialogimg-prompt"></i>'+Msg+'</div>'
        });
        d.showModal();
    }
}