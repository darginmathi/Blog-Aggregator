package main

import (
	"context"
	"fmt"
	"time"

	"github.com/darginmathi/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	user_id := user.ID
	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user_id,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed created successfullt!")

	printFeed(feed, user)
	fmt.Println()
	fmt.Println("=====================================")
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		printFeed(feed, user)
		fmt.Println("=====================================")
	}
	return nil
}

//helper funcs

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf(" * ID:  		%v\n", feed.ID)
	fmt.Printf("* Created:      %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:      %v\n", feed.UpdatedAt)
	fmt.Printf(" * Name: 		%v\n", feed.Name)
	fmt.Printf(" * Url: 		%v\n", feed.Url)
	fmt.Printf(" * User:	 	%v\n", user.Name)
}
