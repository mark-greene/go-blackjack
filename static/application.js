$(document).ready(function() {

  player_hits();
  player_stays();
  player_doubles();
  dealer_hits();

});

function player_hits() {
  $(document).on("click", "form#hit_form input", function() {
    $.ajax({
      type: "Post",
      url: "/game/player/hit"
    }).done(function(msg) {
      $("#game").replaceWith(msg);
    });
    return false;
  });
};

function player_stays() {
  $(document).on("click", "form#stay_form input", function() {
    $.ajax({
      type: "Post",
      url: "/game/player/stay"
    }).done(function(msg) {
      $("#game").replaceWith(msg);
    });
    return false;
  });
};

function player_doubles() {
  $(document).on("click", "form#double_form input", function() {
    $.ajax({
      type: "Post",
      url: "/game/player/double"
    }).done(function(msg) {
      $("#game").replaceWith(msg);
    });
    return false;
  });
};

function dealer_hits() {
  $(document).on("click", "form#dealer_hit input", function() {
    $.ajax({
      type: "Post",
      url: "/game/dealer/hit"
    }).done(function(msg) {
      $("#game").replaceWith(msg);
    });
    return false;
  });
};
