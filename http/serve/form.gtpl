<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>验证表单</title>
    <style type="text/css">
        body {
            padding: 10px 60px;
        }
        dl,dt,dd {
            padding: 0;
            margin: 0;
        }
        dl {
            display: block;
            clear: both;
            overflow: auto;
            margin-bottom: 10px;
        }
        dt,dd {
            float: left;
            color: #333333;
            font-size: 14px;
        }
        dt {
            width: 88px;
        }
        dd input[type="text"] {
            width: 200px;
        }
        p input {
            width: 100px;
            margin-left: 100px;
        }
        i {
            font-size: 12px;
            color: #cccccc;
        }
    </style>
</head>
<body>
<form action="/form" method="post">
    <dl>
        <dt>用户名：</dt>
        <dd>
            <input type="text" name="username" />
            <i>必填字段</i>
        </dd>
    </dl>

    <dl>
        <dt>年龄：</dt>
        <dd>
            <input type="text" name="age" />
            <i>数字</i>
        </dd>
    </dl>

    <dl>
        <dt>真实姓名：</dt>
        <dd>
            <input type="text" name="hanname" />
            <i>中文</i>
        </dd>
    </dl>

    <dl>
        <dt>英文名：</dt>
        <dd>
            <input type="text" name="engname" />
            <i>英文</i>
        </dd>
    </dl>

    <dl>
        <dt>邮箱地址：</dt>
        <dd>
            <input type="text" name="email" />
            <i>电子邮件地址</i>
        </dd>
    </dl>

    <dl>
        <dt>手机号码：</dt>
        <dd>
            <input type="text" name="mobile" />
            <i>手机号码</i>
        </dd>
    </dl>

    <dl>
        <dt>水果：</dt>
        <dd>
            <select name="fruit">
                <option value="apple">apple</option>
                <option value="pear">pear</option>
                <option value="banane">banane</option>
            </select>
            <i>下拉菜单</i>
        </dd>
    </dl>

    <dl>
        <dt>性别：</dt>
        <dd>
            <label><input type="radio" name="gender" value="1">男</label>
            <label><input type="radio" name="gender" value="2">女</label>
            <i>单选按钮</i>
        </dd>
    </dl>

    <dl>
        <dt>爱好：</dt>
        <dd>
            <label><input type="checkbox" name="interest" value="football">足球</label>
            <label><input type="checkbox" name="interest" value="basketball">篮球</label>
            <label><input type="checkbox" name="interest" value="tennis">网球</label>
            <i>复选框</i>
        </dd>
    </dl>

    <dl>
        <dt>身份证号：</dt>
        <dd>
            <input type="text" name="usercard" />
            <i>身份证号码</i>
        </dd>
    </dl>
    <p><input type="submit" value="提交表单"></p>
</form>
</body>
</html>