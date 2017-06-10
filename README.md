# go-blackjack
Port of Ruby/Sinatra game to Go/BeeGo

Using my Ruby/Sinatra blackjack game to learn Golang.  Requires lots more coding and the go template language
takes some getting use to.

```
<% session[:dealer_cards].each_with_index do |card, i| %>
   <% if session[:turn] != 'dealer' && i == 0 %>
     <img src='/images/cards/cover.jpg'>
   <% else %>
     <%= card_image(card) %>
   <% end %>
 <% end %>
```
Vs.
```
{{ range $i, $card := .session.DealerCards }}
  {{ if and (ne $.session.Turn "dealer") (eq $i 0) }}
    <img src='/static/img/cards/cover.jpg'>
  {{ else }}
    <img src='{{ CardImage $card }}' class='card_image'>
  {{ end }}
{{ end }}
```

### To run blackjack
git clone https://github.com/mark-greene/go-blackjack
```
go get github.com/astaxie/beego
bee run
```
