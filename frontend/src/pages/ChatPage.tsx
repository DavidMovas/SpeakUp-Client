import React, { useEffect, useRef, useState } from "react";
import { ChatMessageList } from "@/components/ui/chat/chat-message-list.tsx";
import { ChatBubble, ChatBubbleAvatar, ChatBubbleMessage } from "@/components/ui/chat/chat-bubble.tsx";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { Button } from "@/components/ui/shadcn/button.tsx";
import { Paperclip, Send } from "lucide-react";
import { ChatInput } from "@/components/ui/chat/chat-input.tsx";
import useChatStore from "@/hooks/useChatStore.ts";
import Message from "@/enteties/Message/Message.ts";
import useProfileStore from "@/enteties/Profile/model/store/profileStore.ts";

const MeID = "user_1";

const ChatPage: React.FC = () => {
    const messagesRef = useRef<HTMLDivElement>(null);
    const formRef = useRef<HTMLFormElement>(null);

    const meID = useProfileStore((state) =>  state.data?.id);

    const {messages, addMessage}  = useChatStore();
    const [message, setMessage] = useState("");

    const [isLoading] = useState(false);
    const [input, setInput] = useState("");

    const handleInputChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setInput(event.target.value);
        setMessage(event.target.value);
    };

    useEffect(() => {
        if (messagesRef.current) {
            messagesRef.current.scrollTop = messagesRef.current.scrollHeight;
        }
    }, [messages]);

    const sendMessage = (newMessage: Message) => {
        addMessage(newMessage);
    };

    const handleSend = () => {
        if (meID) {
            const newMessage: Message = {
                id: message.length.toString(),
                chatId: "chat_1",
                senderId: meID,
                message: message.trim(),
                timestamp: new Date(),
            };

            sendMessage(newMessage);
        }

        setInput("");
        setMessage("");
    };

    //const formattedTime = new Date().toLocaleTimeString("en-US", {hour: "numeric", minute: "2-digit", hour12: true,});

    useEffect(() => {
        if (messagesRef.current) {
            messagesRef.current.scrollTop = messagesRef.current.scrollHeight;
        }
    }, [messages]);

    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        handleSend();
    };

    const onKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
        if (e.key === "Enter" && !e.shiftKey) {
            e.preventDefault();
            handleSend();
        }

        if (e.key === "Enter" && e.shiftKey) {
            e.preventDefault();
            setMessage((prev) => prev + "\n");
        }
    };

    return (
        <main className="flex h-screen w-full max-w-3xl flex-col items-center mx-auto">
            <div className="flex-1 w-full overflow-y-auto py-6">
                <ChatMessageList>
                    {/* Initial Message */}

                    {messages.length === 0 && (
                        <div className="w-full flex justify-center">
                            <p className="text-muted-foreground text-sm">
                                Empty
                            </p>
                        </div>
                    )}

                    {/* Messages */}
                    {messages?.map((message) => (
                            <ChatBubble
                                key={message.id}
                                variant={message.senderId == MeID ? "sent" : "received"}
                            >
                                <ChatBubbleAvatar
                                    src=""
                                    fallback={message.senderId[0]}
                                />
                                <ChatBubbleMessage>
                                    {message.message
                                        .split("```")
                                        .map((part: string, index: number) => {
                                            if (index % 2 === 0) {
                                                return (
                                                    <Markdown key={index} remarkPlugins={[remarkGfm]}>
                                                        {part}
                                                    </Markdown>
                                                );
                                            } else {
                                                return (
                                                    <pre className="whitespace-pre-wrap pt-2" key={index}>
                                                    </pre>
                                                );
                                            }
                                        })}

                                </ChatBubbleMessage>
                            </ChatBubble>
                        ))}
                </ChatMessageList>
            </div>

            {/* Form and Footer fixed at the bottom */}
            <div className="w-full px-4 pb-4">
                <form
                    ref={formRef}
                    onSubmit={onSubmit}
                    className="flex items-center gap-3"
                >
                    <Button variant="ghost" size="icon">
                        <Paperclip className="size=4" />
                    </Button>

                    <ChatInput
                        value={input}
                        onKeyDown={onKeyDown}
                        onChange={handleInputChange}
                        placeholder="Type your message here..."
                        className="rounded-lg bg-background border-spacing-1.5 shadow-none focus-visible:ring-0"
                    />

                    <Button
                        disabled={!input || isLoading}
                        type="submit"
                        size="sm"
                        className="ml-auto"
                    >
                        <Send className="size=4" />
                    </Button>
                </form>
            </div>
        </main>
    )
}

export default ChatPage;