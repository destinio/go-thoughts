package api

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

type Thought struct {
	ID          uuid.UUID
	Content     string
	DateCreated time.Time
}

func AddThought(filepath string, thought string) error {
	id := uuid.New()
	fileExists := false

	if len(thought) < 1 {
		fmt.Println("🤤 Please provide a thought")
		return nil
	}

	if _, err := os.Stat(filepath); err == nil {
		fileExists = true
	}

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	newData := Thought{
		ID:          id,
		Content:     thought,
		DateCreated: time.Now(),
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	if !fileExists {
		writer.Write([]string{"ID", "Content", "DateCreated"})
	}

	thoughtData := []string{newData.ID.String(), newData.Content, newData.DateCreated.Format(time.RFC3339)}

	writer.Write(thoughtData)

	writer.Flush()

	err = writer.Error()
	if err != nil {
		return err
	}

	fmt.Println("🤔 Thought added successfully")

	return nil
}

func ListAllThoughts(filepath string) error {
	f, e := os.Open(filepath)
	if e != nil {
		return e
	}
	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for i, r := range records {
		if i != 0 {
			fmt.Printf("%v: %v\n", i, r[1])
		}

	}

	return nil
}
