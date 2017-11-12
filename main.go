package main

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
	coinApi "github.com/miguelmota/go-coinmarketcap"
	"os"
	"strconv"
	"net/http"
)

var mems = map[string]string{
	"tvar":       "Тварь",
	"rostaturka": "Ростатурка",
	"govno":      "Гамно",
}

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Hi there! I'm TelegramBot!"))
}

func main() {

	http.HandleFunc("/", MainHandler)
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	bot, err := tgbotapi.NewBotAPI("490802103:AAEHiF4pl-Vw7ONSV7SEVlK9qoyg2xTFrU4")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.CallbackQuery != nil {
			fmt.Println(update.CallbackQuery)
			mem := update.CallbackQuery.Data
			fmt.Println(mem)
			if mem == "govno" {
				msg := tgbotapi.NewVoiceUpload(update.CallbackQuery.Message.Chat.ID, "govno.ogg")
				bot.Send(msg)
			}

			if mem == "rostaturka" {
				msg := tgbotapi.NewPhotoUpload(update.CallbackQuery.Message.Chat.ID, "rostaturka.jpg")
				bot.Send(msg)
			}
			if mem == "tvar" {
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "тут будет аудиозапись тварь")
				bot.Send(msg)
			}
		}

		if update.Message == nil {
			continue
		}

		fmt.Print(update.Message.Text)
		fmt.Println(update.ChosenInlineResult)

		if update.InlineQuery != nil && update.Message != nil {
			fmt.Println("ttttttttttt")
		} else {

			switch update.Message.Text {
			case "/btc", "/bitcoin":
				coinInfo, err := coinApi.GetCoinData("bitcoin")
				if err != nil {
					log.Println(err)
				} else {

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Курс: "+strconv.FormatFloat(coinInfo.PriceUsd, 'f', -1, 64)+"$")
					bot.Send(msg)

				}

			case "/ltc", "/litecoin":
				coinInfo, err := coinApi.GetCoinData("litecoin")
				if err != nil {
					log.Println(err)
				} else {

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Курс: "+strconv.FormatFloat(coinInfo.PriceUsd, 'f', -1, 64)+"$")
					bot.Send(msg)

				}
			case "/nmc", "/namecoin":
				coinInfo, err := coinApi.GetCoinData("namecoin")
				if err != nil {
					log.Println(err)
				} else {

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Курс: "+strconv.FormatFloat(coinInfo.PriceUsd, 'f', -1, 64)+"$")
					bot.Send(msg)

				}
			case "/ppc", "/peercoin":
				coinInfo, err := coinApi.GetCoinData("peercoin")
				if err != nil {
					log.Println(err)
				} else {

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Курс: "+strconv.FormatFloat(coinInfo.PriceUsd, 'f', -1, 64)+"$")
					bot.Send(msg)

				}
			case "/nvc", "/novacoin":
				coinInfo, err := coinApi.GetCoinData("novacoin")
				if err != nil {
					log.Println(err)
				} else {

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Курс: "+strconv.FormatFloat(coinInfo.PriceUsd, 'f', -1, 64)+"$")
					bot.Send(msg)

				}
			case "/eth", "/ethereum":
				coinInfo, err := coinApi.GetCoinData("ethereum")
				if err != nil {
					log.Println(err)
				} else {

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Курс: "+strconv.FormatFloat(coinInfo.PriceUsd, 'f', -1, 64)+"$")
					bot.Send(msg)

				}
			case "/bch", "/bitcoin-cash":
				coinInfo, err := coinApi.GetCoinData("bitcoin-cash")
				if err != nil {
					log.Println(err)
				} else {

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Курс: "+strconv.FormatFloat(coinInfo.PriceUsd, 'f', -1, 64)+"$")
					bot.Send(msg)

				}
			case "/mems":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите Мем")

				keyboard := tgbotapi.InlineKeyboardMarkup{}
				for key, mem := range mems {
					var row []tgbotapi.InlineKeyboardButton
					btn := tgbotapi.NewInlineKeyboardButtonData(mem, key)
					row = append(row, btn)
					keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
				}

				msg.ReplyMarkup = keyboard
				bot.Send(msg)

			case "/test":

				msg := tgbotapi.NewChatAction(update.Message.Chat.ID, "test")
				bot.Send(msg)

			default:
				fmt.Println("Неизвестная команда.")

			}

		}

	}
}
