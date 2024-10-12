package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"golang.org/x/exp/rand"
)

var rootCmd = &cobra.Command{
	Use:   "item",
	Short: "Add items",
	Long:  "Add items for API reqvest",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			cmd.Help()
			return
		}

		quantity := args[0]

		q, err := strconv.Atoi(quantity)
		if err != nil {
			fmt.Println("Неверное колличество:", err)
			return
		}

		if q != 0 {
			err := createItems(q)
			if err != nil {
				fmt.Println("Ошибка при создании ITEMS", err)
			} else {
				getItems()
			}
		} else {
			fmt.Errorf("Нельза обавть 0 ITEMS")
		}

	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func createItems(quantity int) error {

	for i := 0; i < quantity; i++ {
		structura := generateRandStruct()
		err := Post(structura)
		if err != nil {
			fmt.Errorf("", err)
		}
	}
	return nil

}

func Post(structura reqvest) error {

	url := "http://127.0.0.1:8080/item"
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(structura)
	if err != nil {
		return nil
	}

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Content-Type", "application/json")

	req.Body = ioutil.NopCloser(&buf)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func generateRandStruct() reqvest {
	item := reqvest{
		Caption: "Item" + strconv.Itoa(rand.Intn(100)),
		Weight:  float32(rand.Intn(100)) / 10.0,
		Number:  rand.Intn(100),
	}
	return item
}

func getItems() {

	url := "http://127.0.0.1:8080/item"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	var items []reqvest

	err := decoder.Decode(&items)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Список предметов:")
	for i, item := range items {
		fmt.Printf("%d. Caption: %s, Weight: %.1f, Number: %d\n", i+1, item.Caption, item.Weight, item.Number)
	}
}

type reqvest struct {
	Caption string  `json:"caption"`
	Weight  float32 `json:"weight"`
	Number  int     `json:"number"`
}
