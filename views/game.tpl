<div id='game'>
  {{ if .session.Loser }}
    <div class="alert alert-error">{{ .session.Loser }}</div>
  {{ end }}
  {{ if .session.Winner }}
    <div class="alert alert-success">{{ .session.Winner }}</div>
  {{ end }}

  <h1>Blackjack!</h1>

  <p>Welcome {{ .session.PlayerName }}.</p>

  {{ if .session.PlayAgain }}
    <p>
      <strong>Play again?</strong>
      <a href='/bet' class='btn btn-primary'>Yes</a>
      <a href='/game/over' class='btn'>No</a>
    </p>
  {{ end }}

  <div class="well" id="dealer_cards">
    <h4>Dealer's cards:</h4>
      {{ $turn := .session.Turn }}
      {{ range $i, $card := .session.DealerCards }}
        {{ if and (ne $turn "dealer") (eq $i 0) }}
          <img src='/static/img/cards/cover.jpg'>
        {{ else }}
          <img src='{{ card_image $card }}' class='card_image'>
        {{ end }}
      {{ end }}

      {{ if .session.ShowDealerHitButton }}
      <p>
        <h5>Dealer has {{ calculate_total .session .session.DealerCards }} and will hit.</h5>
        <form id="dealer_hit" action="/game/dealer/hit" method='post' >
          <input type="submit" class="btn btn-primary" value="Click to see dealer card &rarr;" />
        </form>
      </p>
      {{ end }}
  </div>

  <br/>
  <div class="well" id="player_cards">
    <h4>Player's cards:</h4>
      {{  range $i, $card := .session.PlayerCards }}
        <img src='{{ card_image $card }}' class='card_image'>
      {{  end }}

      <h5>
        {{ .session.PlayerName }} has ${{ .session.PlayerPot }} and bet ${{ .session.PlayerBet }}.
      </h5>
  </div>

  <p>
    What would {{ .session.PlayerName }} like to do?
    {{ .session.PlayerName }} has {{ calculate_total .session .session.PlayerCards }}

    {{ if .session.ShowHitStayButton }}
      <form id="hit_form" action="/game/player/hit" method='post' >
        <input type="submit" class="btn btn-success" value="Hit" />
      </form>
      <form id="stay_form" action="/game/player/stay" method='post' >
        <input type="submit" class="btn btn-warning" value="Stay" />
      </form>
    {{ end }}

  </p>
</div>
