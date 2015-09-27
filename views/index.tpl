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
    <div class="container-fluid">
        <header>
            <div class="content">
                {{ .LayoutContent }}
            </div>
        </header>
        <footer>

            {{if .flash.success }}
            <div class="alert alert-success alert-dismissible navbar-fixed-bottom" role="alert">
                {{.flash.success}}
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            {{ else if .flash.notice }}
            <div class="alert alert-info alert-dismissible navbar-fixed-bottom" role="alert">
                {{.flash.notice}}
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            {{ else if .flash.error }}
            <div class="alert alert-danger alert-dismissible navbar-fixed-bottom" role="alert">
                {{.flash.error}}
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            {{ else if .flash.warning }}
            <div class="alert alert-warning alert-dismissible navbar-fixed-bottom" role="alert">
                {{.flash.warning}}
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            {{ end }}
            {{ template "footer.tpl" . }}
        </footer>
        <div class="backdrop"></div>
    </div>
    <script src="/static/js/jquery-1.11.3.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
</body>
</html>
