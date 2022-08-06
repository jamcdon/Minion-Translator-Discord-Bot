package main

import (
    "fmt"
    "strings"
    "database/sql"
    "os/signal"
    "syscall"
    "os"
    "flag"

    "github.com/bwmarrin/discordgo"
)


var (
    Token string
)


func init() {
    flag.StringVar(&Token, "t", "", "Bot Token")
    flag.Parse()
}



func translate(inputString string) string {
    const sqlFile = "translate.db"
    db, _ := sql.Open("sqlite3", sqlFile)
    defer db.Close()


    inputArray := strings.Fields(inputString)
    outputString := ""

    for i:= 1; i < len(inputArray); i++ {
        row := db.QueryRow(`SELECT minionSpeak FROM minionTranslate WHERE english LIKE '`+ inputArray[i] + `';`)
        var minionspeak string //Translation.Minionspeak
        row.Scan(&minionspeak)
        if minionspeak == "" {
            outputString += inputArray[i] + " "
        } else {
            outputString += minionspeak + " "
        }
    }
    outputString = strings.TrimSuffix(outputString, " ")
    return outputString
}

//    fmt.Println(translate("i want a banana"))
func main() {
    // Create a new Discord session using the provided bot token.
    dg, err := discordgo.New("Bot ", Token)
    if err != nil {
        fmt.Println("error creating Discord session,", err)
        return
    }

    // Register the messageCreate func as a callback for MessageCreate events.
    dg.AddHandler(messageCreate)

    // receiving message events
    dg.Identify.Intents = discordgo.IntentsGuildMessages

    //open websocket for discord conn
    err = dg.Open()
    if err != nil {
        fmt.Println("error opening connection,", err)
        return
    }

    // wait here until CTRL-C or term signal received
    fmt.Println("Bot is now running! Bello!")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    //ignores messages created by self
    if m.Author.ID == s.State.User.ID {
        return
    }
    if strings.HasPrefix(m.Content, "!mt"){
        translated := translate(m.Content)
        s.ChannelMessageSend(m.ChannelID, translated)
    }
}