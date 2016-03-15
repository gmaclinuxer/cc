$(document).ready(function() {
    var Level = 2;
    var AppType = 1;

    function format(state) {
        return state.id + "<span class='select-info'>(" + state.text2 + ")</span>";
    }
    function format2(state) {
        return state.id ;
    }
    $('.btn-primary').click(function () {
        if ($(this).hasClass("active")) {
            $(this).removeClass("active");
            return false;
        }
    })

    var UinArr = new Array();
    UinArr.push(Uin);
    $("#Maintainers").select2({
        placeholder: "选择运维人员",
        allowClear: true,
        data: UserList,
        multiple: true,

        allowClear: true,
        //width:'90%',
        formatResult: format,
        formatSelection: format
    }).select2('val', UinArr);


    $("#ProducterList").select2({
        placeholder: "选择产品人员",
        allowClear: true,
        data: UserList,
        multiple: true,
        allowClear: true,
        formatResult: format,
        formatSelection: format
    }).select2('val', UinArr);

    $(".c-panel-content div").click(function(){
        $(this).addClass('buttonMark');
        $(this).siblings().removeClass('buttonMark');
        $(this).parent().parent().removeClass('panelMark');
        $(this).parent().siblings().addClass('headerMark');
        $(this).parent().siblings().find('i').addClass('step-imgMark');

        if($(this).attr('name')=="c-button-game"){
            $("#stepMark").text('step 3：请填写详细的业务信息');
            $(".c-panel-pageTwo").fadeIn();
            $(".c-panel-pageThree").css("display","none");
            $(".c-panel-pageTwo").addClass('panelMark');
            $(".c-panel-pageTwo .c-panel-header").removeClass('headerMark');
            $(".c-panel-pageTwo .c-panel-header i").removeClass('step-imgMark');
            $(".c-panel-pageTwo .c-panel-content div").removeClass('buttonMark');
            AppType = 1;
        }else if($(this).attr('name')=="c-button-notgame"){
            $("#stepMark").text('step 2：请填写详细的业务信息');
            $(".c-panel-pageTwo").css("display","none");
            $(".c-panel-pageThree").fadeIn();
            $("#AppName").val('');
            $("#AppName").focus();
            Level = 2;
            AppType = 0;
        }else if($(this).attr('name')=="c-button-fqff"){
            $(".c-panel-pageThree").fadeIn();
            $("#AppName").val('');
            $("#AppName").focus();
            Level = 3;
        }else if($(this).attr('name')=="c-button-qqqf"){
            $(".c-panel-pageThree").fadeIn();
            $("#AppName").val('');
            $("#AppName").focus();
            Level = 2;
        }
    })

    $(".cancelb").click(function(){
        $(".c-panel-pageTwo,.c-panel-pageThree").css("display","none");
        $(".c-panel-pageOne").addClass('panelMark');
        $(".c-panel-pageOne .c-panel-header").removeClass('headerMark');
        $(".c-panel-pageOne .c-panel-content div").removeClass('buttonMark');

    })


//保存按钮
    $(".btn-primary").click(function () {
        var ApplicationName = $.trim($("#AppName").val());
        if ((ApplicationName.length > 32) || (ApplicationName.length == 0)) {
            var diaCopyMsg = dialog({
                quickClose: true,
                align: 'left',
                padding:'5px 5px 5px 10px',
                skin: 'c-Popuplayer-remind-left',
                content: '<span style="color:#fff">业务名称不合法</span>'
            });
            diaCopyMsg.show($("#AppName").get(0));
            return;
        }
        var Maintainers = $("#Maintainers").select2('val');
        Maintainers.sort();
        $.unique(Maintainers);
        if (Maintainers.length == 0) {
            var diaCopyMsg = dialog({
                quickClose: true,
                align: 'left',
                padding:'5px 5px 5px 10px',
                skin: 'c-Popuplayer-remind-left',
                content: '<span style="color:#fff">请配置运维人员</span>'
            });
            diaCopyMsg.show($("#s2id_Maintainers").get(0));
            return;
        }
        if (Maintainers.length > 24) {
            var diaCopyMsg = dialog({
                quickClose: true,
                align: 'left',
                padding:'5px 5px 5px 10px',
                skin: 'c-Popuplayer-remind-left',
                content: '<span style="color:#fff">请配置24个以下运维人员</span>'
            });
            diaCopyMsg.show($("#s2id_Maintainers").get(0));
            return;
        }
        var ProducterList = $("#s2id_ProducterList").select2('val');
        if (ProducterList.length == 0) {
            var diaCopyMsg = dialog({
                quickClose: true,
                align: 'left',
                padding:'5px 5px 5px 10px',
                skin: 'c-Popuplayer-remind-left',
                content: '<span style="color:#fff">请配置产品人员</span>'
            });
            diaCopyMsg.show($("#s2id_ProducterList").get(0));
            return;
        }
        if (ProducterList.length > 8) {
            var diaCopyMsg = dialog({
                quickClose: true,
                align: 'left',
                padding:'5px 5px 5px 10px',
                skin: 'c-Popuplayer-remind-left',
                content: '<span style="color:#fff">请配置8个以下产品人员</span>'
            });
            diaCopyMsg.show($("#ProducterList").get(0));
            return;
        }
        var LifeCycle = $("#LifeCycle .btn-group .active input").val();
        $.post("/app/add",
            {
                Level: Level,
                Type: AppType,
                ApplicationName: ApplicationName,
                Maintainers: Maintainers,
                ProducterList: ProducterList,
                LifeCycle: LifeCycle
            }
            , function (result) {
                // rere = $.parseJSON(result);
                rere = result;
                if (rere.success == false) {
                    showWindows(rere.errInfo, 'notice');
                }
                else {
                    cookie.set('play_user_guide_imqcloud_index', 0);
                    if (rere.gotopo == 1) {

                        var d = dialog({
                            title: '提示',
                            width: 300,
                            height: 30,
                            okValue: "好",
                            ok: function () {
                                window.location.href = '/host/quickImport';
                                return false;
                            },
                            onclose: function () {
                                window.location.href = '/app/index'
                            },
                            content: '<div class="c-dialogdiv"><i class="c-dialogimg-success"></i>' + '业务建好了，接下来就给它分配一些主机</a>吧' + '</div>'
                        });
                        d.showModal();
                        return;
                    }
                    else {
                        showWindows('新增成功！', 'success');
                        setTimeout(function () {
                            window.location.href = "/app/index";
                        }, 1000);
                    }
                }
            });
    })

    $('#AppName').blur(function (){
        showstep();
    });
    $('#Maintainers').blur(function (){
        showstep();
    });
    $('#s2id_ProducterList').blur(function (){
        showstep();
    });

    function showstep()
    {
        var val=$("#AppName").val();
        var Maintainers = $("#Maintainers").select2('val');
        var ProducterList = $("#s2id_ProducterList").select2('val');
        var header=$('.c-panel-pageThree').find('.c-panel-header');
        if (val && $.trim(val)!='' &&Maintainers.length!=0 &&  ProducterList.length!=0) {
            header.addClass('headerMark');
            header.find('i').addClass('step-imgMark');
        }else{
            header.removeClass('headerMark');
            header.removeClass('step-imgMark');
        }
    }

    function showWindows(Msg, level) {
        if (level == 'success') {

            var d = dialog({
                width: 150,
                content: '<div class="c-dialogdiv2"><i class="c-dialogimg-success"></i>'+Msg+'</div>'
            });
            d.show();
        }
        else if (level == 'error') {
            var d = dialog({
                title: '错误',
                width: 300,
                okValue: "确定",
                ok: function () {
                },
                content: '<div class="c-dialogdiv2"><i class="c-dialogimg-failure"></i>' + Msg + '</div>'
            });
            d.showModal();
        }
        else {
            var d = dialog({
                title: '警告',
                width: 300,
                okValue: "确定",
                ok: function () {
                },
                content: '<div class="c-dialogdiv2"><i class="c-dialogimg-prompt"></i>' + Msg + '</div>'
            });
            d.showModal();
        }
    }
})