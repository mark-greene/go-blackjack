<h3>Make a bet for this round.</h3>

<p> {{ .session.PlayerName }} has ${{ .session.PlayerPot }}</p>

<form action='/bet' method='Post'>
  <input type='text' name='betAmount' />
  <br/>
  <input type='submit' value='Make a bet' class='btn' />
</form>
