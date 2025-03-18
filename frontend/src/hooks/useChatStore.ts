import { create } from "zustand";
import Message from "@/enteties/Message/Message.ts";
import React from "react";


export interface ChatState {
    input: string;
    messages: Message[];
}

interface Actions {
    setInput: (input: string) => void;
    handleInputChange: (
        e:
            | React.ChangeEvent<HTMLInputElement>
            | React.ChangeEvent<HTMLTextAreaElement>,
    ) => void;
    setMessages: (fn: (messages: Message[]) => Message[]) => void;
    addMessage: (message: Message) => void;
}

const useChatStore = create<ChatState & Actions>()((set) => ({
    input: "",

    setInput: (input) => set({ input }),
    handleInputChange: (
        e:
            | React.ChangeEvent<HTMLInputElement>
            | React.ChangeEvent<HTMLTextAreaElement>,
    ) => set({ input: e.target.value }),
    messages: [
        {
            id: "1",
            chatId: "1",
            senderId: "user_1",
            message: "Hello, how are you?",
            timestamp: new Date(),
        },
        {
            id: "2",
            chatId: "1",
            senderId: "user_2",
            message: "I'm doing well, thank you!",
            timestamp: new Date(),
        },
        {
            id: "3",
            chatId: "1",
            senderId: "user_1",
            message: "What's the latest news?",
            timestamp: new Date(),
        },
        {
            id: "4",
            chatId: "1",
            senderId: "user_2",
            message: "It's great! We're working on a new feature.",
            timestamp: new Date(),
        }
    ],

    setMessages: (fn) => set(({ messages }) => ({ messages: fn(messages) })),
    addMessage: (message) => set((state) => ({messages: [...state.messages, message]})),
}));

export default useChatStore;