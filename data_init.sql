
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'todo') THEN
        CREATE DATABASE todo;
    END IF;
END $$;

\c todo;

DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'users') THEN
        CREATE TABLE if not exists users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL UNIQUE,
            password TEXT NOT NULL,
            salt TEXT NOT NULL,
            status VARCHAR(50) DEFAULT 'active',
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            deleted_at TIMESTAMP DEFAULT NULL
        );
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'tasks') THEN
        CREATE TABLE if not exists tasks (
            id SERIAL PRIMARY KEY,
            user_id INT NOT NULL,
            title VARCHAR(255) NOT NULL,
            description TEXT,
            progress VARCHAR(100),
            priority VARCHAR(50),
            status VARCHAR(50) DEFAULT 'active',
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            deleted_at TIMESTAMP DEFAULT NULL,

            CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) 
                ON DELETE CASCADE ON UPDATE CASCADE
        );
    END IF;
END $$;

DO $$
BEGIN
INSERT INTO users ("name","password",salt,status,created_at,updated_at,deleted_at) VALUES
	 ('ducpv','2065c535540f26988fd82659ff3adee6','DXevdSYTNNqMLGYCQyalPNCKpgVsDnhpJMIDqaVcgieXxcLxKm','active','2024-12-23 02:39:09.017','2024-12-23 02:39:09.017',NULL),
	 ('admin','fcbc5b66e238ff1819faa9eb8a4c4cfc','btVLDgCndbCRKGiyVgnLmSMMfJXqeXgVNXYFyagIhPCngMIOVJ','active','2024-12-22 18:32:17.315','2024-12-22 18:32:17.315',NULL);

INSERT INTO tasks (user_id,title,description,progress,priority,status,created_at,updated_at,deleted_at) VALUES
	 (1,'Plan Architecture and Dependencies','Define the architecture (e.g., layered, hexagonal) to ensure simplicity and maintainability. Decide how gRPC services and the database will interact using SQLc/Gorm. Set up a dependency list including Docker, Golang packages, and database.','Done','High','active','2024-12-23 02:42:51.154','2024-12-23 03:01:54.042',NULL),
	 (1,'Set Up Development Environment','Create a Dockerfile for the application and a docker-compose.yml to run the database container, Scaffold a basic Go project structure, including separate folders for proto files, handlers, models, and configs','Done','High','active','2024-12-23 02:49:12.298','2024-12-23 02:49:12.298',NULL),
	 (1,'Implement Authentication Module','Develop gRPC services for user sign-up and login, including password hashing and JWT-based authentication. Test endpoints locally with mock data','Done','High','active','2024-12-23 02:49:03.920','2024-12-23 02:49:03.920',NULL),
	 (1,'Implement CRUD Operations for Tasks','Develop gRPC services for Create, Read, Update, and Delete tasks. Use SQLc/Gorm to interact with the database. Include basic validations (e.g., task ownership).','Done','Medium','active','2024-12-23 02:49:03.920','2024-12-23 02:49:03.920',NULL),
	 (1,'Write a README File','Draft a concise README covering: Steps to run the app using Docker, Sample curl commands for user login and task management, Highlights and limitations of the implementation.','Done','Medium','active','2024-12-23 02:49:03.920','2024-12-23 02:49:03.920',NULL),
	 (1,'Test and Optimize','Conduct manual and automated tests for gRPC services. Optimize Docker configuration for smooth local deployment. Add comments and refactor code for clarity','Done','Medium','active','2024-12-23 02:49:07.788','2024-12-23 02:49:07.788',NULL);
END $$;
