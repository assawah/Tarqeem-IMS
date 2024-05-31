package db

import (
	"database/sql"
	"log"
)

var tableCreationQueries = []string{
	`CREATE TABLE IF NOT EXISTS comments (id integer NOT NULL PRIMARY KEY AUTOINCREMENT, content text NOT NULL, created_at datetime NOT NULL);`,
	`CREATE TABLE IF NOT EXISTS disciplines (id integer NOT NULL PRIMARY KEY AUTOINCREMENT, name text NOT NULL);`,
	`CREATE TABLE IF NOT EXISTS files (id integer NOT NULL PRIMARY KEY AUTOINCREMENT, file_path text NULL, file_name text NULL, file_size integer NULL DEFAULT (0));`,
	`CREATE TABLE IF NOT EXISTS issue_comments (issue_id integer NOT NULL, comment_id integer NOT NULL, PRIMARY KEY (issue_id, comment_id), CONSTRAINT issue_comments_issue_id FOREIGN KEY (issue_id) REFERENCES issues (id) ON DELETE CASCADE, CONSTRAINT issue_comments_comment_id FOREIGN KEY (comment_id) REFERENCES comments (id) ON DELETE CASCADE);`,
	`CREATE TABLE IF NOT EXISTS issue_files (issue_id integer NOT NULL, file_id integer NOT NULL, PRIMARY KEY (issue_id, file_id), CONSTRAINT issue_files_issue_id FOREIGN KEY (issue_id) REFERENCES issues (id) ON DELETE CASCADE, CONSTRAINT issue_files_file_id FOREIGN KEY (file_id) REFERENCES files (id) ON DELETE CASCADE);`,
	`CREATE TABLE IF NOT EXISTS issues (id integer NOT NULL PRIMARY KEY AUTOINCREMENT, title text NOT NULL, description text NOT NULL, creator text NOT NULL DEFAULT ('guest'), status text NOT NULL DEFAULT ('Pending'), date text NOT NULL DEFAULT ('24 Apr 2024'), project_issues integer NULL, CONSTRAINT issues_projects_issues FOREIGN KEY (project_issues) REFERENCES projects (id) ON DELETE SET NULL);`,
	`CREATE TABLE IF NOT EXISTS project_members (project_id integer NOT NULL, user_id integer NOT NULL, PRIMARY KEY (project_id, user_id), CONSTRAINT project_members_project_id FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE, CONSTRAINT project_members_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE);`,
	`CREATE TABLE IF NOT EXISTS projects (id integer NOT NULL PRIMARY KEY AUTOINCREMENT, name text NOT NULL, owner text NOT NULL, location text NOT NULL, type text NOT NULL, project_nature text NOT NULL, delivery_strategies text NOT NULL, state text NOT NULL, contracting_strategies text NOT NULL, dollar_value integer NOT NULL, execution_location text NOT NULL, number_of_top_level_scope_packages integer NOT NULL DEFAULT 0, number_of_joint_venture_partners integer NOT NULL DEFAULT 0, number_of_involved_interface_stakeholders integer NOT NULL DEFAULT 0);`,
	`CREATE TABLE IF NOT EXISTS user_coordinator_of_project (user_id integer NOT NULL, project_id integer NOT NULL, PRIMARY KEY (user_id, project_id), CONSTRAINT user_coordinator_of_project_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE, CONSTRAINT user_coordinator_of_project_project_id FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE);`,
	`CREATE TABLE IF NOT EXISTS user_leader_of_project (user_id integer NOT NULL, project_id integer NOT NULL, PRIMARY KEY (user_id, project_id), CONSTRAINT user_leader_of_project_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE, CONSTRAINT user_leader_of_project_project_id FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE);`,
	`CREATE TABLE IF NOT EXISTS users (id integer NOT NULL PRIMARY KEY AUTOINCREMENT, name text NULL, password text NULL, email text NOT NULL, username text NOT NULL, phone text NULL, created_at datetime NOT NULL, organization text NULL, is_active bool NOT NULL DEFAULT (true), type text NOT NULL);`,
	`CREATE TABLE IF NOT EXISTS interface_point (interface_point_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, interface_categories TEXT, project TEXT, title TEXT, budget_item_number TEXT, disciplines TEXT, area TEXT, system TEXT, raci_matrix TEXT, attach_file TEXT, description TEXT, status TEXT, create_date DATETIME, issue_date DATETIME, close_date DATETIME, recording_comments TEXT, boq TEXT, activities TEXT, interface_scope TEXT, interface_type TEXT);`,
}

func InitDB(db *sql.DB) error {
	// Execute each table creation statement
	for _, query := range tableCreationQueries {
		if _, err := db.Exec(query); err != nil {
			log.Fatalf("Failed to execute query: %v\nQuery: %s", err, query)
			return err
		}
	}

	log.Println("Database initialized successfully.")
	return nil
}
