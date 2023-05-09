var value = "Loading...";

function getValue() {
    axios.get('/send_value')
    .then(function (response) {
        console.log(response); 
        value = response.data;
        console.log("get value"+response); // Hello from Golang!
    })
    .catch(function (error) {
        // handle error
        console.log(error);
    })    
    .finally(function () {
        // always executed
        console.log('I always Execued');
        if (typeof value != "undefined") {
            document.querySelector("#valueDisplay").innerText = value;   
        }
    });   
}


let source;

function StreamMessage() {
    let button = document.getElementById('start-stop');
    if (!source) {
        source = new EventSource('http://localhost:8001/stream_value');
        source.onmessage = e => {
          console.log(e.data);
        };
        button.textContent = 'Stop';
      } else {
        source.close();
        source = null;
        button.textContent = 'Start';
    }
}

const ws = new WebSocket("ws://localhost:8001/ws");
ws.onmessage = e => {
    console.log(e.data);

    var wsJson = JSON.parse(e.data);
    if (wsJson.messageType == 1) {
        if (wsJson.data.motion != "") {
            setMotion(wsJson.data.motion, 2)
        }

        if (wsJson.data.expression != "") {
            setExpression(wsJson.data.expression, 2)
        }
    } else if (wsJson.messageType == 2) {
        console.log("chat: " + wsJson.data)
        if (wsJson.data != "") {
            receiveChatMessage(wsJson.data.content, wsJson.data.role)
        }
    } else if (wsJson.messageType == 3) {
        var audio = new Audio(wsJson.data);
        audio.play();
        audio.addEventListener("ended", function() {
            console.log("Audio ended!");
            setExpression("f01", 1)
          });
    }
}

const setKeyCheckBox = document.getElementById("setKeyShow");
setKeyCheckBox.addEventListener("change", () => {
    let controlBox = document.getElementById("setKey");
    controlBox.hidden = !setKeyCheckBox.checked;
});

function sendKeySet(e) {
    e.preventDefault();  

    let chatKey = document.getElementById("chatKey").value;
    let ttsKey = document.getElementById("ttsKey").value;
    let ttsRegion = document.getElementById("ttsRegion").value;

    fetch('/setKey', {
        method: 'POST', 
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ 
            chatKey, 
            ttsKey,
            ttsRegion 
        })
    })
    .then(res => {
        if (res.ok) {
            alert("Set key successed!")
        } else {
            alert("Error submitting data!");
        }
    });
}

const setTTSVoiceShowBox = document.getElementById("setTTSVoiceShow");
setTTSVoiceShowBox.addEventListener("change", () => {
    let controlBox = document.getElementById("setTTSVoice");
    controlBox.hidden = !setTTSVoiceShowBox.checked;
});

function sendTTSSet(e) {
    e.preventDefault();

    const ttsLanguage = document.getElementById("ttsLanguage").value;
    const ttsGender = document.getElementById("ttsGender").value;
    const ttsVoiceName = document.getElementById("ttsVoiceName").value;

    fetch('/setTTS', {
        method: 'POST', 
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ 
            ttsLanguage, 
            ttsGender,
            ttsVoiceName 
        })
    })
    .then(res => {
        if (res.ok) {
            alert("Set TTS successed!")
        } else {
            alert("Error submitting data!");
        }
    });
}
