<!DOCTYPE html>

<html>
<head>
    <title>Beego</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="/static/css/list.css" media="screen" title="no title" charset="utf-8">
    <link rel="stylesheet" href="/static/css/bootstrap.css" media="screen" title="no title" charset="utf-8">
    <link rel="stylesheet" href="/static/css/bootstrap-theme.css" media="screen" title="no title" charset="utf-8">
</head>

<body>
    <header>
        <div class="container">


            <h1>{{.flash.error}}</h1>
            <h1 class="text-center">{{ .title }}</h1>
            <div class="content">
                {{ .LayoutContent }}
            </div>
        </header>
        <footer>

        </footer>
        <div class="backdrop"></div>
    </div>
    <script src="/static/js/jquery-1.11.3.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
</body>
</html>
