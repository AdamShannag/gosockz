class Event {
    constructor(type, payload) {
        this.type = type;
        this.payload = payload;
    }
}

class BroadcastMessageEvent {
    constructor(message, from) {
        this.message = message;
        this.from = from;
    }
}

class SendMessageEvent {
    constructor(message, from, to) {
        this.message = message;
        this.from = from;
        this.to = to;
    }
}

class NewMessageEvent {
    constructor(message, from, sent) {
        this.message = message;
        this.from = from;
        this.sent = sent;
    }
}

class ChangeChatRoomEvent {
    constructor(name) {
        this.name = name;
    }
}

class ChangeUserChatEvent {
    constructor(name) {
        this.name = name;
    }
}

