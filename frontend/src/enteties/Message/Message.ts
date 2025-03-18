interface Message {
    id: string;
    chatId: string;
    senderId: string;
    message: string;
    timestamp: Date;
}

export default Message;