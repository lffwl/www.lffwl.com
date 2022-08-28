var obj = function (config) {

    //配置
    this.config = {
        //token - key
        token: 'lffwl-bolg', //AJAX默认参数
        ajax: {
            //请求域名
            domain: '/manage',
            isLoad: true,
            isErrorMsg: true,
            loadTitle: '请求中...',
            type: 'POST',
            data: {},
            async: true,
            timeout: 60000,
            errorMsg: '网络异常，请重试',
        }, //加载页面
        loadPage: {
            selector: 'body'
        }, //弹出页面参数
        openParam: {
            area: ['40%', '60%'], zIndex: 0,
        }, //表格默认参数
        table: {
            page: true, height: 'full-230', limits: [10, 20, 30, 40], response: {
                statusCode: 0 //规定成功的状态码，默认：0
            }
        },
    };

    //融合参数
    this.config = $.extend(this.config, config);

    //用户信息
    this.adminInfo = {};

    //多选菜单
    this.tabMenu = {
        //打开的菜单页面lay-id列表
        openMenuTabMap: [],
    };

    //当前页面
    this.page = undefined;

    //打开的页面
    this.openPage = [];

    //设置个人详情
    this.setAdminInfo = function (admin) {
        this.adminInfo = admin;
        //获取权限
        let permissions = [];
        layui.each(admin.auth, function (index, item) {
            permissions = permissions.concat(item.key);
        })
        this.adminInfo.permissions = permissions
    };

    //获取当前页面对象
    this.pageObj = function () {
        return this.openPage[this.page];
    };

    //登录
    this.login = function (param) {
        layui.sessionData(this.config.token, {
            key: 'token', value: param.token
        });
        /*let __this = this;
        //定时刷新token
        let refreshToken = setInterval(function () {
            __this.refreshToken();
            window.clearInterval(refreshToken);
        }, 1000 * (param.expiry_time - 1));*/
    };

    //刷新token
    this.refreshToken = function () {
        let __this = this;
        this.ajax({
            url: '/system/authorized/access-refresh',
            loadTitle: '请稍等，系统处理中...',
            type: 'GET',
            success: function (ajaxData) {
                if (ajaxData.code === 0) {
                    __this.login(ajaxData.data)
                } else {
                    __this.loginOut();
                    layer.msg("刷新身份验证失败", {icon: 5});
                }
            },
            error: function () {
                __this.loginOut();
            }
        });
    };

    //退出登录
    this.loginOut = function () {
        layui.sessionData(this.config.token, null);
        layer.closeAll();
        // this.loadPage({
        //     url: '/page/login/signin',
        // });
        window.location.reload();
    };

    //获取token
    this.getToken = function () {
        return layui.sessionData(this.config.token).token;
    };

    //检查是否登录
    this.checkLogin = function () {
        let token = this.getToken();
        return token === undefined || token === 'undefined' || token === null || token === 'null';
    };

    //AJAX请求
    this.ajax = function (param) {

        //融合配置参数
        param = $.extend({}, this.config.ajax, param);

        if (!param.hasOwnProperty('url')) {
            layer.msg("请求地址不存在", {icon: 5});
            return false;
        }

        let loadIndex, headers = {}, __this = this;

        //检查是否存在token
        if (!this.checkLogin()) {
            headers["Access-Token"] = this.getToken();
        }

        param.url = param.domain + param.url;

        return $.ajax({
            url: param.url,
            type: param.type,
            data: param.data,
            async: param.async,
            timeout: param.timeout,
            headers: headers,
            beforeSend: function (xhr) {
                if (param.isLoad) {
                    loadIndex = layer.msg(param.loadTitle, {
                        icon: 16, shade: 0.01, time: 0
                    });
                }
                //加载前执行方法
                if (typeof param.before == 'function') {
                    param.before(xhr);
                }
            },
            success: function (ajaxData, textStatus, jqXHR) {
                try {
                    ajaxData = JSON.parse(ajaxData);
                } catch (e) {
                    //console.log(ajaxData)
                }
                //身份证验证过期
                if (ajaxData.status == 'jwt') {
                    __this.loginOut();
                }
                //console.log(ajaxData);
                //加载成功执行方法
                if (typeof param.success == 'function') {
                    param.success(ajaxData);
                }
            },
            error: function (xhr, textStatus) {
                if (param.isErrorMsg) {
                    layer.msg(param.errorMsg, {icon: 5});
                }
                //加载错误执行方法
                if (typeof param.error == 'function') {
                    param.error(xhr, textStatus);
                }
            },
            complete: function () {
                if (param.isLoad) {
                    layer.close(loadIndex);
                }
                //加载完成执行的方法
                if (typeof param.complete == 'function') {
                    param.complete();

                }
            },
        });
    };

    //加载页面
    this.loadPage = function (param, selector, title, openParam) {
        selector = $.extend({}, this.config.loadPage, {selector: selector}).selector;
        let __this = this;
        param.success = function (ajaxData) {
            switch (selector) {
                case 'body':
                    $('#L-Body').html(ajaxData);
                    break;
                case 'menu':
                    ajaxData = __this.btnAction(ajaxData);
                    let layId = $.md5(title);
                    layui.element.tabAdd('open-menu-list', {
                        title: title, content: ajaxData, id: layId
                    });
                    layui.element.tabChange('open-menu-list', layId);
                    break;
                case 'open':
                    openParam = $.extend({}, __this.config.openParam, openParam);
                    layer.open({
                        type: 1, title: title || '弹出页面', skin: 'layui-layer-rim', //加上边框
                        area: openParam.area, zIndex: openParam.zIndex, content: ajaxData
                    });
                    break;
                default:
                    $(selector).html(ajaxData);
            }
        };
        if (param.hasOwnProperty('url')) {
            param.url = '/admin' + param.url + '.html?' + Date.parse(new Date());
            param.isLoad = false;
            param.loadTitle = '页面加载中...';
            param.domain = '';
            param.type = 'GET';
        }
        this.ajax(param);
    };

    //加载tab - 页面
    this.loadTabPage = function (param, selector, title) {
        let layId = $.md5(title);
        if (this.tabMenu.openMenuTabMap.includes(layId) === false) {
            this.tabMenu.openMenuTabMap = this.tabMenu.openMenuTabMap.concat(layId);
            this.page = layId;
            this.loadPage(param, selector, title)
        }
        layui.element.tabChange('open-menu-list', layId);
    };

    //删除tab页面
    this.delTabPage = function (layId) {
        delete this.openPage[layId];
        this.tabMenu.openMenuTabMap.splice(this.tabMenu.openMenuTabMap.indexOf(layId), 1);
    }

    //验证按钮权限
    this.btnAction = function (html) {
        let __this = this, delStr = [];
        layui.each($(html), function (index, item) {
            let btnAction = $(item).find("*[data-action]");
            layui.each(btnAction, function (bIndex, bItem) {
                if (__this.adminInfo.permissions.includes($(bItem).attr('data-action')) === false) {
                    delStr = delStr.concat($(bItem).context.outerHTML);
                }
            });
        });
        //压缩代码
        var packer = new Packer;
        html = packer.pack(html, 0, 0);
        layui.each(delStr, function (index, item) {
            item = packer.pack(item, 0, 0);
            html = html.replaceAll(item, "");
        });
        return html;
    }

    //删除行
    this.deleteRow = function (param) {
        let __this = this;
        layer.confirm('确定删除该行？', {
            skin: 'layui-layer-lan', title: '删除操作', btn: ['确定', '取消']
        }, function () {
            param = $.extend({}, param, {
                type: 'Post', loadTitle: '删除中...'
            });
            __this.ajax(param);
        });
    };

    //加载表格
    this.loadTable = function (param, value) {
        let __this = this;

        if (!param.hasOwnProperty('url')) {
            layer.msg("请求地址不存在", {icon: 5});
            return false;
        }

        // url
        param.url = __this.config.ajax.domain + param.url;

        //table参数
        param = $.extend({}, __this.config.table, param);


        //带上token
        let headers = {}
        if (!this.checkLogin()) {
            headers["Access-Token"] = this.getToken();
        }

        //请求head
        param.headers = headers

        //预处理数据
        param.parseData = function (res) {
            let count = 0, data = {};
            if (res.code === 0) {
                count = res.data.total;
                data = res.data.list;

                //config
                if (value) {
                    value.data = res.data;
                    if (res.data.hasOwnProperty('config')) {
                        value.config = res.data.config;
                    }
                    if (res.data.hasOwnProperty('checked')) {
                        value.checked = res.data.checked;
                    }
                } else {
                    __this.pageObj().data = res.data;
                    if (res.data.hasOwnProperty('config')) {
                        __this.pageObj().config = res.config;
                    }
                    if (res.data.hasOwnProperty('checked')) {
                        __this.pageObj().checked = res.checked;
                    }
                }

                //如果需要处理数据
                if (typeof param.handleData == 'function') {
                    data = param.handleData(data);
                }
            }

            return {
                "code": res.code, //解析接口状态
                "msg": res.message, //解析提示文本
                "count": count, //解析数据长度
                "data": data//解析数据列表
            };
        }

        // 渲染后回调
        param.done = function (res) {
            //成功
            if (res.code === 0) {

                //选中处理
                if (value) {
                    if (value.hasOwnProperty("checked")) {
                        layuiTableCheckboxDefault(value.checked, $(".layui-table-view[lay-id=" + param.id + "]").find('.layui-table-main'), 1);
                    }
                } else if (__this.pageObj().hasOwnProperty('checked')) {
                    layuiTableCheckboxDefault(res.data.checked, $(".layui-table-view[lay-id=" + param.id + "]").find('.layui-table-main'), 1);
                }

                if (typeof param.loadOver == 'function') {
                    param.loadOver(res);
                }
            }

        };

        //第一个实例
        layui.table.render(param);

    }

    //创建编辑器
    this.createEditor = function (selector) {
        let __this = this;

        // 富文本
        let E = window.wangEditor
        let editor = new E(selector)

        editor.config.zIndex = 20

        // 挂载highlight插件
        editor.highlight = hljs

        // 配置 server 接口地址
        editor.config.uploadImgServer = '/manage/file/upload'

        //fileName
        editor.config.uploadFileName = 'image'

        //设置token
        editor.config.uploadImgHeaders = {
            Authorization: 'Bearer ' + __this.getToken(),
        };

        editor.create()

        return editor

    }

};

