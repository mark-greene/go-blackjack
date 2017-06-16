<div id='game'>
  {{ if .loser }}
    <div class="alert alert-error">{{ .loser }}</div>
  {{ end }}
  {{ if .winner }}
    <div class="alert alert-success">{{ .winner }}</div>
  {{ end }}

  <h1>Blackjack!</h1>

  <p>Welcome {{ .session.PlayerName }}.</p>

  {{ if .playAgain }}
    <p>
      <strong>Play again?</strong>
      <a href='/bet' class='btn btn-primary'>Yes</a>
      <a href='/game/over' class='btn'>No</a>
    </p>
  {{ end }}

  <div class="well" id="dealer_cards">
    <h4>Dealer's cards:
      {{ if .showDealerHitButton }}
        {{ CalculateTotal .session.DealerCards }}
      {{  end }}
    </h4>
      {{ range $i, $card := .session.DealerCards }}
        {{ if and (ne $.session.Turn "dealer") (eq $i 0) }}
          <img src='/static/img/cards/cover.jpg' class='card_image'>
        {{ else }}
          <img src='{{ CardImage $card }}' class='card_image'>
        {{ end }}
      {{ end }}

      {{ if .showDealerHitButton }}
      <p>
        <h5>Dealer has {{ CalculateTotal .session.DealerCards }} and will hit.</h5>
        <form id="dealer_hit" action="/game/dealer/hit" method='post' >
          <input type="submit" class="btn btn-primary" value="Click to see dealer card &rarr;" />
        </form>
      </p>
      {{ end }}
  </div>

  <br/>
  <div class="well" id="player_cards">
    <h4>Player's cards: {{ CalculateTotal .session.PlayerCards }}</h4>
      {{  range $i, $card := .session.PlayerCards }}
        <img src='{{ CardImage $card }}' class='card_image'>
      {{  end }}

      <h5>
        {{ .session.PlayerName }} has ${{ .session.PlayerPot }} and bet ${{ .session.PlayerBet }}.
      </h5>
  </div>

  <p>
    What would {{ .session.PlayerName }} like to do?

    {{ if .showHitStayButton }}
      <form id="hit_form" action="/game/player/hit" method='post' >
        <input type="submit" class="btn btn-success" value="Hit" />
      </form>
      <form id="stay_form" action="/game/player/stay" method='post' >
        <input type="submit" class="btn btn-warning" value="Stay" />
      </form>
      {{ if AllowDouble .session }}
        <form id="double_form" action="/game/player/double" method='post' >
          <input type="submit" class="btn btn-danger" value="Double" />
        </form>
      {{ end }}
    {{ end }}

  </p>
</div>
