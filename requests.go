package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"google.golang.org/api/sheets/v4"
)

func MakeRequest(speedTest SpeedTest) {

	ctx := context.Background()

	c := getClient(ctx)

	sheetsService, err := sheets.New(c)
	spreadsheetId := "1C7ZgtjrDqrnJ0jYD7iZTdf0i5dpdgmg_fwVvMhpTQx0"

	req := sheets.AppendCellsRequest{
		SheetId: 1,
		Rows:    [4]string{speedTest.Time, speedTest.Up, speedTest.Down, speedtest.Server.Name},
		Fields:  "*"} // List of updates to apply

	dataRange := "A:D"
	valueInputOption := "USER_ENTERED"
	insertDataOption := "INSERT_ROWS"

	vr := sheetsService.Spreadsheets.ValueRange
	vr.Range = "A:D"
	vr.MajorDimension = "ROWS"

	up := strconv.FormatFloat(speedTest.Up, 'f', 0, 64)
	down := strconv.FormatFloat(speedTest.Down, 'f', 0, 64)

	vr.values = [1][4]string{{speedTest.Time, up, down, speedTest.Server.Name}}

	rb := &sheets.ValueRange{}

	resp, err := sheetsService.Spreadsheets.Values.Append(spreadsheetId, dataRange, rb).ValueInputOption(valueInputOption).InsertDataOption(insertDataOption).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp)
}

func getClient(ctx context.Context) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	json.NewEncoder(f).Encode(token)
}
