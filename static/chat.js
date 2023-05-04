var chatBox = document.getElementById("chat");
const user = "user";
const assistent = "assistent";

function receiveChatMessage(msg, role) {
  var newP = document.createElement("p");
  newP.innerHTML = msg;
  newP.className = role
  chatBox.appendChild(newP);
  chatBox.scrollTop = chatBox.scrollHeight;
}

document.getElementById("send").onclick = function() {
  var msg = document.getElementById("message").value; 
  ws.send(msg);

  var newP = document.createElement("p");
  newP.innerHTML = msg;
  newP.className = user;
  chatBox.appendChild(newP);  

  document.getElementById("message").value = "";
  chatBox.scrollTop = chatBox.scrollHeight;
}
