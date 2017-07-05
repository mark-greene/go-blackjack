# go-blackjack
Port of Ruby/Sinatra game to Go/BeeGo

Using my [Ruby/Sinatra blackjack game](https://github.com/mark-greene/ruby-web-blackjack) to learn Golang.  Requires lots more coding and the go template language
takes some getting use to.  Just try to do the following in go (it took me ~80 lines)
```
  suits = ['Clubs', 'Diamonds', 'Hearts', 'Spades']
  ranks = ['2', '3', '4', '5', '6', '7', '8', '9', '10', 'Jack', 'Queen', 'King', 'Ace']
  session[:deck] = suits.product(ranks).shuffle!
```
And what's with the prefix notation?
```
<% session[:dealer_cards].each_with_index do |card, i| %>
   <% if session[:turn] != 'dealer' && i == 0 %>
     <img src='/images/cards/cover.jpg'>
   <% else %>
     <%= card_image(card) %>
   <% end %>
 <% end %>
```
vs.
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
```
go get github.com/mark-greene/go-blackjack
go get github.com/astaxie/beego
bee run
```
Defaults to localhost:8080.  You can change port by `PORT=5000 bee run`.  You can run standalone `./go-blackjack`.
