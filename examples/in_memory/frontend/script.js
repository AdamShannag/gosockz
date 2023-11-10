
const loggedInDiv = document.getElementById("loggedin");
const loggedOutDiv = document.getElementById("loggedout");
loggedInDiv.hidden = true;

let selectedchat = "general";
let userChat = "none";
let username = "";

function routeEvent(event) {
    if (event.type === undefined) {
        alert("no 'type' field in event");
    }
    switch (event.type) {
        case "new_message":
            const messageEvent = Object.assign(new NewMessageEvent, event.payload);
            appendChatMessage(messageEvent);
            break;
        default:
            alert("unsupported message type");
            break;
    }

}

function appendChatMessage(messageEvent) {
    var date = new Date(messageEvent.sent);
    const formattedMsg = `${date.toLocaleString()} ${messageEvent.from}: ${messageEvent.message}`;

    textarea = document.getElementById("chatmessages");
    textarea.innerHTML = textarea.innerHTML + "\n" + formattedMsg;
    textarea.scrollTop = textarea.scrollHeight;
}

function changeChatRoom() {
    var newchat = document.getElementById("chatroom");
    if (newchat != null && newchat.value != selectedchat) {
        selectedchat = newchat.value;
        userChat = "";
        header = document.getElementById("chat-header").innerHTML = "Currently in chat: " + selectedchat;

        let changeEvent = new ChangeChatRoomEvent(selectedchat);
        sendEvent("change_session", changeEvent);
        textarea = document.getElementById("chatmessages");
        textarea.innerHTML = `You changed room into: ${selectedchat}`;
    }
    return false;
}

function selectUser() {
    var userToMessage = document.getElementById("userToMessage");
    if (userToMessage != null && userToMessage.value != userChat) {
        userChat = userToMessage.value;
        selectedchat = "";
        header = document.getElementById("chat-header").innerHTML = "Chatting with: " + userChat;

        textarea = document.getElementById("chatmessages");
        textarea.innerHTML = `You are chatting with: ${userChat}`;
    }
    return false;
}

function sendMessage() {
    var newmessage = document.getElementById("message");

    if (selectedchat !== "") {
        if (newmessage != null) {
            let outgoingEvent = new BroadcastMessageEvent(newmessage.value, username);
            sendEvent("broadcast_message", outgoingEvent)
            newmessage.value = ""
        }
        return false;
    }

    if (newmessage != null) {
        let outgoingEvent = new SendMessageEvent(newmessage.value, username, userChat);
        sendEvent("send_message", outgoingEvent)
        newmessage.value = ""
    }

    return false;
}

function sendEvent(eventName, payload) {
    const event = new Event(eventName, payload);
    conn.send(JSON.stringify(event));
}

function login() {
    const user = document.getElementById("username").value
    let formData = {
        "username": user,
        "password": document.getElementById("password").value
    }

    fetch("login", {
        method: 'post',
        body: JSON.stringify(formData),
        mode: 'cors',
    }).then((response) => {
        if (response.ok) {
            return response.json();
        } else {
            throw 'unauthorized';
        }
    }).then((data) => {
        username = user
        connectWebsocket(data.otp);
        loggedOutDiv.hidden = true;
        loggedInDiv.hidden = false;
    }).catch((e) => { alert(e) });
    return false;
}


function connectWebsocket(otp) {
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws?otp=" + otp + "&username=" + username);

        conn.onopen = function (evt) {
            document.getElementById("connection-header").innerHTML = "Connected to Websocket: true";
        }

        conn.onclose = function (evt) {
            document.getElementById("connection-header").innerHTML = "Connected to Websocket: false";
        }

        conn.onmessage = function (evt) {
            console.log(evt);
            const eventData = JSON.parse(evt.data);
            const event = Object.assign(new Event, eventData);
            routeEvent(event);
        }

    } else {
        alert("Not supporting websockets");
    }
}

window.onload = function () {
    document.getElementById("chatroom-selection").onsubmit = changeChatRoom;
    document.getElementById("user-selection").onsubmit = selectUser;
    document.getElementById("chatroom-message").onsubmit = sendMessage;
    document.getElementById("login-form").onsubmit = login;
};