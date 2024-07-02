
function generateRandomString(len) {
    return Array(len + 1)
        .join((Math.random().toString(36) + "00000000000000000").slice(2, 18))
        .slice(0, len);
}

function sendMessage(ev) {
    ev.preventDefault();
    const input = document.querySelector("#messageInput");
    const messageText = input.value;

    if (!messageText.trim()) {
        alert("Сообщение не может быть пустым!");
        return;
    }

    const message = {
        userId: "user " + generateRandomString(10),
        timestamp: Date.now(),
        text: messageText,
    };

    saveMessage(message);
    input.value = "";
    displayMessages(getAllMessages());
}

function saveMessage(message) {
    const currentMessages = getAllMessages();
    currentMessages.push(message);
    localStorage.setItem("chatMessages", JSON.stringify(currentMessages));

    console.log(localStorage.getItem("chatMessages"));
}

function getAllMessages() {
    const storedMessages = localStorage.getItem("chatMessages");
    return storedMessages ? JSON.parse(storedMessages) : [];
}

// добавить более красивую верстку сообщения
function displayMessages(messages) {
    const chatBox = document.querySelector("#chatMessages");
    chatBox.innerHTML = ""; // Очищаем предыдущие сообщения

    messages.forEach((message) => {
        const messageDiv = document.createElement("div");
        messageDiv.textContent = `${new Date(
            message.timestamp
        ).toLocaleString()} - ${message.userId}: ${message.text}`;
        chatBox.appendChild(messageDiv);
    });

    chatBox.scrollTop = chatBox.scrollHeight;
}

// https://github.com/codesandbox/codesandbox-client/issues/4683
// window.onload = function () {
setTimeout(() => {
    displayMessages(getAllMessages());
    document.querySelector("#sendButton").addEventListener("click", sendMessage);
}, 0);
// };
