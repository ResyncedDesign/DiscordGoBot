<a href="https://docs.resynced.design/" align="center">
    <img src="https://r2.resynced.design/cdn/01JFT00BNQ2R8K4DSVNZKY0R4H.png" align="center" />
</a>

<h1 align="center">🎉 Discord Go Bot Template</h1>

<p align="center">Consider giving this a ⭐ to show your support! <3</p>

---

Welcome to the **Open-Source Discord Go Bot Template**! This project is built using **Go** and comes with everything you need to get started on building your own Go powered discord bot. For a more detailed breakdown look at our **[Documentation.](https://docs.resynced.design/)**

---

## 🌟 Features

-   **🔒 Secure Configurations** – Environment-based `.env` file setup
-   **📦 Modular Design** – Easily extend and customize features


---

## 🚀 Getting Started

1. **Clone the Repository**:

    ```bash
    gh https://github.com/ResyncedDesign/DiscordGoBot.git <folder-name>
    cd <folder-name>
    ```

2. **Install Dependencies**:

    Run one of the following commands to install all required packages:

    ```bash
    go mod tidy
    ```

3. **Set Up Environment Variables**:

    Create a `.env` file in the root directory with the following structure:

    ```plaintext
    TOKEN=your_discord_bot_token
    CLIENTID=your_discord_application_client_id
    GUILDID=your_discord_guild_id
    ```

4. **Build and Start the Bot**:

    Compile the bot and start it using:

    ```bash
    go build -o bot main.go
    ./bot
    ```

---

## ✨ Inviting the Bot to Your Server

To invite the bot to your Discord server, follow these steps:

> [!NOTE]  
> Ensure you select all the intents for the bot to work. 

1. Go to your Discord application's **OAuth2** section in the [Discord Developer Portal](https://discord.com/developers/applications).
2. Generate an invite link using the bot's **Client ID** with the required scopes and permissions. Example link:

    ```plaintext
    https://discord.com/oauth2/authorize?client_id=YOUR_CLIENT_ID&permissions=8&integration_type=0&scope=applications.commands+bot
    ```

3. Replace `YOUR_CLIENT_ID` with the `CLIENT_ID` from your `.env` file.
4. Open the link in a browser and add the bot to your desired server.

---

## 🛠️ Customizing the Bot

To start customizing, edit the files in the `src` directory. Key files include:

-   **`main.go`** – Main entry point for the bot.
-   **`src/commands/categories`** – Directory containing bot commands (add new commands here).
-   **`src/events/categories`** – Directory containing bot events (add new events here).
-   **`src/types`** – Directory containing custom types and interfaces.
-   **`src/config`** – Directory containing configuration files and environment variables.

Changes will take effect after rebuilding the project using:

```bash
go build -o bot main.go
./bot
```

OR you can use the following command to run the bot:

```bash
go run main.go
```

---

## 🤝 Contributing

Contributions are welcome! Feel free to fork the repository and submit pull requests for improvements or new features.

---

## 💡 Suggestions or Feedback?

If you have ideas for improving this bot or encounter any issues, feel free to open an issue or reach out. Happy coding! 🎉
