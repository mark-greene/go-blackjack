<h3>Thanks for playing {{ .session.PlayerName }}</h3>

{{ if gt .session.PlayerPot 0 }}
  <h5>Looks like you're leaving with ${{ .session.PlayerPot }}
{{ else }}
  <h5>Maybe it's time to mortage your house and
    <a href='/new_player'>Start over.</a>
  </h5>
{{ end }}
