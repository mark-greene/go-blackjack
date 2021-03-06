<!DOCTYPE html>
<html>
  <head>
    <title>Blackjack</title>
    <link rel="stylesheet" href="/static/css/bootstrap.min.css">
    <link rel="stylesheet" href="/static/application.css?version=2">
    <script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
    <script src="/static/application.js?version=2"></script>
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
      {{ if .error }}
        <div class="alert alert-error">{{ .error }}</div>
      {{ end }}
      {{ if .success }}
        <div class="alert alert-success">{{ .success }}</div>
      {{ end }}

      {{ .LayoutContent }}

    </div>
  </body>
</html>
