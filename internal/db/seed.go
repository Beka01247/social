package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/Beka01247/social/internal/store"
)

var usernames = []string{
	"beka", "madina", "alikhan", "darmen", "alisher", "aslan", "aldiyar",
	"askar", "eladil", "idris", "daniyal", "daryn", "kairat", "nurzhan",
	"arsen", "asan", "nariman", "arman", "sanzhar",
}

var titles = []string{
	"10 Tips for Boosting Your Productivity in 2025",
	"The Ultimate Guide to Mindfulness and Meditation",
	"Why Remote Work is Here to Stay",
	"How to Start Your Own Side Hustle This Year",
	"The Future of AI: What to Expect by 2030",
	"Top 5 Destinations for Your Next Adventure",
	"Mastering Time Management: A Step-by-Step Guide",
	"The Power of Positive Thinking in Daily Life",
	"Healthy Habits to Transform Your Morning Routine",
	"A Beginner’s Guide to Investing in Cryptocurrency",
	"How to Build a Personal Brand from Scratch",
	"The Art of Minimalism: Declutter Your Life",
	"Why Learning a New Skill Can Change Your Career",
	"Exploring the Benefits of Plant-Based Diets",
	"The Rise of Sustainable Fashion: What You Need to Know",
	"How to Create a Budget That Actually Works",
	"The Science Behind Better Sleep and How to Achieve It",
	"Navigating Career Changes in a Fast-Paced World",
	"The Best Apps for Staying Organized in 2025",
	"How to Cultivate Stronger Relationships in the Digital Age",
}

var contents = []string{
	"Discover practical strategies to skyrocket your productivity in 2025, from time-blocking techniques to leveraging cutting-edge apps for maximum efficiency.",
	"Learn the essentials of mindfulness and meditation with this comprehensive guide, including step-by-step exercises to reduce stress and improve focus.",
	"Explore why remote work continues to dominate, with insights into its benefits, challenges, and tips for thriving in a virtual workplace.",
	"Ready to turn your passion into profit? This post breaks down how to launch a successful side hustle, from ideation to execution.",
	"Get a glimpse into the future of artificial intelligence, exploring predictions and trends that could shape our world by 2030.",
	"From tropical escapes to urban adventures, check out the top five must-visit destinations for an unforgettable travel experience.",
	"Master your schedule with proven time management strategies, including prioritization tips and tools to stay on top of your tasks.",
	"Unlock the benefits of positive thinking and learn how small mindset shifts can lead to big changes in your personal and professional life.",
	"Transform your mornings with these healthy habits, designed to boost energy, improve focus, and set a positive tone for the day.",
	"Dive into the world of cryptocurrency with this beginner’s guide, covering the basics of investing and tips to avoid common pitfalls.",
	"Build a personal brand that stands out with actionable steps, from defining your unique value to growing your online presence.",
	"Embrace minimalism with practical advice on decluttering your space, simplifying your life, and finding joy in less.",
	"Discover why learning a new skill can open doors in your career, with tips on choosing the right skill and mastering it quickly.",
	"Explore the health and environmental benefits of plant-based diets, plus easy tips to transition to a more sustainable way of eating.",
	"Learn how sustainable fashion is reshaping the industry, with tips on making eco-friendly choices without sacrificing style.",
	"Take control of your finances with a step-by-step guide to creating a budget that aligns with your goals and lifestyle.",
	"Struggling to sleep? Uncover the science behind better rest and practical tips to improve your sleep quality tonight.",
	"Navigate career transitions with confidence using expert advice on pivoting industries, upskilling, and seizing new opportunities.",
	"Stay organized in 2025 with our roundup of the best apps for task management, note-taking, and keeping your life on track.",
	"Strengthen your relationships in the digital age with strategies for meaningful communication and staying connected despite busy schedules.",
}

var tags = []string{
	"productivity",
	"mindfulness",
	"remote work",
	"side hustle",
	"artificial intelligence",
	"travel",
	"time management",
	"positive thinking",
	"healthy habits",
	"cryptocurrency",
	"personal branding",
	"minimalism",
	"learning",
	"plant-based diet",
	"sustainable fashion",
	"budgeting",
	"sleep",
	"career change",
	"organization",
	"relationships",
}

var comments = []string{
	"Great tips on productivity! The time-blocking method really helped me get more done this week.",
	"This mindfulness guide is so helpful. The breathing exercises are easy to follow and super calming!",
	"Love the remote work insights. Any tips for staying motivated when working from home long-term?",
	"Starting a side hustle sounds exciting, but I’m nervous about the risks. Any advice for beginners?",
	"Fascinating read on AI! I’m curious about how it’ll impact small businesses in the next few years.",
	"Those travel destinations are now on my list! Which one would you recommend for a solo trip?",
	"The budgeting tips here are spot-on. Finally got a plan that works for my irregular income!",
	"Minimalism is life-changing! Thanks for the practical steps to start decluttering my space.",
	"The plant-based diet tips are awesome. Any favorite recipes you’d recommend for beginners?",
	"Thanks for the career change advice! I feel more confident about switching industries now.",
}

func Seed(store store.Storage, db *sql.DB) error {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Printf("Error creating user: %v", err)
			return err
		}
	}

	tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Printf("Error creating posts: %v", err)
			return err
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Commnets.Create(ctx, comment); err != nil {
			log.Printf("Error creating comments: %v", err)
			return err
		}
	}

	log.Println("seeding complete")
	return nil
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
