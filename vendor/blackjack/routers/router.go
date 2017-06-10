package routers

import (
	"blackjack/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/new_player", &controllers.MainController{}, "get,post:NewPlayer")
	beego.Router("/bet", &controllers.MainController{}, "get,post:Bet")
	beego.Router("/game", &controllers.MainController{}, "get:Game")
	beego.Router("/game/player/hit", &controllers.MainController{}, "post:PlayerHit")
	beego.Router("/game/player/stay", &controllers.MainController{}, "post:PlayerStay")
	beego.Router("/game/dealer", &controllers.MainController{}, "get:Dealer")
	beego.Router("/game/dealer/hit", &controllers.MainController{}, "post:DealerHit")
	beego.Router("/game/compare", &controllers.MainController{}, "get:Compare")
	beego.Router("/game/over", &controllers.MainController{}, "get:Over")
}