package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"math/rand"

	"github.com/kandlagifari/go-nakama-apps/internal/store"
)

var usernames = []string{
	"phoenix", "tiger", "sparrow", "lynx", "falcon", "panther", "orca", "wolf",
	"koala", "eagle", "otter", "raven", "bear", "lion", "puma", "whale",
	"jaguar", "ibis", "cheetah", "shark", "owl", "fox", "hare", "gecko",
	"ferret", "beetle", "python", "hawk", "bison", "moose", "zebra", "elk",
	"antelope", "walrus", "penguin", "mongoose", "viper", "dolphin", "ram", "boar",
	"swallow", "alpaca", "bat", "crane", "frog", "lemur", "mole", "jay",
	"quokka", "tamarin", "gazelle",
}

var titles = []string{
	"Adventures in Coding", "Mastering Microservices", "Docker for Dummies",
	"Cloud Security Essentials", "Scaling with Kubernetes", "Intro to Terraform",
	"Secrets of CI/CD", "Serverless Demystified", "PostgreSQL Tips",
	"The Art of Automation", "Navigating the Cloud", "From Zero to Hero in Go",
	"Cybersecurity for All", "Data Engineering 101", "Building APIs with Gin",
	"Embracing Open Source", "Optimizing Costs in AWS", "Git Best Practices",
	"Monitoring with Prometheus", "Secrets of SRE",
}

var contents = []string{
	"Explore the world of coding with tips and tricks to enhance your skills.",
	"Learn how microservices architecture can improve application scalability.",
	"A beginner-friendly guide to containerizing applications with Docker.",
	"Understand the fundamentals of securing cloud infrastructures effectively.",
	"Discover how Kubernetes simplifies application deployment and scaling.",
	"Step-by-step guidance to master infrastructure automation with Terraform.",
	"Unlock the secrets of building efficient CI/CD pipelines.",
	"Serverless computing made simple: learn how to get started.",
	"Boost your database skills with practical PostgreSQL techniques.",
	"Automate mundane tasks and improve productivity with these insights.",
	"Dive into cloud platforms and navigate their complexities with ease.",
	"A complete guide to mastering the Go programming language.",
	"Protect your digital assets with these essential cybersecurity tips.",
	"Get started with data engineering and learn the tools of the trade.",
	"Learn how to build fast and scalable APIs using the Gin framework.",
	"Discover the benefits of contributing to open-source projects.",
	"Practical ways to reduce AWS costs and optimize spending.",
	"Enhance your version control skills with these Git best practices.",
	"Learn how to monitor and maintain systems using Prometheus.",
	"Understand the core principles of Site Reliability Engineering.",
}

var tags = []string{
	"Technology", "Cloud Computing", "Programming", "Automation", "Security",
	"Database", "Development", "Infrastructure", "Open Source", "Monitoring",
	"Microservices", "Serverless", "SRE", "DevOps", "Cost Optimization",
	"Terraform", "Git", "PostgreSQL", "CI/CD", "Kubernetes",
}

var comments = []string{
	"This is exactly what I needed, thank you!",
	"Such an informative post, appreciate the effort!",
	"I never thought of it that way, very enlightening.",
	"This is incredibly helpful, bookmarking for later.",
	"Great write-up, looking forward to more like this!",
	"Thank you for simplifying such a complex topic.",
	"Insightful and well-explained, kudos!",
	"Learned something new today, much appreciated.",
	"Your posts are always so detailed and helpful, thanks!",
	"Fantastic explanation, cleared up all my doubts.",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("Error creating user:", err)
			return
		}
	}

	tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			RoleID:   1,
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
			Content: titles[rand.Intn(len(contents))],
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
