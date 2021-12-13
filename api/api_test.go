package api

import (
	"assignment/testutils"
	"assignment/utils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetMostUsedWordsJson(t *testing.T) {
	type args struct {
		writer  *httptest.ResponseRecorder
		request *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantRes []Item
	}{
		{name: "should pass if total words are less than 10",
			args: struct {
				writer  *httptest.ResponseRecorder
				request *http.Request
			}{writer: httptest.NewRecorder(), request: testutils.GetPOSTRequest(Request{Text: "apple ball ball cat ball cat"})},
			wantErr: false,
			wantRes: []Item{{Value: "ball", Count: 3}, {Value: "cat", Count: 2}, {Value: "apple", Count: 1}},
		},

		{name: "should pass if total words are more than 10",
			args: struct {
				writer  *httptest.ResponseRecorder
				request *http.Request
			}{writer: httptest.NewRecorder(), request: testutils.GetPOSTRequest(Request{Text: "apple ball ball cat ball cat cat cat cat cat cat apple cat ball cat"})},
			wantErr: false,
			wantRes: []Item{{Value: "cat", Count: 9}, {Value: "ball", Count: 4}, {Value: "apple", Count: 2}},
		},

		{name: "should pass if total words are more than 10",
			args: struct {
				writer  *httptest.ResponseRecorder
				request *http.Request
			}{writer: httptest.NewRecorder(), request: testutils.GetPOSTRequest(Request{Text: "Apart from counting words and characters, our online editor can help you to improve word choice and writing style, and, optionally, help you to detect grammar mistakes and plagiarism. To check word count, simply place your cursor into the text box above and start typing. You'll see the number of characters and words increase or decrease as you type, delete, and edit them. You can also copy and paste text from another program over into the online editor above. The Auto-Save feature will make sure you won't lose any changes while editing, even if you leave the site and come back later. Tip: Bookmark this page now.Knowing the word count of a text can be important. For example, if an author has to write a minimum or maximum amount of words for an article, essay, report, story, book, paper, you name it. WordCounter will help to make sure its word count reaches a specific requirement or stays within a certain limit.In addition, WordCounter shows you the top 10 keywords and keyword density of the article you're writing. This allows you to know which keywords you use how often and at what percentages. This can prevent you from over-using certain words or word combinations and check for best distribution of keywords in your writing.\nIn the Details overview you can see the average speaking and reading time for your text, while Reading Level is an indicator of the education level a person would need in order to understand the words you’re using.Disclaimer: We strive to make our tools as accurate as possible but we cannot guarantee it will always be so."})},
			wantErr: false,
			wantRes: []Item{{Value: "and", Count: 12}, {Value: "the", Count: 11}, {Value: "you", Count: 11}, {Value: "to", Count: 7}, {Value: "of", Count: 6}, {Value: "words", Count: 5}, {Value: "can", Count: 5}, {Value: "a", Count: 5}, {Value: "word", Count: 5}, {Value: "or", Count: 4}},
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetMostUsedWordsJson(tt.args.writer, tt.args.request)

			var resp struct {
				utils.ApiResponse
				Data []Item `json:"data"`
			}
			err := json.NewDecoder(tt.args.writer.Result().Body).Decode(&resp)
			if err != nil {
				t.Fatal(err)
			}

			if !tt.wantErr && !resp.Status {
				t.Fatalf("wanted resp.Status: %v, got: %v. error: %v", !tt.wantErr, resp.Status, resp.Message)
			}

			if !reflect.DeepEqual(tt.wantRes, resp.Data) {
				t.Errorf("wanted resp.Data: %v, got: %v", tt.wantRes, resp.Data)
			}

		})
	}
}
