(function() {
  setCounter = function(n) {
    document.getElementById('counter').innerHTML = n;
  }

  addMessage = function(msg) {
    var newLi = document.createElement("li");
    newLi.innerHTML = msg;
    document.getElementById('msglist').appendChild(newLi);
  }

  var evSrc = new EventSource("/events");

  var messages = 0;

  evSrc.addEventListener("ping", function(e) {
    messages += 1;
    setCounter(messages);
    addMessage(e.data);
  });
})();
