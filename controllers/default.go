package controllers

import (
	"blackjack/lib/blackjack"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
	"strings"
)

type Session struct {
	PlayAgain           bool
	ShowHitStayButton   bool
	ShowDealerHitButton bool
	Winner              template.HTML
	Loser               template.HTML
	Error               template.HTML
	Success             template.HTML
	PlayerName          string
	Turn                string
	PlayerBet           int
	PlayerPot           int
	Deck                blackjack.Deck
	DealerCards         []blackjack.Card
	PlayerCards         []blackjack.Card
}

type Game struct {
	Deck        blackjack.Deck
	DealerCards []blackjack.Card
	PlayerCards []blackjack.Card
}

func CalculateTotal(s *Session, cards []blackjack.Card) int {
	total := 0
	aces := 0
	for _, card := range cards {
		beego.Debug(card)
		if card.Rank == "Ace" {
			total += 11
			aces++
		} else {
			v, err := strconv.Atoi(card.Rank)
			if err != nil {
				total += 10
			} else {
				total += v
			}
		}
		beego.Debug(total)
	}

	for i := 0; i < aces; i++ {
		if total <= blackjack.BLACKJACK {
			break
		}
		total -= 10
	}

	return total
}
func CardImage(card blackjack.Card) string {

	return fmt.Sprintf("/static/img/cards/%s_%s.jpg", strings.ToLower(card.Suit), strings.ToLower(card.Rank))
}
func Winner(s *Session, msg string) template.HTML {
	s.PlayAgain = true
	s.ShowHitStayButton = false
	s.PlayerPot = s.PlayerPot + s.PlayerBet
	s.Winner = template.HTML(fmt.Sprintf("<strong>%s won!</strong> %s", s.PlayerName, msg))
	return s.Winner
}
func Loser(s *Session, msg string) template.HTML {
	s.PlayAgain = true
	s.ShowHitStayButton = false
	s.PlayerPot = s.PlayerPot - s.PlayerBet
	s.Loser = template.HTML(fmt.Sprintf("<strong>%s loses</strong> %s", s.PlayerName, msg))
	return s.Loser
}
func Tie(s *Session, msg string) template.HTML {
	s.PlayAgain = true
	s.ShowHitStayButton = false
	s.Winner = template.HTML(fmt.Sprintf("<strong>It's a tie!</strong> %s", msg))
	return s.Winner
}

func init() {
	beego.AddFuncMap("calculate_total", CalculateTotal)
	beego.AddFuncMap("card_image", CardImage)
	beego.AddFuncMap("Winner", Winner)
	beego.AddFuncMap("Loser", Loser)
	beego.AddFuncMap("Tie", Tie)
}

type MainController struct {
	beego.Controller
}

// Prepare runs after Init before request function execution.
func (c *MainController) Prepare() {
	c.Layout = "layout.tpl"
	s := c.GetSession("session")
	if s == nil {
		s = &Session{}
		c.SetSession("session", s)
	}
	c.Data["session"] = s
	s.(*Session).ShowHitStayButton = true
	s.(*Session).ShowDealerHitButton = false
	s.(*Session).Success = ""
	s.(*Session).Error = ""
	s.(*Session).Winner = ""
	s.(*Session).Loser = ""
}

func (c *MainController) Get() {
	s := c.GetSession("session").(*Session)
	if s.PlayerName == "" {
		c.Redirect("/new_player", 302)
		return
	}
	c.Data["session"] = s
	c.TplName = "game.tpl"
}

func (c *MainController) NewPlayer() {
	s := c.GetSession("session").(*Session)
	if c.Ctx.Input.Method() == "POST" {
		s.PlayerName = c.GetString("playerName")
		if s.PlayerName != "" {
			c.Redirect("/bet", 302)
			return
		}
		s.Error = "Name is required"
	}

	s.PlayerPot = blackjack.INITIAL_POT
	c.Data["session"] = s
	c.TplName = "new_player.tpl"
}

func (c *MainController) Bet() {
	s := c.GetSession("session").(*Session)
	if c.Ctx.Input.Method() == "POST" {
		betAmount, _ := strconv.Atoi(c.GetString("betAmount"))
		if betAmount == 0 {
			s.Error = "Must place a bet"
		} else if betAmount > s.PlayerPot {
			s.Error = template.HTML(fmt.Sprintf("Bet must be less that %d.", s.PlayerPot))
		} else {
			s.PlayerBet = betAmount
			c.Data["session"] = s
			c.Redirect("/game", 302)
			return
		}
	}

	s.PlayerBet = 0
	if s.PlayerPot <= 0 {
		c.Redirect("/game/over", 302)
		return
	}

	c.Data["session"] = s
	c.TplName = "bet.tpl"
}

func (c *MainController) Game() {
	s := c.GetSession("session").(*Session)

	s.Turn = s.PlayerName

	deck := blackjack.NewBlackJackDeck()
	deck.Shuffle()
	s.PlayerCards = nil
	s.DealerCards = nil
	s.PlayerCards = append(s.PlayerCards, deck.Deal())
	s.DealerCards = append(s.DealerCards, deck.Deal())
	s.PlayerCards = append(s.PlayerCards, deck.Deal())
	s.DealerCards = append(s.DealerCards, deck.Deal())
	s.Deck = deck
	c.Data["session"] = s
	c.TplName = "game.tpl"
}

func (c *MainController) PlayerHit() {
	s := c.GetSession("session").(*Session)

	s.PlayerCards = append(s.PlayerCards, s.Deck.Deal())
	total := CalculateTotal(s, s.PlayerCards)
	if total == blackjack.BLACKJACK {
		Winner(s, "You hit Blackjack!")
	} else if total > blackjack.BLACKJACK {
		Loser(s, "You busted!")
	}
	c.Data["session"] = s
	c.Layout = ""
	c.TplName = "game.tpl"
}

func (c *MainController) PlayerStay() {
	s := c.GetSession("session").(*Session)
	s.Success = template.HTML(fmt.Sprintf("%s has chosen to stay.", s.PlayerName))
	c.Data["session"] = s
	c.Redirect("/game/dealer", 302)
}

func (c *MainController) Dealer() {
	s := c.GetSession("session").(*Session)

	s.Turn = "dealer"
	total := CalculateTotal(s, s.DealerCards)
	if total == blackjack.BLACKJACK {
		Loser(s, "Dealer hit blackjack.")
	} else if total > blackjack.BLACKJACK {
		Winner(s, "Dealer busted!")
	} else if total >= 17 {
		c.Redirect("/game/compare", 302)
	} else {
		s.ShowDealerHitButton = true
	}

	c.Data["session"] = s
	c.Layout = ""
	c.TplName = "game.tpl"
}

func (c *MainController) DealerHit() {
	s := c.GetSession("session").(*Session)

	s.DealerCards = append(s.DealerCards, s.Deck.Deal())

	c.Data["session"] = s
	c.Redirect("/game/dealer", 302)
}

func (c *MainController) Compare() {
	s := c.GetSession("session").(*Session)

	s.ShowHitStayButton = false

	playerTotal := CalculateTotal(s, s.PlayerCards)
	dealerTotal := CalculateTotal(s, s.DealerCards)

	if playerTotal > dealerTotal {
		Winner(s, fmt.Sprintf("You stayed at %d and the dealer stayed at %d.", playerTotal, dealerTotal))
	} else if playerTotal < dealerTotal {
		Loser(s, fmt.Sprintf("You stayed at %d and the dealer stayed at %d.", playerTotal, dealerTotal))
	} else {
		Tie(s, fmt.Sprintf("You and the dealer stayed at %d.", playerTotal))
	}

	c.Data["session"] = s
	c.Layout = ""
	c.TplName = "game.tpl"
}

func (c *MainController) Over() {
	s := c.GetSession("session").(*Session)
	if s.PlayerPot <= 0 {
		s.Error = template.HTML(fmt.Sprintf("<strong>%s, You are busted!</strong>", s.PlayerName))
	}
	c.TplName = "game_over.tpl"
}
