<!DOCTYPE html>
<html>
  <head>
    <title>My First App</title>
    <link rel="stylesheet" href="/static/css/bootstrap.min.css">
    <link rel="stylesheet" href="/static/application.css">
    <script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
    <script src="/static/application.js"></script>
  </head>
  <body style="padding-top: 50px;">

    <div class="navbar navbar-fixed-top">
      <div class="navbar-inner">
        <div class="container">
            <ul class="nav">
              <li><a href="/new_player">Start Over</a></li>
            </ul>
        </div>
      </div>
    </div>

    <div class="container">
      {{ if .session.Error }}
        <div class="alert alert-error">{{ .session.Error }}</div>
      {{ end }}
      {{ if .session.Success }}
        <div class="alert alert-success">{{ .session.Success }}</div>
      {{ end }}

      {{ .LayoutContent }}

    </div>
  </body>
</html>
