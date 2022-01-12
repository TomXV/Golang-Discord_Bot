package main
import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
  "strings"

  "github.com/bwmarrin/discordgo"
)

func main() {
  var TOKEN string
  fmt.Println("Token?")
  fmt.Scan(&TOKEN)
  dg, err := discordgo.New("Bot " + TOKEN)
  if err != nil {
    fmt.Println("error:start\n", err)
    return
  }

  //on message
  dg.AddHandler(messageCreate)

  err = dg.Open()
  if err != nil {
    fmt.Println("error:wss\n", err)
    return
  }
  fmt.Println("BOT Running...")

  //シグナル受け取り可にしてチャネル受け取りを待つ（受け取ったら終了）
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-sc
  dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Author.Bot {
    return
  }
  nick := m.Author.Username
  member, err := s.State.Member(m.GuildID, m.Author.ID)
  if err == nil && member.Nick != "" {
    nick = member.Nick
  }
  fmt.Println("< "+m.Content+" by "+nick)

  if m.Content == "ping" {
    s.ChannelMessageSend(m.ChannelID, "pong")
    fmt.Println("> pong")
  }
  if strings.Contains(m.Content,"www") {
    s.ChannelMessageSend(m.ChannelID, "lol")
    fmt.Println("> lol")
  }
}